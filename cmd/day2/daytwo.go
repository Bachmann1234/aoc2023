package main

import (
	"dev/mattbachmann/aoc2023/internal"
	"os"
	"strconv"
	"strings"
)

type TotalStones struct {
	red   int
	green int
	blue  int
}

type Round struct {
	red   int
	green int
	blue  int
}

type Game struct {
	id     int
	rounds []Round
}

func parseRounds(rounds []string) []Round {
	var resultRounds []Round
	for _, round := range rounds {
		blue := 0
		red := 0
		green := 0
		colors := strings.Split(round, ", ")
		for _, color := range colors {
			colorParts := strings.Split(color, " ")
			number, err := strconv.Atoi(colorParts[0])
			internal.Check(err)
			if colorParts[1] == "green" {
				green = number
			}
			if colorParts[1] == "red" {
				red = number
			}
			if colorParts[1] == "blue" {
				blue = number
			}
		}
		resultRounds = append(resultRounds, Round{
			red:   red,
			green: green,
			blue:  blue,
		})
	}
	return resultRounds
}

func parseGame(line string) Game {
	lineParts := strings.Split(line, ": ")
	gameId, err := strconv.Atoi(strings.Split(lineParts[0], " ")[1])
	internal.Check(err)
	return Game{
		gameId,
		parseRounds(strings.Split(lineParts[1], "; ")),
	}
}

func gameIsPossible(totalStones TotalStones, game Game) bool {
	for _, round := range game.rounds {
		if round.blue > totalStones.blue || round.red > totalStones.red || round.green > totalStones.green {
			return false
		}
	}
	return true
}

func sumIdsPossibleGames(totalStones TotalStones, games []Game) int {
	sum := 0
	for _, game := range games {
		if gameIsPossible(totalStones, game) {
			sum += game.id
		}
	}
	return sum
}

func powerMaxes(games []Game) (ret int) {
	ret = 0
	for _, game := range games {
		mins := TotalStones{0, 0, 0}
		for _, round := range game.rounds {
			if round.red > mins.red {
				mins.red = round.red
			}
			if round.green > mins.green {
				mins.green = round.green
			}
			if round.blue > mins.blue {
				mins.blue = round.blue
			}
		}
		ret += mins.red * mins.green * mins.blue
	}
	return ret
}

func main() {
	lines := internal.ReadFileToLines(os.Args[1])
	var games []Game
	for _, line := range lines {
		games = append(games, parseGame(line))
	}
	println(sumIdsPossibleGames(TotalStones{12, 13, 14}, games))
	println(powerMaxes(games))
}
