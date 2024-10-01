package custom_types

type ConfigStruct struct {
	DrainedPrivateKey  string   `json:"drainedPrivateKey"`
	SaveFundsToAddress string   `json:"saveFundsToAddress"`
	GweiMultiplier     float32  `json:"gweiMultiplier"`
	RpcURLs            []string `json:"rpcUrls"`
}
