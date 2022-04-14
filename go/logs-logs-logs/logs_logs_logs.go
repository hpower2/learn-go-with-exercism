package logs

import "strings"

// Application identifies the application emitting the given log.
func Application(log string) string {
	var tempVal string
	for _, val := range log{
		if val == 10071 {
			tempVal = "recommendation"
			break
		}
		if val == 128269{
			tempVal = "search"
			break 
		}
		if val == 9728{
			tempVal = "weather"
			break 
		}
	}
	if len(tempVal) > 0{
		return tempVal
	}else{
		return "default"
	}
	
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var newSen string
	newSen = log
	for _, val := range log{
		if val == oldRune{
			newSen = strings.Replace(log, string(oldRune), string(newRune), -1)
		}else{
			continue
		}
	}
	return newSen
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	var counter int
	for _, val := range log{
		if (val >=97 && val <= 122) || (val >= 65 && val <= 90){
			counter++
		}
	}
	if counter <= limit{
		return true
	}
	return false
}
