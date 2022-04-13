package hamming

import "errors"

func Distance(a, b string) (int, error) {
	var counter int

	if len(a) != len(b){
		return -1, errors.New("Please input the same length")
	}
	for in, _ := range a{
		if a[in] != b[in]{
			counter = counter + 1
		}
	}
	return counter, nil
}
