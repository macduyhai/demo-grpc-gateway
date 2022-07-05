package main

import (
	"context"
	"errors"
	"log"
	_ "log"
	"net"
	"time"

	gatewaypb "github.com/macduyhai/demo-grpc-gateway/gatewayproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type server struct {
	gatewaypb.UnimplementedGrpcGatewayServer
}

func (s *server) CheckPing(ctx context.Context, req *gatewaypb.PingRequest) (*gatewaypb.PingResponse, error) {
	log.Printf("Receive message %s\n", req.GetMessage())
	return &gatewaypb.PingResponse{
		Message: "PongPong",
	}, nil

}
func (s *server) CreateUser(ctx context.Context, req *gatewaypb.CreateUserRequest) (*gatewaypb.CreateUserResponse, error) {
	headers, err := metadata.FromIncomingContext(ctx)
	if err != true {
		log.Printf("Get header error\n")
		return nil, errors.New("Get header error")
	}
	log.Println("Get header ok", "header", headers)
	log.Println("Create user success")
	return &gatewaypb.CreateUserResponse{
		Code:     1,
		Messsage: "create success",
		Data: &gatewaypb.CreateUserResponse_Data{
			UserId: time.Now().Unix(),
		},
	}, nil

}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:8989")
	if err != nil {
		log.Fatalf("Init listen error:%v", err)
	}

	s := grpc.NewServer()
	gatewaypb.RegisterGrpcGatewayServer(s, &server{})
	log.Println("Grpc gateway is running ...")
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Error while run Serve:%v", err)
	}

}
