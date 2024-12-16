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
	utils.RespondWithErrorIfNotNil(ctx, 404, "Game not found", err)
	ctx.JSON(200, game)
}

func UpdateGame(ctx *gin.Context) {
	id := ctx.Param("id")
	var game models.Game

	err := ctx.ShouldBindBodyWithJSON(&game)
	utils.RespondWithErrorIfNotNil(ctx, http.StatusBadRequest, "Invalid data", err)

	err = db.DB.Where("id = ?", id).Updates(&game).Error
	utils.RespondWithErrorIfNotNil(ctx, http.StatusInternalServerError, "Error while updating game", err)

	ctx.JSON(http.StatusOK, gin.H{"message": "Game updated succesfully", "game": game})
}
func DeleteGame(ctx *gin.Context) {
	id := ctx.Param("id")
	var game models.Game

	err := ctx.ShouldBindBodyWithJSON(&game)
	utils.RespondWithErrorIfNotNil(ctx, http.StatusBadRequest, "Invalid Data", err)

	err = db.DB.Where("id = ?", id).Delete(&game).Error
	utils.RespondWithErrorIfNotNil(ctx, http.StatusInternalServerError, "Error while deleting game", err)

	ctx.JSON(http.StatusOK, gin.H{"message": "Game deleted succesfully"})
}
