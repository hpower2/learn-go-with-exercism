package main

import (
	"fmt"
	"regexp"
	"strings"
)

func RemoveEndOfLineText(text string) string{
	var sl1 []string
	re := regexp.MustCompile(`(end-of-line)\d+`)
	sl1 = re.Split(text, -1)
	mySen := strings.Join(sl1 , " ")
    return mySen
}


func main() {
    var text string
	text = "[INF] end-of-line23033 Network Failure end-of-line27"
	fmt.Println(RemoveEndOfLineText(text))
}
