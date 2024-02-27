package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stjnvc/wallet-api/internal/api/v1/service"
	"strconv"
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
	balance, err := h.walletService.GetWalletBalance(walletID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"wallet_id": walletID,
		"balance":   balance,
	})
}

func (h *WalletHandler) Credit(c *gin.Context) {
	walletID, _ := strconv.Atoi(c.Param("wallet_id"))
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
