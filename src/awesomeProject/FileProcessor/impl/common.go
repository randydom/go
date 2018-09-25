package impl

import (
	"strings"
	"unicode"
	"awesomeProject/FileProcessor"
)

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

