package diffsquares

import "math"

func SquareOfSum(n int) int {
	var counter int
	for i := 1; i <= n; i++{
		counter += i
	}
	return counter * counter
}

func SumOfSquares(n int) int {
	var counter int
	for i:= 1; i<= n ; i++{
		counter += i*i
	}
	return counter
}

func Difference(n int) int {
	return int(math.Abs(float64(SquareOfSum(n))-float64(SumOfSquares(n))))
}
