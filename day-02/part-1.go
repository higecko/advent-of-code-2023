package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readInput() []string {
  //Read file into array
  content, _ := os.ReadFile("input.txt")
  lines := strings.Split(string(content), "\n")
  return lines
}

func SubmatchToInt(match []string) int {
  num, _ := strconv.Atoi(match[1])
  return num
}

func main() {
  input := readInput()

  sum := 0
  
  bag := map[string]int {
    "red": 12,
    "green": 13,
    "blue": 14,
  }

  colorRegex := make(map[string]*regexp.Regexp)

  for color, _ := range(bag) {
    colorRegex[color] = regexp.MustCompile(fmt.Sprintf(`(\d+) %s`, color))
  }

  for _, line := range(input) {
    gameIdRegex := regexp.MustCompile(`Game (\d+):`)
    gameId := SubmatchToInt(gameIdRegex.FindStringSubmatch(line))

    possible := true
    bagcheckloop:
    for color, count := range(bag) {
      gameCountMatches := colorRegex[color].FindAllStringSubmatch(line, -1)
      
      for _, gameCountMatch := range(gameCountMatches) {
        gameCount := SubmatchToInt(gameCountMatch)
        if gameCount > count {
          possible = false
          break bagcheckloop
        }
      }
    }

    if possible {
      sum += gameId
    }
  }

  fmt.Println("Sum of possible game IDs", sum)
}