//go:generate protoc -I ../../ --go_out ../  --go_opt paths=source_relative  --go-grpc_out ../  --go-grpc_opt paths=source_relative  --grpc-gateway_out ../  --grpc-gateway_opt logtostderr=true  --grpc-gateway_opt paths=source_relative  ../../proto/echo.proto
package main

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"grpcServer/internal/consul"
	pb "grpcServer/proto"
	"log"
	"net"
)

var port = 50001
var ip = "127.0.0.1"

type server struct {
	pb.UnimplementedEchoServiceServer
}

func (s *server) Echo(ctx context.Context, in *pb.StringMessage) (*pb.StringMessage, error) {
	log.Printf("Received: %v", in.GetValue())
	//return in, nil
	return &pb.StringMessage{Value: "." + in.GetValue()}, nil
}
func RegisterToConsul() {
	consul.RegitserService("127.0.0.1:8500", &consul.ConsulService{
		IP:   ip,
		Port: port,
		Tag:  []string{"echo", "hello world"},
		Name: "echoService",
	})
}

//health
type HealthImpl struct{}

// Check 实现健康检查接口，这里直接返回健康状态，这里也可以有更复杂的健康检查策略，比如根据服务器负载来返回
func (h *HealthImpl) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	fmt.Print("health checking\n")
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (h *HealthImpl) Watch(req *grpc_health_v1.HealthCheckRequest, w grpc_health_v1.Health_WatchServer) error {
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":"+cast.ToString(port))
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterEchoServiceServer(s, &server{})
	grpc_health_v1.RegisterHealthServer(s, &HealthImpl{})
	RegisterToConsul()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}

}
