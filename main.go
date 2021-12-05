package main

import (
	"log"
	"os"

	"github.com/Phelickz/go-sql/src/database"
	"github.com/Phelickz/go-sql/src/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	//getting environment

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	//set router
	router := gin.New()
	router.Use(gin.Logger())

	//connect to database
	database.ConnectDb()

	//initialize routes
	routes.AccessCredentials(router)
	routes.PensionGistRoute(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Success"})
	})

	router.Run(":" + port)

}
