package main

import (
	"log"
	"bufio"
	"os"
	"fmt"
	"strings"
	"unicode"
	"io/ioutil"
		"time"
)

type document interface {
	Count() string
}

type letter struct{
	numLetter int
	numSymbol int
}

type word struct{
	numWord int
	vocabulary []string
}

type sentence struct{
	numSentence int
	listSentence []string
}

type paragraph struct{
	numParagraph int
	listParagraphs []string
}

type report struct{
	l letter
	w word
	s sentence
	p paragraph
	fileName string
}
// The main program will read the test case text and create the report
func main() {

	dir :=  "././Shakespere/" //"./test/"

	files, err := ioutil.ReadDir(dir)
	if err != nil{
		log.Fatal(err)
		fmt.Println(err)
	}

	reportList := make(map[string]report)

	/*
			Start Countdown
	 */

	 start := time.Now()

	/*
			Open a channel
	 */

	 rep := make(chan os.File)
	 con := make(chan os.File)

	 go func() {
	 	for _, file := range files{
			content, err := os.Open(dir+file.Name()) // The “test_case.txt” is a sample text to test the code.
			if err != nil {
				log.Fatal(err)
			}
			rep <- *content
		}
	 }()

	 go func() {

	 	for c := range rep{
	 		con <- c

			scanner := bufio.NewScanner(con)
		}

	 }()

	for _, file := range files{
		/****************/
		// Read the sample test file
		content, err := os.Open(dir+file.Name()) // The “test_case.txt” is a sample text to test the code.
		if err != nil {
			log.Fatal(err)
		}
		//defer content.Close()

		scanner := bufio.NewScanner(content)

		//p := paragraph{numParagraph:0,listParagraphs:[]string{},}
		//s := sentence{numSentence:0,listSentence:[]string{},}
		//w := word{numWord:0, vocabulary: []string{},}
		//l := letter{numLetter:0, numSymbol:0,}

		r := report{
			// initialisation of struct variables
			p:paragraph{numParagraph:0,listParagraphs:[]string{},},
			s:sentence{numSentence:0,listSentence:[]string{},},
			w:word{numWord:0, vocabulary: []string{},},
			l:letter{numLetter:0, numSymbol:0,},
			fileName: file.Name(),
		}

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

				r.p.listParagraphs = append(r.p.listParagraphs, line)
				r.p.numParagraph ++

				f := func(c rune) bool{
					return unicode.IsSpace(c)
				}
				subFields := strings.FieldsFunc(line,f) // built in function to break a string in whitespace
				tempSentence := "" // the tempSentence will concatenate fields of words until it meets a sentence stop.
				for _, sf := range subFields{
					tempSentence = tempSentence + " " + sf

					if findSentence(sf) || sf==subFields[len(subFields)-1]{ // The method findSentence returns a boolean to either the field is the last word of a sentence or not
						r.s.numSentence++
						r.s.listSentence = append(r.s.listSentence, tempSentence) // Then the sentence is appended to the struct list

						tempSentence = ""
					}

					r.findWord(sf)

				}
			}



		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		reportList[r.fileName] = r

		fmt.Println(r.fileName)
		printReport(r.p)
		printReport(r.s)
		printReport(r.w)
		printReport(r.l)

		//for _, sts := range s.listSentence{
		//	fmt.Println(sts)
		//}
		//fmt.Println("\nThe original text is:")
		//for _, par := range reportList[file.Name()].p.listParagraphs{
		//	fmt.Println(par)
		//}
		content.Close()
	}

	end := time.Now()
	elapsed := end.Sub(start)

	fmt.Printf("The concrete program took: %s\n", elapsed.String())
	fmt.Printf("It started: %s\n It ended: %s\n", start.String(), end.String())

}

func printReport(d document) {
	fmt.Println(d.Count())
}

func (s sentence) Count() string{
	return fmt.Sprintf("The number of sentences are: %d\n", s.numSentence)
}

func (w word) Count() string{
	return fmt.Sprintf("The number of words are: %d\n", w.numWord)
}

func (l letter) Count() string{
	return fmt.Sprintf("The number of letters & numbers are: %d\n", l.numLetter)
}

func (p paragraph) Count() string{
	return fmt.Sprintf("The number of paragraphs are: %d\n", p.numParagraph)
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

func findWordLetter(w string) string{

	var str string

	for _, tok := range w{
		if unicode.IsLetter(tok) || unicode.IsNumber(tok){
			str += string(tok)
		}
	}

	return str
}

func checkCombineWord(w string) bool {
	if strings.Contains(w, "-"){
		return true
	}
	return false
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


func (r *report) findWord(sf string) {
	tmpWord := findWordLetter(sf) // This method cleans the field word from any symbol, in order to count the letters and add the word to the vocabulary
	r.addWord(tmpWord, sf)
}

func (r *report) addWord(tmpWord string, sf string){
	r.w.vocabulary = append(r.w.vocabulary, tmpWord)

	if len(tmpWord)>0{
		r.w.numWord ++
	}
	if checkCombineWord(tmpWord){ // The method checkCombineWord returns a boolean if the field word is a combined word with dash
		r.w.numWord ++
	}
	r.l.numLetter += len(tmpWord)
	r.l.numSymbol += len(strings.TrimSpace(sf))

	if strings.Contains(sf,`“`) || strings.Contains(sf,`”`){
		r.l.numSymbol -= 2
	}
}


