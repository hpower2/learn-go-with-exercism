package main

import (
	"fmt"
	"math"
	// "strings"
)




func main() {
	var newVal []int
	state := false
	stateVal := 0
    for in, val := range s{
        if val >= 48 && val <= 57{
			newVal = append(newVal, int(val-48))
			stateVal = in
		}
		if val == 45 && stateVal == 0{
			state = math.Max()
		}
    }
	l := len(newVal)-1
	myItoa := 0
	for in, val :=range newVal{
		myItoa = myItoa + val * int(math.Pow10(l-in))
	}
	if state == true{
		myItoa = -1 * myItoa
	}
}