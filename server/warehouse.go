package server

import (
	"context"

	protos "github.com/Anisia-Klimenko/gRPC_golang_21school/protos/warehouse"
	"github.com/hashicorp/go-hclog"
)

// Currency is a gRPC server it implements the methods defined by the CurrencyServer interface
type Warehouse struct {
	log hclog.Logger
}

// NewCurrency creates a new Currency server
func NewWarehouse(l hclog.Logger) *Warehouse {
	return &Warehouse{l}
}

// GetRate implements the CurrencyServer GetRate method and returns the currency exchange rate
// for the two given currencies.
func (c *Warehouse) GetItem(ctx context.Context, rr *protos.UUID) (*protos.GetItemResponse, error) {
	data := "kek"
	
	return &protos.GetItemResponse{Name: data}, nil
}

func (c *Warehouse) SetItem(ctx context.Context, rr *protos.UUID) (*protos.OperationResultResponse, error) {
	data := "kek"
	
	return &protos.OperationResultResponse{Msg: data} , nil
}

func (c *Warehouse) DeleteItem(ctx context.Context, rr *protos.UUID) (*protos.OperationResultResponse, error) {
	data := "kek"
	
	return &protos.OperationResultResponse{Msg: data}, nil
}


