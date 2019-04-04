package connectdb

import "strings"

func GetStringInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.LastIndex(str, end)
	return str[s:e]
}
