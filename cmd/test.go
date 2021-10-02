package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/teramono/utilities/pkg/request"
	"github.com/teramono/utilities/pkg/response"
)

func main() {
	fmt.Println(response.SourceAPIServer)
	fmt.Println(request.NewRequestFromContext(&gin.Context{}))
}
