//go:generate protoc -I ../../ --go_out ../  --go_opt paths=source_relative  --go-grpc_out ../  --go-grpc_opt paths=source_relative  --grpc-gateway_out ../  --grpc-gateway_opt logtostderr=true  --grpc-gateway_opt paths=source_relative  ../../proto/echo.proto
package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpcServer/proto"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedEchoServiceServer
}

func (s *server) Echo(ctx context.Context, in *pb.StringMessage) (*pb.StringMessage, error) {
	log.Printf("Received: %v", in.GetValue())
	//return in, nil
	return &pb.StringMessage{Value: "." + in.GetValue()}, nil
}
func main() {
	lis, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterEchoServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}

}
