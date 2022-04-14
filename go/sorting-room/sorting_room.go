package sorting

import (
	"fmt"
	"math"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	round := math.Round(f*10.0)/10.0
	return fmt.Sprintf("This is the number %.1f", round)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	switch fnb.(type) {
	case FancyNumber:
		i, _ := strconv.Atoi(fnb.Value())
		return i
	default:
		return 0
	}
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	myNum := ExtractFancyNumber(fnb)
	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(myNum))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch val := i.(type){
	case float64:
		return DescribeNumber(val)
	case int:
		return DescribeNumber(float64(val))
	case NumberBox:
		return DescribeNumberBox(val)
	case FancyNumberBox:
		return DescribeFancyNumberBox(val)
	default:
		return "Return to sender"
	}
}
