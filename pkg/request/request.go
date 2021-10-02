package request

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

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

func (req *Request) Send() (*http.Response, error) {
	// Http client with 10s timeout.
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// Construct a new request.
	newReq, err := http.NewRequest(req.Method, req.URL.Host+req.URL.URI, bytes.NewReader(req.Body))
	if err != nil {
		return nil, err
	}

	// Set headers.
	for key, value := range req.Headers {
		newReq.Header.Set(key, strings.Join(value, ","))
	}

	return client.Do(newReq)
}
