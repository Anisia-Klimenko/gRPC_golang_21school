package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"os"

	//database "github.com/Anisia-Klimenko/gRPC_golang_21school/database"
	protos "github.com/Anisia-Klimenko/gRPC_golang_21school/protos/warehouse"
)

// Warehouse Currency is a gRPC database it implements the methods defined by the Currencydatabase interface
type Warehouse struct {
	// log hclog.Logger
}

// NewWarehouse NewCurrency creates a new Currency database
func NewWarehouse() *Warehouse {
	return &Warehouse{}
}

var Backups = map[string]string{
	"main":    "backup/data.json",
	"backup1": "backup/replica1.json",
	"backup2": "backup/replica2.json",
}

// GetItem GetRate implements the Currencydatabase GetRate method and returns the currency exchange rate
// for the two given currencies.
func (c *Warehouse) GetItem(ctx context.Context, rr *protos.ItemRequest) (*protos.Item, error) {
	var db []protos.Item
	var err error
	for _, file := range Backups {
		f, _ := ioutil.ReadFile(file)
		err = json.Unmarshal(f, &db)
		if err == nil {
			break
		}
	}
	if err != nil {
		return &protos.Item{}, errors.New("backups broken")
	}
	for _, elem := range db {
		if elem.UUID == rr.UUID {
			return &elem, nil
		}
	}
	return &protos.Item{}, errors.New("element not found")
}

func (c *Warehouse) SetItem(ctx context.Context, rr *protos.Item) (*protos.OperationResultResponse, error) {
	var err error
	var newId uuid.UUID
	if len(rr.UUID) == 0 {
		newId = uuid.New()
	} else {
		newId, err = uuid.Parse(rr.UUID)
		if err != nil {
			return &protos.OperationResultResponse{Msg: "error: key is not a proper uuid4"}, err
		}
	}
	var newElem = protos.Item{UUID: newId.String(), Content: rr.Content}
	text, _ := json.Marshal(newElem)
	for _, file := range Backups {
		f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		_, err = f.WriteString(string(text))
		if err == nil {
			break
		}
		f.Close()
	}
	if err != nil {
		return &protos.OperationResultResponse{Msg: "error: " + err.Error()}, err
	}
	return &protos.OperationResultResponse{Msg: "created (2 replicas)"}, nil
}

func (c *Warehouse) DeleteItem(ctx context.Context, rr *protos.Item) (*protos.OperationResultResponse, error) {
	var db []protos.Item
	var err error
	for _, file := range Backups {
		f, _ := ioutil.ReadFile(file)
		err = json.Unmarshal(f, &db)
		if err == nil {
			break
		}
	}
	if err != nil {
		return &protos.OperationResultResponse{Msg: "backups broken"}, err
	}
	for index, elem := range db {
		if elem.UUID == rr.UUID {
			db = append(db[:index], db[index+1:]...)
			break
		}
	}
	newDb, _ := json.Marshal(db)
	for _, file := range Backups {
		f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		err = f.Truncate(0)
		_, err = f.Seek(0, 0)
		_, err = fmt.Fprintf(f, "%s", string(newDb))
		if err == nil {
			break
		}
	}
	return &protos.OperationResultResponse{Msg: "deleted (2 replicas)"}, nil
}
