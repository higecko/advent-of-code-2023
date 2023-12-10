package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
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
	input := readInput()
	sum := 0

  numberMatchRegex := regexp.MustCompile(`\d+`)
  symbolMatchRegex := regexp.MustCompile(`[^\d\.]`)

	for lineIndex, line := range(input) {
    numberMatches := numberMatchRegex.FindAllStringIndex(line, -1)
    fmt.Println("line matches", lineIndex, numberMatches)

    for _, numberMatch := range(numberMatches) {
      isPart := false

      surroundSubstring := line[max(0, numberMatch[0] - 1):min(len(line),numberMatch[1] + 1)]
      if symbolMatchRegex.MatchString(surroundSubstring) {
        isPart = true || isPart
      }

      aboveLine := input[max(0, lineIndex - 1)]
      aboveLineSubstring := aboveLine[max(0, numberMatch[0] - 1):min(len(line),numberMatch[1] + 1)]
      if symbolMatchRegex.MatchString(aboveLineSubstring) {
        isPart = true || isPart
      }

      belowLine := input[min(len(input) - 1, lineIndex + 1)]
      belowLineSubstring := belowLine[max(0, numberMatch[0] - 1):min(len(line),numberMatch[1] + 1)]
      if symbolMatchRegex.MatchString(belowLineSubstring) {
        isPart = true || isPart
      }

      fmt.Println(surroundSubstring, isPart)
      if (isPart) {
        numberString := line[numberMatch[0]:numberMatch[1]]
        partNumber, _ := strconv.Atoi(numberString)
        sum += partNumber
      }
    }
  }

  fmt.Println("The sum is ", sum)
}