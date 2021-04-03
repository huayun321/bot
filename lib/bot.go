package lib

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/huayun321/bot/global"
	"github.com/huayun321/bot/pancake/factory"
	"github.com/kokardy/listing"
	"log"
	"math/big"
)

type swapConfig struct {
	Amount *big.Int
	Profit *big.Int
	Price  *big.Int
	Cost   *big.Int
}

type Bot struct {
	rpc             *ethclient.Client
	ws              *ethclient.Client
	factoryInstance *factory.Factory
	allPairs        map[string]*BPair
	swapConfig      *swapConfig
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

	b.swapConfig = parseSwapSetting()
	return b
}

func (b *Bot) Run() {
	b.generatePath()
	//b.allPairsRun()
	//select {}
}

func (b *Bot) generatePath() {
	//symbol := global.TokesSetting.Symbol
	//address := global.TokesSetting.Address
	//ta := common.HexToAddress(address)

	symbols := make([]string, 0)
	addresses := make([]common.Address, 0)
	//path := make([]common.Address, 0)

	for k, v := range global.TokesSetting.Others {
		symbols = append(symbols, k)
		addresses = append(addresses, common.HexToAddress(v))
	}
	ss := listing.StringReplacer(symbols)
	fmt.Println("Permutations")
	for perm := range listing.Permutations(ss, 2, false, 5) {
		fmt.Println(perm.(listing.StringReplacer))
	}

	log.Println(b.allPairs)
}

func (b *Bot) allPairsRun() {
	for n, p := range b.allPairs {
		log.Println(n, "run...")
		p.run()
	}
}

func parseSwapSetting() *swapConfig {
	amountIn, ok := new(big.Int).SetString(global.SwapSetting.Amount, 10)
	if !ok {
		log.Fatal(errors.New("parse amountIn error"))
	}
	profit, ok := new(big.Int).SetString(global.SwapSetting.Profit, 10)
	if !ok {
		log.Fatal(errors.New("parse profit error"))
	}
	price, ok := new(big.Int).SetString(global.SwapSetting.Price, 10)
	if !ok {
		log.Fatal(errors.New("parse price error"))
	}
	cost, ok := new(big.Int).SetString(global.SwapSetting.Cost, 10)
	if !ok {
		log.Fatal(errors.New("parse cost error"))
	}
	sc := &swapConfig{
		Amount: amountIn,
		Profit: profit,
		Price:  price,
		Cost:   cost,
	}
	return sc
}
