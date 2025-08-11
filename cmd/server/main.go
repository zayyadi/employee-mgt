package main

import (
	"employee-management/internal/database"
	"employee-management/internal/logging"
	"employee-management/internal/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	// Initialize logger
	logger := logging.InitLogger()

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize database connection
	db, err := database.Initialize()
	if err != nil {
		logger.WithError(err).Fatal("Failed to connect to database")
	}
	defer db.Close()

	// Initialize and start the server
	srv := server.NewServer(db, logger)
	if err := srv.Run(); err != nil {
		logger.WithError(err).Fatal("Failed to start server")
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
