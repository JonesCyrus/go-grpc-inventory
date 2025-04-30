package services

import (
	"context"
	"fmt"
	"go-grpc-inventory/pkg/models"
	"go-grpc-inventory/pkg/pb"
	"go-grpc-inventory/pkg/utils"

	"github.com/jinzhu/copier"
)

func (s *InventoryService) FindRecords(ctx context.Context, req *pb.EmptyRequest) (*pb.FindRecordsResponse, error) {
	// Simulate fetching records from a database or other source
	records := []*pb.InventoryResponse{
		{Name: "Record 1", Description: "Description 1"},
		{Name: "Record 2", Description: "Description 2"},
	}

	return &pb.FindRecordsResponse{Records: records}, nil
}

func (s *InventoryService) CreateInventory(ctx context.Context, req *pb.CreateInventoryRequest) (*pb.EmptyResponse, error) {
	var inventory models.Inventory

	err := copier.CopyWithOption(&inventory, req, copier.Option{IgnoreEmpty: true})
	if err != nil {
		return nil, err
	}

	if err := s.DB.Create(&inventory).Error; err != nil {
		return nil, err
	}

	return nil, err
}

func (s *InventoryService) GetInventoryList(c context.Context, req *pb.PaginationRequest) (*pb.GetInventoryResponse, error) {
	var inventories []models.Inventory
	pagination := &models.Pagination{}

	var totalRecords int64

	err := copier.CopyWithOption(pagination, req, copier.Option{IgnoreEmpty: true})
	if err != nil {
		fmt.Println("error", err)
		return nil, err
	}

	if err := s.DB.Model(&models.Inventory{}).Count(&totalRecords).Scopes(utils.Paginate(pagination)).Find(&inventories).Error; err != nil {
		return nil, err
	}

	pbInventories := make([]*pb.InventoryResponse, len(inventories))

	for i, inventory := range inventories {
		pbInventory := &pb.InventoryResponse{
			Id:          inventory.ID.String(),
			Name:        inventory.Name,
			Description: inventory.Description,
		}

		pbInventories[i] = pbInventory
	}

	return &pb.GetInventoryResponse{Records: pbInventories, Meta: &pb.MetaData{
		TotalRecords: totalRecords,
		Page:         int64(pagination.Page),
		Size:         int64(pagination.Size),
	}}, nil
}
