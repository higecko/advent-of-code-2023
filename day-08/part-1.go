package main

import (
	"fmt"
	"os"
	// "slices"
	// "strconv"

	// "slices"
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
	return readFile("input-test2.txt")
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

	currentNode := "AAA"
	i := 0
	steps := 0

	for currentNode != "ZZZ" {
		fmt.Println(currentNode)
		currentNode = nodes[currentNode][rune(directions[i])]
		i++
		steps++
		if i == len(directions) {
			i = 0
		}
	}
	fmt.Println(steps)
}