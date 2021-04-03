package lib

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/huayun321/bot/pancake/pair"
	"log"
	"math/big"
	"strings"
	"sync"
)

type BPair struct {
	rw          sync.Mutex
	bot         *Bot
	name        string
	address     common.Address
	tokenA      common.Address
	tokenB      common.Address
	contractAbi abi.ABI
	reserve0    *big.Int
	reserve1    *big.Int
	consumer    []*Chain
}

func NewBPair(bot *Bot, name, tokenA, tokenB string) (*BPair, error) {
	ta := common.HexToAddress(tokenA)
	tb := common.HexToAddress(tokenB)
	bp := &BPair{
		bot:      bot,
		name:     name,
		consumer: make([]*Chain, 0),
	}

	pairAddress, err := bot.factoryInstance.GetPair(&bind.CallOpts{}, ta, tb)
	if err != nil {
		log.Println("new pair name", name, " err ", err)
		return nil, err
	}
	bp.address = pairAddress

	err = bp.sortTokens(ta, tb)
	if err != nil {
		return nil, err
	}

	err = bp.checkReserve()
	if err != nil {
		return nil, err
	}

	contractAbi, err := abi.JSON(strings.NewReader(pair.PairABI))
	if err != nil {
		return nil, err
	}
	bp.contractAbi = contractAbi

	return bp, nil
}

func (bp *BPair) run() {
	go bp.subEvent()
}

func (bp *BPair) subEvent() {
	eventSignature := []byte("Sync(uint112,uint112)")
	hash := crypto.Keccak256Hash(eventSignature)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{bp.address},
		Topics:    [][]common.Hash{{hash}},
	}
	logs := make(chan types.Log)
	sub, err := bp.bot.ws.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			fmt.Println(bp.name, " ==== got an event =====")
			log.Fatal(err)
		case vLog := <-logs:
			go bp.handleEvent(vLog)
		}
	}
}

func (bp *BPair) handleEvent(event types.Log) {
	fmt.Println(bp.name, " ==== get an event =====")
	ps := &pair.PairSync{}
	err := bp.contractAbi.UnpackIntoInterface(ps, "Sync", event.Data)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("--------- handleEvent ", bp.name, ps)
	bp.updateReserve(ps)
	for _, v := range bp.consumer {
		v.pipe <- 1
	}
}

func (bp *BPair) updateReserve(ps *pair.PairSync) {
	bp.rw.Lock()
	defer bp.rw.Unlock()
	bp.reserve0 = ps.Reserve0
	bp.reserve1 = ps.Reserve1
}

func (bp *BPair) checkReserve() error {
	instance, err := pair.NewPair(bp.address, bp.bot.rpc)
	if err != nil {
		return err
	}
	rs, err := instance.GetReserves(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if rs.Reserve0.Cmp(big.NewInt(0)) <= 0 || rs.Reserve1.Cmp(big.NewInt(0)) <= 0 {
		return errors.New("reserve0 or reserve1 is 0")
	}
	bp.reserve0 = rs.Reserve0
	bp.reserve1 = rs.Reserve1
	return nil
}

func (bp *BPair) sortTokens(tokenA, tokenB common.Address) error {
	//check tokenA = t0
	instance, err := pair.NewPair(bp.address, bp.bot.rpc)
	if err != nil {
		return err
	}
	token0, err := instance.Token0(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if tokenA == token0 {
		bp.tokenA = tokenA
		bp.tokenB = tokenB
	} else {
		bp.tokenA = tokenB
		bp.tokenB = tokenA
	}
	return nil
}

func (bp *BPair) getReserve(tokenA common.Address) (*big.Int, *big.Int) {
	bp.rw.Lock()
	defer bp.rw.Unlock()
	if tokenA == bp.tokenA {
		return bp.reserve0, bp.reserve1
	} else {
		return bp.reserve1, bp.reserve0
	}
}

func (bp *BPair) addConsumer(chain *Chain) {
	bp.consumer = append(bp.consumer, chain)
}
