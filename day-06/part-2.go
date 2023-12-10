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
	numberRegex := regexp.MustCompile(`\d+`)

	timesStr := numberRegex.FindAllString(input[0], -1)
	distancesStr := numberRegex.FindAllString(input[1], -1)

	var timesBuilder, distancesBuilder strings.Builder

	for i := 0; i < len(timesStr); i++ {
		timesBuilder.WriteString(timesStr[i])
		distancesBuilder.WriteString(distancesStr[i])
	}

	time, _ := strconv.ParseUint(timesBuilder.String(), 10, 64)
	distance, _ := strconv.ParseUint(distancesBuilder.String(), 10, 64)

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
	
	var waysToWin uint64
	if time % 2 == 0 {
		waysToWin = (start - timeDiff) * 2 + 1
	} else {
		waysToWin = (start - timeDiff + 1) * 2
	}

	fmt.Println(waysToWin)
	fmt.Println("done")
}
