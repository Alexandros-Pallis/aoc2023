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

func getNodes(lines []string) map[string]Location {
	locationList := make(map[string]Location)
	for _, line := range lines {
		line = strings.ReplaceAll(line, " ", "")
		firstSplit := strings.Split(line, "=")
		key := firstSplit[0]
		leftRight := strings.Split(firstSplit[1], ",")
		left := strings.Replace(leftRight[0], "(", "", 1)
		right := strings.Replace(leftRight[1], ")", "", 1)
		locationList[key] = Location{Key: key, Left: left, Right: right}
	}
	return locationList
}

func filterNodesThatEndsWith(nodes map[string]Location, needle string) []Location {
	var result []Location
	for key, node := range nodes {
		if !strings.HasSuffix(key, needle) {
			continue
		}
		result = append(result, node)
	}
	return result
}

func checkValues(values []bool) bool {
	for _, value := range values {
		if !value {
			return false
		}
	}
	return true
}

func main() {
	contents, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(contents), "\n")
	directionsString := lines[0]
	directions := strings.Split(directionsString, "")
	lines = lines[2 : len(lines)-1]
	nodes := getNodes(lines)
	currentNodes := filterNodesThatEndsWith(nodes, "A")
	steps := 0
	index := 0
	for {
        if index == len(directions) {
            index = 0
        }
		direction := directions[index]
		nextNodes := make([]Location, 0)
		checkValuesSlice := make([]bool, 0)
		for _, node := range currentNodes {
			if direction == "L" {
				nextNode := nodes[node.Left]
				checkValuesSlice = append(checkValuesSlice, strings.HasSuffix(nextNode.Key, "Z"))
				nextNodes = append(nextNodes, nextNode)
			} else {
				nextNode := nodes[node.Right]
				checkValuesSlice = append(checkValuesSlice, strings.HasSuffix(nextNode.Key, "Z"))
				nextNodes = append(nextNodes, nextNode)
			}
		}
		currentNodes = nextNodes
		steps++
        index++
		if checkValues(checkValuesSlice) {
			break
		}
	}

	fmt.Println(steps)
}
