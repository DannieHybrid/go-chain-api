package handlers

import (
	"context"
	"log"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

const infuraURL = "https://mainnet.infura.io/v3/YOUR_INFURA_KEY"

// ERC20 ABI snippet for balanceOf
const erc20ABI = `[{"constant":true,"inputs":[{"name":"_owner","type":"address"}],
"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"type":"function"}]`

func GetTokenBalance(c *gin.Context) {
	address := c.Param("address")
	token := c.Param("token")

	if !common.IsHexAddress(address) || !common.IsHexAddress(token) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Ethereum or token address"})
		return
	}

	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	parsedABI, err := abi.JSON(strings.NewReader(erc20ABI))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ABI parse failed"})
		return
	}

	contractAddr := common.HexToAddress(token)
	callData, err := parsedABI.Pack("balanceOf", common.HexToAddress(address))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ABI pack failed"})
		return
	}

	result, err := client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contractAddr,
		Data: callData,
	}, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Contract call failed"})
		return
	}

	var balance = new(big.Int)
	err = parsedABI.UnpackIntoInterface(&balance, "balanceOf", result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unpack failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"address": address,
		"token":   token,
		"balance": balance.String(),
	})
}
