package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
    result := 0
	contents, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
        number, ok := getCalibrationValue(line)
        if(!ok) {
            continue
        }
        result = result + number
	}

    fmt.Print(result)
}

func getCalibrationValue(line string) (int, bool) {
	var accumulatedValue strings.Builder
	chars := strings.Split(line, "")
	for i := 0; i < len(chars); i++ {
		_, err := strconv.Atoi(chars[i])
		if err != nil {
			continue
		}
		accumulatedValue.WriteString(chars[i])
		break
	}
	for i := len(chars) - 1; i >= 0; i-- {
		_, err := strconv.Atoi(chars[i])
		if err != nil {
			continue
		}
		accumulatedValue.WriteString(chars[i])
		break
	}
	resultString := accumulatedValue.String()
	result, err := strconv.Atoi(resultString)
    if(err != nil) {
        return 0, false
    }
	return result, true
}
