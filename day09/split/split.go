package split

import "strings"

//切割字符串
// example:
// abc,b => [a c]
func Split(str string, sep string) []string {
	// b := strings.Split(str, sep)
	// return b
	var realStr []string
	index := strings.Index(str, sep)
	for index > 0 {
		realStr = append(realStr, str[:index])
		str = str[index+1:]
		index = strings.Index(str, sep)
	}
	realStr = append(realStr, str)
	return realStr
}
