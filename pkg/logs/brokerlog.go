package logs

import (
	"fmt"
	"log"

	"github.com/teramono/utilities/pkg/broker"
)

type BrokerClient interface {
	LogsVersion() uint
	Publish(subject string, msg []byte) error
}

func Panicf(client BrokerClient, format string, v ...interface{}) {
	subject := broker.GetSubjectWithId(client.LogsVersion(), "create", "logs", "panic")
	message := fmt.Sprintf(format, v...)

	if err := client.Publish(subject, []byte(message)); err != nil {
		log.Printf("unable to publish error log: %v", err)
	}

	log.Panicf(format, v...)
}

func Printf(client BrokerClient, format string, v ...interface{}) {
	subject := broker.GetSubjectWithId(client.LogsVersion(), "create", "logs", "panic")
	message := fmt.Sprintf(format, v...)

	if err := client.Publish(subject, []byte(message)); err != nil {
		log.Printf("unable to publish error log: %v", err)
	}

	log.Printf(format, v...)
}
