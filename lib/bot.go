package lib

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/huayun321/bot/pancake/factory"
	"github.com/huayun321/bot/setting"
	"log"
)

type Bot struct {
	rpc             *ethclient.Client
	ws              *ethclient.Client
	factoryInstance *factory.Factory
}

func NewBot(serverConfig *setting.ServerSetting, contractConfig *setting.ContractSetting) *Bot {
	b := &Bot{}
	rpc, err := ethclient.Dial(serverConfig.RPC)
	if err != nil {
		log.Fatal(err)
	}
	b.rpc = rpc
	ws, err := ethclient.Dial(serverConfig.WS)
	if err != nil {
		log.Println("not connected")
		log.Fatal(err)
	}
	b.ws = ws
	address := common.HexToAddress(contractConfig.Factory)
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
