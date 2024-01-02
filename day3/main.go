package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func main() {
	contents, err := os.ReadFile("input.txt")
	var directions []Point
	directions = append(directions, Point{0, 1})
	directions = append(directions, Point{0, -1})
	directions = append(directions, Point{-1, 0})
	directions = append(directions, Point{1, 0})
	directions = append(directions, Point{-1, 1})
	directions = append(directions, Point{1, 1})
	directions = append(directions, Point{-1, -1})
	directions = append(directions, Point{1, -1})
	var arr [][]string
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(contents), "\n")
	lines = lines[:len(lines)-1]
	for _, line := range lines {
		values := strings.Split(line, "")
		tmp := make([]string, 0)
		tmp = append(tmp, values...)
		arr = append(arr, tmp)
	}
	sum := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			val := arr[i][j]
			if !isNumeric(val) || isDot(val) {
				continue
			}
			lastKPos := j
			var sb strings.Builder
			isAdjacent := false
			for k := j; k < len(arr[i]); k++ {
				if !isNumeric(arr[i][k]) {
					break
				}
				for _, point := range directions {
					yIndex := i + point.Y
					xIndex := k + point.X
					if yIndex < 0 || yIndex >= len(arr) {
						continue
					}
					if xIndex < 0 || xIndex >= len(arr[i]) {
						continue
					}
					adjacentValue := arr[yIndex][xIndex]
					if !isNumeric(adjacentValue) && !isDot(adjacentValue) {
						isAdjacent = true
					}
				}
				sb.WriteString(arr[i][k])
				lastKPos = k
			}
			if isAdjacent {
				num, err := strconv.Atoi(sb.String())
				if err != nil {
					log.Fatal(err)
				}
				sum = sum + num
				j = lastKPos
			}
		}
	}
	fmt.Println(sum)
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true
}

func isDot(s string) bool {
	return s == "."
}
