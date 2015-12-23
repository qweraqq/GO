package hw8

//import "fmt"
import "errors"



// Job
type Arith int


func (t *Arith) Plus(args *DoJobArgs, reply *DoJobReply) error {
	var t1,t2 int;
	switch args.A.(type){
		case int32:
			t1 = 1
		case float64:
			t1 = 2
		case int64:
			t1 = 3
		default:
			t1 = 0		
	}
	switch args.B.(type){
		case int32:
			t2 = 1
		case float64:
			t2 = 2
		case int64:
			t2 = 3
		default:
			t2 = 0		
	}
	switch {
		case t1==1 && t2==1:
			var r int32
			//fmt.Printf("int32... rpc\n")
			a,_ := args.A.(int32)
			b,_ := args.B.(int32)
			r = a + b
			reply.T = "int32"
			reply.R = r

		case t1==2 && t2==2:
			var r float64
			//fmt.Printf("float64... rpc\n")
			a,_ := args.A.(float64)
			b,_ := args.B.(float64)
			r = a + b
			reply.T = "float64"
			reply.R = r

		case t1==3 && t2==3:
			var r int64
			//fmt.Printf("int64... rpc\n")
			a,_ := args.A.(int64)
			b,_ := args.B.(int64)
			r = a + b
			reply.T = "int64"
			reply.R = r
		default:
			//fmt.Printf("error... rpc\n")
			return errors.New("plus RPC error:wrong args")
	}
			
	return nil
}


func (t *Arith) Minus(args *DoJobArgs, reply *DoJobReply) error {
	var t1,t2 int;
	switch args.A.(type){
		case int32:
			t1 = 1
		case float64:
			t1 = 2
		case int64:
			t1 = 3
		default:
			t1 = 0		
	}
	switch args.B.(type){
		case int32:
			t2 = 1
		case float64:
			t2 = 2
		case int64:
			t2 = 3
		default:
			t2 = 0		
	}
	switch {
		case t1==1 && t2==1:
			var r int32
			//fmt.Printf("int32... rpc\n")
			a,_ := args.A.(int32)
			b,_ := args.B.(int32)
			r = a - b
			reply.T = "int32"
			reply.R = r

		case t1==2 && t2==2:
			var r float64
			//fmt.Printf("float64... rpc\n")
			a,_ := args.A.(float64)
			b,_ := args.B.(float64)
			r = a - b
			reply.T = "float64"
			reply.R = r

		case t1==3 && t2==3:
			var r int64
			//fmt.Printf("int64... rpc\n")
			a,_ := args.A.(int64)
			b,_ := args.B.(int64)
			r = a - b
			reply.T = "int64"
			reply.R = r
		default:
			//fmt.Printf("error... rpc\n")
			reply.T = "wrong args"
			return nil//errors.New("minus RPC error:wrong args")
	}
			
	return nil
}

func (t *Arith) Multiply(args *DoJobArgs, reply *DoJobReply) error {
	var t1,t2 int;
	switch args.A.(type){
		case int32:
			t1 = 1
		case float64:
			t1 = 2
		case int64:
			t1 = 3
		default:
			t1 = 0		
	}
	switch args.B.(type){
		case int32:
			t2 = 1
		case float64:
			t2 = 2
		case int64:
			t2 = 3
		default:
			t2 = 0		
	}
	switch {
		case t1==1 && t2==1:
			var r int32
			//fmt.Printf("int32... rpc\n")
			a,_ := args.A.(int32)
			b,_ := args.B.(int32)
			r = a * b
			reply.T = "int32"
			reply.R = r
		case t1==2 && t2==2:
			var r float64
			//fmt.Printf("float64... rpc\n")
			a,_ := args.A.(float64)
			b,_ := args.B.(float64)
			r = a * b
			reply.T = "float64"
			reply.R = r
		case t1==3 && t2==3:
			var r int64
			//fmt.Printf("int64... rpc\n")
			a,_ := args.A.(int64)
			b,_ := args.B.(int64)
			r = a * b
			reply.T = "int64"
			reply.R = r
		default:
			reply.T = "wrong args"
			
			//fmt.Printf("error... rpc\n")
			return nil//errors.New("multiply RPC error:wrong args")
	}
			
	return nil
}

func (t *Arith) Divide(args *DoJobArgs, reply *DoJobReply) error {
	var t1,t2 int;
	switch args.A.(type){
		case int32:
			t1 = 1
		case float64:
			t1 = 2
		case int64:
			t1 = 3
		default:
			t1 = 0		
	}
	switch args.B.(type){
		case int32:
			t2 = 1
		case float64:
			t2 = 2
		case int64:
			t2 = 3
		default:
			t2 = 0		
	}
	switch {
		case t1==1 && t2==1:
			var r int32
			//fmt.Printf("int32... rpc\n")
			a,_ := args.A.(int32)
			b,_ := args.B.(int32)

			if b == 0{
				reply.T = "divide by zero"
				return nil
			}
			r = a / b
			reply.T = "int32"
			reply.R = r

		case t1==2 && t2==2:
			var r float64
			//fmt.Printf("float64... rpc\n")
			a,_ := args.A.(float64)
			b,_ := args.B.(float64)

			if b == 0{
				reply.T = "divide by zero"
				return nil
			}
			r = a / b
			reply.T = "float64"
			reply.R = r

		case t1==3 && t2==3:
			var r int64
			//fmt.Printf("int64... rpc\n")
			a,_ := args.A.(int64)
			b,_ := args.B.(int64)
			if b == 0{
				reply.T = "divide by zero"
				return nil //errors.New("divide RPC error:divide by zero")
			}

			r = a / b
			reply.T = "int64"
			reply.R = r
		default:
			//fmt.Printf("error... rpc\n")
			reply.T = "wrong args"
			return nil//errors.New("divide RPC error:wrong args")
	}
			
	return nil
}


