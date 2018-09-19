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

	r := &FileProcessor.Report{
		P:FileProcessor.Paragraph{NumParagraph:0,ListParagraphs:[]string{}},
		S:FileProcessor.Sentence{NumSentence:0, ListSentence:[]string{}},
		W:FileProcessor.Word{NumWord:0, Vocabulary:[]string{}},
		L:FileProcessor.Letter{NumLetter:0, NumSymbol:0},
		FileName: file,
		ThreadId: b.ThreadId,
	}

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

			line = formatText(line) // This method add and delete a space where is needed

			r.P.ListParagraphs = append(r.P.ListParagraphs, line)
			r.P.NumParagraph ++

			f := func(c rune) bool{
				return unicode.IsSpace(c)
			}
			subFields := strings.FieldsFunc(line,f) // built in function to break a string in whitespace
			tempSentence := "" // the tempSentence will concatenate fields of words until it meets a sentence stop.
			for _, sf := range subFields{
				tempSentence = tempSentence + " " + sf

				if findSentence(sf) || sf==subFields[len(subFields)-1]{ // The method findSentence returns a boolean to either the field is the last word of a sentence or not
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

	out <- r

}


func formatText(line string) string{

	if strings.Contains(line, "http"){
		return line
	}

	spaceMap := map[string]string{
		".(" : ". (", ".[" : ". [", "?(" : "? (", "?[" : "? [",
		"!(" : "! (", "![" : "! [", "," : ", ", "." : ". ", ":" : ": ",
		"?" : "? ", "!" : "! ", ";" : "; ", ")": ") ", "]": "] ",
	}

	replaceMap := map[string]string{
		`. )` : `.)`, `! )` : `!)`, `? )` : `?)`,
		`. ]` : `.]`, `! ]` : `!]`, `? ]` : `?]`,
		`. "` : `."`, `! ”` : `!”`, `? "` : `?"`,
		`. . .` : `...`,`. .` : `..`,`) ,` : `),`,
		`] ,` : `],`,`" ,` : `",`,`” ,` : `”,`,
	}

	for key,value := range spaceMap{
		line = strings.Replace(line, key, value, -1)
	}
	for key, value := range replaceMap{
		line = strings.Replace(line, key,  value, -1)
	}
	return line
}

// This function will determine the number of sentences in a line
func findSentence(l string) bool{

	sents := []string{`.`,`?`,`!`,`:`,`…`}

	if strings.HasPrefix(l, "http"){
		return false
	}else{
		for _,sen := range sents{
			if strings.Contains(l, sen){
				if checkIfEndingIsCorrect(l){ // The name of the method is obvious. The method is right below.
					return true
				}
			}
		}
	}
	return false
}

func checkIfEndingIsCorrect(l string) bool{
	enders := []string{`."`,`.'`,`.)`,`.]`,`?'`,`?"`,`?)`,`!"`,`!'`,`!)`,`,`}

	for _, end := range enders{
		if strings.Contains(l,end){
			return false
		}
	}
	return true
}




