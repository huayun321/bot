package lib

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/huayun321/bot/global"
	"github.com/huayun321/bot/pancake/factory"
	"github.com/kokardy/listing"
	"log"
	"math/big"
	"strings"
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
	allPairs        map[common.Address]*BPair
	allChains       map[string]*Chain
	swapConfig      *swapConfig
}

func NewBot() *Bot {
	b := &Bot{
		allPairs:  make(map[common.Address]*BPair),
		allChains: make(map[string]*Chain),
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
	paths := b.generatePath()
	b.setupChain(paths)
	log.Println(b.allPairs)
	log.Println(b.allChains)
	log.Println("all pair length ", len(b.allPairs))
	log.Println("all chain length ", len(b.allChains))

	b.allChainsRun()
	b.allPairsRun()
	select {}
}

func (b *Bot) generatePath() [][]string {
	symbol := global.TokensSetting.Symbol

	symbols := make([]string, 0)
	for k := range global.TokensSetting.Others {
		symbols = append(symbols, k)
	}

	paths := make([][]string, 0)
	ss := listing.StringReplacer(symbols)
	for perm := range listing.Permutations(ss, 2, false, 5) {
		path := make([]string, 0)
		path = append(path, symbol)
		for _, s := range perm.(listing.StringReplacer) {
			path = append(path, s)
		}
		path = append(path, symbol)
		paths = append(paths, path)
	}
	log.Println("generatePath", paths)
	return paths
}

func (b *Bot) setupChain(paths [][]string) {
	allMap := global.TokensSetting.Others
	allMap[global.TokensSetting.Symbol] = global.TokensSetting.Address
	//convert path
	for _, path := range paths {
		//setup pair
		ps, err := b.setupPair(path)
		if err != nil {
			log.Println("setupChain pair ", err)
			continue
		}
		for k, v := range ps {
			if bp, ok := b.allPairs[v.address]; ok {
				ps[k] = bp
			}
		}
		//setup chain
		chainName := strings.Join(path, "->")
		paths := make([]common.Address, 0)
		for _, v := range path {
			paths = append(paths, common.HexToAddress(allMap[v]))
		}
		c := newChain(chainName, paths, ps, b.swapConfig)
		b.addPairs(ps)
		b.addChains(c)
	}
}

func (b *Bot) addPairs(bps []*BPair) {
	for _, bp := range bps {
		if _, ok := b.allPairs[bp.address]; ok {
			continue
		} else {
			b.allPairs[bp.address] = bp
		}
	}
}

func (b *Bot) addChains(c *Chain) {
	b.allChains[c.name] = c
}

func (b *Bot) setupPair(path []string) ([]*BPair, error) {
	allMap := global.TokensSetting.Others
	allMap[global.TokensSetting.Symbol] = global.TokensSetting.Address
	pairs := make([]*BPair, 0)
	for i := 0; i < len(path)-1; i++ {
		s := []string{path[i], path[i+1]}
		name := strings.Join(s, "-")
		bp, err := NewBPair(b, name, allMap[path[i]], allMap[path[i+1]])
		if err != nil {
			log.Println("pair name err ", name, err)
			return nil, err
		}
		pairs = append(pairs, bp)
	}
	return pairs, nil
}

func (b *Bot) allPairsRun() {
	for _, p := range b.allPairs {
		log.Println(p.name, "run...")
		p.run()
	}
}

func (b *Bot) allChainsRun() {
	for _, c := range b.allChains {
		log.Println(c.name, "run...")
		c.run()
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
