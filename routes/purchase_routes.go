package routes

import (
	"github.com/DolaniDolani/dolan-gaming/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterPurchaseRoutes(router *gin.Engine) {
	router.POST("purchases/add", controllers.CreatePurchase)
	router.GET("purchases", controllers.GetAllPurchases)
}
