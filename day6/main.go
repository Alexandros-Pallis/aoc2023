package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	RaceTime       int
	RecordDistance int
}

func NewRace(t int, d int) *Race {
	return &Race{RaceTime: t, RecordDistance: d}
}
func (r Race) getWaysRecordCanBeBeat() int {
	counter := 0
	for i := r.RaceTime; i >= 0; i-- {
		remaining := r.RaceTime - i
		distance := i * remaining
		if distance <= r.RecordDistance {
			continue
		}
		counter++
	}
	return counter
}

func main() {
	contents, _ := os.ReadFile("input.txt")
	sum := 0
	races := getRaces(string(contents))
	for _, race := range races {
		if sum == 0 {
			sum = sum + race.getWaysRecordCanBeBeat()
            continue
		}
        sum = sum * race.getWaysRecordCanBeBeat()
	}
	fmt.Println(sum)
}

func getRaces(contents string) []Race {
	res := make([]Race, 0)
	lines := strings.Split(contents, "\n")
	timeValues := strings.Split(lines[0], ":")[1]
	distanceValues := strings.Split(lines[1], ":")[1]
	time := getValues(timeValues)
	distance := getValues(distanceValues)
	if len(time) != len(distance) {
		log.Fatal("races amount mismatch")
	}
	for i := range time {
		res = append(res, *NewRace(time[i], distance[i]))
	}
	return res
}
func getValues(values string) []int {
	values = strings.Trim(values, " ")
	strValues := strings.Split(values, " ")
	res := make([]int, 0)
	for i := range strValues {
		if strValues[i] == "" {
			continue
		}
		num, _ := strconv.Atoi(strValues[i])
		res = append(res, num)
	}
	return res
}
