package hw8

import (
	"testing"
	"fmt"
	"time"	
)

var c chan int


func RunMaster(master_addr string,master_port string){
	
	master := new(Master)
	master.Workers = make(map[string]*WorkerInfo)
	master.RunMaster(master_addr, master_port)
	
}

func RunWorker(master_addr string,master_port string,worker_addr string,worker_port string){
	w := new(Worker)
	w.RunWorker(master_addr+":"+master_port,worker_addr+":"+worker_port,0)
}

func RunClient(master_addr string,master_port string,op int){
	master_reply := new(MasterReply)
	dummy := new(RequestWorker)
	do_job_args := &DoJobArgs{float64(8.5),float64(op)}
	do_job_reply := new(DoJobReply)

	for ok := call(master_addr+":"+master_port,"Master.HandOverJobs", dummy, &master_reply);ok == false;{
		DPrintf("retrying master...\n")
		time.Sleep(time.Second)
		ok = call(master_addr+":"+master_port,"Master.HandOverJobs", dummy, &master_reply)
	}
	DPrintf("worker is at %s\n",master_reply.Addr)
	
	for i:=0;i<1500000;i=i+1{
		for ok := call(master_reply.Addr,"Arith.Divide", do_job_args, &do_job_reply);ok == false;{
			fmt.Printf("retrying worker...\n")
			time.Sleep(time.Second)
			ok = call(master_reply.Addr,"Arith.Divide", do_job_args, &do_job_reply)
		}
		if i%100000==0{
			DPrintf("Arith: %s %f*%f=%f\n", do_job_reply.T, do_job_args.A, do_job_args.B, do_job_reply.R)
		}
	}
	c<-1
	
}
func TestBasic(t *testing.T){
	c = make(chan int)
	var Workerport [4]string
	Workerport = [4]string{"11235","11236","11237","11238"}
	go func(){
		RunMaster("","11234")
		c<-1
	}()
	<-c  // wait master
	
	for i:=0;i<NumWorkers;i=i+1 {
		go RunWorker("","11234","",Workerport[i])
	}

	for i:=0;i<NumWorkers;i=i+1 {
		go RunClient("","11234",i)
	}
	
	for i:=0;i<NumWorkers;i=i+1 {
		<-c
	}
	

}





