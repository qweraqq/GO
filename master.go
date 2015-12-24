package hw8


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
	l net.Listener
}

func (t *Master) Register(args *RegisterArgs, res *RegisterReply) error {
	DPrintf("Register: worker %s\n", args.Worker)
	//t.registerChannel <- args.Worker
	t.Workers[args.Worker] = &WorkerInfo{args.Worker,1}
	res.OK = true
	return nil
}


// TODO:assigning alg
func (t *Master) HandOverJobs(_ *RequestWorker, res *MasterReply) error {
	DPrintf("Master is assigning worker to the client ... \n")
	//t.registerChannel <- args.Worker
	if len(t.Workers) == 0{
		return errors.New("no worker available now,try again...")
	}
	for k := range t.Workers{
		DPrintf("checking worker %s ... \n",k)
		if (t.Workers[k].status == 1) {
			t.Workers[k].status = 2
			res.Addr = t.Workers[k].address
			return nil
		}		
			
	}	
	
	return errors.New("no worker available now,try again...")
}

func  (t *Master) RunMaster(MasterAddr string,MasterPort string) {
	// Your code here
	
	DPrintf("Master is being inited...\n")
	// start listening workers to reg

	
	rpc.Register(t)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", "localhost:"+MasterPort)
	if e != nil {
		DPrintf("master listen error: %s\n",e)
		log.Fatal("listen error:", e)
	}
	
	go http.Serve(l, nil)
	
	DPrintf("Master has been inited\n")
	
}

