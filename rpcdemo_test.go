package rpcdemo

import (
	"testing"
	"fmt"
	"log"
	"net/rpc"
	"time"
)


func TestServer(t *testing.T) {
	fmt.Printf("Test: Server ...\n")

	go RunServer("0.0.0.0","1234",0)
	time.Sleep(100)
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	args := &Args{"float64",float64(8.5),float64(0)}
	reply := new(Reply)
	err = client.Call("Arith.Divide", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %f*%f=%f", args.A, args.B, reply.R)
	fmt.Printf("  ... Test Passed\n")
}


