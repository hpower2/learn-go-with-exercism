package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Valid(id string) bool {
	var newVal int
	var newVal2 int
	var myCounter int
	var noSpaces string = strings.ReplaceAll(id, " ", "")
	var noZero string
	noZero = noSpaces
	if len(id) <= 1{
		re := regexp.MustCompile(`^0`)
		noZero = re.ReplaceAllString(noSpaces, "")
	}
	if len(id) >= 3{
		for in, val := range noZero{
			newVal, _ = strconv.Atoi(string(val))
			if in % 2 == 0{
				newVal2 = newVal * 2
				if newVal2 > 9{
					newVal2 = newVal2 - 9
				}
				myCounter += newVal2
			}else{
				myCounter += newVal
			}
		}
		fmt.Println(myCounter)
	}
	return myCounter % 10 == 0
}
func main() {
    var text string
	text = "095 245 88"
	fmt.Println(Valid(text))
}
