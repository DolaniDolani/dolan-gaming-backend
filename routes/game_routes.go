package routes

import (
	"github.com/DolaniDolani/dolan-gaming/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterGameRoutes(router *gin.Engine) {
	router.POST("games/add", controllers.AddGame)
	router.GET("games", controllers.GetAllGames)
	router.GET("games/:id", controllers.GetGameByID)
	router.PUT("games/:id", controllers.UpdateGame)
	router.DELETE("games/:id", controllers.DeleteGame)
}
