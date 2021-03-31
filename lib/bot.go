package lib

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/huayun321/bot/global"
	"github.com/huayun321/bot/pancake/factory"
	"log"
	"strings"
)

type Bot struct {
	rpc             *ethclient.Client
	ws              *ethclient.Client
	factoryInstance *factory.Factory
	allPairs        map[string]*BPair
}

func NewBot() *Bot {
	b := &Bot{
		allPairs: make(map[string]*BPair),
	}
	rpc, err := ethclient.Dial(global.ServerSetting.RPC)
	if err != nil {
		log.Fatal(err)
	}
	b.rpc = rpc
	ws, err := ethclient.Dial(global.ServerSetting.WS)
	if err != nil {
		log.Println("not connected")
		log.Fatal(err)
	}
	b.ws = ws
	address := common.HexToAddress(global.ContractSetting.Factory)
	instance, err := factory.NewFactory(address, rpc)
	if err != nil {
		log.Fatal(err)
	}
	b.factoryInstance = instance
	return b
}

func (b *Bot) Run() {
	b.generatePath()
	b.allPairsRun()
	select {}
}

func (b *Bot) generatePath() {
	symbol := global.TokesSetting.Symbol
	address := global.TokesSetting.Address
	for k, v := range global.TokesSetting.Others {
		name := strings.Join([]string{symbol, k}, "-")
		ta := common.HexToAddress(address)
		tb := common.HexToAddress(v)
		p := NewBPair(b, name, ta, tb)
		b.allPairs[name] = p
	}
	log.Println(b.allPairs)
}

func (b *Bot) allPairsRun() {
	for n, p := range b.allPairs {
		log.Println(n, "run...")
		p.run()
	}
}
