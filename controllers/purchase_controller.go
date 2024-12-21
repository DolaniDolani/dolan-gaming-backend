package controllers

import (
	"fmt"
	"net/http"

	"github.com/DolaniDolani/dolan-gaming/db"
	"github.com/DolaniDolani/dolan-gaming/models"
	"github.com/DolaniDolani/dolan-gaming/utils"
	"github.com/gin-gonic/gin"
)

func CreatePurchase(ctx *gin.Context) {
	var purchase models.Purchase

	err := ctx.ShouldBindBodyWithJSON(&purchase)
	if utils.RespondWithErrorIfNotNil(ctx, http.StatusBadRequest, "Invalid data", err) {
		return
	}
	err = purchaseCheck(purchase)
	if utils.RespondWithErrorIfNotNil(ctx, http.StatusBadRequest, "Invalid data", err) {
		return
	}

	err = db.DB.Create(&purchase).Error
	if utils.RespondWithErrorIfNotNil(ctx, http.StatusInternalServerError, "Error during purchase creation", err) {
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Purchase created successfully", "purchase": purchase})
}

func purchaseCheck(purchase models.Purchase) error {

	if purchase.Date.IsZero() || len(purchase.Games) == 0 {
		return fmt.Errorf("date and games can't be null")
	}
	return nil
}

func GetAllPurchases(ctx *gin.Context) {
	var purchases []models.Purchase
	err := db.DB.Find(&purchases).Error
	if utils.RespondWithErrorIfNotNil(ctx, 500, "Error while fetching purchases", err) {
		return
	}

	ctx.JSON(200, purchases)
}

func GetPurchaseById(ctx *gin.Context) {
	id := ctx.Param("id")

	var purchase models.Purchase

	err := db.DB.First(&purchase, "id = ?", id).Error
	if utils.RespondWithErrorIfNotNil(ctx, 404, "Purchase not found", err) {
		return
	}

	ctx.JSON(200, purchase)
}

func UpdatePurchase(ctx *gin.Context) {
	id := ctx.Param("id")

	var purchase models.Purchase

	err := ctx.ShouldBindBodyWithJSON(&purchase)
	if utils.RespondWithErrorIfNotNil(ctx, http.StatusBadRequest, "Invalid data", err) {
		return
	}

	err = db.DB.Where("id = ?", id).Updates(&purchase).Error
	if utils.RespondWithErrorIfNotNil(ctx, http.StatusInternalServerError, "Error while updating game", err) {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Purchase updated succesfully", "purchase": purchase})
}
