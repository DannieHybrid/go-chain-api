package utils

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

// NewClient creates and returns an Ethereum client
func NewClient(rpcURL string) *ethclient.Client {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}
	return client
}
