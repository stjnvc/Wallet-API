package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stjnvc/wallet-api/internal/api/v1/route"
	"github.com/stjnvc/wallet-api/internal/db"
	"github.com/stjnvc/wallet-api/internal/util"
	"log"
)

func main() {
	// Connect to the database
	if err := db.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Defer the closure of the database connection
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
	}()

	// Migrate database tables
	if err := db.Migrate(db.DB); err != nil {
		log.Fatalf("Failed to migrate database tables: %v", err)
	}

	if err := db.Seed(); err != nil {
		log.Fatalf("Failed to seed database %v", err)
	} else {
		log.Println("Database seeded.")
	}

	// Initialize Redis
	err := util.InitRedisClient()
	if err != nil {
		logrus.Error("Failed to initialize redis cache:", err)
	}

	router := gin.New()

	route.Setup(router, db.DB)

	router.Run(":8080")
}
