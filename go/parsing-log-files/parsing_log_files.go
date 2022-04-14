package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
    var sl []string
    re := regexp.MustCompile(`<[~=*-]*>`)
    sl = re.Split(text, -1)
    return sl
}

func CountQuotedPasswords(lines []string) int {
    var counter int
	for _, val := range lines{
		re := regexp.MustCompile(`(?i)".*password.*"`)
        match := re.MatchString(val)
        if match{
            counter++
        }
	}
    return counter
}

func RemoveEndOfLineText(text string) string{
	re := regexp.MustCompile(`end-of-line\d*`)
    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re1 := regexp.MustCompile(`User\s+(\w+)`)
	ret := []string{}
	for _, val := range lines{
		founds := re1.FindStringSubmatch(val)
		if founds != nil{
			val = fmt.Sprintf("[USR] %s %s", founds[1], val)
		}
		ret = append(ret, val)
	}
	return ret
}
