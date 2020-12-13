package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	max := 0
	available := map[int]bool{}
	for i := 0; i <= 8*127+7; i++ {
		available[i] = true
	}
	for _, line := range readFile("./input.txt") {
		id := seatID(line)
		if id > max {
			max = id
		}

		delete(available, id)
	}

	log.Printf("Part 1 => %v", max)

	for id := range available {
		if _, empty := available[id+1]; empty {
			continue
		}
		if _, empty := available[id-1]; empty {
			continue
		}
		log.Printf("Part 2 => %v", id)
		break
	}

}

func seatID(s string) int {
	row := [2]int{0, 127}
	col := [2]int{0, 7}

	for _, x := range s {
		switch x {
		case 'F':
			row[1] = (row[1] + row[0]) / 2
		case 'B':
			row[0] = (row[1]+row[0])/2 + 1
		case 'R':
			col[0] = (col[1]+col[0])/2 + 1
		case 'L':
			col[1] = (col[1] + col[0]) / 2
		}
	}
	return 8*row[0] + col[0]
}

func readFile(path string) []string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(data), "\n")
}
