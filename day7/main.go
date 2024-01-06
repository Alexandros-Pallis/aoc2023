// parts of code inspired by bsadia
// https://github.com/bsadia/aoc_goLang/blob/master/day07/main.go
package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type Hand struct {
	Cards    string
	Bid      int
	HandType int
}

func NewHand(cards string, bid int, hType int) Hand {
	return Hand{
		Cards:    cards,
		Bid:      bid,
		HandType: hType,
	}
}

var typeMap = map[string]int{
	"FIVE_KIND":  7,
	"FOUR_KIND":  6,
	"FULL_HOUSE": 5,
	"THREE_KIND": 4,
	"TWO_PAIR":   3,
	"ONE_PAIR":   2,
	"HIGH_CARD":  1,
}

var cardsMap = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

func getHandType(cards string) int {
	cardMap := map[rune]int{}
	list := []int{}
	for _, card := range cards {
		cardMap[card]++
	}
	for _, value := range cardMap {
		list = append(list, value)
	}
	sort.Ints(list)
	highest := 0
	if len(list) > 0 {
		highest = list[len(list)-1]
	}
	secondHighest := 0
	if len(list) > 1 {
		secondHighest = list[len(list)-2]
	}
	htype := 0
	switch highest {
	case 5:
		htype = typeMap["FIVE_KIND"]
	case 4:
		htype = typeMap["FOUR_KIND"]
	case 3:
		if secondHighest == 2 {
			htype = typeMap["FULL_HOUSE"]
		}
		htype = typeMap["THREE_KIND"]
	case 2:
		htype = typeMap["TWO_PAIR"]
	default:
		htype = typeMap["ONE_PAIR"]
	}
	return htype
}

func main() {
	contents, _ := os.ReadFile("test.txt")
	lines := strings.Split(string(contents), "\n")
	lines = lines[:len(lines)-1]
	hands := getHands(lines)
	slices.SortFunc(hands, func(a, b Hand) int {
		if a.HandType < b.HandType {
			return 1
		}
		if a.HandType > b.HandType {
			return -1
		}
		return 0
	})
	fmt.Println(hands)
}

func getHands(lines []string) []Hand {
	res := make([]Hand, 0)
	for _, line := range lines {
		cards := strings.Split(line, " ")[0]
		bid, _ := strconv.Atoi(strings.Split(line, " ")[1])
		handType := getHandType(cards)
		hand := NewHand(cards, bid, handType)
		res = append(res, hand)
	}
	return res
}
