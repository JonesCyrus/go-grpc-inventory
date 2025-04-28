package main

import (
	"fmt"
	"go-grpc-inventory/pkg/config"
	"go-grpc-inventory/pkg/pb"
	"go-grpc-inventory/pkg/services"
	"log"
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func getPort(vi *viper.Viper) string {
	port := vi.GetString("port")
	if port == "" {
		return "8001"
	}

	return port
}

func startServer(svc *services.InventoryService, port string) (*grpc.Server, net.Listener, error) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to listen: %w", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterInventoryProtoServiceServer(grpcServer, svc)

	return grpcServer, lis, nil
}

func main() {
	vi := config.InitEnv()

	port := getPort(vi)

	svc := services.NewInventoryService()

	grpcServer, lis, err := startServer(svc, port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	log.Printf("Server is listening on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
