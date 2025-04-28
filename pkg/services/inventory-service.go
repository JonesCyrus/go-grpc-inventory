package services

import "go-grpc-inventory/pkg/pb"

type InventoryService struct {
	pb.UnimplementedInventoryProtoServiceServer
}

func NewInventoryService() *InventoryService {
	return &InventoryService{}
}
