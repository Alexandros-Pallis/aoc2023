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
			ratio := getStarRatio(directions, arr, i, j)
			if ratio == 0 {
				continue
			}
			sum = sum + ratio
		}
	}
	fmt.Println(sum)
}

func getStarRatio(directions []Point, arr [][]string, i int, j int) int {
	adjacentCounter := 0
	partOne := 0
	partTwo := 0
	for _, point := range directions {
		xIndex := j + point.X
		yIndex := i + point.Y
		if yIndex < 0 || xIndex > len(arr) {
			continue
		}
		if xIndex < 0 || yIndex > len(arr[i]) {
			continue
		}
		val := arr[yIndex][xIndex]
		if !isNumber(val) {
			continue
		}
		left := walkLeft(arr, point, yIndex, xIndex)
		right := walkRight(arr, point, yIndex, xIndex)
		res := fmt.Sprint(left + val + right)
        num, err := strconv.Atoi(res)
        if(err != nil) {
            log.Fatal(err)
        }
		adjacentCounter++
	}
	if adjacentCounter != 2 || partOne == 0 || partTwo == 0 {
		return 0
	}
	return partOne * partTwo
}


func walkRight(arr [][]string, p Point, y int, x int) string {
	str := ""
	for k := x + 1; k < len(arr[y]); k++ {
		val := arr[y][k]
		if !isNumber(val) {
			break
		}
		str = fmt.Sprint(str + val)
	}
	return str
}

func walkLeft(arr [][]string, p Point, y int, x int) string {
	str := ""
	for k := x - 1; k >= 0; k-- {
		val := arr[y][k]
		if !isNumber(val) {
			break
		}
		str = fmt.Sprint(val + str)
	}
	return str
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
