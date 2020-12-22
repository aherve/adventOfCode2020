package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type (
	shipStr struct {
		Direction int
		EastPos   int
		NorthPos  int
		Waypoint  *WaypointStr
	}

	WaypointStr struct {
		EastPos  int
		NorthPos int
	}
)

func newShip() shipStr {
	return shipStr{
		Direction: 0,
		EastPos:   0,
		NorthPos:  0,
		Waypoint:  &WaypointStr{10, 1},
	}
}

func main() {
	ship1, ship2 := newShip(), newShip()
	for _, instr := range slurpFile("./input.txt") {
		ship1.part1Navigate(instr)
		ship2.part2Navigate(instr)
	}
	log.Printf("part1 => %v", ship1.ManhattanDist())
	log.Printf("part2 => %v", ship2.ManhattanDist())
}

func (ship *shipStr) part2Navigate(instr string) {
	command, arg := readInstruction(instr)

	switch command {
	case 'N':
		ship.Waypoint.NorthPos += arg
	case 'S':
		ship.Waypoint.NorthPos -= arg
	case 'E':
		ship.Waypoint.EastPos += arg
	case 'W':
		ship.Waypoint.EastPos -= arg
	case 'L':
		ship.Waypoint.rotate(arg)
	case 'R':
		ship.Waypoint.rotate(-1 * arg)
	case 'F':
		ship.NorthPos += arg * ship.Waypoint.NorthPos
		ship.EastPos += arg * ship.Waypoint.EastPos
	}
}

func (wp *WaypointStr) rotate(degres int) {
	for i := 0; i < ((degres+360)%360)/90; i++ { // convert to a positive number of left turns
		wp.EastPos, wp.NorthPos = -wp.NorthPos, wp.EastPos // apply left turn
	}
}

func readInstruction(input string) (byte, int) {
	command := input[0]
	argStr := input[1:]
	arg, err := strconv.Atoi(argStr)
	if err != nil {
		log.Fatal(err)
	}
	return command, arg
}

func (ship *shipStr) part1Navigate(instr string) {
	command, arg := readInstruction(instr)

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
