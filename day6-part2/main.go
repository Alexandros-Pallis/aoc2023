package main

import (
	"fmt"
	// "log"
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
    res := 0
    maxHold := 0
	minHold := 0
    // find lowest hold time to beat record
    for holdTime := 1; holdTime <= r.RaceTime; holdTime++ {
        remaining := r.RaceTime - holdTime;
        record := holdTime * remaining
        if record <= r.RecordDistance {
            continue
        }
        minHold = holdTime - 1 // minus one because lowest value needs to be included in substraction later
        break;
    }
    // find highest hold time to beat record
    for holdTime := r.RaceTime - 1; holdTime >= 0; holdTime-- {
        remaining := r.RaceTime - holdTime;
        record := holdTime * remaining
        if record <= r.RecordDistance {
            continue
        }
        maxHold = holdTime
        break;
    }
    res = maxHold - minHold
	return res
}

func main() {
	contents, _ := os.ReadFile("input.txt")
	race := getRaces(string(contents))
    fmt.Println(race.getWaysRecordCanBeBeat())
}

func getRaces(contents string) Race {
	lines := strings.Split(contents, "\n")
    timeStr := strings.Split(lines[0], ":")[1]
    timeStr = strings.ReplaceAll(timeStr, " ", "")
    distanceStr := strings.Split(lines[1], ":")[1]
    distanceStr = strings.ReplaceAll(distanceStr, " ", "")
    time, _ := strconv.Atoi(timeStr)
    distance, _ := strconv.Atoi(distanceStr)
    return *NewRace(time, distance)
}
