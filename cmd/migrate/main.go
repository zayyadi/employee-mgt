package main

import (
	"employee-management/internal/migrate"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Get database connection details from environment variables
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "password")
	dbname := getEnv("DB_NAME", "employee_management")

	// Create database connection string
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbname)

	// Check command line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: migrate [up|down]")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "up":
		if err := migrate.RunMigrations(connStr); err != nil {
			log.Fatal("Failed to run migrations:", err)
		}
		fmt.Println("Migrations applied successfully")
	case "down":
		if err := migrate.RunDownMigrations(connStr); err != nil {
			log.Fatal("Failed to rollback migrations:", err)
		}
		fmt.Println("Migrations rolled back successfully")
	default:
		fmt.Println("Unknown command. Usage: migrate [up|down]")
		os.Exit(1)
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
