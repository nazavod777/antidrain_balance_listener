package core

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/valyala/fasthttp"
	"log"
	"main/custom_types"
	"main/utils"
	"math/big"
)

func ListenBalance(providerClient *fasthttp.Client,
	configData custom_types.ConfigStruct,
	drainedAccountAddress common.Address,
	drainedAccountKeyFormatted *ecdsa.PrivateKey,
	rpcUrl string) {

	chainId, _ := utils.GetChainID(providerClient, rpcUrl)
	chainIdFormatted := big.NewInt(int64(chainId))
	saveFundsToAddressFormatted := common.HexToAddress(configData.SaveFundsToAddress)

	for {
		accountBalanceInt, _ := utils.GetBalance(providerClient, rpcUrl, drainedAccountAddress.String())
		accountBalance := big.NewInt(int64(accountBalanceInt))
		log.Printf("%s | Current Balance: %.18f ETH", drainedAccountAddress.String(), float64(accountBalanceInt)/1e18)

		currentGwei, _ := utils.GetGwei(providerClient, rpcUrl)
		gasPrice := big.NewInt(int64(float32(currentGwei) * configData.GweiMultiplier))

		gasLimit := uint64(21000)
		gasCost := new(big.Int).Mul(gasPrice, big.NewInt(int64(gasLimit)))

		if accountBalance.Cmp(gasCost) < 0 {
			log.Printf("Insufficient balance for gas fees")
			continue
		}

		transferAmount := new(big.Int).Sub(accountBalance, gasCost)
		accountNonce, _ := utils.GetNonce(providerClient, rpcUrl, drainedAccountAddress.String())

		transferTransaction := types.LegacyTx{
			Nonce:    accountNonce,
			GasPrice: gasPrice,
			Gas:      gasLimit,
			To:       &saveFundsToAddressFormatted,
			Value:    transferAmount,
			Data:     nil,
		}

		signedTransaction, err := types.SignTx(types.NewTx(&transferTransaction), types.NewEIP155Signer(chainIdFormatted), drainedAccountKeyFormatted)
		if err != nil {
			log.Printf("Failed to sign transfer transaction: %v", err)
			continue
		}

		rawTransactionHex, _ := signedTransaction.MarshalBinary()
		rawTransactionData := "0x" + common.Bytes2Hex(rawTransactionHex)

		transferTransactionHash, err := utils.SendTransaction(providerClient, rpcUrl, rawTransactionData)

		if err != nil {
			continue
		}
		log.Printf("Transfer transaction hash: %s | Amount: %.18f ETH", transferTransactionHash, float64(transferAmount.Int64())/1e18)
	}
}
