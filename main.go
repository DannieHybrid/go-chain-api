package main

import (
	"github.com/DannieHybrid/go-chain-api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/balance/:address", handlers.GetBalance)
	r.GET("/token/:token/:address", handlers.GetTokenBalance) // ✅ token first, address second

	r.Run(":8080")
}
