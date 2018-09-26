package main

import (
	"awesomeProject/FileProcessor"
	"awesomeProject/ServerFileProcessor"
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
	"time"
)

func MultiplyCall(client *rpc.Client, args *SFProcessor.Args) {

	var reply int

	mulCall := client.Go("Arith.Multiply", args, &reply, nil)
	replyCall := <-mulCall.Done

	if replyCall.Error != nil {
		log.Fatal("arith error: ", replyCall.Error)
	}

	fmt.Printf("Arith: %d * %d = %d\n", args.A, args.B, reply)

}

func DivisionCall(client *rpc.Client, args *SFProcessor.Args) {
	//Asynchronous call
	quotient := new(SFProcessor.Quotient)

	divCall := client.Go("Arith.Division", args, &quotient, nil)
	replyCall := <-divCall.Done

	// will be equal to divCall
	if replyCall.Error != nil {
		log.Fatal("arith error: ", replyCall.Error)
	}

	fmt.Printf("Arith: %d / %d = %d, the reminder is: %d\n", args.A, args.B, quotient.Quo, quotient.Rem)
}

func SendNewMessage(client *rpc.Client, m *SFProcessor.Message) bool{
	var reply string

	msgCall := client.Go("Diavlos.NewMessage", m, &reply, nil)
	replyCall := <- msgCall.Done

	if replyCall.Error != nil {
		log.Fatal("Diavlos error: ", replyCall.Error)
	}

	fmt.Println("The server replied:\n", reply)

	if strings.Contains(reply, "Good bye"){
		fmt.Println("Client: -Good bye from me also.")
		return true
	}

	return false
}

func NewFileProcessor(client *rpc.Client, m *SFProcessor.Message){

	var report FileProcessor.Report

	rptCall := client.Go("Diavlos.FileProcessor", m, &report, nil)
	replyCall := <- rptCall.Done

	if replyCall.Error != nil {
		log.Fatal("Diavlos error: ", replyCall.Error)
	}

	fmt.Printf("\n*****\nThread Id: %d text: %s\nNumber of:\nparagraphs: %d\nsentences: %d\nword: %d\nletters: %d\n", report.ThreadId,
		report.FileName, report.P.NumParagraph, report.S.NumSentence, report.W.NumWord, report.L.NumLetter)
}

func main() {

	var A, B int
	var T string
	var flag bool

	for{
		//connect to this socket
		addr := "192.168.100.162:8024"
		client, err := rpc.Dial("tcp", addr)

		if err != nil {
			log.Fatal("dialing: ", err)
		}

		fmt.Print("Enter the first integer number: ")
		fmt.Scan(&A)

		fmt.Print("Enter the second integer number: ")
		fmt.Scan(&B)

		args := &SFProcessor.Args{A, B,}

		MultiplyCall(client, args)

		DivisionCall(client, args)

		fmt.Print("Do you want to send a text? ")

		T, _ = bufio.NewReader(os.Stdin).ReadString('\n')

		msg := &SFProcessor.Message{T}

		flag = SendNewMessage(client, msg)
		if flag{
			break
		}
	}

	fmt.Println("Last call to server: ", time.Now())
}