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
			if !isStar(arr[i][j]) {
				continue
			}
			ratio := getRatio(arr, directions, i, j)
			sum = sum + ratio
		}
	}
	fmt.Println(sum)
}

func getRatio(arr [][]string, directions []Point, i int, j int) int {
	one := 0
	two := 0
	partsCount := 0
	for _, p := range directions {
		yIndex := i + p.Y
		xIndex := j + p.X
		if yIndex < 0 || yIndex >= len(arr) || xIndex < 0 || xIndex >= len(arr[i]) {
			continue
		}
		val := arr[yIndex][xIndex]
        if !isNumber(val) {
            continue
        }
		if instancesAreOverlapping(arr, p, xIndex, yIndex) {
			continue
		}
		res := ""
		if p.X == -1 {
			res = fmt.Sprint(walkLeft(arr, yIndex, xIndex-1) + val)
		} else if p.X == 1 {
			res = fmt.Sprint(val + walkRight(arr, yIndex, xIndex+1))
		} else {
			res = fmt.Sprint(walkLeft(arr, yIndex, xIndex-1) + val + walkRight(arr, yIndex, xIndex+1))
		}
		num, err := strconv.Atoi(res)
		if err != nil {
			log.Fatal(err)
		}
		if one == 0 {
			one = num
		} else {
			two = num
		}
		partsCount++
	}
	if partsCount != 2 {
		return 0
	}
	return one * two
}

func walkLeft(arr [][]string, i int, j int) string {
	var sb strings.Builder
	for k := j; k >= 0; k-- {
		if !isNumber(arr[i][k]) {
			break
		}
		sb.WriteString(arr[i][k])
	}
	return reverse(sb.String())
}
func walkRight(arr [][]string, i int, j int) string {
	var sb strings.Builder
	for k := j; k < len(arr[i]); k++ {
		if !isNumber(arr[i][k]) {
			break
		}
		sb.WriteString(arr[i][k])
	}
	return sb.String()
}

func instancesAreOverlapping(arr [][]string, p Point, x int, y int) bool {
	if p.X == -1 && p.Y == -1 && isNumber(arr[y][x+1]) {
		return true
	}
	if p.X == 1 && p.Y == -1 && isNumber(arr[y][x-1]) {
		return true
	}
	if p.X == -1 && p.Y == 1 && isNumber(arr[y][x+1]) {
		return true
	}
	if p.X == 1 && p.Y == 1 && isNumber(arr[y][x-1]) {
		return true
	}
	return false
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true
}

func isStar(s string) bool {
	return s == "*"
}
