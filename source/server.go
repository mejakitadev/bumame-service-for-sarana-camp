// @title			Sarana Boilerplate API
// @version		2.0
// @description	This is a sample server with v1 and v2 endpoints.
// @host			localhost:8200
// @BasePath		/
package main

import (
	"flag"
	"fmt"
	"os"
	"sarana-dafa-ai-service/config"
	"sarana-dafa-ai-service/controller"
	_ "sarana-dafa-ai-service/docs"
	"sarana-dafa-ai-service/service"
	"sarana-dafa-ai-service/storage"
	"sarana-dafa-ai-service/storage/env"

	"github.com/gofiber/swagger"
	"gorm.io/gorm"
)

func main() {
	// Load data from os/file/string to var
	storage.InitStorage()
	config.InitCasbin()

	db := config.NewDatabase()
	validate := config.NewValidator() // Changed from 'val' to 'validate' to match usage

	// Additional CLI Command
	if cliMigrate(db) {
		return
	}

	// Fiber Definition
	app := config.NewFiber()
	config.SetRecover(app)
	config.SetCORS(app)
	config.SetAccessLogger(app)

	bumameAuthService := service.NewBumameAuthService(db)
	bumameB2BProductService := service.NewBumameB2BProductService(db)
	bumameB2BPasienService := service.NewBumameB2BPasienService(db)
	bumameB2BDokterService := service.NewBumameB2BDokterService(db)

	// Controller
	// Initialize controllers
	bumameAuthController := controller.NewBumameAuthController(
		bumameAuthService,
		validate, // Fixed variable name
	)

	bumameB2BProductController := controller.NewBumameB2BProductController(
		bumameB2BProductService,
		validate, // Fixed variable name
	)

	bumameB2BPasienController := controller.NewBumameB2BPasienController(
		bumameB2BPasienService,
		validate, // Fixed variable name
	)

	bumameB2BDokterController := controller.NewBumameB2BDokterController(
		bumameB2BDokterService,
		validate, // Fixed variable name
	)

	// Register routes
	config.BumameAuthRouter(app, bumameAuthController)
	config.BumameB2BProductRouter(app, bumameB2BProductController)
	config.BumameB2BPasienRouter(app, bumameB2BPasienController)
	config.BumameB2BDokterRouter(app, bumameB2BDokterController)

	// Only enable Swagger in non-production environments
	if os.Getenv(env.APP_ENV) == "dev" {
		app.Get("/swagger/*", swagger.HandlerDefault)
	}

	app.Listen(":" + os.Getenv(env.APP_PORT))
}

// CLI Migrate
// Usage: go run server.go -migrate
func cliMigrate(db *gorm.DB) bool {
	migrate := flag.Bool("migrate", false, "Run database migrations")
	flag.Parse()

	if *migrate {
		fmt.Println("Running database migrations...")
		config.MigrateTable(db)
		fmt.Println("Database migrations completed!")
		return true
	}
	return false
}
