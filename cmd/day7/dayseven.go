package main

import (
	"dev/mattbachmann/aoc2023/internal"
	"os"
	"sort"
	"strings"
)

type Play struct {
	hand  string
	wager int
}

type Rank int

const FiveOfAKind = 7
const FourOfAKind = 6
const FullHouse = 5
const ThreeOfAKind = 4
const TwoPair = 3
const OnePair = 2
const HighCard = 1

const Joker = 'J'

func getCardRank(card rune, joker bool) int {
	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case Joker:
		if joker {
			return 1
		} else {
			return 11
		}
	case 'T':
		return 10
	case '9':
		return 9
	case '8':
		return 8
	case '7':
		return 7
	case '6':
		return 6
	case '5':
		return 5
	case '4':
		return 4
	case '3':
		return 3
	case '2':
		return 2
	default:
		panic("UNKNOWN CARD RANK")
	}
}

func cardsByLabels(hand string) map[rune]int {
	cardCounts := make(map[rune]int)
	for _, card := range []rune(hand) {
		count, ok := cardCounts[card]
		if ok {
			cardCounts[card] = count + 1
		} else {
			cardCounts[card] = 1
		}
	}
	return cardCounts
}

func computeMaxValueKey(cardsByLabel map[rune]int) int32 {
	maxValueKey := Joker
	for key, value := range cardsByLabel {
		if value >= cardsByLabel[maxValueKey] {
			maxValueKey = key
		}
	}
	return maxValueKey
}

func applyJoker(cardsByLabel map[rune]int) (map[rune]int, []int, Rank) {
	cardsByLabel[Joker] = cardsByLabel[Joker] - 1
	if cardsByLabel[Joker] == 0 {
		delete(cardsByLabel, Joker)
	}
	maxValueKey := computeMaxValueKey(cardsByLabel)
	cardsByLabel[maxValueKey] = cardsByLabel[maxValueKey] + 1
	counts := computeCounts(cardsByLabel)
	return cardsByLabel, counts, handRankWithoutJoker(cardsByLabel, counts)
}

func handRankWithoutJoker(cardsByLabel map[rune]int, counts []int) Rank {
	if len(cardsByLabel) == 1 {
		return FiveOfAKind
	} else if len(cardsByLabel) == 2 {
		if internal.IntSliceContains(counts, 4) {
			return FourOfAKind
		} else {
			return FullHouse
		}
	} else if len(cardsByLabel) == 3 {
		if internal.IntSliceContains(counts, 3) {
			return ThreeOfAKind
		} else {
			return TwoPair
		}
	} else if len(cardsByLabel) == 4 {
		return OnePair
	} else if len(cardsByLabel) == 5 {
		return HighCard
	}
	panic("UGH")
}

func computeCounts(cardsByLabel map[rune]int) []int {
	var counts []int
	for _, value := range cardsByLabel {
		counts = append(counts, value)
	}
	return counts
}

func computeHandRank(hand string, joker bool) (ret Rank) {
	cardsByLabel := cardsByLabels(hand)
	counts := computeCounts(cardsByLabel)
	jokers, ok := cardsByLabel[Joker]
	if !ok {
		jokers = 0
	}
	ret = handRankWithoutJoker(cardsByLabel, counts)
	var newRank Rank
	if joker {
		maxValueKey := computeMaxValueKey(cardsByLabel)
		if maxValueKey == Joker {
			switch ret {
			case FiveOfAKind:
				ret = FiveOfAKind
				break
			case FourOfAKind:
				ret = FiveOfAKind
				break
			case FullHouse:
				ret = FiveOfAKind
				break
			case ThreeOfAKind:
				ret = FourOfAKind
				break
			case TwoPair:
				ret = FourOfAKind
				break
			case OnePair:
				ret = ThreeOfAKind
			case HighCard:
				ret = OnePair

			}
		} else {
			for i := 0; i < jokers; i++ {
				cardsByLabel, counts, newRank = applyJoker(cardsByLabel)
				if newRank > ret {
					ret = newRank
				} else {
					break
				}
			}
		}

	}
	return ret
}

func parseHands() (ret []Play) {
	lines := internal.ReadFileToLines(os.Args[1])
	for _, line := range lines {
		parts := strings.Fields(line)
		ret = append(ret, Play{
			parts[0],
			internal.ParseInt(parts[1]),
		})
	}
	return ret
}

func computeWagers(plays []Play, joker bool) (ret int) {
	sort.Slice(plays, func(i, j int) bool {
		handOne := plays[i].hand
		handTwo := plays[j].hand
		handOneRank := computeHandRank(handOne, joker)
		handTwoRank := computeHandRank(handTwo, joker)
		if handOneRank == handTwoRank {
			cardsInHandOne := []rune(handOne)
			cardsInHandTwo := []rune(handTwo)
			for n := 0; n < len(cardsInHandOne); n++ {
				handOneCardRank := getCardRank(cardsInHandOne[n], joker)
				handTwoCardRank := getCardRank(cardsInHandTwo[n], joker)
				if handOneCardRank != handTwoCardRank {
					return handOneCardRank < handTwoCardRank
				}
			}
			return false
		} else {
			return handOneRank < handTwoRank
		}
	})

	for index, play := range plays {
		ret += play.wager * (index + 1)
	}
	return ret
}

func main() {
	hands := parseHands()
	println(computeWagers(hands, false))
	println(computeWagers(hands, true))
}
