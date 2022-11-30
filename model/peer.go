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
	Id     int32
	Port   int
	Client msg.MessageClient
	Server msg.MessageServer
}

func (p *Peer) String() string {
	return fmt.Sprintf("[ID: %03d, PORT: %05d]", p.Id, p.Port)
}

func (p *Peer) Send(receiverId int32, message string, a ...any) (*msg.MessageSendResponse, error) {
	return p.Client.MessageSend(context.Background(), &msg.MessageSendRequest{
		Sender:   p.Id,
		Receiver: receiverId,
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
