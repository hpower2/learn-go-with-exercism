package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	myString := strconv.Itoa(n)
	var arrayNum []int
	for _, val := range myString{
		myNumber := int(val) - 48
		arrayNum = append(arrayNum, myNumber)
	}
	numDigits := len(arrayNum)
	var tempSum float64
	for _, val := range arrayNum{
		tempSum += math.Pow(float64(val), float64(numDigits))
	}
	if int(tempSum) == n{
		return true
	}else{
		return false
	}
}
