package main

import (
	"log"
	"os"

	"github.com/DolaniDolani/dolan-gaming/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error during loading of .env file: %v", err)
	}

	db.ConnectDB()

	serverPort := os.Getenv("SERVER_PORT")

	server := gin.Default()

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "It just works",
		})
	})

	server.GET("/test-db", func(ctx *gin.Context) {
		version := " "
		err := db.Pool.QueryRow(ctx, "SELECT version()").Scan(&version)
		if err != nil {
			log.Println(err)
			ctx.JSON(500, gin.H{"error": "Database access error"})
			return
		}
		ctx.JSON(200, gin.H{"database_version": version})
	})

	server.Run(":" + serverPort)

}
