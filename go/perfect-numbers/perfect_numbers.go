package perfect

import (
	"errors"
	"fmt"
)

// Define the Classification type here.
type Classification int

const (
	ClassificationPerfect Classification = iota
	ClassificationAbundant
	ClassificationDeficient
  )

var ErrOnlyPositive = errors.New("Only positive integers allowed")

func myFactor(n int64) []int64{
	var temp []int64
	var i int64
	for i = 1; i <= n; i++{
		if n % i == 0{
			if n/i == i {
				temp = append(temp, i)
			}else{
				temp = append(temp, n/i)
			}
		}
	}
	if len(temp) == 0{
		return temp
	}
	return temp[1:]
}

func Classify(n int64) (Classification, error) {
	var tempval []int64	
	var Value int64
	tempval = myFactor(n)
	for _, val := range tempval{
		Value = Value + val
	}
	fmt.Println(Value)
	if n < 0{
		return -1, ErrOnlyPositive
	}
	if n == 1{
		return ClassificationDeficient, nil
	}
	if n < 1{
		return ClassificationDeficient, ErrOnlyPositive
	}
	if Value > n {
		return ClassificationAbundant, nil
	}
	if Value < n{
		return ClassificationDeficient, nil
	}
	if Value == n{
		return ClassificationPerfect, nil
	}
	return ClassificationDeficient, ErrOnlyPositive
}
