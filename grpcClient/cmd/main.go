//go:generate protoc -I ../../ --go_out ../  --go_opt paths=source_relative  --go-grpc_out ../  --go-grpc_opt paths=source_relative  --grpc-gateway_out ../  --grpc-gateway_opt logtostderr=true  --grpc-gateway_opt paths=source_relative  ../../proto/echo.proto
package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpcClient/proto"
	"log"
	"os"
	"time"
)

const address = "localhost:50001"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEchoServiceClient(conn)

	string := "hello"
	if len(os.Args) > 1 {
		string = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Echo(ctx, &pb.StringMessage{Value: string})
	if err != nil {
		log.Fatalf("echo failed : %v", err)
	}
	log.Printf("get echo: %v", r.GetValue())

}
