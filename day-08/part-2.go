package main

import (
	"fmt"
	"os"
	"slices"
	// "strconv"
	"regexp"
	"strings"
)

func readFile(filename string) []string {
	content, _ := os.ReadFile(filename)
	lines := strings.Split(string(content), "\n")
	return lines
}

func readInput() []string {
	return readFile("input.txt")
}

func readTestInput() []string {
	return readFile("input-test-part2.txt")
}

func main() {
	input := readInput()
	var directions string
	nodes := make(map[string]map[rune]string)

	nodeRegex := regexp.MustCompile(`([A-Z]{3}) = \(([A-Z]{3}), ([A-Z]{3})\)`)

	for idx, line := range(input) {
		if idx == 0 {
			directions = line
			continue
		}

		if line == "" {
			continue
		}

		nodeMatch := nodeRegex.FindAllStringSubmatch(line, -1)

		nodes[nodeMatch[0][1]] = map[rune]string {
			'L': nodeMatch[0][2],
			'R': nodeMatch[0][3],
		} 
	}

	currentNodes := []string{}
	for nodeKey := range(nodes) {
		if rune(nodeKey[2]) == 'A' {
			currentNodes = append(currentNodes, nodeKey)
		}
	}

	i := 0
	steps := uint64(0)
	nodeSteps := make(map[int]uint64)
	fmt.Println(currentNodes)

	for {
		for idx, currentNode := range(currentNodes) {
			if rune(currentNode[2]) != 'Z' {
				currentNodes[idx]= nodes[currentNode][rune(directions[i])]
				if rune(currentNodes[idx][2]) == 'Z' {
					nodeSteps[idx] = steps + 1
				}	
			}
		}
		i++
		steps++
		if i == len(directions) {
			i = 0
		}
		fmt.Println(currentNodes)
		if slices.ContainsFunc(currentNodes, func(node string) bool { return rune(node[2]) != 'Z'	}) == false {
			break
		}
	}

	allSteps := []uint64{}
	for _, val := range(nodeSteps) {
		allSteps = append(allSteps, val)
	}

	max := slices.Max(allSteps)
	multiplier := 1
	lcm := max * uint64(multiplier)

	for slices.ContainsFunc(allSteps, func(e uint64) bool { return (lcm % e) != 0}) {
		multiplier++
		lcm = max * uint64(multiplier)
	}

	fmt.Println(lcm)
}