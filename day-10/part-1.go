package main

import (
	"fmt"
	"os"
	"strings"
	"slices"
	// "strconv"
	// "regexp"
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

	// Find S
	// Using v for vertical (Y) and h for horizontal (X). 
	v := 0
	var h int
	for v = 0; v < len(input); v++ {
		h = strings.Index(input[v], "S")
		if h != -1 {
			break
		}
	}

	//Holds direction of last move
	var d rune

	//Figure out which way we can go to start
	if slices.Contains([]byte{'7','-','J'}, input[v][h + 1]) {
		h += 1
		d = 'W'
	} else if slices.Contains([]byte{'L','-','F'}, input[v][h - 1]) {
		h -= 1
		d = 'E'
	} else if slices.Contains([]byte{'7','|','F'}, input[v - 1][h]) {
		v -= 1
		d = 'N'
	} else if slices.Contains([]byte{'J','|','L'}, input[v + 1][h]) {
		v += 1
		d = 'S'
	} else {
		panic("We shouldn't be here")
	}

	//Start walking the pipe
	steps := 1
	for input[v][h] != 'S' {

		//Figure out which direction to go
		switch input[v][h] {
		case '|':
			if d == 'N' {
				v -= 1
			} else {
				v += 1
			}

		case '-':
			if d == 'W' {
				h += 1
			} else {
				h -= 1
			}

		case 'L':
			if d == 'S' {
				h += 1
				d = 'W'
			} else {
				v -= 1
				d = 'N'
			}

		case 'J':
			if d == 'S' {
				h -= 1
				d = 'E'
			} else {
				v -= 1
				d = 'N'
			}

		case '7':
			if d == 'W' {
				v += 1
				d = 'S'
			} else {
				h -= 1
				d = 'E'
			}

		case 'F':
			if d == 'E' {
				v += 1
				d = 'S'
			} else {
				h += 1
				d = 'W'
			}

		default:
			panic("What direction???")
		}

		steps++
	}

	fmt.Println(steps/2)
}