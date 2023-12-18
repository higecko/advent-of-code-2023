package main

import (
	"fmt"
	"os"
	"strings"

	// "slices"
	// "strconv"

	"slices"
	"strconv"
	// "regexp"
	// "strings"
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

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func rowdiff(row []int) []int {
	res := []int{}
	for i := 0; i < len(row) - 1; i++ {
		res = append(res, row[i] - row[i + 1])
	}
	return res
}

func main() {
	input := readInput()
	lines := [][]int{}
	sum := 0

	for _, line := range input {
		digits := strings.Split(line, " ")
		digitLine := []int{}
		for _, digit := range(digits) {
			digitLine = append(digitLine, atoi(digit))
		}
		slices.Reverse(digitLine)
		lines = append(lines, digitLine)
	}

	for _, line := range(lines) {
		stack := [][]int{}
		
		for {
			if slices.ContainsFunc(line, func(e int) bool { return e != line[0]}) {
				stack = append(stack, line)
				line = rowdiff(line)
			} else {
				break
			}
		}

		diff := line[0]

		for len(stack) > 0 {
			fmt.Println(stack)
			elem := stack[len(stack) - 1]
			stack = stack[0:len(stack) - 1]

			diff = elem[0] + diff
		}

		sum += diff
	}


	fmt.Println(sum)
}