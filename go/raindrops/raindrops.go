package raindrops

import (
	"strconv"
	"strings"
)

func Convert(number int) string {
	var tempConvert []string
	if number % 3 == 0{
		tempConvert = append(tempConvert, "Pling")
	}
	if number % 5 == 0{
		tempConvert = append(tempConvert, "Plang")
	}
	if number % 7 == 0{
		tempConvert = append(tempConvert, "Plong")
	}
	if tempConvert == nil{
		return strconv.Itoa(number)
	}else{
		return strings.Join(tempConvert, "")
	}
}
