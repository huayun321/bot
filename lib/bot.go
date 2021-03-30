package lib

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/huayun321/bot/pancake/factory"
	"log"
)

type Bot struct {
	rpc             *ethclient.Client
	ws              *ethclient.Client
	factoryInstance *factory.Factory
}

func NewBot() *Bot {
	b := &Bot{}
	rpc, err := ethclient.Dial("https://bsc-dataseed.binance.org")
	if err != nil {
		log.Fatal(err)
	}
	b.rpc = rpc
	ws, err := ethclient.Dial("wss://bsc-ws-node.nariox.org:443")
	if err != nil {
		log.Println("not connected")
		log.Fatal(err)
	}
	b.ws = ws
	address := common.HexToAddress("0xBCfCcbde45cE874adCB698cC183deBcF17952812")
	instance, err := factory.NewFactory(address, rpc)
	if err != nil {
		log.Fatal(err)
	}
	b.factoryInstance = instance
	return b
}

func (b *Bot) Run() {
	bp := NewBPair(b)
	bp.run()
	select {}
}
