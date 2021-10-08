package messages

import (
	"fmt"

	"github.com/teramono/utilities/pkg/request"
)

type ErrorMessage struct {
	Message string `json:"message"`
	Code    uint   `json:"code"`
}

var (
	InvalidWorkspaceIDHeader = ErrorMessage{
		Message: fmt.Sprintf("invalid or missing %s header value", request.WorkspaceIDHeader),
		Code:    110,
	}
	InvalidWorkspaceNameHeader = ErrorMessage{
		Message: fmt.Sprintf("invalid or missing %s header value", request.WorkspaceNameHeader),
		Code:    120,
	}
	InvalidAuthorizationHeader = ErrorMessage{
		Message: fmt.Sprintf("invalid or missing %s header value", request.AuthorizationHeader),
		Code:    130,
	}
	ValidationError = ErrorMessage{
		Message: "invalid user input",
		Code:    140,
	}
	InvalidBodyJson = ErrorMessage{
		Message: "invalid body json",
		Code:    150,
	}
	InvalidMessageJson = ErrorMessage{
		Message: "invalid message json",
		Code:    160,
	}
	UnableToGetBody = ErrorMessage{
		Message: "unable to get body bytes",
		Code:    160,
	}
)

func UnableToPublish(subject string) ErrorMessage {
	return ErrorMessage{
		Message: fmt.Sprintf("unable to publish `%s`", subject),
		Code:    200,
	}
}

func UnableToFindWorkspace(name string) ErrorMessage {
	return ErrorMessage{
		Message: fmt.Sprintf("unable to find workspace `%s`", name),
		Code:    210,
	}
}

func UnableToCreateWorkspace(name string) ErrorMessage {
	return ErrorMessage{
		Message: fmt.Sprintf("unable to create workspace `%s`", name),
		Code:    220,
	}
}

func (msg ErrorMessage) String() string {
	return fmt.Sprintf("ERR%s %s", msg.Code, msg.Message)
}
