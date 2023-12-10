package main

import (
	"fmt"
	"os"
	// "slices"
	"strconv"
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
	return readFile("input-test.txt")
}

func main() {
	input := readInput()
	prod := 1
	numberRegex := regexp.MustCompile(`\d+`)

	timesStr := numberRegex.FindAllString(input[0], -1)
	distancesStr := numberRegex.FindAllString(input[1], -1)

	for i := 0; i < len(timesStr); i++ {

		time, _ := strconv.Atoi(timesStr[i])
		distance, _ := strconv.Atoi(distancesStr[i])

		timeDiff := time / 2
		start := timeDiff

		for {
			// fmt.Println(timeDiff)
			if timeDiff < 1 {
				panic("We should not be here!!!!!")
			}

			if timeDiff * (time - timeDiff) <= distance {
				timeDiff++
				break
			}

			timeDiff--
		}
		
		var waysToWin int
		if time % 2 == 0 {
			waysToWin = (start - timeDiff) * 2 + 1
		} else {
			waysToWin = (start - timeDiff + 1) * 2
		}

		prod *= waysToWin
	}
	fmt.Println(prod)
	fmt.Println("done")
}
