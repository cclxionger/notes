package split

import (
	"strings"
)

func Split(s, sep string) (ret []string) {
	index := strings.Index(s, sep)
	for index > -1 {
		ret = append(ret, s[:index])
		s = s[index+1:]
		index = strings.Index(s, sep)
	}
	ret = append(ret, s)
	return
}
