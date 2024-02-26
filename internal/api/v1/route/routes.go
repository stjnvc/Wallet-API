package route

import (
	"github.com/gin-gonic/gin"
	"github.com/stjnvc/wallet-api/internal/api/v1/handler"
)

func Setup(router *gin.Engine) {
	router.GET("/api/v1/wallets/:wallet_id/balance", handler.BalanceHandler)
	router.GET("/api/v1/wallets/:wallet_id/credit", handler.CreditHandler)
	router.POST("/api/v1/wallets/:wallet_id/debit", handler.DebitHandler)
}
