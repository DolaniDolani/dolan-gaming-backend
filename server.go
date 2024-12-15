package main

import (
	"log"
	"os"

	"github.com/DolaniDolani/dolan-gaming/db"
	"github.com/DolaniDolani/dolan-gaming/routes"

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

	routes.RegisterGameRoutes(server)
	routes.RegisterTestRoutes(server)

	server.Run(":" + serverPort)

}
