package messages

import (
	"fmt"

	"github.com/teramono/utilities/pkg/request"
)

type ErrorMessage string

var (
	InvalidWorkspaceIDHeader   = ErrorMessage(fmt.Sprintf("Invalid or missing %s header value", request.WorkspaceIDHeader))
	InvalidWorkspaceNameHeader = ErrorMessage(fmt.Sprintf("Invalid or missing %s header value", request.WorkspaceNameHeader))
)
