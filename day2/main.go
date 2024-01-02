package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	// "strconv"
	"strings"
)

type Game struct {
	id int
}

func NewGame(id int) *Game {
	g := new(Game)
	g.id = id
	return g
}

type Cube struct {
	color     string
	available int
	total     int
}

func (c Cube) isLimitExceeded(amount int) bool {
	return c.available-amount < 0
}

func NewCube(color string, total int) *Cube {
	c := new(Cube)
	c.color = color
	c.available = total
	c.total = total
	return c
}

func main() {
	sum := 0
	red := NewCube("red", 12)
	green := NewCube("green", 13)
	blue := NewCube("blue", 14)
	contents, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(contents), "\n")
	lines = lines[:len(lines)-1]
	for _, line := range lines {
        id, possible := handleGame(line, *red, *green, *blue)
		if !possible {
			continue
		}
		sum = sum + id
	}

	fmt.Print(sum)
}

func handleGame(line string, red Cube, green Cube, blue Cube) (int , bool) {
	isPossible := true
	split := strings.Split(line, ": ")
	game := split[0]
	gameId, err := strconv.Atoi(strings.Split(game, " ")[1])
	if err != nil {
		log.Fatal(err)
	}
	rounds := strings.Split(split[1], "; ")
	for _, round := range rounds {
		redCount := 0
		greenCount := 0
		blueCount := 0
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
				redCount = redCount + amount
			case "green":
				greenCount = greenCount + amount
			case "blue":
				blueCount = blueCount + amount
			default:
				log.Fatal("not supported")
				break
			}
			if red.isLimitExceeded(redCount) ||
				green.isLimitExceeded(greenCount) ||
				blue.isLimitExceeded(blueCount) {
				isPossible = false
			}
		}
	}
    return gameId, isPossible
}
