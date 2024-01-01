package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
    res := 0
	tokens := getTokens()
	contents, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
        num, ok := findDigitsHandler(line, tokens)
        if(!ok) {
            continue
        }
        res = res + num
	}
    fmt.Print(res)
}

func findDigitsHandler(line string, tokens map[string]int) (int, bool) {
	first := findFirstDigit(line, tokens)
	last := findLastDigit(line, tokens)
    resultString := fmt.Sprintf("%v%v", first, last)
    res, err := strconv.Atoi(resultString)
    if(err != nil) {
        return 0, false
    }
	return res, true
}

func findLastDigit(line string, tokens map[string]int) string {
	chars := strings.Split(line, "")
	for i := len(chars) - 1; i >= 0; i-- {
		char := chars[i]
		if isNumeric(char) {
			return char
		}
		for j := i; j > i-5; j-- {
			if j < 0 || i < j {
				break
			}
			possibleToken := strings.Join(chars[j:i+1], "")
			if isToken(possibleToken, tokens) {
				return fmt.Sprintf("%v", tokens[possibleToken])
			}
		}
	}
	return ""
}

func findFirstDigit(line string, tokens map[string]int) string {
	chars := strings.Split(line, "")
	for i, char := range chars {
		if isNumeric(char) {
			return char
		}
		for j := i; j < i+5; j++ {
			if j >= len(chars) {
				break
			}
			possibleToken := strings.Join(chars[i:j+1], "")
			if isToken(possibleToken, tokens) {
				return fmt.Sprintf("%v", tokens[possibleToken])
			}
		}
	}
	return ""
}

func isNumeric(char string) bool {
	_, err := strconv.Atoi(char)
	if err != nil {
		return false
	}
	return true
}

func isToken(s string, tokens map[string]int) bool {
	result := false
	for token := range tokens {
		if s != token {
			continue
		}
		result = true
		break
	}
	return result
}

func getTokens() map[string]int {
	return map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
}
