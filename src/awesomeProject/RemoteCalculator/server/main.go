package main

import (
	"awesomeProject/RemoteCalculator/Calculator/impl"
	pb "awesomeProject/RemoteCalculator/remoteCalcpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"math/rand"
	"net"
)

const (
	port = ":8024"
)

var operations = map[int32]string{1:"sum", 2: "sub", 3: "mul", 4: "div",}


//
type (
	server struct{}
)

// CalculateOperation implements remoteCalcpb.
func (s *server) CalculateOperation(ctx context.Context, in *pb.Operation) ( *pb.Outcome, error) {
	var oc, rm int32

	oc, rm = impl.NewBasicCalculator(in.Operator, in.Numbers.First, in.Numbers.Second).Operate()
	fmt.Println(oc,rm)

	return &pb.Outcome{Result: oc, Remainder: rm,}, nil
}

func (s *server) CalculateManyOperations(oper *pb.Operation, stream pb.Calculator_CalculateManyOperationsServer) error {

	var randomOperator, numA, numB int32

	for i:=0; i<10; i++ {
		randomOperator = rand.Int31n(4)+1
		oper.Operator = operations[randomOperator]

		numA = rand.Int31n(100)
		numB = rand.Int31n(100)

		oper.Numbers.First = numA
		oper.Numbers.Second = numB

		oc, rm := impl.NewBasicCalculator(oper.Operator, oper.Numbers.First, oper.Numbers.Second).Operate()
		fmt.Printf("the operation is: %s\nThe first number: %d\nThe second number: %d\n\n",oper.Operator, oper.Numbers.First, oper.Numbers.Second)

		stream.Send(&pb.Outcome{Result: oc, Remainder: rm})

	}

	return nil
}

func main() {

	fmt.Println("The server is listening..")

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterCalculatorServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}