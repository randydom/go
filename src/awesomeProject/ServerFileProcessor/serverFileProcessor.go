package SFProcessor

import (
	"errors"
	"strings"
)

type Args struct{
	A, B int
}

type Quotient struct{
	Quo, Rem int
}

type Message struct{
	Msg string
}

type(
	Arith int
	Diavlos string
)


func (t *Arith) Multiply( args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Division( args *Args, quo *Quotient) error {

	if args.B == 0{
		return errors.New("divide by zero is not easy.")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func (t *Diavlos) NewMessage( m *Message, reply *string) error {


	if strings.TrimSpace(strings.ToUpper(m.Msg)) == "EXIT" {
		*reply = "Good bye then. See you next time"
	}else{
		*reply = "Your message was in uppercase letters:\n" + strings.ToUpper(m.Msg)
	}

	return nil
}

