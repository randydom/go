package impl

import (
	"awesomeProject/FileProcessor"
	"os"
	"log"
	"bufio"
	"unicode"
	"strings"
)

type(
	BasicFileProcessor struct{ThreadId int}
)

func NewBasicFileProcessor(threadId int) FileProcessor.ProcessFromPath{

	return &BasicFileProcessor{threadId}
}

func (b *BasicFileProcessor) FromFile(file string, out FileProcessor.OutChannel) {

	r := createNewReport(file, b.ThreadId)

	content, err := os.Open(file) // The “test_case.txt” is a sample text to test the code.
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(content)

	// First of all scan for paragraphs
	// A paragraph is added when a new line is read by the scanner
	/*
		Before add a paragraph the program will format each paragraph. Then:
		1. if newline && not len()==0 then add a new paragraph.

		The second step is to break the paragraph into fields of words. Then, each word will be examined if contains an ender delimeter
		to add a sentence.

		The last step is to trim any symbol from each word and add the word to the vocabulary, as well as count the letters of each word.
	 */
	for scanner.Scan(){
		line := scanner.Text()
		//fmt.Printf("Raw Line: %s\n", line)
		if len(line)>0{

			//line = formatText(line) // This method add and delete a space where is needed

			r.P.ListParagraphs = append(r.P.ListParagraphs, line)
			r.P.NumParagraph ++

			f := func(c rune) bool{
				return unicode.IsSpace(c)
			}
			subFields := strings.FieldsFunc(line,f) // built in function to break a string in whitespace
			tempSentence := "" // the tempSentence will concatenate fields of words until it meets a sentence stop.
			post := ""
			for i, sf := range subFields{
				tempSentence = tempSentence + " " + sf

				if i == (len(subFields) -1){
					post = "EOF"
				}else{
					post = subFields[i+1]
				}

				if isNewSentence(sf, post){ // The method findSentence returns a boolean to either the field is the last word of a sentence or not
					r.S.NumSentence++
					r.S.ListSentence = append(r.S.ListSentence, tempSentence) // Then the sentence is appended to the struct list

					tempSentence = ""
				}

				r.FindWord(sf)

			}
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	content.Close()

	out <- &r

}


