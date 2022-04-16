package reverse

import (
	"strings"
)

func Reverse(input string) string {
	var tempWords []string
	var newWords []string
	for _, val := range input{
		tempWords = append(tempWords,string(val))
		newWords = append(newWords, "")
	}
	stringLength := len(tempWords)
	// fmt.Println(tempWords)
	for in, val1 := range tempWords{
		newWords[stringLength - in - 1] = val1
		// fmt.Println(newWords)
	}
	return strings.Join(newWords, "")
}