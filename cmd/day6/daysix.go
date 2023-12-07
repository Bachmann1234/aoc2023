package main

import (
	"dev/mattbachmann/aoc2023/internal"
	"os"
	"strings"
)

type Race struct {
	distance int
	time     int
}

func howManyWaysToWinRace(race Race) int {
	count := 0
	winning := false
	for i := 1; i < race.time; i++ {
		newDistance := i * (race.time - i)
		if newDistance > race.distance {
			winning = true
			count++
		} else {
			if winning {
				break
			}
		}
	}
	return count
}

func partOne(times []string, distances []string) (ret int) {
	var races []Race
	for index, time := range times {
		races = append(races, Race{
			distance: internal.ParseInt(distances[index]),
			time:     internal.ParseInt(time),
		})
	}
	ret = 1
	for _, race := range races {
		ret *= howManyWaysToWinRace(race)
	}
	return ret
}

func partTwo(times []string, distances []string) (ret int) {
	race := Race{
		time:     internal.ParseInt(strings.Join(times, "")),
		distance: internal.ParseInt(strings.Join(distances, "")),
	}
	return howManyWaysToWinRace(race)

}

func main() {
	lines := internal.ReadFileToLines(os.Args[1])
	times := strings.Fields(lines[0])[1:]
	distances := strings.Fields(lines[1])[1:]
	println(partOne(times, distances))
	println(partTwo(times, distances))

}
