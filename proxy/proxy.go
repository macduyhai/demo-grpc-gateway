package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	gw "github.com/macduyhai/demo-grpc-gateway/gatewayproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:8989", "gRPC server endpoint")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gw.RegisterGrpcGatewayHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	log.Println("Proxy is running ...")
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
