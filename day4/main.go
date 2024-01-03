package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	contents, err := os.ReadFile("input.txt")
	lines := strings.Split(string(contents), "\n")
	lines = lines[:len(lines)-1]
	if err != nil {
		log.Fatal(err)
	}
	winning := getWinningNumbers(lines)
	cardNumbers := getCardNumbers(lines)
	fmt.Println(getSum(cardNumbers, winning))
}

func getSum(playerCards [][]int, winningCards [][]int) int {
	sum := 0
	for i, card := range playerCards {
		firstTime := true
		cardSum := 0
		for _, num := range card {
			if !inArray[int](num, winningCards[i]) {
				continue
			}
			if firstTime {
				cardSum = cardSum + 1
			} else {
				cardSum = cardSum * 2
			}
			firstTime = false
		}
		sum = sum + cardSum
	}
	return sum
}

func getWinningNumbers(lines []string) [][]int {
	result := make([][]int, 0)
	for _, line := range lines {
		values := strings.Split(line, ":")[1]
        values = strings.Split(values, "|")[0]
		values = strings.Trim(values, " ")
		cardNumbers := strings.Split(values, " ")
		cardArr := make([]int, 0)
		for _, cardNum := range cardNumbers {
			cardNum = strings.Trim(cardNum, " ")
            if cardNum == " " || cardNum == "" {
                continue
            }
			num, err := strconv.Atoi(cardNum)
			if err != nil {
				log.Fatal(err)
			}
			cardArr = append(cardArr, num)
		}
		result = append(result, cardArr)
	}
	return result
}

func getCardNumbers(lines []string) [][]int {
	result := make([][]int, 0)
	for _, line := range lines {
		rightSide := strings.Split(line, "|")[1]
		rightSide = strings.Trim(rightSide, " ")
		values := strings.Split(rightSide, " ")
		cardArr := make([]int, 0)
		for _, v := range values {
			v = strings.Trim(v, " ")
            if v == " " || v == "" {
                continue
            }
			num, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			cardArr = append(cardArr, num)
		}
		result = append(result, cardArr)
	}
	return result
}

func inArray[T comparable](needle T, haystack []T) bool {
	res := false
	for _, v := range haystack {
		if v == needle {
			res = true
		}
	}
	return res
}
