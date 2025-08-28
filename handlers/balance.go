package handlers

import (
	"net/http"

	"github.com/DannieHybrid/go-chain-api/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

// GetBalance handles GET /balance/:address
func GetBalance(c *gin.Context) {
	address := c.Param("address")
	if !common.IsHexAddress(address) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Ethereum address"})
		return
	}

	client := utils.NewClient("https://eth.llamarpc.com") // free RPC for testing

	balance, err := client.BalanceAt(c, common.HexToAddress(address), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"address": address,
		"balance": balance.String(),
	})
}
