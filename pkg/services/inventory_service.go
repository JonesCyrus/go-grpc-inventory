package services

import (
	"go-grpc-inventory/pkg/pb"

	"gorm.io/gorm"
)

type InventoryService struct {
	pb.UnimplementedInventoryProtoServiceServer
	DB *gorm.DB
}

func NewInventoryService(db *gorm.DB) *InventoryService {
	return &InventoryService{DB: db}
}
