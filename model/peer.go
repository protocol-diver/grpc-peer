package model

import (
	"context"
	"fmt"
	"net"

	"github.com/protocol-diver/grpc-peer/logger"
	"github.com/protocol-diver/grpc-peer/msg"
	"google.golang.org/grpc"
)

type Peer struct {
	Id      int32
	Port    int
	Clients map[int]msg.MessageClient
	Server  msg.MessageServer
	Dial    Dial
}

func (p *Peer) String() string {
	return fmt.Sprintf("[ID: %03d, PORT: %05d]", p.Id, p.Port)
}

func (p *Peer) Send(senderPort, receiverPort int, message string, a ...any) (*msg.MessageSendResponse, error) {
	if _, ok := p.Clients[receiverPort]; !ok {
		client, err := p.Dial.NewClient(p.Port)
		if err != nil {
			return nil, err
		}
		p.Clients[receiverPort] = client
	}
	return p.Clients[receiverPort].MessageSend(context.Background(), &msg.MessageSendRequest{
		Sender:   int32(senderPort),
		Receiver: int32(receiverPort),
		Message:  fmt.Sprintf(message, a...),
	})
}

func (p *Peer) Listen() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", p.Port))
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	msg.RegisterMessageServer(grpcServer, p.Server)

	logger.Log(p.Id, "listening at %d", p.Port)
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}
