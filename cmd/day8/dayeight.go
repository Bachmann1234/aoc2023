package main

import (
	"dev/mattbachmann/aoc2023/internal"
	"os"
	"strings"
	"time"
)

const Right = 'R'
const Left = 'L'

type Node struct {
	label string
	left  *Node
	right *Node
}

func parseTree(mapLines []string) map[string]*Node {
	labelToNode := make(map[string]*Node)
	for _, line := range mapLines {
		lineParts := strings.Split(line, " = ")
		nodeParts := strings.Split(lineParts[1], ", ")
		labelToNode[lineParts[0]] = &Node{
			lineParts[0],
			&Node{
				nodeParts[0][1:],
				nil,
				nil,
			},
			&Node{
				nodeParts[1][:len(nodeParts[1])-1],
				nil,
				nil,
			},
		}
	}

	for _, line := range mapLines {
		lineParts := strings.Split(line, " = ")
		nodeParts := strings.Split(lineParts[1], ", ")
		label := lineParts[0]
		left := nodeParts[0][1:]
		right := nodeParts[1][:len(nodeParts[1])-1]
		node := labelToNode[label]
		leftNode, ok := labelToNode[left]
		if !ok {
			panic("Could not get left")
		}
		rightNode, ok := labelToNode[right]
		if !ok {
			panic("Could not get right")
		}
		node.left = leftNode
		node.right = rightNode

	}

	return labelToNode
}

func partOne(directions []rune, mapLines []string) (ret int) {
	labelToNode := parseTree(mapLines)
	position := labelToNode["AAA"]
	if position == nil {
		return -1
	}
	for {
		for _, direction := range directions {
			if direction == Right {
				position = position.right
			} else if direction == Left {
				position = position.left
			} else {
				panic("WHAT")
			}
			ret += 1
			if position.label == "ZZZ" {
				return ret
			}
		}
	}
	return -1
}

type Ghost struct {
	directions     []rune
	position       *Node
	directionIndex int
	moves          int
}

func findCycle(c chan int, ghost *Ghost) {
	count := 0
	lastDifference := 0
	var matches []int
	for {
		for _, direction := range ghost.directions {
			if direction == Right {
				ghost.position = ghost.position.right
			} else if direction == Left {
				ghost.position = ghost.position.left
			} else {
				panic("WHAT")
			}
			count += 1
			if ghost.position.label[len(ghost.position.label)-1] == 'Z' {
				matches = append(matches, count)
				if len(matches) > 1 {
					difference := matches[len(matches)-1] - matches[len(matches)-2]
					if difference == lastDifference {
						c <- difference
						return
					}
					lastDifference = difference
				}
			}
		}
		time.Sleep(1 * time.Millisecond)
	}
}

func partTwo(directions []rune, mapLines []string) (ret int) {
	labelToNode := parseTree(mapLines)
	var ghosts []*Ghost
	for key, _ := range labelToNode {
		if key[len(key)-1] == 'A' {
			ghosts = append(ghosts, &Ghost{
				directions:     directions,
				position:       labelToNode[key],
				directionIndex: 0,
				moves:          0,
			})
		}
	}
	channel := make(chan int)
	var results []int
	for _, ghost := range ghosts {
		go findCycle(channel, ghost)
	}

	for i := 0; i < len(ghosts); i++ {
		v := <-channel
		results = append(results, v)
	}
	ret = results[0]
	for {
		if allDivideEvenly(ret, results) {
			return ret
		} else {
			ret += results[0]
		}
	}
}

func allDivideEvenly(result int, values []int) bool {
	for _, value := range values {
		if result%value != 0 {
			return false
		}
	}
	return true
}

func main() {
	lines := internal.ReadFileToLines(os.Args[1])
	directions := []rune(lines[0])
	println(partOne(directions, lines[1:]))
	println(partTwo(directions, lines[1:]))
}
