package main

import "fmt"
import "os"
import "strings"
import "strconv"

var numbers = [...]string{
  "fake",
  "one", "two", "three",
  "four", "five", "six",
  "seven", "eight", "nine" }


func assignIfNumber(a *int, b byte) error {
	var num int
	var err error
	if num, err = strconv.Atoi(string(b)); err == nil {
		*a = num
	}
	return err
}

func findLastNumericValue(str string) (int, int) {
  num := -1
  i := 0
  for i = len(str) -1 ; i > -1; i-- {
    var err error
    if err = assignIfNumber(&num, str[i]); err == nil {
      break
    }
  }

  return i, num
}

func findFirstNumericValue(str string) (int, int) {
  num := -1
  i := 0
  for i = 0; i < len(str); i++ {
    var err error
    if err = assignIfNumber(&num, str[i]); err == nil {
      break
    }
  }

  return i, num
}

func findFirstStringValue(str string) (int, int) {
  num := -1
  index := -1
  for i := 1; i < len(numbers); i++ {
    if in := strings.Index(str, numbers[i]); in > -1 {
      if index == -1 || in < index {
        index = in
        num = i
      }
    }
  }
  
  return index, num
}

func findLastStringValue(str string) (int, int) {
  num := -1
  index := -1
  for i := 1; i < len(numbers); i++ {
    if in := strings.LastIndex(str, numbers[i]); in > -1 {
      if in > index {
        index = in
        num = i
      }
    }
  }
  
  return index, num
}

func main() {
	//Read file into array
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")

  sum := 0
	for _, line := range lines {
    fi, f := findFirstNumericValue(line)
    if fsi, fs := findFirstStringValue(line); fsi > -1 && fsi < fi {
      f = fs
    }

    si, s := findLastNumericValue(line)
    if ssi, ss := findLastStringValue(line); ssi > si {
      s = ss
    }
		
    if f > 0 {
      sum += f * 10 + s
    } else {
      fmt.Println("ERROR!!!!", f, s, line)
    }
	}

  fmt.Println("The sum is", sum)
}