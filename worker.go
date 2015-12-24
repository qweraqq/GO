package hw8

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
	l net.Listener
}


// Tell the master we exist and ready to work
func Register(master string, me string) {
	args := &RegisterArgs{}
	args.Worker = me
	var reply RegisterReply
	for ok := call(master, "Master.Register", args, &reply);ok == false;{
		DPrintf("Register RPC %s register error,retrying...\n", master)
		time.Sleep(time.Second)
		ok = call(master, "Master.Register", args, &reply)
	}
	DPrintf("Worker %s has been registered \n", me)
}


// Set up a connection with the master, register with the master,
// and wait for jobs from the client
func (wk *Worker)RunWorker(MasterAddress string, workerAddress string,
	nRPC int) {
	DPrintf("Runing Worker %s\n", workerAddress)
	//wk := new(Worker)
	wk.name = workerAddress
	wk.nRPC = nRPC
	
	arith := new(Arith)
	rpc.Register(arith)
	//rpc.HandleHTTP()
	l, e := net.Listen("tcp", workerAddress)
	if e != nil {
		DPrintf("RunWorker: worker %s error %s\n", workerAddress,e)
		log.Fatal("RunWorker: worker ", workerAddress, " error: ", e)
	}
	wk.l = l
	Register(MasterAddress, workerAddress)
	go http.Serve(l, nil)
	
	

	DPrintf("Worker %s is running...\n", workerAddress)
}
