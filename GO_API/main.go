package main

import (
	"example.com/m/db"
	"example.com/m/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8082") //localhost:8082

}
