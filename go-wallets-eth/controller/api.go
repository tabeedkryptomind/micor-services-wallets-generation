package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"eth-wallet/wallets"
)


func GenerateWallets(ctx *gin.Context) {
	w, err := wallets.GetAddress()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot create wallet",
			"status":  http.StatusInternalServerError,
		})
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status": http.StatusOK,
		"wallet": w,
	})
}