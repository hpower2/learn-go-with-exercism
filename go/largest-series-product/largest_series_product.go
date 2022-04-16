package lsproduct

import (
	"errors"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	mymap := make(map[int]int)
	var currentLarge int
	var prevLarge int

	if len(digits) < span{
		return 0, errors.New("Cannot parse the func because digits fewer than span")
	}else if span < 0 {
		return 0, errors.New("Span cannot be a negative number")
	}

	for in, val := range digits{
		if val < 48 || val > 57{
			return 0, errors.New("cannot be a character in the digits")
		}
		mymap[in] = int(val) - 48
	}
	for i := 0; i <= len(digits)-span; i++{
		prevLarge = currentLarge
		currentLarge = 1
		for j := i ; j < i + span; j++{
			currentLarge *= mymap[j]
		}
		if currentLarge < prevLarge{
			currentLarge = prevLarge
		}
	}
	return int64(currentLarge), nil
}