package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/DolaniDolani/dolan-gaming/db"
	"github.com/DolaniDolani/dolan-gaming/models"
	"github.com/DolaniDolani/dolan-gaming/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	err := db.DB.Preload("Games").Find(&purchases).Error
	if utils.RespondWithErrorIfNotNil(ctx, 500, "Error while fetching purchases", err) {
		return
	}

	ctx.JSON(200, purchases)
}

func GetPurchaseById(ctx *gin.Context) {
	id := ctx.Param("id")

	var purchase models.Purchase

	err := db.DB.Preload("Games").First(&purchase, "id = ?", id).Error
	if utils.RespondWithErrorIfNotNil(ctx, 404, "Purchase not found", err) {
		return
	}

	ctx.JSON(200, purchase)
}
func UpdatePurchase(ctx *gin.Context) {
	id := ctx.Param("id")

	var purchase models.Purchase

	// Leggi il body della richiesta
	if err := ctx.ShouldBindJSON(&purchase); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dati non validi", "details": err.Error()})
		return
	}

	// Assegna l'ID passato nell'URL al modello
	purchase.ID = idAsInt(id)

	// Aggiorna il Purchase e i giochi in una transazione
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		// Aggiorna il Purchase
		if err := tx.Model(&purchase).Where("id = ?", purchase.ID).Updates(purchase).Error; err != nil {
			return err
		}

		// Aggiorna i giochi associati
		for _, game := range purchase.Games {
			game.PurchaseID = purchase.ID
			if game.ID == 0 {
				// Nuovo gioco
				if err := tx.Create(&game).Error; err != nil {
					return err
				}
			} else {
				// Gioco esistente
				if err := tx.Model(&game).Where("id = ?", game.ID).Updates(&game).Error; err != nil {
					return err
				}
			}
		}

		return nil
	})

	// Gestisci errori di transazione
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Errore durante l'aggiornamento", "details": err.Error()})
		return
	}

	// Risposta di successo
	ctx.JSON(http.StatusOK, gin.H{"message": "Purchase aggiornato con successo", "purchase": purchase})
}

func idAsInt(id string) int64 {
	intID, _ := strconv.ParseInt(id, 10, 64)
	return intID
}

func DeletePurchase(ctx *gin.Context) {
	id := ctx.Param("id")

	var purchase models.Purchase

	err := db.DB.Where("id = ?", id).Delete(&purchase).Error
	if utils.RespondWithErrorIfNotNil(ctx, http.StatusInternalServerError, "Error while deleting purchase", err) {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Purchase deleted successfully", "purchase": purchase})
}
