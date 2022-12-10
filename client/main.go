package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

// ~$ ./warehouse-cli -H 127.0.0.1 -P 8765
// Connected to a database of Warehouse 13 at 127.0.0.1:8765
// Known nodes:
// 127.0.0.1:8765
// 127.0.0.1:9876
// 127.0.0.1:8697

func knownHosts(ports []int) {
	for _, value := range ports {
		_, err := net.Listen("tcp", fmt.Sprintf(":%d", value))
		if err == nil {
			fmt.Printf("%s:%d\n", "127.0.0.1", value)
		}
	}
}

func main() {

	fHost := *flag.String("host", "127.0.0.1", "--host 127.0.0.1")
	fPort := *flag.String("port", "8765", "--port 8765")

	if flag.NFlag() != 2 {
		log.Fatalln("Usage: ./warehouse-cli -H 127.0.0.1 -P 8765")
	}

	ports := []int{8765, 9876, 8697}
	var opts []grpc.DialOption
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", fHost, fPort), opts...)
	if err != nil {
		fmt.Println("Can't connect to a database\nKnown nodes:")
		knownHosts(ports)
		defer conn.Close()
	} else {
		
	}
}
// go func() {
// 	knownHosts(ports)
// }()
