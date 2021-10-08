package broker

import (
	"github.com/nats-io/nats.go"
)

type BrokerClient struct { // TODO: Should come with CommonSetup
	*nats.Conn
}

func NewBrokerClient(url string) (BrokerClient, error) {
	natsConn, err := nats.Connect(url)
	if err != nil {
		return BrokerClient{}, err
	}

	return BrokerClient{natsConn}, nil
}
