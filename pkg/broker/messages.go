package broker

import (
	"encoding/json"
	"fmt"

	"github.com/nats-io/nats.go"
)

type M map[string]interface{}

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

func Json(headers M, data []byte) ([]byte, error) {
	return json.Marshal(M{
		"headers": headers,
		"data":    data,
	})
}
