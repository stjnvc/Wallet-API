package route

import (
	"github.com/gin-gonic/gin"
	"github.com/stjnvc/wallet-api/internal/api/v1/handler"
	"github.com/stjnvc/wallet-api/internal/api/v1/repository"
	"github.com/stjnvc/wallet-api/internal/api/v1/service"
	"gorm.io/gorm"
)

func Setup(router *gin.Engine, db *gorm.DB) {

	walletRepo := repository.NewWalletRepository(db)
	walletService := service.NewWalletService(walletRepo)
	walletHandler := handler.NewWalletHandler(walletService)

	routeGroup := router.Group("/api/v1/wallets")
	{
		routeGroup.GET("/:wallet_id/balance", walletHandler.GetBalance)
		routeGroup.POST("/:wallet_id/credit", walletHandler.Credit)
		routeGroup.POST("/:wallet_id/debit", walletHandler.Debit)
	}
}
