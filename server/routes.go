package server

import (
	"main/server/gateway"
	"main/server/handler"
)

func ConfigureRoutes(server *Server) {
	// use cors middleware
	server.engine.Use(gateway.CORSMiddleware())

	server.engine.POST("/players", handler.CreatePlayer)
	server.engine.GET("/players", handler.GetPlayer)
	server.engine.PUT("/players/:id", handler.UpdatePlayer)
	server.engine.DELETE("/players/:id", handler.DeletePlayer)
	server.engine.GET("/players/rank/:val", handler.GetPlayerByRank)
	server.engine.GET("/players/random", handler.GetRandomPlayer)

}
