package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"

	// "slices"
	// "strconv"
	// "regexp"
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

const (
	FiveOfAKind string = "FiveOfAKind"
	FourOfAKind string = "FourOfAKind"
	FullHouse string = "FullHouse"
	ThreeOfAKind string = "ThreeOfAKind"
	TwoPair string = "TwoPair"
	OnePair string = "OnePair"
	HighCard string = "HighCard"
)

var cardRanks = map[rune]int {
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'J': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

var handRanks = map[string]int {
	HighCard: 0,
	OnePair: 1,
	TwoPair: 2,
	ThreeOfAKind: 3,
	FullHouse: 4,
	FourOfAKind: 5,
	FiveOfAKind: 6,
}

type Hand struct {
	cards string
	bet int
	hand string
}

func ParseHand(cards string) string {
	counts := make(map[rune]int)

	for _, char := range(cards) {
		if _, ok := counts[char]; ok {
			counts[char]++
		} else {
			counts[char] = 1
		}
	}

	var hand string
	switch len(counts) {
	case 1:
		// One type of card in map has to be 5 of a kind
		hand = FiveOfAKind
	case 2:
		// 2 cards is either 4 of a kind or full house
		for _, val := range(counts) {
			if val == 2 || val == 3 {
				hand = FullHouse
			} else {
				hand = FourOfAKind
			}
			break
		}
	case 3:
		// 3 cards is 2 pair or 3 of a kind
		for _, val := range(counts) {
			switch val {
			case 1:
				continue
			case 2:
				hand = TwoPair
				break
			case 3:
				hand = ThreeOfAKind
				break
			}
		}
	case 4:
		// 4 cards is a pair
		hand = OnePair
	case 5:
		// 5 cards is high card
		hand = HighCard
	}

	return hand
}

func HandCmp(a,b Hand) int {
	cmp := handRanks[a.hand] - handRanks[b.hand]
	if cmp == 0 {
		for i := 0; i < len(a.cards); i++ {
			cmp = cardRanks[rune(a.cards[i])] - cardRanks[rune(b.cards[i])]
			if cmp != 0 {
				break
			}
		}
	}
	return cmp
}


func main() {
	input := readInput()
	
	hands := []Hand{}

	sum := 0

	for _, line := range(input) {
		parts := strings.Split(line, " ")
		bet, _ := strconv.Atoi(parts[1])
		hands = append(hands, Hand{parts[0], bet, ParseHand(parts[0])})
	}

	slices.SortFunc(hands, HandCmp)

	for idx, hand := range(hands) {
		sum += hand.bet * (idx + 1)
	}
	
	fmt.Println(sum)

}