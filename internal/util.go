package internal

import (
	"os"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFileToLines(path string) (ret []string) {
	data, err := os.ReadFile(path)
	Check(err)
	for _, line := range strings.Split(string(data), "\n") {
		if line != "" {
			ret = append(ret, line)
		}
	}
	return ret
}

func ToSudoSet(items []string) (ret map[string]bool) {
	ret = make(map[string]bool)
	for _, item := range items {
		ret[item] = true
	}
	return ret
}

type IntStack []int

func (s *IntStack) Push(v int) {
	*s = append(*s, v)
}

func (s *IntStack) Pop() int {
	ret := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return ret
}
