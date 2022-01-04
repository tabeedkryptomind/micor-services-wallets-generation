package main

import (
	"github.com/gin-gonic/gin"

	"eth-wallet/controller"
)

func main() {

	route := gin.Default()
	route.GET("/api/eth/create-wallet/", controller.GenerateWallets)
	route.Run(":8080")
}
