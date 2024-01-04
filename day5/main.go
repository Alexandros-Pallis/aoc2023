package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Seed struct {
	Number   int
	Location int
}

func (s Seed) getLocation(
	seedToSoil map[int]int,
	soilToFertilizer map[int]int,
	fertilizerToWater map[int]int, waterToLight map[int]int,
	lightToTemperature map[int]int, temperatureToHumidity map[int]int, humidityToLocation map[int]int) int {
	soil := getFromSource(s.Number, seedToSoil)
	fertilizer := getFromSource(soil, soilToFertilizer)
	water := getFromSource(fertilizer, fertilizerToWater)
	light := getFromSource(water, waterToLight)
	temperature := getFromSource(light, lightToTemperature)
	humidity := getFromSource(temperature, temperatureToHumidity)
	location := getFromSource(humidity, humidityToLocation)
	return location
}

func getFromSource(n int, m map[int]int) int {
	val, ok := m[n]
	if ok {
		return val
	}
	return n
}

func getRanges(maps []Map) map[int]int {
    desinationSource := make(map[int]int)
	for _, m := range maps {
		for i := 0; i <= m.Range; i++ {
			desinationSource[m.DestinationStart+i] = m.SourceStart + i
		}
	}
	return desinationSource
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
	seedToSoil := getRanges(getMap(lines, "seed-to-soil"))
	soilToFertilizer := getRanges(getMap(lines, "soil-to-fertilizer"))
	fertilizerToWater := getRanges(getMap(lines, "fertilizer-to-water"))
	waterToLight := getRanges(getMap(lines, "water-to-light"))
	lightToTemperature := getRanges(getMap(lines, "light-to-temperature"))
	temperatureToHumidity := getRanges(getMap(lines, "temperature-to-humidity"))
	humidityToLocation := getRanges(getMap(lines, "humidity-to-location"))
	for _, seed := range seeds {
		seed.Location = seed.getLocation(
            seedToSoil,
            soilToFertilizer,
			fertilizerToWater,
            waterToLight,
            lightToTemperature,
            temperatureToHumidity,
			humidityToLocation)
	}
    fmt.Println(seeds)
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
