package setting

type ServerSetting struct {
	RPC string
	WS  string
}

type ContractSetting struct {
	Factory string
	Router  string
}

type TokensSetting struct {
	Symbol  string
	Address string
	Others  map[string]string
}

type SwapSetting struct {
	Amount string
	Profit string
	Price  string
	Max    string
	Cost   string
	Limit  string
	Rate   string
	Dead   string
}

type AccountSetting struct {
	Public  string
	Private string
}
