package main

import (
	"fmt"
	"os"
	"strings"
)

type Location struct {
	Key   string
	Left  string
	Right string
}

func main() {
	contents, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(contents), "\n")
	directions := strings.Split(lines[0], "")
	locationList := lines[2 : len(lines)-1]
	get(locationList, directions)
}

func get(locationList []string, directions []string) {
	locationMap := make(map[string]Location)
	firstKey := "AAA"
	lastKey := "ZZZ"
	for i := range locationList {
		split := strings.Split(locationList[i], " = ")
		key := split[0]
		valuesStr := split[1]
		values := strings.Split(valuesStr, ", ")
		left := strings.ReplaceAll(values[0], "(", "")
		right := strings.ReplaceAll(values[1], ")", "")
		location := Location{Key: key, Left: left, Right: right}
		locationMap[key] = location
	}
	index := 0
	maxIndex := len(directions)
	step := 0
	location := locationMap[firstKey]
	for location.Key != lastKey {
		step++
		direction := directions[index]
		if direction == "L" {
			location = locationMap[location.Left]
		} else {
			location = locationMap[location.Right]
		}
        index++
		if index == maxIndex {
			index = 0
		}
	}
	fmt.Println(step)
}
