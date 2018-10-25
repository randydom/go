package main

import (
	"awesomeProject/RemoteCalculator/Calculator/impl"
	"fmt"
	"testing"
)

func TestOperation(t *testing.T){

	var testCases = []struct{
		M string
		F int32
		S int32
		Expected int32
		Remainder int32
	} {
		{M: "sum", F: 15, S: 7, Expected: 22, Remainder: 0},
		{M: "sum", F: 7, S: 15, Expected: 22, Remainder: 0},
		{M: "sum", F: -15, S: 7, Expected: -8, Remainder: 0},
		{M: "sub", F: 15, S: 7, Expected: 8, Remainder: 0},
		{M: "sub", F: -15, S: 7, Expected: -22, Remainder: 0},
		{M: "sub", F: -15, S: -7, Expected: -8, Remainder: 0},
		{M: "mul", F: 15, S: 7, Expected: 105, Remainder: 0},
		{M: "mul", F: -15, S: 7, Expected: -105,},
		{M: "mul", F: 11000000, S: 9000000, Expected: 99000000, Remainder: 0},
		{M: "mul", F: 12000000, S: 9000000, Expected: 108000000, Remainder: 0},
		{M: "div", F: 15, S: 7, Expected: 2, Remainder: 1},
		{M: "div", F: 15, S: -7, Expected: -2, Remainder: 1},
		{M: "div", F: 84, S: 6, Expected: 14, Remainder: 0},
	}

	for _, testcase := range testCases {

		t.Run(testcase.M, func(t *testing.T) {

			result, remainder := impl.NewBasicCalculator(testcase.M, testcase.F, testcase.S).Operate()
			if result != testcase.Expected {
				t.Error(fmt.Printf("%s %d with %d expected %d but instead got: %d and %d", testcase.M, testcase.F, testcase.S, testcase.Expected, result, remainder))
			}
		})
	}

	//t.Run("")
}



//func testSumFunc(f int32, s int32, r int32) func(t *testing.T) {
//	return func(t *testing.T) {
//		result := Sum(f,s)
//		if result != r{
//			t.Error(fmt.Printf("Summing %d + %d expected %d but instead got: %d", f,s,r,result))
//		}
//	}
//}


