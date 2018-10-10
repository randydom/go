package main

import (
	pb "awesomeProject/RemoteCalculator/remoteCalcpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:8024"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	var opr string
	var f, s int32

	fmt.Println("Choose an operation from {sum, sub, mul, div} and enter two integers with spaces")
	fmt.Scanf("%s %d %d", &opr, &f, &s)

	c := pb.NewCalculatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	r, err := c.CalculateOperation(ctx, &pb.Operation{Operator: opr, Numbers: &pb.Akeraios{First: f, Second: s,},})


	if err != nil {
		log.Fatalf("could not operate the numbers: %v", err)
	}

	log.Printf("\nThe outcome is: %d\nThe remainder: %d\n", r.Result, r.Remainder)

}
