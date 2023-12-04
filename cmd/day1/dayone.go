package main

import (
	"dev/mattbachmann/aoc2023/internal"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func findDigit(line string) int {
	var result string
	for i := 0; i < len(line); i++ {
		c := line[i]
		if unicode.IsDigit(rune(c)) {
			result = string(c)
			break
		}
	}
	for i := len(line) - 1; i >= 0; i-- {
		c := line[i]
		if unicode.IsDigit(rune(c)) {
			result += string(c)
			break
		}
	}
	numericResult, err := strconv.Atoi(result)
	internal.Check(err)
	if err != nil {
		println("warning, returning 0")
		return 0
	}
	return numericResult
}

func main() {
	lines := internal.ReadFileToLines(os.Args[1])
	replacer := strings.NewReplacer(
		"one", "o1e",
		"two", "t2o",
		"three", "t3e",
		"four", "f4r",
		"five", "f5e",
		"six", "s6x",
		"seven", "s7n",
		"eight", "e8t",
		"nine", "n9e",
	)
	partOneTotal := 0
	partTwoTotal := 0
	for _, line := range lines {
		partOneTotal += findDigit(line)
		replacedLine := line
		originalLine := line
		for {
			replacedLine = replacer.Replace(originalLine)
			if replacedLine == originalLine {
				break
			} else {
				originalLine = replacedLine
			}
		}
		partTwoTotal += findDigit(replacedLine)
	}

	println(fmt.Sprintf("Part One Total: %d", partOneTotal))
	println(fmt.Sprintf("Part Two Total: %d", partTwoTotal))
}
