package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	protos "github.com/Anisia-Klimenko/gRPC_golang_21school/protos/warehouse"
	"github.com/google/uuid"
)

type DB struct {
	UUID protos.UUID
	Elem protos.GetItemResponse
}

type File struct {
	id       int
	filePath string
}

var Backups = map[string]string{
	"main":    "backup/data.json",
	"backup1": "backup/replica1.json",
	"backup2": "backup/replica2.json",
}

func GetItemFromBackup(uuid protos.UUID) (protos.GetItemResponse, error) {
	var db []DB
	var err error
	for _, file := range Backups {
		f, _ := ioutil.ReadFile(file)
		err = json.Unmarshal(f, &db)
		if err == nil {
			break
		}
	}
	if err != nil {
		return protos.GetItemResponse{}, errors.New("backups broken")
	}
	for _, elem := range db {
		if elem.Elem.Name == uuid.Value {
			return elem.Elem, nil
		}
	}
	return protos.GetItemResponse{}, errors.New("element not found")
}

func SetItemToBackup(elem string) protos.OperationResultResponse {
	var err error
	var newElem = DB{protos.UUID{Value: uuid.New().String()}, protos.GetItemResponse{Name: elem}}
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
		return protos.OperationResultResponse{Msg: "error: " + err.Error()}
	}
	return protos.OperationResultResponse{Msg: "created (2 replicas)"}
}

func DeleteItemToBackup(uuid string) protos.OperationResultResponse {
	var db []DB
	var err error
	for _, file := range Backups {
		f, _ := ioutil.ReadFile(file)
		err = json.Unmarshal(f, &db)
		if err == nil {
			break
		}
	}
	if err != nil {
		return protos.OperationResultResponse{Msg: "backups broken"}
	}
	for index, elem := range db {
		if elem.Elem.Name == uuid {
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
	return protos.OperationResultResponse{Msg: "deleted (2 replicas)"}
}
