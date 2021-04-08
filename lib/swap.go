package lib

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/huayun321/bot/global"
	"github.com/huayun321/bot/pancake/router"
	"log"
	"math/big"
	"sync"
	"time"
)

type Swap struct {
	b          *Bot
	cid        *big.Int
	sc         *swapConfig
	nonce      uint64
	privateKey *ecdsa.PrivateKey
	public     common.Address
	gasPrice   *big.Int
	router     *router.Router
	tl         *TxLimiter
	limit      int64
}

type TxLimiter struct {
	sync.Mutex
	lastTx int64
}

func (tl *TxLimiter) check() bool {
	tl.Lock()
	defer tl.Unlock()
	now := time.Now().Unix()
	if now-tl.lastTx <= 0 {
		return false
	}
	tl.lastTx = now
	return true
}

func newSwap(b *Bot) *Swap {
	cid, err := b.rpc.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA(global.AccountSetting.Private)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := b.rpc.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress(global.ContractSetting.Router)
	instance, err := router.NewRouter(address, b.rpc)
	if err != nil {
		log.Fatal(err)
	}
	s := &Swap{
		b:          b,
		cid:        cid,
		sc:         b.swapConfig,
		nonce:      nonce,
		privateKey: privateKey,
		public:     fromAddress,
		gasPrice:   b.swapConfig.Price,
		router:     instance,
		tl:         &TxLimiter{},
		limit:      200,
	}
	return s
}

func (s *Swap) pendingNonce() {
	nonce, err := s.b.rpc.PendingNonceAt(context.Background(), s.public)
	if err != nil {
		log.Fatal(err)
	}
	s.nonce = nonce
}

func (s *Swap) startTx(amountIn, amountOut *big.Int, path []common.Address) {
	if ok := s.tl.check(); !ok {
		return
	}
	//if atomic.LoadInt64(&s.limit) <= 0 {
	//	log.Println(s.limit)
	//	return
	//}
	//atomic.AddInt64(&s.limit, -1)
	auth, err := bind.NewKeyedTransactorWithChainID(s.privateKey, s.cid)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(s.nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(360000) // in units
	auth.GasPrice = s.gasPrice

	deadLine := time.Now().Unix() + 600
	dlb := big.NewInt(deadLine)

	go s.pendingNonce()

	tx, err := s.router.SwapExactTokensForTokens(auth, amountIn, amountOut, path, s.public, dlb)
	if err != nil {
		log.Println("SwapExactTokensForTokens", err)
		return
	}
	log.Println(tx.Hash())
}
