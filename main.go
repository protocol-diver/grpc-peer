package main

import (
	"errors"
	"math/rand"
	"os"
	"time"

	"github.com/protocol-diver/grpc-peer/logger"
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

func main() {
	pre()

	ports := []int{3100, 3101, 3102, 3103, 3104, 3105, 3106, 3107}

	for _, port := range ports {
		go func(port int) {
			p := peer.NewPeer(port)
			time.Sleep(time.Second * 3) // wait 3s to start every peer listening

			for i := int64(0); ; i++ {
				// select target port
				targetPort := port
				for ; targetPort == port; targetPort = ports[rand.Intn(len(ports))] {
				}

				res, err := p.Send(port, targetPort, "p%d-%d-%03d", p.Port, targetPort, i)
				if err != nil {
					logger.Log(p.Id, err.Error())
				} else {
					logger.Log(p.Id, res.String())
				}
				i++

				// sleep 1.5s ~ 3s
				ms := rand.Intn(1500) + 1500
				time.Sleep(time.Millisecond * time.Duration(ms))
			}
		}(port)
	}

	time.Sleep(time.Second * 30)
}
