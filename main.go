package main

import (
	"log"
	"main/server"


)

// @title Gin Demo App
// @version 1.0
// @description This is a demo version of Gin app.
// @BasePath /
func main() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// connection := db.InitDB()
	// db.Transfer(connection)
	app := server.NewServer(nil)
	server.ConfigureRoutes(app)

	if err := app.Run("8000"); err != nil {
		log.Print(err)
	}
}
