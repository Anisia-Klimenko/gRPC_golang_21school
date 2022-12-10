package main

import (
	"context"

	protos "github.com/Anisia-Klimenko/gRPC_golang_21school/protos/warehouse"
	server "github.com/Anisia-Klimenko/gRPC_golang_21school/server"
)

// Warehouse Currency is a gRPC server it implements the methods defined by the CurrencyServer interface
type Warehouse struct {
	// log hclog.Logger
}

// NewWarehouse NewCurrency creates a new Currency server
func NewWarehouse() *Warehouse {
	return &Warehouse{}
}

// GetItem GetRate implements the CurrencyServer GetRate method and returns the currency exchange rate
// for the two given currencies.
func (c *Warehouse) GetItem(ctx context.Context, rr *protos.UUID) (*protos.GetItemResponse, error) {
	response, err := server.GetItemFromBackup(rr)
	if err != nil {
		return &protos.GetItemResponse{}, err
	}

	return &response, nil
}

func (c *Warehouse) SetItem(ctx context.Context, rr string) (*protos.OperationResultResponse, error) {
	response := server.SetItemToBackup(rr)
	//if err != nil {
	//	data = err
	//} else {
	//	data = "The elem was created\n"
	//}
	return &response, nil
}

func (c *Warehouse) DeleteItem(ctx context.Context, rr *protos.UUID) (*protos.OperationResultResponse, error) {
	response := server.DeleteItemFromBackup(rr)
	//if err != nil {
	//	data = err
	//} else {
	//	data = "The item was deleted\n"
	//}

	return &response, nil
}
