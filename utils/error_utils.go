package utils

import "github.com/gin-gonic/gin"

func RespondWithError(ctx *gin.Context, code int, message string, details error) {
	ctx.JSON(code, gin.H{
		"error":   message,
		"details": details.Error(),
	})
}

func RespondWithErrorIfNotNil(ctx *gin.Context, code int, message string, details error) {
	if details != nil {
		RespondWithError(ctx, code, message, details)
		return
	}
}
