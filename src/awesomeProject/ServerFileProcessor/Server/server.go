package main

import (
	"awesomeProject/ServerFileProcessor"
	"fmt"
	"log"
	"net"
	"net/rpc"
)


func main() {

	fmt.Println("Launching server...")

	arith := new(SFProcessor.Arith)

	diav := new(SFProcessor.Diavlos)

	rpc.Register(arith)
	rpc.Register(diav)
	//rpc.HandleHTTP()

	//listen on all interfaces
	ln, e := net.Listen("tcp", ":8024")

	if e != nil {
		log.Fatal("Listen error: ", e)
	}

	//close Listener whenever we stop
	defer ln.Close()

	rpc.Accept(ln)


}