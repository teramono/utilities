package response

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/teramono/utilities/pkg/messages"
)

func SetUserErrorResponse(ctx *gin.Context, errorMessages ...messages.ErrorMessage) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"errors": errorMessages,
	})
}

func SetUserValidationErrorResponse(ctx *gin.Context, validErrs validator.ValidationErrors) {
	var errorMessages []interface{}
	for _, err := range validErrs {
		errorMessages = append(errorMessages, messages.ErrorMessage{
			Code:    messages.ValidationError.Code,
			Message: err.Error(),
		})
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"errors": errorMessages,
	})
}

func SetServerErrorResponse(ctx *gin.Context, errorMessages ...messages.ErrorMessage) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"errors": errorMessages,
	})
}

func CreateErrorResponseBody(errorMessages ...messages.ErrorMessage) map[string][]messages.ErrorMessage {
	return map[string][]messages.ErrorMessage{
		"errors": errorMessages,
	}
}

func CreateErrorResponseBodyBytes(errorMessages ...messages.ErrorMessage) ([]byte, error) {
	body := CreateErrorResponseBody(errorMessages...)
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return []byte{}, err
	}

	return bodyBytes, nil
}
