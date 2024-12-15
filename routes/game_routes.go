package routes

import (
	"github.com/DolaniDolani/dolan-gaming/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterGameRoutes(router *gin.Engine) {
	router.POST("games/add", controllers.AddGame)
	router.GET("games/get/all", controllers.GetAllGames)
	router.GET("games/get/:id", controllers.GetGameByID)
}
