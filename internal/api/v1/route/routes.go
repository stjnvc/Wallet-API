package route

import (
	"github.com/gin-gonic/gin"
	"github.com/stjnvc/wallet-api/internal/api/v1/handler"
	"github.com/stjnvc/wallet-api/internal/api/v1/middleware"
	"github.com/stjnvc/wallet-api/internal/api/v1/repository"
	"github.com/stjnvc/wallet-api/internal/api/v1/service"
	"gorm.io/gorm"
)

func Setup(router *gin.Engine, db *gorm.DB) {

	walletRepo := repository.NewWalletRepository(db)
	walletService := service.NewWalletService(walletRepo)
	walletHandler := handler.NewWalletHandler(walletService)

	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository)
	authHandler := handler.NewAuthHandler(authService)

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"api":     "Wallet API",
			"version": "1",
			"message": "Explore other routes defined in README",
		})
	})

	apiRoutes := router.Group("/api/v1")
	{
		authRoutes := apiRoutes.Group("/auth")
		{
			authRoutes.POST("/login", authHandler.Login)
			authRoutes.POST("/register", authHandler.Register)
		}

		walletRoutes := apiRoutes.Group("/wallets")
		walletRoutes.Use(middleware.AuthMiddleware())
		{
			walletRoutes.GET("/:wallet_id/balance", walletHandler.GetBalance)
			walletRoutes.POST("/:wallet_id/credit", walletHandler.Credit)
			walletRoutes.POST("/:wallet_id/debit", walletHandler.Debit)
		}
	}
}
