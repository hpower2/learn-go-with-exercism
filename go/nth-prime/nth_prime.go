package prime
import "math"
func isPrime(prime int) bool {
	if prime%2 == 0 {
		return prime == 2
	}
	m := int(math.Sqrt(float64(prime))) + 1
	for t := 3; t < m; t += 2 {
		if prime%t == 0 {
			return false
		}
	}
	return true
}
func Nth(n int) (int, bool) {
	if n <= 0 {
		return 0, false
	}
	for p := 2; ; p++ {
		if isPrime(p) {
			n--
			if n == 0 {
				return p, true
			}
		}
	}
}