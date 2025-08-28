package services

import (
	"context"
	_ "embed"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

//go:embed erc20.abi.json
var erc20ABI string

// GetERC20Balance fetches the balance of an ERC20 token for an address
func GetERC20Balance(client *ethclient.Client, tokenAddress string, walletAddress string) (*big.Int, error) {
	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(erc20ABI))
	if err != nil {
		return nil, err
	}

	// Bind contract
	contract := bind.NewBoundContract(
		common.HexToAddress(tokenAddress),
		parsedABI,
		client,
		client,
		client,
	)

	// Call balanceOf
	var out []interface{}
	err = contract.Call(
		&bind.CallOpts{Context: context.Background()},
		&out,
		"balanceOf",
		common.HexToAddress(walletAddress),
	)
	if err != nil {
		return nil, err
	}

	// Decode result
	if len(out) == 0 {
		return big.NewInt(0), nil
	}

	balance, ok := out[0].(*big.Int)
	if !ok {
		log.Println("⚠️ unexpected type from contract.Call:", out[0])
		return big.NewInt(0), nil
	}

	return balance, nil
}
