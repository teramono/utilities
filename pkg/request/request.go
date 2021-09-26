package request

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type Request struct {
	URL     URL
	Body    []byte
	Method  string
	Headers map[string][]string
	Version string
}

type URL struct {
	Host   string
	Path   string
	URI    string
	Scheme string
}

// NewRequestFromContext ...
func NewRequestFromContext(ctx *gin.Context) (Request, error) {
	// URL
	url := URL{
		Host:   ctx.Request.Host,
		Path:   ctx.Request.URL.Path,
		URI:    ctx.Request.RequestURI,
		Scheme: ctx.Request.URL.Scheme,
	}

	// Method
	method := ctx.Request.Method

	// Body
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return Request{}, err
	}

	// Headers
	var headers = ctx.Request.Header

	return Request{
		URL:     url,
		Body:    body,
		Method:  method,
		Headers: headers,
	}, nil
}
