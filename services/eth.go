package services

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetETHBalance(address string) (string, error) {
	rpcURL := os.Getenv("ALCHEMY_RPC_URL")
	if rpcURL == "" {
		return "", fmt.Errorf("ALCHEMY_RPC_URL not set")
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Println("Error connecting to RPC:", err)
		return "", err
	}

	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return "", err
	}

	ethValue := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))
	return fmt.Sprintf("%f ETH", ethValue), nil
}
