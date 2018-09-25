package FileProcessor

import (
	"log"
	"net/url"
	"strings"
	"unicode"
)

type (

	 /*************************/
	/****** Structures *******/

	Letter struct{
		NumLetter int
		NumSymbol int
	}

	Word struct{
		NumWord int
		Vocabulary []string
	}

	Sentence struct{
		NumSentence int
		ListSentence []string
	}

	Paragraph struct{
		NumParagraph int
		ListParagraphs []string
	}

	Report struct{
		L Letter
		W Word
		S Sentence
		P Paragraph
		FileName string
		ThreadId int
	}

	FileCounter struct{
		Buff int
	}

	 /*************************/
	/*******  Channel ********/

	OutChannel chan *Report
	OutLineChannel chan *Report

	 /*************************/
	/****** Interfaces *******/

	ProcessFromPath interface {
		FromFile(file string, out OutChannel)
	}

	Document interface {
		Count(file string) *Report
	}

	ProcessFromLine interface {
		FromLine(line string, out OutLineChannel)
	}
)


func (r *Report) FindWord(sf string) {

	tmpWord := FindWordLetter(sf) // This method cleans the field word from any symbol, in order to count the letters and add the word to the vocabulary
	r.AddWord(tmpWord, sf)
}

func (r *Report) AddWord(tmpWord string, sf string){

	r.W.Vocabulary = append(r.W.Vocabulary, tmpWord)

	if len(tmpWord)>0{
		if strings.HasPrefix(tmpWord,"http"){

			u, err := url.Parse(sf)
			if err != nil {
				log.Fatal(err)
			}

			r.W.NumWord += len(u.Query())*2 + len(strings.Split(u.Host, ".")) + 1

		}else{
			r.W.NumWord ++
		}

	}
	if CheckCombineWord(sf){ // The method checkCombineWord returns a boolean if the field word is a combined word with dash
		r.W.NumWord ++
	}
	r.L.NumLetter += len(tmpWord)
	r.L.NumSymbol += len(strings.TrimSpace(sf))

}

func FindWordLetter(w string) string{

	var str string

	for _, tok := range w{
		if unicode.IsLetter(tok) || unicode.IsNumber(tok){
			str += string(tok)
		}
	}

	return str
}

func CheckCombineWord(w string) bool {
	if strings.Contains(w, "-"){
		return true
	}
	return false
}

