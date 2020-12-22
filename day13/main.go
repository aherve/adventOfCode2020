package main

import (
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	data := slurpFile("input.txt")
	start, err := strconv.Atoi(data[0])
	if err != nil {
		log.Fatal(err)
	}
	buses := []int{}
	for _, val := range strings.Split(data[1], ",") {
		if id, err := strconv.Atoi(val); err == nil && val != "x" {
			buses = append(buses, id)
		}
	}

	part1(start, buses)

}

func part1(start int, buses []int) {
	minWait := math.MaxInt32
	var rightID int
	for _, id := range buses {
		floatStart := float64(start)
		floatID := float64(id)
		floatMin := math.Round(floatID * math.Ceil(floatStart/floatID))
		newMin := int(floatMin)
		if newMin < minWait {
			minWait = newMin
			rightID = id
		}
	}
	log.Printf("Part 1 => %v", rightID*(minWait-start))
}

func slurpFile(path string) []string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	lines = lines[0 : len(lines)-1]
	return lines
}
