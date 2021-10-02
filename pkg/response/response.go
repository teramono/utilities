package response

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teramono/utilities/pkg/messages"
)

func ReadBodyBytes(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}

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
