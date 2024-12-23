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
	err := ctx.ShouldBindJSON(&game)

	if utils.RespondWithErrorIfNotNil(ctx, http.StatusBadRequest, "Invalid data", err) {
		return
	}

	err = db.DB.Create(&game).Error
	if utils.RespondWithErrorIfNotNil(ctx, 500, "Error while inserting game", err) {
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"usage": "Game inserted succesfully", "game": game})
}

func GetAllGames(ctx *gin.Context) {
	var games []models.Game
	err := db.DB.Find(&games).Error
	if utils.RespondWithErrorIfNotNil(ctx, 500, "Error while fetching games", err) {
		return
	}

	ctx.JSON(200, games)
}

func GetGameByID(ctx *gin.Context) {
	id := ctx.Param("id")
	var game models.Game

	err := db.DB.First(&game, "id = ?", id).Error
	if utils.RespondWithErrorIfNotNil(ctx, 404, "Game not found", err) {
		return
	}

	ctx.JSON(200, game)
}

func UpdateGame(ctx *gin.Context) {
	id := ctx.Param("id")
	var game models.Game

	err := ctx.ShouldBindBodyWithJSON(&game)
	if utils.RespondWithErrorIfNotNil(ctx, http.StatusBadRequest, "Invalid data", err) {
		return
	}

	err = db.DB.Where("id = ?", id).Updates(&game).Error
	if utils.RespondWithErrorIfNotNil(ctx, http.StatusInternalServerError, "Error while updating game", err) {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Game updated succesfully", "game": game})
}
func DeleteGame(ctx *gin.Context) {
	id := ctx.Param("id")
	var game models.Game

	err := db.DB.Where("id = ?", id).Delete(&game).Error
	if utils.RespondWithErrorIfNotNil(ctx, http.StatusInternalServerError, "Error while deleting game", err) {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Game deleted successfully"})
}

func GetGameByPurchaseId(ctx *gin.Context) {
	id := ctx.Param("id")
	var games []models.Game
	err := db.DB.Where("purchase_id = ?", id).Find(&games).Error
	if utils.RespondWithErrorIfNotNil(ctx, http.StatusInternalServerError, "Error while fetching games", err) {
		return
	}
	if len(games) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No game where found for this purchase"})
		return
	}
	ctx.JSON(http.StatusOK, games)
}
