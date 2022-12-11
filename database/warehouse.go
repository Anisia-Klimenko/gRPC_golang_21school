package main

import (
	"context"

	protos "github.com/Anisia-Klimenko/gRPC_golang_21school/protos/warehouse"
	database "github.com/Anisia-Klimenko/gRPC_golang_21school/database"
)

// Warehouse Currency is a gRPC database it implements the methods defined by the Currencydatabase interface
type Warehouse struct {
	// log hclog.Logger
}

// NewWarehouse NewCurrency creates a new Currency database
func NewWarehouse() *Warehouse {
	return &Warehouse{}
}

// GetItem GetRate implements the Currencydatabase GetRate method and returns the currency exchange rate
// for the two given currencies.
func (c *Warehouse) GetItem(ctx context.Context, rr *protos.Item) (*protos.Item, error) {
	response, err := database.GetItemFromBackup(rr)
	if err != nil {
		return &protos.Item{}, err
	}

	return &protos.GetItemResponse{Name: data}, nil
}

func (c *Warehouse) SetItem(ctx context.Context, id *protos.Item, data *protos.Item) (*protos.OperationResultResponse, error) {
	response, err := database.SetItemToBackup(id, data)
	if err != nil {
		data = err
	}else {
		data = "The elem was created\n"}
	return &protos.OperationResultResponse{Msg: data} , nil
}

func (c *Warehouse) DeleteItem(ctx context.Context, rr *protos.Item) (*protos.OperationResultResponse, error) {
	response := database.DeleteItemFromBackup(rr)
	//if err != nil {
	//	data = err
	//} else {
	//	data = "The item was deleted\n"
	//}

	return &protos.OperationResultResponse{Msg: data}, nil
}
