package main

import (
	"context"
	"log"
	"net"

	pb "github.com/muhrizqiardi/trygogrpc/start_grpc/proto"
	"google.golang.org/grpc"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

type Server struct {
	pb.UnimplementedFactorialServiceServer
}

func (s *Server) FindFactorial(ctx context.Context, in *pb.FindFactorialRequest) (*pb.FindFactorialReply, error) {
	log.Println("`FindFactorial` called with input:", in.Input)

	return &pb.FindFactorialReply{Result: int32(factorial(int(in.Input)))}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterFactorialServiceServer(s, &Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
