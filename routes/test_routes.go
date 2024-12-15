package routes

import (
	"github.com/DolaniDolani/dolan-gaming/db"
	"github.com/DolaniDolani/dolan-gaming/utils"
	"github.com/gin-gonic/gin"
)

func RegisterTestRoutes(router *gin.Engine) {
	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "It just works",
		})
	})

	router.GET("/test-db", func(ctx *gin.Context) {
		sqlDB, err := db.DB.DB()
		utils.RespondWithErrorIfNotNil(ctx, 500, "Error while testing database", err)

		err = sqlDB.Ping()
		utils.RespondWithErrorIfNotNil(ctx, 500, "Error while pinging database", err)

		ctx.JSON(200, gin.H{"message": "Database connected succesfully"})
	})
}
