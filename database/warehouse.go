package main

import (
	"context"
	protos "github.com/Anisia-Klimenko/gRPC_golang_21school/protos/warehouse"
)

// Currency is a gRPC server it implements the methods defined by the CurrencyServer interface
type Warehouse struct {
	// log hclog.Logger
}

// NewCurrency creates a new Currency server
func NewWarehouse() *Warehouse {
	return &Warehouse{}
}

// GetRate implements the CurrencyServer GetRate method and returns the currency exchange rate
// for the two given currencies.
func (c *Warehouse) GetItem(ctx context.Context, rr *protos.UUID) (*protos.GetItemResponse, error) {
	data, err := GetItemFromBackup(rr)
	if err != nil {
		data = err}
	return &protos.GetItemResponse{Name: data}, nil
}

func (c *Warehouse) SetItem(ctx context.Context, rr *protos.UUID) (*protos.OperationResultResponse, error) {
	data, err := SetItemFromBackup(rr)
	if err != nil {
		data = err
	}else {
		data = "The elem was created\n"}
	return &protos.OperationResultResponse{Msg: data} , nil
}

func (c *Warehouse) DeleteItem(ctx context.Context, rr *protos.UUID) (*protos.OperationResultResponse, error) {
	data, err := DeleteItemFromBackup(rr)
	if err != nil {
		data = err;
	}else {
		data = "The item was deleted\n"}

	return &protos.OperationResultResponse{Msg: data}, nil
}
