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
	input := readInput()
	sum := 0

  numberRegex := regexp.MustCompile(`\d+`)

  cards := make(map[int]int)

	for idx, line := range(input) {
    if _, ok := cards[idx]; ok {
      cards[idx]++
    } else {
      cards[idx] = 1
    }

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
    
    for i := 1; i <= counter; i++ {
      nextIndex := idx + i
      if _, ok := cards[nextIndex]; ok {
        cards[nextIndex] += cards[idx]
      } else {
        cards[nextIndex] = cards[idx]
      }
    }
  }
  for _, v := range(cards) {
    sum += v
  }
  fmt.Println("The sum is ", sum)
}