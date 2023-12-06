package internal

import (
	"os"
	"strconv"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ParseInt(s string) int {
	ret, err := strconv.Atoi(s)
	Check(err)
	return ret
}

func SplitFile(path string, sep string) (ret []string) {
	data, err := os.ReadFile(path)
	Check(err)
	for _, line := range strings.Split(string(data), sep) {
		if line != "" {
			ret = append(ret, line)
		}
	}
	return ret
}

func ReadFileToLines(path string) []string {
	return SplitFile(path, "\n")
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
