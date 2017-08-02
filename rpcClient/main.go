package main

import (
	"log"
	"net/rpc/jsonrpc"
)

type Agrs struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

// http  rpc  方式
func main() {
	//if len(os.Args) != 2 {
	//	log.Println("Usage: ", os.Args[0], "server")
	//	os.Exit(1)
	//}
	//
	//serverAddress := os.Args[1]
	//client, err := rpc.DialHTTP("tcp",serverAddress+ ":1234")

	// http  rpc  方式
	//client, err := rpc.DialHTTP("tcp", ":1234")
	//tcp rpc 方式
	//client,err:=rpc.Dial("tcp",":1234")

	//json rcp 方式
	client,err:=jsonrpc.Dial("tcp",":1234")
	if err != nil {
		log.Fatal("rpc dial fail")
	}

	args := Agrs{12, 5}
	var reply int
	err = client.Call("Math.Multiply", args, &reply)

	if err != nil {
		log.Fatal("Math.Mutiply call fail")
	}

	log.Printf("Math: %d*%d=%d\n", args.A, args.B, reply)

	var quotient Quotient
	err = client.Call("Math.Divide", args, &quotient)

	if err != nil {
		log.Fatal("Math.Divide call fail")
	}
	log.Printf("Math: %d/%d=%d remainder %d\n", args.A, args.B, quotient.Quo, quotient.Rem)
}

// tcp  rpc  方式
