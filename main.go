package main

import (
	"github.com/gin-gonic/gin"
	"simple_rest_api/db"
	"simple_rest_api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
