package main

import (
	pb "awesomeProject/RemoteCalculator/remoteCalcpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

const (
	address = "127.0.0.1:8024"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	var opr string
	var f, s int32

	fmt.Println("You have client_version2 and you will execute 10 random operations.")


	c := pb.NewCalculatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	req := &pb.Operation{
		Operator: opr,
		Numbers: &pb.Akeraios{
			First: f, Second: s,
		},
	}

	stream, err := c.CalculateManyOperations(ctx, req)

	if err != nil {
		log.Fatalf("could not operate the numbers: %v", err)
	}

	for{
		feature, err := stream.Recv()

		if err == io.EOF{
			break // We reached the end of file
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", c, err)
		}
		log.Printf("\nThe outcome is: %d\nThe remainder: %d\n", feature.Result, feature.Remainder)
	}

}
