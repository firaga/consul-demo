//go:generate protoc -I ../../ --go_out ../  --go_opt paths=source_relative  --go-grpc_out ../  --go-grpc_opt paths=source_relative  --grpc-gateway_out ../  --grpc-gateway_opt logtostderr=true  --grpc-gateway_opt paths=source_relative  ../../proto/echo.proto
package main

import (
	"context"
	"google.golang.org/grpc"
	"grpcClient/internal/consul"
	pb "grpcClient/proto"
	"log"
	"os"
	"time"
)

const target = "consul://127.0.0.1:8500/echoService"

func init() {
	consul.Init()
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	conn, err := grpc.DialContext(ctx, target, grpc.WithBlock(), grpc.WithInsecure(), grpc.WithBalancerName("round_robin"))
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
