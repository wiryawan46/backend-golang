package main

import (
	"backend-golang/config"
	"backend-golang/database"
	"backend-golang/routes"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize database
	database.InitDB()

	// setup router

	r := routes.SetupRouter()

	// Start server
	r.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
