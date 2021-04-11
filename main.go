package main

import (
	"github.com/huayun321/bot/global"
	"github.com/huayun321/bot/lib"
	setting "github.com/huayun321/bot/setting"
	"log"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
}

func main() {
	b := lib.NewBot()
	b.Run()
	//log.Println(global.SwapSetting)
}

func setupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Contract", &global.ContractSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Tokens", &global.TokensSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Swap", &global.SwapSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Account", &global.AccountSetting)
	if err != nil {
		return err
	}
	return nil
}
