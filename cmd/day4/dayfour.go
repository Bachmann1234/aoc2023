package main

import (
	"dev/mattbachmann/aoc2023/internal"
	"math"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	id             int
	winningNumbers map[string]bool
	cardNumbers    map[string]bool
}

func readCards(lines []string) (ret []Card) {
	for _, line := range lines {
		lineParts := strings.Split(line, ": ")
		id := strings.Fields(lineParts[0])[1]
		idNum, err := strconv.Atoi(id)
		internal.Check(err)
		game := strings.Split(lineParts[1], " | ")
		winningNumbers := strings.Fields(game[0])
		cardNumbers := strings.Fields(game[1])
		ret = append(
			ret,
			Card{
				idNum,
				internal.ToSudoSet(winningNumbers),
				internal.ToSudoSet(cardNumbers),
			},
		)
	}
	return ret
}

func computeMatches(cards []Card) (ret map[int]int) {
	ret = make(map[int]int)
	for _, card := range cards {
		matches := 0
		for number := range card.cardNumbers {
			_, ok := card.winningNumbers[number]
			if ok {
				matches += 1
			}
		}
		ret[card.id] = matches
	}
	return ret
}

func sumWinningCards(cardMatches map[int]int) (ret int) {
	for _, matches := range cardMatches {
		if matches > 0 {
			ret += int(math.Pow(2, float64(matches-1)))
		}
	}
	return ret
}

func scoreCards(cardMatches map[int]int) (ret int) {
	var stack internal.IntStack
	for key := range cardMatches {
		stack.Push(key)
		ret += 1
	}
	for {
		if len(stack) == 0 {
			break
		}
		cardId := stack.Pop()
		matches, ok := cardMatches[cardId]
		if !ok {
			panic("there should always be a card")
		}
		for i := 0; i < matches; i++ {
			prizeCardId := cardId + i + 1
			_, ok := cardMatches[prizeCardId]
			if ok {
				stack.Push(prizeCardId)
				ret += 1
			}
		}
	}
	return ret
}

func main() {
	cards := readCards(internal.ReadFileToLines(os.Args[1]))
	matches := computeMatches(cards)
	println(sumWinningCards(matches))
	println(scoreCards(matches))
}
