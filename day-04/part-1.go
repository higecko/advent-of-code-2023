package main

import (
	"fmt"
	"os"
	"regexp"
	// "strconv"
	"strings"
  "math"
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

  numberRegex := regexp.MustCompile(`\d+`)

	for _, line := range(input) {
    colonPos := strings.Index(line, ":")
    pipePos := strings.Index(line, "|")
    winningNumbers := numberRegex.FindAllString(line[colonPos + 1:pipePos - 1], -1)
    myNumbers := numberRegex.FindAllString(line[pipePos+1:], -1)
    
    table := make(map[string]bool)
    for _, num := range(winningNumbers) {
      table[num] = false
    }

    counter := 0
    
    for _, num := range(myNumbers) {
      val, ok := table[num]
      if (ok && val == false) {
        counter++
        table[num] = true
      }
    }
    
    if counter > 0 {
      sum += int(math.Pow(2,float64((counter - 1))))
    }
  }

  fmt.Println("The sum is ", sum)
}