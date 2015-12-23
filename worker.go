package hw8

import "fmt"
import "log"
import "net/rpc"
import "net"
import "net/http"
import "time"

// Worker is a server waiting for DoJob or Shutdown RPCs

type Worker struct {
	name   string
	nRPC   int
	nJobs  int
	l      net.Listener
}


// Tell the master we exist and ready to work
func Register(master string, me string) {
	args := &RegisterArgs{}
	args.Worker = me
	var reply RegisterReply
	for ok := call(master, "Master.Register", args, &reply);ok == false;{
		fmt.Printf("Register: RPC %s register error,retrying...\n", master)
		time.Sleep(time.Second)
		ok = call(master, "Master.Register", args, &reply)
	}
	fmt.Printf("Register: RPC %s register succeed\n", master)
}


// Set up a connection with the master, register with the master,
// and wait for jobs from the master
func (wk *Worker)RunWorker(MasterAddress string, port string,
	nRPC int) {
	DPrintf("RunWorker %s\n", port)
	//wk := new(Worker)
	wk.name = port	
	wk.nRPC = nRPC
	
	arith := new(Arith)
	rpc.Register(arith)
	//rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":"+port)
	if e != nil {
		DPrintf("RunWorker: worker %s error %s\n", port,e)
		log.Fatal("RunWorker: worker ", port, " error: ", e)
	}
	//wk.l = l
	Register(MasterAddress, ":"+port)
	go http.Serve(l, nil)
	
	

	DPrintf("Worker %s is running...\n", port)
}
