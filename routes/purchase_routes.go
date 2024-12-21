package routes

import (
	"github.com/DolaniDolani/dolan-gaming/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterPurchaseRoutes(router *gin.Engine) {
	router.POST("purchases/add", controllers.CreatePurchase)
	router.PUT("purchases/:id", controllers.UpdatePurchase)
	router.GET("purchases/:id", controllers.GetPurchaseById)
	router.GET("purchases", controllers.GetAllPurchases)
}
