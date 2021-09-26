package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teramono/utilities/pkg/messages"
)

func SetErrorResponseWithSourceType(ctx *gin.Context, src ResponseSource, errorMessages ...messages.ErrorMessage) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"source": src,
		"body": gin.H{
			"errors": errorMessages,
		},
	})
}

func SetErrorResponse(ctx *gin.Context, errorMessages ...messages.ErrorMessage) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"errors": errorMessages,
	})
}
