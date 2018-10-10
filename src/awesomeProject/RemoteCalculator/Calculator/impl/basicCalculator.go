package impl

import (
	"awesomeProject/RemoteCalculator/Calculator"
	"fmt"
)

type (

	BasicCalculator struct{
		Opr string
		F int32
		S int32
	}

)

func NewBasicCalculator(o string, nf int32, ns int32) Calculator.MyCalculator{

	return &BasicCalculator{Opr: o, F: nf, S: ns,}

}

func (b *BasicCalculator) Operate() (int32, int32){

	// The nSet contains the info for the outcome.
	nSet := Calculator.NumberSet{First: b.F, Second: b.S, Outcome: 0, Remainder: 0,}

	// This block of code determines which operation should the program use to calculate the correct outcome
	if b.Opr == "sum" {
		nSet.Outcome =  Sum(nSet.First, nSet.Second)
	}else if b.Opr == "mul" {
		nSet.Outcome =  Multiply(nSet.First, nSet.Second)
	}else if b.Opr == "sub" {
		nSet.Outcome =  Substract(nSet.First, nSet.Second)
	}else if b.Opr == "div" {
		nSet.Outcome, nSet.Remainder =  Divide(nSet.First, nSet.Second)
	}else {
		fmt.Println("Sorry, but I cannot find this operation.") // If the operation is wrong then a msg will be displayed.
	}

	return nSet.Outcome, nSet.Remainder
}

//func PrintOutcome(o string, f int32, s int32){
//	fmt.Printf("The operation is: %s\nThe outcome is: %d\nand the remainder: %d\n", o, f, s)
//}