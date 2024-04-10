package main

import (
	"github.com/PrasannaAdiga/my-first-go/db"
	"github.com/PrasannaAdiga/my-first-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
