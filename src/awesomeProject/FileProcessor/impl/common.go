package impl

import (
	"strings"
	"unicode"
	"awesomeProject/FileProcessor"
)

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

//// This function will determine the number of sentences in a line
//func findSentence(l string) bool{
//
//	sents := []string{`.`,`?`,`!`,`:`,`…`}
//
//	if strings.HasPrefix(l, "http"){
//		return false
//	}else{
//		for _,sen := range sents{
//			if strings.Contains(l, sen){
//				if checkIfEndingIsCorrect(l){ // The name of the method is obvious. The method is right below.
//					return true
//				}
//			}
//		}
//	}
//	return false
//}
//
//func checkIfEndingIsCorrect(l string) bool{
//	enders := []string{`."`,`.'`,`.)`,`.]`,`?'`,`?"`,`?)`,`!"`,`!'`,`!)`,`,`}
//
//	for _, end := range enders{
//		if strings.Contains(l,end){
//			return false
//		}
//	}
//	return true
//}

func isNewSentence(pre string, post string) bool{

	sents := []string{`.`,`?`,`!`,`:`,`…`}

	if strings.EqualFold(post, "EOF"){
		return true
	}else{
		for _,sen := range sents{
			if strings.HasSuffix(pre, sen){
				if checkWordIsUpper(post){
					return true
				}
			}
		}

	}

	return false
}

func checkWordIsUpper(word string) bool{

	for _, rn := range word{
		if unicode.IsUpper(rn){
			return true
		}else{
			return false
		}
	}
	return false
}

func createNewReport(name string, id int) FileProcessor.Report{
	return FileProcessor.Report{
		P:FileProcessor.Paragraph{NumParagraph:0,ListParagraphs:[]string{}},
		S:FileProcessor.Sentence{NumSentence:0, ListSentence:[]string{}},
		W:FileProcessor.Word{NumWord:0, Vocabulary:[]string{}},
		L:FileProcessor.Letter{NumLetter:0, NumSymbol:0},
		FileName: name,
		ThreadId: id,
	}
}

