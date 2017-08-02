package main

import (
	"net/rpc"
	"errors"
	"log"
	"net/http"
	"net"
	"os"
	"net/rpc/jsonrpc"
)

type Agrs struct {
	A, B int
}

type Math int

func (m *Math)Multiply(args *Agrs, reply *int) error {
	*reply = args.A * args.B
	return nil;
}

type Quotient struct {
	Quo, Rem int
}

func (m *Math)Divide(args *Agrs, reply  *Quotient) error {
	if args.B == 0 {
		return errors.New("args B can be zero")
	}
	reply.Quo = args.A / args.B
	reply.Rem = args.A % args.B
	return nil;
}

func main1() {
	math := new(Math)
	rpc.Register(math)
	rpc.HandleHTTP()

	log.Println("rpc server is starting")
	log.Fatal(http.ListenAndServe(":1234", nil))
}

// tcp rpc  调用
func main() {
	math := new(Math)
	rpc.Register(math)

	tcpAddress, err := net.ResolveTCPAddr("tcp", ":1234")
	if err != nil {
		log.Println("rpc fail")
		os.Exit(1)
	}
	listener, err := net.ListenTCP("tcp", tcpAddress)
	if err != nil {
		log.Println("rpc fail")
		os.Exit(1)
	}

	log.Println("rpc server is starting")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		// tcp 方式rpc   调用
		//rpc.ServeConn(conn)
		//tcp 方式调用
		jsonrpc.ServeConn(conn)
	}

}


