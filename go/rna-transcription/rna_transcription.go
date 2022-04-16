package strand

import "strings"



func ToRNA(dna string) string {
	myMap := map[string]string{
		"G" : "C",
		"C" : "G",
		"T" : "A",
		"A" : "U",
	}
	var emptyArr []string 
	for _, val := range dna{
		for key, elements := range myMap{
			if string(val) == key{
				emptyArr = append(emptyArr, elements)
			}
		}
	}
	myRna := strings.Join(emptyArr, "")
	return myRna
}
