package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"testing"

	pb "github.com/muhrizqiardi/trygogrpc/start_grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func BenchmarkGRPC(b *testing.B) {
	conn, err := grpc.Dial("localhost:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewFactorialServiceClient(conn)
	arg, _ := strconv.Atoi(os.Args[1])

	for i := 0; i < b.N; i++ {
		r, err := c.FindFactorial(context.Background(), &pb.FindFactorialRequest{Input: int32(arg)})
		if err != nil {
			log.Fatalf("could not call: %v", err)
		}
		log.Println("factorial of", arg, "is", r.Result)
	}
}
