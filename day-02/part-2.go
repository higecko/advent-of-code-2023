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
  
  colors := [...]string {
    "red",
    "green",
    "blue",
  }

  colorRegex := make(map[string]*regexp.Regexp)

  for _, color := range(colors) {
    colorRegex[color] = regexp.MustCompile(fmt.Sprintf(`(\d+) %s`, color))
  }

  for _, line := range(input) {
    cubeSetPower := 1
    for _, color := range(colors) {
      colorCountMatches := colorRegex[color].FindAllStringSubmatch(line, -1)
      minColorCount := 0
      for _, colorCountMatch := range(colorCountMatches) {
        colorCount := SubmatchToInt(colorCountMatch)
        minColorCount = max(minColorCount, colorCount)
      }
      cubeSetPower *= minColorCount
    }

    sum += cubeSetPower
  }

  fmt.Println("Sum of powers of cube sets", sum)
}