package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Seed struct {
	number int
}

type Map struct {
	SourceStart      int
	DestinationStart int
	Range            int
}

func main() {
	contents, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(contents), "\n")
	seeds := getSeeds(lines)
	seedToSoil := getMap(lines, "seed-to-soil")
	soilToFertilizer := getMap(lines, "soil-to-fertilizer")
	fertilizerToWater := getMap(lines, "fertilizer-to-water")
	waterToLight := getMap(lines, "water-to-light")
	lightToTemperature := getMap(lines, "light-to-temperature")
	temperatureToHumidity := getMap(lines, "temperature-to-humidity")
	humidityToLocation := getMap(lines, "humidity-to-location")

	fmt.Println(seeds)
	fmt.Println(seedToSoil)
	fmt.Println(soilToFertilizer)
	fmt.Println(fertilizerToWater)
	fmt.Println(waterToLight)
	fmt.Println(lightToTemperature)
	fmt.Println(temperatureToHumidity)
	fmt.Println(humidityToLocation)
}

func getMap(lines []string, search string) []Map {
	res := make([]Map, 0)
	mapLines := make([]string, 0)
	done := false
	for i, line := range lines {
		if done {
			break
		}
		if !strings.HasPrefix(line, search) {
			continue
		}
		for j := i + 1; j < len(lines); j++ {
			if lines[j] == "" {
				done = true
				break
			}
			mapLines = append(mapLines, lines[j])
		}
	}
	for _, line := range mapLines {
		line = strings.Trim(line, " ")
		values := strings.Split(line, " ")
		destinationStart, _ := strconv.Atoi(values[0])
		sourceStart, _ := strconv.Atoi(values[1])
		r, _ := strconv.Atoi(values[2])
		res = append(res, Map{
			DestinationStart: destinationStart,
			SourceStart:      sourceStart,
			Range:            r,
		})
	}
	return res
}

func getSeeds(lines []string) []Seed {
	res := make([]Seed, 0)
	seedsLine := strings.Split(lines[0], ":")[1]
	seedsLine = strings.Trim(seedsLine, " ")
	seeds := strings.Split(seedsLine, " ")
	for _, seed := range seeds {
		val, err := strconv.Atoi(seed)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, Seed{number: val})
	}
	return res
}
