package main

import (
	"bufio"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/valyala/fasthttp"
	"log"
	"main/core"
	"main/utils"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Unknown Error: %s\n", r)
			fmt.Printf("Press Enter to Exit...")
			fmt.Scanln() // Ожидаем ввода
		}
	}()

	configData := utils.ReadConfig()

	if strings.HasPrefix(configData.DrainedPrivateKey, "0x") {
		configData.DrainedPrivateKey = configData.DrainedPrivateKey[2:]
	}

	log.Printf("The Config File Was Successfully Read, Total RPC URL's: %d", len(configData.RpcURLs))

	fmt.Println("\nPress Enter to start...")
	_, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
	if err != nil {
		log.Panicf("Error reading input: %s", err)
	}

	drainedAccountKeyFormatted, _ := crypto.HexToECDSA(configData.DrainedPrivateKey)
	drainedAccountPublicKey := drainedAccountKeyFormatted.PublicKey
	drainedAccountAddress := crypto.PubkeyToAddress(drainedAccountPublicKey)

	providerClient := &fasthttp.Client{
		ReadTimeout:         5 * time.Second,
		WriteTimeout:        5 * time.Second,
		MaxConnsPerHost:     100,
		MaxIdleConnDuration: 0,
	}

	var wg sync.WaitGroup
	for _, rpcUrl := range configData.RpcURLs {
		wg.Add(1)
		go func(rpcUrl string) {
			defer wg.Done()
			core.ListenBalance(providerClient, configData, drainedAccountAddress, drainedAccountKeyFormatted, rpcUrl)
		}(rpcUrl)
	}

	wg.Wait()
}
