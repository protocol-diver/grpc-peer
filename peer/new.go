package peer

import (
	"math/rand"

	"github.com/protocol-diver/grpc-peer/logger"
	"github.com/protocol-diver/grpc-peer/model"
	"github.com/protocol-diver/grpc-peer/msg"
)

type Dial struct{}

func (d *Dial) NewClient(port int) (msg.MessageClient, error) {
	return NewClient(port)
}

func NewPeer(port int) *model.Peer {
	peer := &model.Peer{
		Port:    port,
		Dial:    &Dial{},
		Server:  &Server{},
		Clients: make(map[int]msg.MessageClient),
	}
	for {
		colorIdx := rand.Intn(len(colors))
		color := colors[colorIdx]
		if using[color] == nil {
			using[color] = &port
			peer.Id = int32(color)
			break
		}
	}

	go peer.Listen()
	logger.Log(peer.Id, "Peer created: %s", peer.String())

	return peer
}
