package main

import (
	"github.com/zhongqiuwood/sample/grpc"
	"time"
)

func main() {
	peer, _  := p2p.NewPeer("peer0", "127.0.0.1:10000")

	var addresses []string
	addresses = append(addresses, "127.0.0.1:10001")

	peer.ChatWithSomePeers(addresses);

	for ;; {
		time.Sleep(100 * time.Millisecond)
	}
}
