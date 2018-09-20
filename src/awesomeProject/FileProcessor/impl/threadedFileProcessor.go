package impl

import (
	"awesomeProject/FileProcessor"
	"sync"
	"os"
	"log"
	"bufio"
)

type (
	ThreadedFileProcessor struct{ThreadId int}
	ThreadedLineProcessor struct{threadId int}
)

var(
	outLineChan = make(FileProcessor.OutLineChannel,50)
	chanReportList = []FileProcessor.OutLineChannel{nil}
)

func NewThreadedFileProcessor(threadId int) FileProcessor.ProcessFromPath{

	return &ThreadedFileProcessor{threadId}
}

func NewThreadedLineProcessor(thrId int) FileProcessor.ProcessFromLine{

	return &ThreadedLineProcessor{thrId}
}

func (b *ThreadedFileProcessor) FromFile(file string, out FileProcessor.OutChannel){

	r := createNewReport(file, b.ThreadId)

	content, err := os.Open(file) // The “test_case.txt” is a sample text to test the code.
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(content)

	go b.processLine(scanner)

	for output := range outLineChan{
		chanReportList = append(chanReportList, output)
	}

}

func (b *ThreadedFileProcessor) processLine(scanner *bufio.Scanner) {

	wg := sync.WaitGroup{}
	id := 0
	for scanner.Scan() {
		wg.Add(1)
		id++
		ln := scanner.Text()

		go func(line string, index int) {
			defer func() {
				wg.Done()
			}()

			NewThreadedLineProcessor(index).FromLine(line, outLineChan)

		}(ln, id)
	}

	wg.Wait()
	close(outLineChan)

}

func (t *ThreadedLineProcessor) FromLine(line string, out FileProcessor.OutLineChannel){



	//TODO complete the method FromLine
}