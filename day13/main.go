package main

import (
	"io/ioutil"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data := slurpFile("input.txt")
	start, err := strconv.Atoi(data[0])
	if err != nil {
		log.Fatal(err)
	}
	part1(start, data[1])
	part2(data[1])
}

func part2(schedule string) {
	buses := [][2]int{}
	for i, val := range strings.Split(schedule, ",") {
		if id, err := strconv.Atoi(val); err == nil {
			buses = append(buses, [2]int{i, id})
		}
	}

	// sort to iterate quicker
	sort.Slice(buses, func(i, j int) bool { return buses[i][1] > buses[j][1] })

	acc := 1
	val := 0
	for _, bus := range buses {
		for (val+bus[0])%bus[1] != 0 {
			val += acc
		}
		acc *= bus[1]
	}

	log.Printf("Part2 => %v", val)
}

func test(value int, buses [][2]int) bool {
	for _, bus := range buses {
		if (value+bus[0])%bus[1] != 0 {
			return false
		}
	}
	return true
}

func part1(start int, schedule string) {
	buses := []int{}
	for _, val := range strings.Split(schedule, ",") {
		if id, err := strconv.Atoi(val); err == nil {
			buses = append(buses, id)
		}
	}
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
