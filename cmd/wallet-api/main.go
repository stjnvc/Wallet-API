package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stjnvc/wallet-api/internal/api/v1/route"
	"github.com/stjnvc/wallet-api/internal/db"
	"github.com/stjnvc/wallet-api/internal/migration"
)

func main() {

	// Connect to the database
	if err := db.Connect(); err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// Migrate database tables
	if err := migration.Migrate(db.DB); err != nil {
		panic("Failed to migrate database tables")
	}

	router := gin.Default()
	route.Setup(router)
	router.Run(":8080")
}
