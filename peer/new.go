package peer

import (
	"math/rand"

	"github.com/protocol-diver/grpc-peer/logger"
	"github.com/protocol-diver/grpc-peer/match"
	"github.com/protocol-diver/grpc-peer/model"
)

func NewPeer(port int) *model.Peer {
	peer := &model.Peer{
		Port:   port,
		Server: &Server{},
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
	client, err := NewClient(port)
	if err != nil {
		return nil
	}
	peer.Client = client
	logger.Log(peer.Id, "Peer created: %s", peer.String())

	match.Set(peer)
	return peer
}
