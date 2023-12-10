package main

import (
	"dev/mattbachmann/aoc2023/internal"
	"fmt"
	"math"
	"os"
	"strings"
)

type Sequence []int

func parseSequences() []Sequence {
	lines := internal.ReadFileToLines(os.Args[1])
	var sequences []Sequence
	for _, line := range lines {
		var sequence []int
		for _, value := range strings.Fields(line) {
			sequence = append(sequence, internal.ParseInt(value))
		}
		sequences = append(sequences, sequence)
	}
	return sequences
}

func isAllZero(sequence Sequence) bool {
	for _, value := range sequence {
		if value != 0 {
			return false
		}
	}
	return true
}

func sequenceFromDifferences(sequence Sequence) (ret Sequence) {
	for i := 1; i < len(sequence); i++ {
		ret = append(ret, sequence[i]-sequence[i-1])
	}
	return ret
}

func traceSequenceToZeroes(sequence Sequence) []Sequence {
	sequences := []Sequence{sequence}
	for {
		sequences = append(sequences, sequenceFromDifferences(sequences[len(sequences)-1]))
		if isAllZero(sequences[len(sequences)-1]) {
			break
		}
	}

	return sequences
}

func findNextValue(sequence Sequence) (ret int) {
	sequences := traceSequenceToZeroes(sequence)
	ret = math.MinInt
	for i := len(sequences) - 1; i >= 0; i-- {
		if ret == math.MinInt {
			ret = 0
		} else {
			ret = sequences[i][len(sequences[i])-1] + ret
		}
	}
	for _, seq := range sequences {
		var seqStrings []string
		for _, i := range seq {
			seqStrings = append(seqStrings, fmt.Sprintf("%d", i))
		}

	}
	return ret
}
func findPreviousValue(sequence Sequence) (ret int) {
	var reversedSequence Sequence
	for i := len(sequence) - 1; i >= 0; i-- {
		reversedSequence = append(reversedSequence, sequence[i])
	}
	return findNextValue(reversedSequence)

}

func partOne(sequences []Sequence) (ret int) {
	for _, sequence := range sequences {
		newValue := findNextValue(sequence)
		ret += newValue
	}
	return ret

}

func partTwo(sequences []Sequence) (ret int) {
	for _, sequence := range sequences {
		newValue := findPreviousValue(sequence)
		ret += newValue
	}
	return ret

}

func main() {
	sequences := parseSequences()

	println(partOne(sequences))
	println(partTwo(sequences))
}
