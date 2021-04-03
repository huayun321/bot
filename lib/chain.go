package lib

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
)

type Chain struct {
	name   string
	path   []common.Address
	pairs  []*BPair
	pipe   chan int
	config *swapConfig
	want   *big.Int
	swap   *Swap
}

func newChain(name string, path []common.Address, pairs []*BPair, config *swapConfig, swap *Swap) *Chain {
	c := &Chain{
		name:   name,
		path:   path,
		pairs:  pairs,
		pipe:   make(chan int, 10),
		config: config,
		swap:   swap,
	}
	want := calculateWant(len(pairs), config)
	c.want = want
	return c
}

func (c *Chain) subEvent() {
	//i := 0
	for range c.pipe {
		//if i >= 1 {
		//	return
		//}
		go c.handleEvent()
		//i++
	}
}

func (c *Chain) run() {
	c.subscribe()
	go c.subEvent()
}

func (c *Chain) handleEvent() {
	//if c.name == "busd->wbnb->egg->busd" {
	//	return
	//}
	out, err := c.getAmountsOut(c.config.Amount)
	if err != nil {
		log.Println(err)
		return
	}
	ok := c.checkProfit(out)
	if ok {
		c.swap.startTx(c.config.Amount, out.Add(c.config.Amount, c.want), c.path)
	}
}

func (c *Chain) checkProfit(amountOut *big.Int) bool {
	result := new(big.Int)
	result.Sub(amountOut, c.config.Amount)
	//log.Println(c.name, c.config.Amount, amountOut, result, c.want)

	if result.Cmp(c.want) >= 0 {
		log.Println(c.name, c.config.Amount, amountOut, result, c.want)
		//swap
		return true
	}
	return false
}

func calculateWant(pairLength int, config *swapConfig) *big.Int {
	// contract * cost * price + profit
	want := new(big.Int)
	want.Mul(config.Cost, config.Price)
	want.Mul(want, big.NewInt(int64(pairLength)))
	want.Add(want, config.Profit)
	return want
}

func getAmountOut(amountIn, reserveIn, reserveOut *big.Int) (*big.Int, error) {
	if reserveIn == nil || reserveOut == nil || amountIn == nil {
		return nil, errors.New("nil args error")
	}
	//log.Println("===== amountOut rin rout ", reserveIn, " ",  reserveOut)
	if amountIn.Cmp(big.NewInt(0)) <= 0 {
		return nil, errors.New("amountIn 0")
	}
	if reserveIn.Cmp(big.NewInt(0)) <= 0 {
		return nil, errors.New("reserveIn 0")
	}
	if reserveOut.Cmp(big.NewInt(0)) <= 0 {
		return nil, errors.New("reserveOut 0")
	}
	amountInWithFee := new(big.Int)
	numerator := new(big.Int)
	denominator := new(big.Int)
	amountOut := new(big.Int)
	amountInWithFee.Mul(amountIn, big.NewInt(998))
	numerator.Mul(amountInWithFee, reserveOut)
	denominator.Add(new(big.Int).Mul(reserveIn, big.NewInt(1000)), amountInWithFee)
	amountOut.Div(numerator, denominator)
	return amountOut, nil
}

func (c *Chain) getAmountsOut(amountIn *big.Int) (*big.Int, error) {
	result := new(big.Int).Set(amountIn)
	var err error
	for i := 0; i < len(c.path)-1; i++ {
		r0, r1 := c.pairs[i].getReserve(c.path[i])
		//log.Println("getAmountsOut chain ", c.name, " pair ", c.pairs[i].name, " address ", c.pairs[i].address)
		//log.Println("getAmountsOut chain result r0 r1 ", result, r0, r1)
		result, err = getAmountOut(result, r0, r1)
		if err != nil {
			return nil, err
		}
		//log.Print("out i ", i, " -- ", result, "---")
	}
	return result, err
}

func (c *Chain) subscribe() {
	for _, p := range c.pairs {
		//log.Println(c.name, "sub pair", p.name)
		p.addConsumer(c)
	}
}
