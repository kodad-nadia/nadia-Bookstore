package main

import (
	"nadia/bookstore/database"
	"nadia/bookstore/routes"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatal("Database connection failed:", err)
	}

	r := gin.Default()
	routes.RegisterBookRoutes(r)
	r.Run(":8081")
}
