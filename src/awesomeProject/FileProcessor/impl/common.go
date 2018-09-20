package impl

import "strings"

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
