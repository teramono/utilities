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
		Code:    111,
	}
	InvalidWorkspaceIDAndNameHeader = ErrorMessage{
		Message: fmt.Sprintf("invalid or missing %s and %s header value", request.WorkspaceIDHeader, request.WorkspaceNameHeader),
		Code:    112,
	}
	InvalidAuthorizationHeader = ErrorMessage{
		Message: fmt.Sprintf("invalid or missing %s header value", request.AuthorizationHeader),
		Code:    113,
	}
	ValidationError = ErrorMessage{
		Message: "invalid user input",
		Code:    120,
	}
	InvalidMessageJson = ErrorMessage{
		Message: "invalid message json",
		Code:    130,
	}
	InvalidBodyJson = ErrorMessage{
		Message: "invalid body json",
		Code:    131,
	}
	InvalidReplyJson = ErrorMessage{
		Message: "invalid reply json",
		Code:    132,
	}
	InvalidReplyDataJson = ErrorMessage{
		Message: "invalid reply data json",
		Code:    133,
	}
	UnableToGetBody = ErrorMessage{
		Message: "unable to get body bytes",
		Code:    135,
	}
	UnableToGetReplyFromSubscriber = ErrorMessage{
		Message: "unable to get reply from subscriber",
		Code:    140,
	}
	ProblemProcessingRequest = ErrorMessage{
		Message: "encountered problem while processing request",
		Code:    150,
	}
	UnableToConstructServerErrorResponse = ErrorMessage{
		Message: "unable to construct server error response",
		Code:    151,
	}
	UnableToSendServerErrorResponse = ErrorMessage{
		Message: "unable to send server error response",
		Code:    152,
	}
	UnableToGetDataFromMessage = ErrorMessage{
		Message: "unable to get data from recieved message",
		Code:    300,
	}
	UnableToGetCanonicalWorkspacePath = ErrorMessage{
		Message: "unable to get canonical workspace path",
		Code:    301,
	}
	UnableToFetchSurlManifest = ErrorMessage{
		Message: "unable to fetch surl manifest",
		Code:    302,
	}
	UnableToRunAuthScriptSuccesfully = ErrorMessage{
		Message: "unable to run all auth script successfully",
		Code:    303,
	}
	UnableToRunMiddlewareScriptsSuccessfully = ErrorMessage{
		Message: "unable to run all middleware scripts successfully",
		Code:    304,
	}
	UnableToRunIndexScriptSuccessfully = ErrorMessage{
		Message: "unable to run all index script successfully",
		Code:    305,
	}
	UnableToAuthenticate = ErrorMessage{
		Message: "unable to fetch all middleware scripts",
		Code:    400,
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

func WrongReturnType(expectedType string, object string) ErrorMessage {
	return ErrorMessage{
		Message: fmt.Sprintf(
			"wrong type returned, expected type `%s` from script `%s`",
			expectedType,
			object,
		),
		Code: 230,
	}
}

func (msg ErrorMessage) String() string {
	return fmt.Sprintf("ERR%d %s", msg.Code, msg.Message)
}
