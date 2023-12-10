package main

import (
	"fmt"
	"os"
	"regexp"
	// "strconv"
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
  return readFile("input-test.txt")
}

func main() {
	input := readTestInput()
	sum := 0

  // numberMatchRegex := regexp.MustCompile(`\d+`)
  // symbolMatchRegex := regexp.MustCompile(`[^\d\.]`)
  starMatchRegex := regexp.MustCompile(`\*{1}`)

	for lineIndex, line := range(input) {
    starMatches := starMatchRegex.FindAllStringIndex(line, -1)
    fmt.Println("star matches", lineIndex, starMatches)

    for _, starMatch := range(starMatches) {
      var first, second int

      toBeAssigned := &first

      break
    }
  }

  fmt.Println("The sum is ", sum)
}