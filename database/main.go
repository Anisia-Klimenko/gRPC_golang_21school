package main

import (
	"fmt"
	"log"
	"net"
	"time"

	// database "github.com/Anisia-Klimenko/gRPC_golang_21school/database"
	protos "github.com/Anisia-Klimenko/gRPC_golang_21school/protos/warehouse"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// pb "google.golang.org/grpc/examples/helloworld/helloworld"

func main() {

	gs := grpc.NewServer()
	protos.RegisterWarehouseServer(gs, &Warehouse{})
	reflection.Register(gs)
	var sem = make(chan int, 2)

	ports := []int{8765, 9876, 8697}
	for _, value := range ports {
		sem <- 1
		l, err := net.Listen("tcp", fmt.Sprintf(":%d", value))
		if err == nil {
			fmt.Printf("running node on : 127.0.0.1:%d\n", value)
			gs.Serve(l)
		}
		go fmt.Println()
		<-sem
	}
	close(sem)
	log.Fatalln("Did't found free port for db instance")
}

func Print(value chan int) {
	a := value
	fmt.Println("Key = ", a)
	time.Sleep(1 * time.Second)
}