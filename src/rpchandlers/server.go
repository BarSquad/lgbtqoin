package main

import (
	"net"

	pb "github.com/BarSquad/lgbtqoin/src/rpchandlers/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	listener, err := net.Listen("tcp", ":5300")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterReverseServer(grpcServer, &server{})
	grpcServer.Serve(listener)
}

type server struct {
	pb.UnimplementedReverseServer
}

func (s *server) Do(c context.Context, request *pb.Request) (response *pb.Response, err error) {
	n := 0
	// Сreate an array of runes to safely reverse a string.
	runes := make([]rune, len(request.Message))

	for _, r := range request.Message {
		runes[n] = r
		n++
	}

	// Reverse using runes.
	runes = runes[0:n]

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	output := string(runes)
	response = &pb.Response{
		Message: output,
	}

	return response, nil
}
