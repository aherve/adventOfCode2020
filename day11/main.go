package main

import (
	"io/ioutil"
	"log"
	"strings"
)

type worldStr struct {
	seats [][]rune
	iMax  int
	jMax  int
}

func main() {
	world := newWorld(readFile("./input.txt"))

	var changed = true
	for changed {
		world, changed = world.next()
	}

	// now count seats
	res := 0
	for _, s := range world.seats {
		res += count('#', s)
	}

	log.Printf("Part1 => %v", res)

}

func (world worldStr) next() (worldStr, bool) {
	var changed bool
	seats := make([][]rune, len(world.seats))
	nextWorld := worldStr{
		iMax:  world.iMax,
		jMax:  world.jMax,
		seats: seats,
	}

	for i := range world.seats {
		nextWorld.seats[i] = make([]rune, len(world.seats[i]))
		for j := range world.seats[i] {
			nextWorld.seats[i][j] = world.seats[i][j]
			seat := world.seats[i][j]
			adjacent := world.adjacent(i, j)

			if seat == 'L' && count('#', adjacent) == 0 {
				nextWorld.seats[i][j] = '#'
				changed = true
			}

			if seat == '#' && count('#', adjacent) > 3 {
				nextWorld.seats[i][j] = 'L'
				changed = true
			}

		}
	}

	return nextWorld, changed

}

func count(r rune, list []rune) (res int) {
	for _, rr := range list {
		if rr == r {
			res++
		}
	}
	return res
}

func (world worldStr) adjacent(i, j int) []rune {
	res := []rune{}
	if i > 0 && j > 0 {
		res = append(res, world.seats[i-1][j-1])
	}
	if i > 0 {
		res = append(res, world.seats[i-1][j])
	}
	if i > 0 && j < world.jMax {
		res = append(res, world.seats[i-1][j+1])
	}
	if j > 0 {
		res = append(res, world.seats[i][j-1])
	}
	if j < world.jMax {
		res = append(res, world.seats[i][j+1])
	}
	if i < world.iMax && j > 0 {
		res = append(res, world.seats[i+1][j-1])
	}
	if i < world.iMax {
		res = append(res, world.seats[i+1][j])
	}
	if i < world.iMax && j < world.jMax {
		res = append(res, world.seats[i+1][j+1])
	}
	return res
}

func newWorld(input []string) worldStr {
	seats := make([][]rune, len(input))

	world := worldStr{
		seats: seats,
	}

	for i, line := range input {
		world.seats[i] = make([]rune, len(line))
		for j, r := range line {
			world.seats[i][j] = r
		}
	}
	world.iMax = len(world.seats) - 1
	world.jMax = len(world.seats[0]) - 1

	return world
}

func readFile(path string) []string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	lines = lines[0 : len(lines)-1]
	log.Printf("DEBG '%s'", lines[len(lines)-1])
	return lines
}
