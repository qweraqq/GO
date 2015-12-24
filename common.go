package hw8

import "fmt"
import "net/rpc"

const Debug = 1

func DPrintf(format string, a ...interface{}) (n int, err error) {
	if Debug > 0 {
		n, err = fmt.Printf(format, a...)
	}
	return
}
// four kinds of jobs
const (
	Add = "Add"
	Minus = "Minus"
	Multiply = "Multiply"
	Divide = "Divide"
	MasterPort = "11234"
	NumWorkers = 4
	
	
)


// RPC arguments and replies.  Field names must start with capital letters,
// otherwise RPC will break.
type DoJobArgs struct {
	//T string
	A, B interface{}
}

type DoJobReply struct {
	T string
	R interface{}
}


type RequestWorker struct{
	Dummy int
}

type MasterReply struct {
	Addr string
}

type ShutdownArgs struct {
}

type ShutdownReply struct {
	Njobs int
	OK    bool
}

type RegisterArgs struct {
	Worker string
}

type RegisterReply struct {
	OK bool
}


//
// call() sends an RPC to the rpcname handler on server srv
// with arguments args, waits for the reply, and leaves the
// reply in reply. the reply argument should be the address
// of a reply structure.
//
// call() returns true if the server responded, and false
// if call() was not able to contact the server. in particular,
// reply's contents are valid if and only if call() returned true.
//
// you should assume that call() will time out and return an
// error after a while if it doesn't get a reply from the server.
//
func call(srv string, rpcname string,
	args interface{}, reply interface{}) bool {
	c, errx := rpc.DialHTTP("tcp", srv)
	if errx != nil {
		return false
	}
	defer c.Close()

	err := c.Call(rpcname, args, reply)
	if err == nil {
		return true
	}

	DPrintf("rpc %s error:%s\n",rpcname,err)
	
	return false
}
