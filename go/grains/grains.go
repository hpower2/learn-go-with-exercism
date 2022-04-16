package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number > 0 && number <= 64{
		myVal := math.Pow(2, float64(number-1))
		return uint64(myVal), nil
	}else{
		return 0, errors.New("Wrong input")
	}
}

func Total() uint64 {
	total, _ := Square(65)
	return total - 1
}
