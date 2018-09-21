package impl

import (
	"awesomeProject/FileProcessor"
	"unicode"
	"strings"
)

type (
	ThreadedLineProcessor struct{fileName string; threadId int}
)

func NewThreadedLineProcessor(file string, thrId int) FileProcessor.ProcessFromLine{

	return &ThreadedLineProcessor{fileName: file, threadId: thrId,}
}

func (t *ThreadedLineProcessor) FromLine(line string, out FileProcessor.OutLineChannel){

	localR := createNewReport(t.fileName, t.threadId)

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
			localR.S.NumSentence++
			localR.S.ListSentence = append(localR.S.ListSentence, tempSentence) // Then the sentence is appended to the struct list

			tempSentence = ""
		}

		localR.FindWord(sf)

	}

	out <- &localR

}