package broker

import (
	"encoding/json"
	"fmt"

	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
)


// TODO: Having different types of MsgData: MsgDataJson, MsgDataBytes
type MsgData struct {
	URL        URL                 `json:"url"`
	Headers    map[string][]string `json:"headers"`
	Data       string              `json:"data"`
	StatusCode int                 `json:"statusCode"`
}

type URL struct {
	Path     string `json:"path"`
	URI      string `json:"uri"`
	Location string `json:"location"`
}

type Header map[string][]string

func ToString(msg *nats.Msg) string {
	headers := ""

	for k, v := range msg.Header {
		headers += fmt.Sprintf("%s: \"%s\"", k, v)
	}

	return fmt.Sprintf(
		"msg { data: \"%s\", subject: \"%s\", headers: {%s}, reply: \"%s\" }",
		string(msg.Data),
		msg.Subject,
		headers,
		msg.Reply,
	)
}

func JsonFromMsgData(url URL, headers Header, data string, statusCode int) ([]byte, error) {
	return json.Marshal(MsgData{
		URL:        url,
		Headers:    headers,
		Data:       data,
		StatusCode: statusCode,
	})
}

func NewURLFromCtx(ctx *gin.Context) URL {
	url := location.Get(ctx)
	return URL{
		Path:     ctx.Request.URL.Path,
		URI:      ctx.Request.RequestURI,
		Location: url.Host,
	}
}

func NewMsgData(msg *nats.Msg) (MsgData, error) {
	msgData := MsgData{}
	err := json.Unmarshal(msg.Data, &msgData)
	return msgData, err
}
