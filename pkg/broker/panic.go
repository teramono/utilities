package broker

import (
	"fmt"
	"os"

	"github.com/nats-io/nats.go"
)

// Recover recovers from a panic.
func Recover() {
	if r := recover(); r != nil {
		fmt.Fprintf(os.Stderr, "Panic recovered: %v", r)
	}
}

// Wraps a handler function so that panic is handled before it reaches the parent function.
func PanicWrap(msg *nats.Msg, hanlderFunc func(*nats.Msg)) {
	// Panic handler.
	defer Recover()

	hanlderFunc(msg)
}
