package main

import (
	"employee-management/internal/database"
	"employee-management/internal/logging"
	"employee-management/internal/server"
)

// @title Employee Management System API
// @version 1.0
// @description This is an employee management system with payroll functionality.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://github.com/your-repo/license

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	// Initialize logger
	logger := logging.InitLogger()

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
