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
)

type BPair struct {
	bot      *Bot
	name     string
	address  common.Address
	tokenA   common.Address
	tokenB   common.Address
	reserve0 *big.Int
	reserve1 *big.Int
}

func NewBPair(bot *Bot, name string, tokenA, tokenB common.Address) *BPair {
	bp := &BPair{
		bot:    bot,
		name:   name,
		tokenA: tokenA,
		tokenB: tokenB,
	}

	pairAddress, err := bot.factoryInstance.GetPair(&bind.CallOpts{}, tokenA, tokenB)
	if err != nil {
		log.Fatal(err)
	}
	bp.address = pairAddress
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
	contractAbi, err := abi.JSON(strings.NewReader(string(pair.PairABI)))
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			fmt.Println(bp.name, " ==== got an event =====")
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(bp.name, " ==== get an event =====")
			ps := &pair.PairSync{}
			err := contractAbi.UnpackIntoInterface(ps, "Sync", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(ps) // pointer to event log
		}
	}
}
