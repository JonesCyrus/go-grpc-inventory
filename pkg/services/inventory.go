package services

import (
	"context"
	"go-grpc-inventory/pkg/pb"
)

func (s *InventoryService) FindRecords(ctx context.Context, req *pb.EmptyRequest) (*pb.FindRecordsResponse, error) {
	// Simulate fetching records from a database or other source
	records := []*pb.InventoryResponse{
		{Name: "Record 1", Description: "Description 1"},
		{Name: "Record 2", Description: "Description 2"},
	}

	return &pb.FindRecordsResponse{Records: records}, nil
}
