package peer

import (
	"fmt"

	"github.com/protocol-diver/grpc-peer/msg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient(port int) (msg.MessageClient, error) {
	conn, err := grpc.Dial(fmt.Sprintf(":%d", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return msg.NewMessageClient(conn), nil
}
