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
	'J': 0,
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
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

func CheckJokerCount(counts map[rune]int, jokerCount int) bool {
	ret := false
	if jc, ok := counts['J']; ok && jc == jokerCount {
		ret = true
	}
	return ret
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
				// If we have a pair of Jokers, this can turn into four of a kind
				if CheckJokerCount(counts, 2) || CheckJokerCount(counts, 3) {
					hand = FiveOfAKind
				}
			} else {
				hand = FourOfAKind
				// if we have a Joker, this becomes 5 of a kind
				if CheckJokerCount(counts, 1) || CheckJokerCount(counts, 4) {
					hand = FiveOfAKind
				}
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
				if CheckJokerCount(counts, 1) {
					hand = FullHouse
				} else if CheckJokerCount(counts, 2) {
					hand = FourOfAKind
				}
				break
			case 3:
				hand = ThreeOfAKind
				if CheckJokerCount(counts, 1) || CheckJokerCount(counts, 3) {
					hand = FourOfAKind
				}
				break
			}
		}
	case 4:
		// 4 cards is a pair
		hand = OnePair
		if CheckJokerCount(counts, 1) || CheckJokerCount(counts, 2){
			hand = ThreeOfAKind
		}
	case 5:
		// 5 cards is high card
		hand = HighCard
		if CheckJokerCount(counts, 1) {
			hand = OnePair
		}
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