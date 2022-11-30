package main

import (
	"errors"
	"math/rand"
	"os"
	"time"

	"github.com/protocol-diver/grpc-peer/logger"
	"github.com/protocol-diver/grpc-peer/model"
	"github.com/protocol-diver/grpc-peer/peer"
)

func pre() {
	rand.Seed(time.Now().Unix())

	if _, err := os.Stat("./data"); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if err := os.Mkdir("./data", 0755); err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}
}

func run(from, to *model.Peer, ms int) {
	i := int64(0)
	for {
		time.Sleep(time.Millisecond * time.Duration(ms))
		res, err := from.Send(to.Id, "p%d-%d-%03d", from.Id, to.Id, i)
		if err != nil {
			logger.Log(from.Id, err.Error())
		} else {
			logger.Log(from.Id, res.String())
		}
		i++
	}
}

func main() {
	pre()

	p0 := peer.NewPeer(3100)
	p1 := peer.NewPeer(3101)
	p2 := peer.NewPeer(3102)
	p3 := peer.NewPeer(3103)

	go run(p0, p1, 1800)
	go run(p1, p2, 2000)
	go run(p2, p3, 1500)
	go run(p3, p0, 3000)

	time.Sleep(time.Second * 30)
}
