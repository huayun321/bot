package lib

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/huayun321/bot/pancake/pair"
	"log"
	"strings"
)

type BPair struct {
	bot *Bot
}

func NewBPair(bot *Bot) *BPair {
	bp := &BPair{bot: bot}
	return bp
}

func (bp *BPair) run() {
	bp.subEvent()
}

func (bp *BPair) subEvent() {
	contractAddress := common.HexToAddress("0xa527a61703d82139f8a06bc30097cc9caa2df5a6")
	eventSignature := []byte("Sync(uint112,uint112)")
	hash := crypto.Keccak256Hash(eventSignature)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
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
			fmt.Println("==== got an event =====")
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println("==== get an event =====")
			ps := &pair.PairSync{}
			err := contractAbi.UnpackIntoInterface(ps, "Sync", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(ps) // pointer to event log
		}
	}
}
