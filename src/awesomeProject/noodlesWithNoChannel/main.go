package main

import (
	"path/filepath"
	"os"
	"strings"
	"time"
	"fmt"
	"awesomeProject/FileProcessor/impl"
	"awesomeProject/FileProcessor"
	)

var(
	fileMap []string
	reportList = []*FileProcessor.Report{}//= make(map[int]FileProcessor.Report)
	now = time.Now()
)


func main(){

	buildFileMap()

	processFiles()

	printReport()

	fmt.Println("\nThe time was:")
	fmt.Println(time.Now().Sub(now).Seconds())
}

func buildFileMap() {

	filepath.Walk("./Test/", func(path string, info os.FileInfo, err error) error { 	//"./Shakespere/"

		if strings.HasSuffix(path, ".txt") {
			fileMap = append(fileMap, path)
		}

		return nil
	})
}

func processFiles() {

	for index, fp := range fileMap {
		reportList = append(reportList, impl.NewConcreteFileProcessor(index).Count(fp))
	}

}


func printReport() {

	for _,rep := range reportList{
		fmt.Printf("\n*****\nThread Id: %d text: %s\nNumber of:\nparagraphs: %d\nsentences: %d\nword: %d\nletters: %d\n", rep.ThreadId,
			rep.FileName, rep.P.NumParagraph , rep.S.NumSentence, rep.W.NumWord, rep.L.NumLetter)
	}

}
