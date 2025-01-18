package grpc

import (
	"context"
)

type InventoryServer struct{}

func (s *InventoryServer) GetProperties(ctx context.Context, req *Empty) (*PropertyList, error) {
	return &PropertyList{}, nil
}

func (s *InventoryServer) CreateTransaction(ctx context.Context, req *Transaction) (*TransactionResponse, error) {
	return &TransactionResponse{Success: true}, nil
}
