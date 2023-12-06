package main

import (
	"dev/mattbachmann/aoc2023/internal"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type MapEntry struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}

type SeedRange struct {
	start  int
	length int
}
type Almanac struct {
	seeds                 []int
	seedRanges            []SeedRange
	seedToSoil            []MapEntry
	soilToFertilizer      []MapEntry
	fertilizerToWater     []MapEntry
	waterToLight          []MapEntry
	lightToTemperature    []MapEntry
	temperatureToHumidity []MapEntry
	humidityToLocation    []MapEntry
}

func readSeeds(line string) (ret []int) {
	for _, field := range strings.Fields(line) {
		if field != "seeds:" {
			val, err := strconv.Atoi(field)
			internal.Check(err)
			ret = append(ret, val)
		}
	}
	return ret
}

func readSeedRanges(line string) (ret []SeedRange) {
	start := -1
	for _, field := range strings.Fields(line) {
		if field != "seeds:" {
			val, err := strconv.Atoi(field)
			internal.Check(err)
			if start == -1 {
				start = val
			} else {
				ret = append(ret, SeedRange{
					start:  start,
					length: val,
				})
				start = -1

			}
		}
	}
	return ret
}

func readMapEntry(m string) MapEntry {
	parts := strings.Fields(m)
	return MapEntry{
		destinationRangeStart: internal.ParseInt(parts[0]),
		sourceRangeStart:      internal.ParseInt(parts[1]),
		rangeLength:           internal.ParseInt(parts[2]),
	}
}

func readMap(m string) (ret []MapEntry) {
	parts := strings.Split(m, "\n")
	for i := 1; i < len(parts); i++ {
		ret = append(ret, readMapEntry(parts[i]))
	}
	return ret
}

func readAlmanac() Almanac {
	lines := internal.SplitFile(os.Args[1], "\n\n")
	return Almanac{
		seeds:                 readSeeds(lines[0]),
		seedRanges:            readSeedRanges(lines[0]), // Yes, same line
		seedToSoil:            readMap(lines[1]),
		soilToFertilizer:      readMap(lines[2]),
		fertilizerToWater:     readMap(lines[3]),
		waterToLight:          readMap(lines[4]),
		lightToTemperature:    readMap(lines[5]),
		temperatureToHumidity: readMap(lines[6]),
		humidityToLocation:    readMap(lines[7]),
	}
}

func mapValue(value int, mapEntries []MapEntry) int {
	for _, entry := range mapEntries {
		if value >= entry.sourceRangeStart && value <= entry.sourceRangeStart+entry.rangeLength {
			index := value - entry.sourceRangeStart
			return entry.destinationRangeStart + index
		}
	}
	return value

}

func mapKey(key int, mapEntries []MapEntry) int {
	for _, entry := range mapEntries {
		if key >= entry.destinationRangeStart && key <= entry.destinationRangeStart+entry.rangeLength {
			index := key - entry.destinationRangeStart
			return entry.sourceRangeStart + index
		}
	}
	return key
}

func seedToLocationRangeFast(almanac Almanac) int {
	sort.Slice(almanac.humidityToLocation, func(i, j int) bool {
		return almanac.humidityToLocation[i].destinationRangeStart < almanac.humidityToLocation[j].destinationRangeStart
	})

	for _, entry := range almanac.humidityToLocation {
		for i := 0; i < entry.rangeLength; i++ {
			location := entry.destinationRangeStart + i
			humid := mapKey(location, almanac.humidityToLocation)
			temperature := mapKey(humid, almanac.temperatureToHumidity)
			light := mapKey(temperature, almanac.lightToTemperature)
			water := mapKey(light, almanac.waterToLight)
			fertilizer := mapKey(water, almanac.fertilizerToWater)
			soil := mapKey(fertilizer, almanac.soilToFertilizer)
			seed := mapKey(soil, almanac.seedToSoil)
			for _, seedRange := range almanac.seedRanges {
				if seed >= seedRange.start && seed <= seedRange.start+seedRange.length {
					return location
				}
			}
		}
	}
	return -1
}

func seedToLocationRange(almanac Almanac) (ret int) {
	ret = math.MaxInt
	values := make(chan []int, len(almanac.seedRanges))
	for _, seedRange := range almanac.seedRanges {
		seedRange := seedRange
		go func() {
			var locations []int
			for i := 0; i < seedRange.length; i++ {
				soil := mapValue(i+seedRange.start, almanac.seedToSoil)
				fertilizer := mapValue(soil, almanac.soilToFertilizer)
				water := mapValue(fertilizer, almanac.fertilizerToWater)
				light := mapValue(water, almanac.waterToLight)
				temp := mapValue(light, almanac.lightToTemperature)
				humid := mapValue(temp, almanac.temperatureToHumidity)
				locations = append(locations, mapValue(humid, almanac.humidityToLocation))
			}
			values <- locations
		}()

	}
	for i := 0; i < len(almanac.seedRanges); i++ {
		for _, v := range <-values {
			if v < ret {
				ret = v
			}
		}
	}
	return ret
}

func seedToLocation(almanac Almanac) (ret int) {
	ret = math.MaxInt
	for _, seed := range almanac.seeds {
		soil := mapValue(seed, almanac.seedToSoil)
		fertilizer := mapValue(soil, almanac.soilToFertilizer)
		water := mapValue(fertilizer, almanac.fertilizerToWater)
		light := mapValue(water, almanac.waterToLight)
		temp := mapValue(light, almanac.lightToTemperature)
		humid := mapValue(temp, almanac.temperatureToHumidity)
		location := mapValue(humid, almanac.humidityToLocation)
		if location < ret {
			ret = location
		}
	}
	return ret
}

func main() {
	almanac := readAlmanac()
	println(seedToLocation(almanac))
	println(seedToLocationRangeFast(almanac))
	//println(seedToLocationRange(almanac))
}
