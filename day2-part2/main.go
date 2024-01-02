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
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(contents), "\n")
	lines = lines[:len(lines)-1]
    sum := 0
	for _, line := range lines {
        red, green, blue := handleGame(line)
        power := red * green * blue
        sum = sum + power
	}
    fmt.Print(sum)
}

func handleGame(line string) (int, int, int) {
	split := strings.Split(line, ": ")
	rounds := strings.Split(split[1], "; ")
	red := 0
	green := 0
	blue := 0
	for _, round := range rounds {
		cubes := strings.Split(round, ", ")
		for _, cube := range cubes {
			cubeContent := strings.Split(cube, " ")
			amount, err := strconv.Atoi(cubeContent[0])
			color := cubeContent[1]
			if err != nil {
				log.Fatal(err)
			}
			switch color {
			case "red":
				if amount > red {
					red = amount
				}
			case "green":
				if amount > green {
					green = amount
				}
			case "blue":
				if amount > blue {
					blue = amount
				}
			default:
				log.Fatal("not supported")
				break
			}
		}
	}
	return red, green, blue
}
