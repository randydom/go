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
	//ThreadedLineProcessor struct{fileName string; threadId int}
)

var(
	outLineChan = make(FileProcessor.OutLineChannel,10)
	chanReportList = []*FileProcessor.Report{nil}
)

func NewThreadedFileProcessor(threadId int) FileProcessor.ProcessFromPath{

	return &ThreadedFileProcessor{threadId}
}

//func NewThreadedLineProcessor(file string, thrId int) FileProcessor.ProcessFromLine{
//
//return &ThreadedLineProcessor{fileName: file, threadId: thrId,}
//}

func (b *ThreadedFileProcessor) FromFile(file string, out FileProcessor.OutChannel){
	  /*********************************/
	 /*********** 1st Stage ***********/
	/*********************************/
	r := createNewReport(file, b.ThreadId)

	content, err := os.Open(file) // The “test_case.txt” is a sample text to test the code.
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(content)

	  /********************************/
	 /********** 2nd Stage ***********/
	/********************************/
	go b.processLine(scanner, &r)

	  /********************************/
	 /********** 3rd Stage ***********/
	/********************************/
	for output := range outLineChan{

		r.S.NumSentence += output.S.NumSentence
		r.W.NumWord += output.W.NumWord
		r.L.NumLetter += output.L.NumLetter

	}

	content.Close()

	out <- &r
}

func (b *ThreadedFileProcessor) processLine(scanner *bufio.Scanner, r *FileProcessor.Report) {

	wg := sync.WaitGroup{}
	id := 0
	for scanner.Scan() {
		wg.Add(1)
		id++
		ln := scanner.Text()

		if len(ln)>0{

			r.P.ListParagraphs = append(r.P.ListParagraphs, ln)
			r.P.NumParagraph ++

			go func(line string, index int, fn string) {
				defer func() {
					wg.Done()
				}()

				NewThreadedLineProcessor(fn, index).FromLine(line, outLineChan)

			}(ln, id, r.FileName)

		}

	}

	wg.Wait()
	close(outLineChan)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

