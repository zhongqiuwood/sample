package main

import (
	"flag"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "github.com/zhongqiuwood/sample/grpc/protos"
	"github.com/zhongqiuwood/sample/grpc"
)


func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "127.0.0.1:10001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}


	log.Printf("listen on: 127.0.0.1:10001")

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	peer, _  := p2p.NewPeer("peer1", "127.0.0.1:10001")

	pb.RegisterPeerServer(grpcServer, peer)
	grpcServer.Serve(lis)
}
