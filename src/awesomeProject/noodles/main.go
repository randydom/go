package main

import (
	"awesomeProject/FileProcessor/impl"
	"awesomeProject/FileProcessor"
	"path/filepath"
	"os"
	"sync"
	"fmt"
	"strings"
	"time"
)

var (
	outChan = make(FileProcessor.OutChannel, 10)
	fileMap []string
	now = time.Now()
)
func main() {


	buildFileMap()

	go processFiles()

	for output := range outChan{

		fmt.Printf("\n*****\nThread Id: %d text: %s\nNumber of:\nparagraphs: %d\nsentences: %d\nword: %d\nletters: %d\n", output.ThreadId,
			output.FileName, output.P.NumParagraph ,output.S.NumSentence, output.W.NumWord, output.L.NumLetter)
	}
	fmt.Println("\nThe time was:")
	fmt.Println(time.Now().Sub(now).Seconds())
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