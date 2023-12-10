package main

import "fmt"
import "os"
import "strings"
import "strconv"

func assignIfNumber(a *int, b byte) error {
	var num int
	var err error
	if num, err = strconv.Atoi(string(b)); err == nil {
		*a = num
	}
	return err
}

func main() {
	//Read file into array
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")

  sum := 0

	for _, line := range lines {
		f, s := -1, -1
		i := 0
		j := len(line) - 1
		for i <= j && (f < 0 || s < 0) {
			// fmt.Println(i, string(line[i]), j, string(line[j]))
			if f < 0 {
				if err := assignIfNumber(&f, line[i]); err != nil {
					i++
				}
			}

			if s < 0 {
				if err := assignIfNumber(&s, line[j]); err != nil {
					j--
				}
			}
		}
		
    if f > 0 {
      sum += f * 10 + s
    } else {
      fmt.Println("ERROR!!!!", f, s, line)
    }
	}

  fmt.Println("The sum is", sum)
}