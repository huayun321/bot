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
	lastTx   int64
	lastPath string
}

func (tl *TxLimiter) check(path string) bool {
	tl.Lock()
	defer tl.Unlock()
	now := time.Now().Unix()
	if tl.lastPath == path && now-tl.lastTx <= 0 {
		return false
	}
	tl.lastTx = now
	tl.lastPath = path
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

	go s.pendingNonce()

	return s
}

func (s *Swap) pendingNonce() {
	t := time.NewTicker(100 * time.Millisecond)
	for range t.C {
		nonce, err := s.b.rpc.PendingNonceAt(context.Background(), s.public)
		if err != nil {
			log.Fatal(err)
		}
		s.nonce = nonce
	}
}

func (s *Swap) startTx(amountIn, amountOut, price *big.Int, path []common.Address, pathName, uid string) {
	if ok := s.tl.check(pathName); !ok {
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
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = s.sc.Limit // in units
	auth.GasPrice = price

	deadLine := time.Now().Unix() + s.sc.Dead
	dlb := big.NewInt(deadLine)
	tx, err := s.router.Hua(auth, amountIn, amountOut, path, dlb)
	if err != nil {
		log.Println(uid, pathName, "SwapExactTokensForTokens", err, price)
		return
	}
	log.Println(uid, pathName, tx.Hash(), price)
}
