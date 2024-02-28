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
	return &WalletHandler{
		walletService: walletService,
	}
}

func (h *WalletHandler) GetBalance(c *gin.Context) {
	walletID, _ := strconv.Atoi(c.Param("wallet_id"))

	logrus.Infof("GetBalance request for wallet ID: %d", walletID)

	cacheKey := fmt.Sprintf("wallet:%d:balance", walletID)
	balanceStr, err := util.RedisClient.Get(c, cacheKey).Result()

	if err == nil {
		balance, _ := strconv.ParseFloat(balanceStr, 64)
		logrus.Errorf("Error fetching balance for wallet ID %d: %s", walletID, err)
		c.JSON(200, gin.H{
			"wallet_id": walletID,
			"balance":   balance,
		})

		return
	}

	balance, err := h.walletService.GetWalletBalance(walletID)
	logrus.Infof("Balance fetched successfully for wallet ID %d", walletID)

	if err != nil {
		logrus.Infof("Error fetching balance from database for wallet ID %d", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	util.RedisClient.Set(c, cacheKey, fmt.Sprintf("%.2f", balance), time.Hour)
	logrus.Infof("Balance fetched successfully for wallet ID %d", walletID)
	c.JSON(200, gin.H{
		"wallet_id": walletID,
		"balance":   balance,
	})
}

func (h *WalletHandler) Credit(c *gin.Context) {
	walletID, _ := strconv.Atoi(c.Param("wallet_id"))
	logrus.Infof("Credit request for wallet ID: %d", walletID)

	// Clear cache when updating balance
	cacheKey := fmt.Sprintf("wallet:%d:balance", walletID)
	util.RedisClient.Del(c, cacheKey)

	amount, _ := strconv.ParseFloat(c.PostForm("amount"), 64)
	if err := h.walletService.CreditWallet(walletID, amount); err != nil {
		logrus.Errorf("Error crediting wallet ID %d: %s", walletID, err.Error())
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
	logrus.Infof("Debit request for wallet ID: %d", walletID)

	cacheKey := fmt.Sprintf("wallet:%d:balance", walletID)
	util.RedisClient.Del(c, cacheKey)

	amount, _ := strconv.ParseFloat(c.PostForm("amount"), 64)
	if err := h.walletService.DebitWallet(walletID, amount); err != nil {
		logrus.Errorf("Error debiting wallet ID %d: %s", walletID, err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"wallet_id": walletID,
		"message":   "Wallet debited successfully",
	})
}
