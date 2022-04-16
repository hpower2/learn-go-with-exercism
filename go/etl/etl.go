package etl

import (
	"fmt"
	"strings"
)

func Transform(in map[int][]string) (out map[string]int) {
	var alpha []string
	var poin []int
	out = make(map[string]int)
	for key, val := range in{
		for _, alphabet := range val{
			out[strings.ToLower(alphabet)] = key
			// alpha = append(alpha, alphabet)
			// poin = append(poin, key)
		}
	}
	
	fmt.Println(alpha, poin)
	return out
}
