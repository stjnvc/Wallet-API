package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stjnvc/wallet-api/internal/api/v1/service"
	"github.com/stjnvc/wallet-api/internal/util"
	"strconv"
	"time"
)

type WalletHandler struct {
	walletService *service.WalletService
}

func NewWalletHandler(walletService *service.WalletService) *WalletHandler {

	// Initialize Redis client
	err := util.InitRedisClient()
	if err != nil {
		logrus.Error("Err loging client")
	}

	return &WalletHandler{
		walletService: walletService,
	}
}

func (h *WalletHandler) GetBalance(c *gin.Context) {
	walletID, _ := strconv.Atoi(c.Param("wallet_id"))

	cacheKey := fmt.Sprintf("wallet:%d:balance", walletID)
	balanceStr, err := util.RedisClient.Get(c, cacheKey).Result()

	if err == nil {
		balance, _ := strconv.ParseFloat(balanceStr, 64)
		c.JSON(200, gin.H{
			"wallet_id": walletID,
			"balance":   balance,
		})

		return
	}

	balance, err := h.walletService.GetWalletBalance(walletID)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	util.RedisClient.Set(c, cacheKey, fmt.Sprintf("%.2f", balance), time.Hour)

	c.JSON(200, gin.H{
		"wallet_id": walletID,
		"balance":   balance,
	})
}

func (h *WalletHandler) Credit(c *gin.Context) {
	walletID, _ := strconv.Atoi(c.Param("wallet_id"))

	// Clear cache when updating balance
	cacheKey := fmt.Sprintf("wallet:%d:balance", walletID)
	util.RedisClient.Del(c, cacheKey)

	amount, _ := strconv.ParseFloat(c.PostForm("amount"), 64)
	if err := h.walletService.CreditWallet(walletID, amount); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"wallet_id": walletID,
		"message":   "Wallet credited successfully",
	})
}

func (h *WalletHandler) Debit(c *gin.Context) {
	walletID, _ := strconv.Atoi(c.Param("wallet_id"))

	cacheKey := fmt.Sprintf("wallet:%d:balance", walletID)
	util.RedisClient.Del(c, cacheKey)

	amount, _ := strconv.ParseFloat(c.PostForm("amount"), 64)
	if err := h.walletService.DebitWallet(walletID, amount); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"wallet_id": walletID,
		"message":   "Wallet debited successfully",
	})
}
