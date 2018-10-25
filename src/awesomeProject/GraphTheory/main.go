package main

import (
	"awesomeProject/GraphTheory/Graph/impl"
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

var (
	filePath []string
)

func main(){

	buildFileMap()

	scanElements(filePath[0])

}
// helper
func buildFileMap() {


	filepath.Walk("./GraphTheory/", func(path string, info os.FileInfo, err error) error {

		if strings.HasSuffix(path, ".txt"){
			filePath = append(filePath, path)
		}


		return nil
	})
}
// helper
func scanText(file string) []string{

	var content []string

	con, err := os.Open(file)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(con)

	for scanner.Scan() {

		line := scanner.Text()

		content = append(content, line)

	}

	return content
}

//helper
func scanList(line string) []string{

	var content []string

	if len(line) > 0 {

		f := func(c rune) bool{
			return unicode.IsSpace(c)
		}

		subFields := strings.FieldsFunc(line,f)

		for _, el := range subFields{

			content = append(content, el)

		}
	}

	return content
}

func scanElements(file string) {

	line := scanText(file)

	for _,l := range line{

		el := scanList(l)

	}


}

func scanRelations(file string) {



}

