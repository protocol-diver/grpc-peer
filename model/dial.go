package model

import "github.com/protocol-diver/grpc-peer/msg"

type Dial interface {
	NewClient(port int) (msg.MessageClient, error)
}
