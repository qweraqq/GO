package hw8

import (
	"testing"
	"fmt"
	"time"	
)

var c chan int
var arith *Arith
func TestServer(t *testing.T) {
	c = make(chan int)
	arith = new(Arith)
	fmt.Printf("Test: Server ...\n")
	m := new(Master)
	m.Workers = make(map[string]*WorkerInfo)
	w1 := new(Worker)
	go m.RunMaster(":","11234")
	time.Sleep(100000)
	go w1.RunWorker("0.0.0.0:11234","12135",0)
	time.Sleep(100000)
	
	
	reply := new(MasterReply)
	dummy := new(RequestWorker)
	for ok := call(":11234","Master.HandOverJobs", dummy, &reply);ok == false;{
		fmt.Printf("retrying master...\n")
		time.Sleep(time.Second)
		ok = call(":11234","Master.HandOverJobs", dummy, &reply)
	}


	DPrintf("worker is at port %s\n",reply.Addr)
	args := &DoJobArgs{float64(8.5),float64(0)}
	reply2 := new(DoJobReply)
	for ok := call(reply.Addr,"Arith.Divide", args, &reply2);ok == false;{
		fmt.Printf("retrying worker...\n")
		time.Sleep(time.Second)
		ok = call(reply.Addr,"Arith.Divide", args, &reply2)
	}
	fmt.Printf("Arith: %s %f*%f=%f", reply2.T, args.A, args.B, reply2.R)
	
	fmt.Printf("  ... Test Passed\n")
}



