package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stjnvc/wallet-api/internal/api/v1/route"
)

func main() {
	router := gin.Default()
	route.Setup(router)
	router.Run(":8080")
}
