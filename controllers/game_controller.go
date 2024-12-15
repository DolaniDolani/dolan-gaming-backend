package controllers

import (
	"net/http"

	"github.com/DolaniDolani/dolan-gaming/db"
	"github.com/DolaniDolani/dolan-gaming/models"
	"github.com/DolaniDolani/dolan-gaming/utils"
	"github.com/gin-gonic/gin"
)

func AddGame(ctx *gin.Context) {
	var game models.Game
	if err := ctx.ShouldBindJSON(&game); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data", "details": err.Error()})
		return
	}

	err := db.DB.Create(&game)
	utils.RespondWithErrorIfNotNil(ctx, 500, "Error while inserting game", err.Error)

	ctx.JSON(http.StatusCreated, gin.H{"usage": "Game inserted succesfully", "game": game})
}

func GetAllGames(ctx *gin.Context) {
	var games []models.Game
	err := db.DB.Find(&games).Error
	utils.RespondWithErrorIfNotNil(ctx, 500, "Error while fetching games", err)

	ctx.JSON(200, games)
}

func GetGameByID(ctx *gin.Context) {
	id := ctx.Param("id")
	var game models.Game

	err := db.DB.First(&game, "id = ?", id).Error
	utils.RespondWithErrorIfNotNil(ctx, 500, "Error while fetching game", err)
	ctx.JSON(200, game)
}
