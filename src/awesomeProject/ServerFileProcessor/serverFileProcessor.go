package SFProcessor

import (
	"awesomeProject/FileProcessor"
	"awesomeProject/FileProcessor/impl"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	fileMap []string
	outChan = make(FileProcessor.OutChannel, 10)
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
		return errors.New("divide by zero is not easy")
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

func (t *Diavlos) FileProcessor(m *Message, reply *[]*FileProcessor.Report) error {

	outChan = make(FileProcessor.OutChannel, 10)
	var reportList []*FileProcessor.Report

	buildFileMap()

	go processFiles()

	fmt.Println("Print the output channels")

	for output := range outChan {
		reportList = append(reportList,output)

		fmt.Printf("\n*****\nThread Id: %d text: %s\nNumber of:\nparagraphs: %d\nsentences: %d\nword: %d\nletters: %d\n", output.ThreadId,
			output.FileName, output.P.NumParagraph ,output.S.NumSentence, output.W.NumWord, output.L.NumLetter)
	}

	*reply = reportList

	return nil
}

func processFiles() {
	wg := sync.WaitGroup{}

	for i, fp := range fileMap {
		wg.Add(1)

		go func(tfp string, index int) {
			defer func () {
				wg.Done()
			}()

			impl.NewBasicFileProcessor(index).FromFile(tfp, outChan)

		}(fp,i)

	}

	wg.Wait()
	close(outChan)
}


func buildFileMap(){


	filepath.Walk("./Shakespere/", func(path string, info os.FileInfo, err error) error {

		if strings.HasSuffix(path, ".txt"){
			fileMap = append(fileMap, path)
		}

		return nil
	})



}