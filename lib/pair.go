package lib

import (
	"context"
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
	rw          sync.RWMutex
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

func NewBPair(bot *Bot, name string, path [2]common.Address) *BPair {
	bp := &BPair{
		bot:  bot,
		name: name,
	}

	pairAddress, err := bot.factoryInstance.GetPair(&bind.CallOpts{}, path[0], path[1])
	if err != nil {
		log.Fatal(err)
	}
	bp.address = pairAddress
	bp.sortTokens(path[0], path[1])

	contractAbi, err := abi.JSON(strings.NewReader(string(pair.PairABI)))
	if err != nil {
		log.Fatal(err)
	}
	bp.contractAbi = contractAbi

	return bp
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
	bp.rw.Lock()
	defer bp.rw.Unlock()
	bp.reserve0 = ps.Reserve0
	bp.reserve1 = ps.Reserve1
	//todo send to chain
	for _, v := range bp.consumer {
		v.pipe <- 1
	}
}

func (bp *BPair) sortTokens(tokenA, tokenB common.Address) {
	//check tokenA = t0
	instance, err := pair.NewPair(bp.address, bp.bot.rpc)
	if err != nil {
		log.Fatal(err)
	}
	token0, err := instance.Token0(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	if tokenA == token0 {
		bp.tokenA = tokenA
		bp.tokenB = tokenB
	} else {
		bp.tokenA = tokenB
		bp.tokenB = tokenA
	}
}

func (bp *BPair) getReserve(tokenA common.Address) (*big.Int, *big.Int) {
	bp.rw.RLock()
	defer bp.rw.RUnlock()
	if tokenA == bp.tokenA {
		return bp.reserve0, bp.reserve1
	} else {
		return bp.reserve1, bp.reserve0
	}
}
