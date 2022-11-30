package peer

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/protocol-diver/grpc-peer/msg"
)

type Server struct {
	msg.MessageServer
}

func (m *Server) MessageSend(ctx context.Context, req *msg.MessageSendRequest) (*msg.MessageSendResponse, error) {
	res := &msg.MessageSendResponse{
		Sender:   req.Sender,
		Receiver: req.Receiver,
		Message:  req.Message,
	}

	f, err := os.OpenFile(fmt.Sprintf("./data/%d.csv", req.Sender), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	if err != nil {
		res.Error = err.Error()
		return res, err
	}
	defer f.Close()

	x := fmt.Sprintf(
		"timstamp: %d, sender: %d, receiver: %d, message: %s\n",
		time.Now().Unix(),
		req.Sender,
		req.Receiver,
		req.Message,
	)
	f.WriteString(x)
	b := []byte{}
	f.Read(b)
	res.Id = int32(bytes.Count(b, []byte{10})) // [10] means "\n"

	return res, nil
}
