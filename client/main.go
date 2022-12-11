package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	protos "github.com/Anisia-Klimenko/gRPC_golang_21school/protos/warehouse"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ~$ ./warehouse-cli -H 127.0.0.1 -P 8765
// Connected to a database of Warehouse 13 at 127.0.0.1:8765
// Known nodes:
// 127.0.0.1:8765
// 127.0.0.1:9876
// 127.0.0.1:8697

func printKnownHosts() {
	fmt.Println("Known nodes:")
	for _, value := range ports {
		_, err := net.Listen("tcp", fmt.Sprintf(":%d", value))
		if err != nil {
			fmt.Printf("%s:%d\n", "127.0.0.1", value)
		}
	}
}

var ports = []int{8765, 9876, 8697}

type RequestType int

const (
	GET    RequestType = 1
	SET    RequestType = 2
	DELETE RequestType = 3
)

type Request struct {
	Type RequestType
	Body any
}

func connectIfPossible(host string, port string) (conn *grpc.ClientConn, err error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err = grpc.Dial(fmt.Sprintf("%s:%s", host, port), opts...)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Connected to a database of Warehouse 13 at %s:%s\n", host, port)
		printKnownHosts()
	} else {
		for _, p := range ports {
			conn, err = grpc.Dial(fmt.Sprintf("%s:%d", host, p), opts...)
			if err == nil {
				fmt.Println("Requested node is node available")
				fmt.Printf("Connected to a database of Warehouse 13 at %s:%d\n", host, p)
				return
			}
		}
	}
	return
}

func init() {
	log.SetPrefix("Error: ")
	log.SetFlags(0)
}

func main() {

	fHost := flag.String("H", "127.0.0.1", "-H 127.0.0.1")
	fPort := flag.String("P", "8765", "-P 8765")
	flag.Parse()

	if flag.NFlag() != 2 {
		fmt.Println(flag.NFlag())
		log.Fatalln("Usage: ./warehouse-cli -H 127.0.0.1 -P 8765")
	}

	for {
		conn, err := connectIfPossible(*fHost, *fPort)
		if err != nil {
			log.Fatalln("This and all other nodes are not available")
		} else {
			client := protos.NewWarehouseClient(conn)
			client.GetItem(context.Background(), &protos.ItemRequest{})
			time.Sleep(time.Duration(5) * time.Second)
		}
	}
	// }
	// defer conn.Close()
	// } else {
	// 	conn, err = grpc.Dial(fmt.Sprintf("%s:%s", fHost, fPort), opts...)
	// 	conn, err = grpc.Dial(fmt.Sprintf("%s:%s", fHost, fPort), opts...)

	// r := Request{GET, protos.ItemRequest{UUID: "asdas"}}
	// 	// response := protos.Item{}
	// }
}

// go func() {
// 	knownHosts(ports)
// }()
