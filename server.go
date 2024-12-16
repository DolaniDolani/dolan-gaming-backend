package main

import (
	"log"
	"os"

	"github.com/DolaniDolani/dolan-gaming/db"
	"github.com/DolaniDolani/dolan-gaming/routes"
	"github.com/gin-contrib/cors"

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

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"}, // Domini autorizzati (frontend Angular)
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	routes.RegisterGameRoutes(server)
	routes.RegisterTestRoutes(server)

	server.Run(":" + serverPort)

}
