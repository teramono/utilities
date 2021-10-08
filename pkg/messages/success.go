package messages

import "fmt"

type SuccessMessage string

var (
	LoggingIn = SuccessMessage("Logging in...")
	ScriptStarted = SuccessMessage("Script started...")
)

func Created(name string) SuccessMessage {
	return SuccessMessage(fmt.Sprintf("%s created successfully", name))
}

func (msg SuccessMessage) String() string {
	return string(msg)
}
