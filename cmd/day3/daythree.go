package main

import (
	"dev/mattbachmann/aoc2023/internal"
	"os"
	"strconv"
	"unicode"
)

type Element struct {
	value  string
	startX int
	endX   int
	y      int
}

type Schematic struct {
	potentialPartNumbers []Element
	partNumbers          []Element
	symbols              []Element
	gearRatios           []int
}

func parseSchematic(lines []string) Schematic {
	schematic := Schematic{
		potentialPartNumbers: []Element{},
		partNumbers:          []Element{},
		symbols:              []Element{},
		gearRatios:           []int{},
	}
	for y, line := range lines {
		parsedNumber := ""
		startX := -1
		for x, char := range line {
			if char == '.' {
				if parsedNumber != "" {
					schematic.potentialPartNumbers = append(schematic.potentialPartNumbers, Element{
						startX: startX,
						endX:   x - 1,
						y:      y,
						value:  parsedNumber,
					})
					startX = -1
					parsedNumber = ""
				}

			} else if unicode.IsDigit(char) {
				parsedNumber = parsedNumber + string(char)
				if startX == -1 {
					startX = x
				}
			} else {
				schematic.symbols = append(schematic.symbols, Element{
					startX: x,
					endX:   x - 1,
					y:      y,
					value:  string(char),
				})
				if parsedNumber != "" {
					schematic.potentialPartNumbers = append(schematic.potentialPartNumbers, Element{
						startX: startX,
						endX:   x - 1,
						y:      y,
						value:  parsedNumber,
					})
					startX = -1
					parsedNumber = ""
				}
			}
		}
		if parsedNumber != "" {
			schematic.potentialPartNumbers = append(schematic.potentialPartNumbers, Element{
				startX: startX,
				endX:   len(lines[len(lines)-1]) - 1,
				y:      y,
				value:  parsedNumber,
			})
		}
	}
	schematic.partNumbers = getPartNumbers(schematic)
	schematic.gearRatios = getGearRatios(schematic)
	return schematic
}

func getGearRatios(schematic Schematic) (ret []int) {
	for _, symbol := range schematic.symbols {
		if symbol.value == "*" {
			var adjacentParts []Element
			for _, part := range schematic.partNumbers {
				if isAdjacent(symbol, part) {
					adjacentParts = append(adjacentParts, part)
				}
			}
			if len(adjacentParts) == 2 {
				partOne, err := strconv.Atoi(adjacentParts[0].value)
				internal.Check(err)
				partTwo, errTwo := strconv.Atoi(adjacentParts[1].value)
				internal.Check(errTwo)
				ret = append(ret, partOne*partTwo)
			}
		}
	}
	return ret
}

func isAdjacent(element Element, potential Element) bool {
	if element.y == potential.y || element.y == (potential.y-1) || element.y == (potential.y+1) {
		for index := potential.startX; index <= potential.endX; index++ {
			if element.startX == (index-1) || element.startX == index+1 || element.startX == index {
				return true
			}
		}
	}
	return false
}

func isPartNumberValid(potentialPart Element, symbols []Element) bool {
	for _, symbol := range symbols {
		if isAdjacent(symbol, potentialPart) {
			return true
		}
	}
	return false
}

func getPartNumbers(schematic Schematic) (ret []Element) {
	for _, potentialPart := range schematic.potentialPartNumbers {
		if isPartNumberValid(potentialPart, schematic.symbols) {
			ret = append(ret, potentialPart)
		}
	}
	return ret
}

func getNumericValueForElement(element Element) int {
	value, err := strconv.Atoi(element.value)
	internal.Check(err)
	return value
}

func sumPartNumbers(schematic Schematic) (ret int) {
	ret = 0
	for _, element := range getPartNumbers(schematic) {
		ret += getNumericValueForElement(element)
	}
	return ret
}

func sumGearRatios(schematic Schematic) (ret int) {
	ret = 0
	for _, gearRatio := range schematic.gearRatios {
		ret += gearRatio
	}
	return ret
}

func main() {
	schematic := parseSchematic(internal.ReadFileToLines(os.Args[1]))
	println(sumPartNumbers(schematic))
	println(sumGearRatios(schematic))
}
