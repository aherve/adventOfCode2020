package main

import (
	"io/ioutil"
	"log"
	"strings"
)

type (
	worldStr struct {
		seats [][]rune
	}

	direction [2]int
)

var directions = [8]direction{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func main() {
	part1()
	part2()
}

func part2() {
	world := newWorld(readFile("./input.txt"))

	// iterate
	for world.part2Next() {
	}

	// count seats
	res := 0
	for _, s := range world.seats {
		res += count('#', s)
	}

	log.Printf("Part2 => %v", res)
}

func part1() {
	world := newWorld(readFile("./input.txt"))

	for world.part1Next() {
	}

	// now count seats
	res := 0
	for _, s := range world.seats {
		res += count('#', s)
	}

	log.Printf("Part1 => %v", res)
}

func (world *worldStr) part2Next() bool {
	var changed bool
	seats := make([][]rune, len(world.seats))
	nextWorld := worldStr{
		seats: seats,
	}

	for i := range world.seats {
		nextWorld.seats[i] = make([]rune, len(world.seats[i]))
		for j := range world.seats[i] {
			seat := world.seats[i][j]
			nextWorld.seats[i][j] = seat
			visible := []rune{}
			// look for seats
			for _, dir := range directions {
				visible = append(visible, world.lookFrom(i, j, dir))
			}

			// apply changes if any
			if seat == 'L' && count('#', visible) == 0 {
				nextWorld.seats[i][j] = '#'
				changed = true
			}

			if seat == '#' && count('#', visible) > 4 {
				nextWorld.seats[i][j] = 'L'
				changed = true
			}

		}

	}

	*world = nextWorld
	return changed
}

func (world *worldStr) part1Next() bool {
	var changed bool
	seats := make([][]rune, len(world.seats))
	nextWorld := worldStr{
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

	*world = nextWorld
	return changed
}

func count(r rune, list []rune) (res int) {
	for _, rr := range list {
		if rr == r {
			res++
		}
	}
	return res
}

func (world worldStr) lookFrom(i, j int, dir direction) rune {
	u, v := i+dir[0], j+dir[1]
	for world.contains(u, v) {
		if world.seats[u][v] == '#' {
			return '#'
		}
		if world.seats[u][v] == 'L' {
			return 'L'
		}
		u, v = u+dir[0], v+dir[1]
	}
	return '.'
}

func (world worldStr) adjacent(i, j int) []rune {
	res := []rune{}
	for _, dir := range directions {
		if u, v := i+dir[0], j+dir[1]; world.contains(u, v) {
			res = append(res, world.seats[u][v])
		}
	}
	return res
}

func (world worldStr) contains(i, j int) bool {
	return i >= 0 && j >= 0 && i < len(world.seats) && j < len(world.seats[0])
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

	return world
}

func readFile(path string) []string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	lines = lines[0 : len(lines)-1]
	return lines
}
