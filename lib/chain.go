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
	config swapConfig
	want   *big.Int
}

func newChain(name string, path []common.Address, pairs []*BPair, config swapConfig) *Chain {
	c := &Chain{
		name:   name,
		path:   path,
		pairs:  pairs,
		pipe:   make(chan int, 10),
		config: config,
	}
	want := calculateWant(len(pairs), config)
	c.want = want
	return c
}

func (c *Chain) subEvent() {
	for v := range c.pipe {
		log.Println(v)
		go c.handleEvent()
	}
}

func (c *Chain) handleEvent() {
	out, err := c.getAmountsOut(c.config.Amount)
	if err != nil {
		log.Fatal(err)
		return
	}
	c.checkProfit(out)
}

func (c *Chain) checkProfit(amountOut *big.Int) {
	if amountOut.Cmp(c.want) > 0 {
		log.Println(c.name, c.config.Amount, amountOut)
		//swap
	}
}

func calculateWant(pairLength int, config swapConfig) *big.Int {
	// contract * cost * price + profit
	want := new(big.Int)
	want.Mul(config.Cost, config.Price)
	want.Mul(want, big.NewInt(int64(pairLength)))
	want.Add(want, config.Profit)
	return want
}

func getAmountOut(amountIn, reserveIn, reserveOut *big.Int) (*big.Int, error) {
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
	log.Println("===== amountOut", amountOut)
	return amountOut, nil
}

func (c *Chain) getAmountsOut(amountIn *big.Int) (*big.Int, error) {
	result := new(big.Int).Set(amountIn)
	var err error
	log.Println("getAmountsOut ", len(c.path))

	for i := 0; i < len(c.path)-1; i++ {
		log.Println("getAmountsOut i ", i)
		r0, r1 := c.pairs[i].getReserve(c.path[i])
		result, err = getAmountOut(result, r0, r1)
	}
	return result, err
}
