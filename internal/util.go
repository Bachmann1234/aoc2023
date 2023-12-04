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
