package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stjnvc/wallet-api/internal/api/v1/route"
	"github.com/stjnvc/wallet-api/internal/db"
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

	router := gin.Default()

	route.Setup(router, db.DB)

	router.Run(":8080")
}
