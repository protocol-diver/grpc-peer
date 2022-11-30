package match

import "github.com/protocol-diver/grpc-peer/model"

var matcher = make(map[int32]*model.Peer)

func Get(id int32) *model.Peer {
	if p, ok := matcher[id]; !ok {
		return nil
	} else {
		return p
	}
}

func Set(p *model.Peer) {
	matcher[p.Id] = p
}
