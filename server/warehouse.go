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
func (c *Warehouse) GetRate(ctx context.Context, rr *protos.GetIt) (*protos.RateResponse, error) {
	c.log.Info("Handle request for GetRate", "base", rr.GetBase(), "dest", rr.GetDestination())
	return &protos.RateResponse{Rate: 0.5}, nil
}
