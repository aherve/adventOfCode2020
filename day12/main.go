package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type shipStr struct {
	Direction int
	EastPos   int
	NorthPos  int
}

func newShip() shipStr {
	return shipStr{
		Direction: 0,
		EastPos:   0,
		NorthPos:  0,
	}
}

func main() {
	ship := newShip()
	for _, instr := range slurpFile("./input.txt") {
		ship.do(instr)
	}
	log.Printf("part1 => %v", ship.ManhattanDist())
}

func (ship *shipStr) do(instr string) {
	command := instr[0]
	argStr := instr[1:]
	arg, err := strconv.Atoi(argStr)
	if err != nil {
		log.Fatal(err)
	}

	switch command {
	case 'N':
		ship.NorthPos += arg
	case 'S':
		ship.NorthPos -= arg
	case 'E':
		ship.EastPos += arg
	case 'W':
		ship.EastPos -= arg
	case 'L':
		ship.Direction = (ship.Direction + arg) % 360
	case 'R':
		ship.Direction = (ship.Direction + 360 - arg) % 360
	case 'F':
		switch ship.Direction {
		case 0:
			ship.EastPos += arg
		case 90:
			ship.NorthPos += arg
		case 180:
			ship.EastPos -= arg
		case 270:
			ship.NorthPos -= arg
		}
	}
}

func (ship shipStr) ManhattanDist() (dist int) {
	if d := ship.EastPos; d > 0 {
		dist += d
	} else {
		dist -= d
	}

	if d := ship.NorthPos; d > 0 {
		dist += d
	} else {
		dist -= d
	}
	return dist
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
