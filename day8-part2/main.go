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
	locationsMap := get(locationList, directions)
	locationsStart := filterLocations(locationsMap, "A")
    fmt.Println(walk(locationsMap, locationsStart, directions))
}

func walk(locMap map[string]Location, nodes []Location, directions []string) int {
    step := 0
    idx := 0
    maxIdx := len(directions)
    for !allNodesEndWith(nodes, "Z") {
        direction := directions[idx]
        idx++
        step++
        for i, loc := range nodes {
            if direction == "L" {
                nodes[i] = locMap[loc.Left]
            } else {
                nodes[i] = locMap[loc.Right]
            }
        }
        if idx == maxIdx {
            idx = 0
        }
    }
    return step
}

func allNodesEndWith(nodes []Location, needle string) bool {
	var result []Location
	for _, node := range nodes {
		if strings.HasSuffix(node.Key, needle) {
			result = append(result, node)
		}
	}
	return len(nodes) == len(result)
}

func filterLocations(locations map[string]Location, needle string) []Location {
	var result []Location
	for key, loc := range locations {
		if strings.HasSuffix(key, needle) {
			result = append(result, loc)
		}
	}
	return result
}

func get(locationList []string, directions []string) map[string]Location {
	locationMap := make(map[string]Location)
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
	return locationMap
}
