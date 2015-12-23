package hw8


import "fmt"
import "net/rpc"
import "net"
import "log"
import "net/http"
import "errors"
type WorkerInfo struct {
	address string
	// You can add definitions here.
	status int //0-error,1-idle,2-busy
}

type Master struct{

	// Map of registered workers that you need to keep up to date
	Workers map[string]*WorkerInfo 	
}

func (t *Master) Register(args *RegisterArgs, res *RegisterReply) error {
	DPrintf("Register: worker %s\n", args.Worker)
	//t.registerChannel <- args.Worker
	t.Workers[args.Worker] = &WorkerInfo{args.Worker,1}
	res.OK = true
	return nil
}

func (t *Master) HandOverJobs(_ *RequestWorker, res *MasterReply) error {
	DPrintf("Master is HandOverJobs \n")
	//t.registerChannel <- args.Worker
	if len(t.Workers) == 0{
		return errors.New("no worker available now,try again...")
	}
	for k := range t.Workers{
		
		if (t.Workers[k].status == 1) {
			res.Addr = t.Workers[k].address
		}else{
			return errors.New("no worker available now,try again...")
		}	
	}	
	
	return nil
}

func  (t *Master) RunMaster(MasterAddr string,MasterPort string) {
	// Your code here
	
	fmt.Printf("Master being inited...\n")
	// start listening workers to reg

	
	rpc.Register(t)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", "localhost:"+MasterPort)
	if e != nil {
		DPrintf("master listen error: %s\n",e)
		log.Fatal("listen error:", e)
	}
	
	go http.Serve(l, nil)
	
	fmt.Printf("Master inited done...\n")
	
}

