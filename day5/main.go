package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Seed struct {
	Number   int
	Location int
}

func (s Seed) getLocation(
	seedToSoil []Map,
	soilToFertilizer []Map,
	fertilizerToWater []Map, waterToLight []Map,
	lightToTemperature []Map, temperatureToHumidity []Map, humidityToLocation []Map) int {
	soil := getFromSource(s.Number, seedToSoil)
	fertilizer := getFromSource(soil, soilToFertilizer)
	water := getFromSource(fertilizer, fertilizerToWater)
	light := getFromSource(water, waterToLight)
	temperature := getFromSource(light, lightToTemperature)
	humidity := getFromSource(temperature, temperatureToHumidity)
	location := getFromSource(humidity, humidityToLocation)
	return location
}

func getFromSource(n int, maps []Map) int {
	res := 0
	isInRange := false
	for _, m := range maps {
		if n < m.SourceStart || n > m.SourceStart+m.Range-1 {
			continue
		}
		isInRange = true
		diff := n - m.SourceStart
		res = m.DestinationStart + diff
	}
	if !isInRange {
		return n
	}
	return res
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
    minLocation := math.MaxInt
	for i, seed := range seeds {
        location := seed.getLocation(
			seedToSoil,
			soilToFertilizer,
			fertilizerToWater,
			waterToLight,
			lightToTemperature,
			temperatureToHumidity,
			humidityToLocation)
        seeds[i].Location = location
        if location < minLocation {
            minLocation = location
        }
	}
    fmt.Println(minLocation)
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
		res = append(res, Seed{Number: val})
	}
	return res
}
