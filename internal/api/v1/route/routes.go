package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stjnvc/wallet-api/internal/api/v1/handlers"
)

func Setup(router *gin.Engine) {
	router.GET("/api/v1/wallets/:wallet_id/balance", handlers.BalanceHandler)
	router.GET("/api/v1/wallets/:wallet_id/credit", handlers.CreditHandler)
	router.POST("/api/v1/wallets/:wallet_id/debit", handlers.DebitHandler)
}
