package isogram

import "strings"

func IsIsogram(word string) bool {
	// hashMap := make(map[string]int)
	var myCount int
	newWord := strings.ToLower(word)
	if word == ""{
		return true
	}
	for _, val := range newWord{
		myCount = strings.Count(newWord, string(val))
		if myCount > 1 && val >= 97 && val <= 122{
			return false
		}
	}
	return true
}