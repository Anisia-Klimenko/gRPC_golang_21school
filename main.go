package main

import (
	"fmt"

	database "github.com/Anisia-Klimenko/gRPC_golang_21school/database"
	protos "github.com/Anisia-Klimenko/gRPC_golang_21school/protos/warehouse"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

// pb "google.golang.org/grpc/examples/helloworld/helloworld"

func main() {
	log := hclog.Default()

	// create a new gRPC server, use WithInsecure to allow http connections
	gs := grpc.NewServer()
	fmt.Println("sada")

	// create an instance of the Currency server
	c := database.NewWarehouse(log)

	// register the currency server
	protos.RegisterWarehouseServer(gs, c)
}
