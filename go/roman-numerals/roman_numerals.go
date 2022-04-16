package romannumerals

import (
	"errors"
	"strings"
)

func converToInt(in int) []int{
	val := in
	var intVal []int
	for i:= 0; ; i++{
		if val >= 1000{
			intVal = append(intVal, 1000)
			val -= 1000
		}else if val >= 900 && val < 1000{
			intVal = append(intVal, 900)
			val -= 900
		}else if val >= 500 && val < 900{
			intVal = append(intVal, 500)
			val -= 500
		}else if val >= 400 && val < 500{
			intVal = append(intVal, 400)
			val -= 400
		}else if val >= 100 && val < 400{
			intVal = append(intVal, 100)
			val -= 100
		}else if val >= 90 && val < 100{
			intVal = append(intVal, 90)
			val -= 90
		}else if val >= 50 && val < 90{
			intVal = append(intVal, 50)
			val -= 50
		}else if val >= 40 && val < 50{
			intVal = append(intVal, 40)
			val -= 40
		}else if val >= 10 && val < 40{
			intVal = append(intVal, 10)
			val -= 10
		}else if val == 9{
			intVal = append(intVal, 9)
			val -= 9
		}else if val >= 5 && val < 9{
			intVal = append(intVal, 5)
			val -= 5
		}else if val == 4{
			intVal = append(intVal, 4)
			val -=4
		}else if val >= 1 && val < 4{
			intVal = append(intVal, 1)
			val -=1
		}
		if val == 0{
			break
		}
	}
	return intVal
}

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3000 {
		return "", errors.New("Cannot convert negative number / Number too high")
	}
	myRoman := map[int]string{
		1 : "I",
		4 : "IV",
		5 : "V",
		9 : "IX",
		10 : "X",
		40 : "XL",
		50 : "L",
		90 : "XC",
		100 : "C",
		400 : "CD",
		500 : "D",
		900 : "CM",
		1000 : "M",
	}
	myRomanInt := converToInt(input)
	var roman []string

	for _, romanInt := range myRomanInt{
		roman = append(roman, myRoman[romanInt]) 
	}
	// fmt.Println(roman)
	return strings.Join(roman, ""), nil
}
