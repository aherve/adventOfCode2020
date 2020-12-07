package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	world := readFile("input.txt")
	fmt.Printf("part1 => %v\n", treeCount(1, 3, world))

	part2 := treeCount(1, 1, world) * treeCount(1, 3, world) * treeCount(1, 5, world) * treeCount(1, 7, world) * treeCount(2, 1, world)

	fmt.Printf("part2 => %v\n", part2)
}

func treeCount(istep, jstep int, world []string) int {
	depth := len(world)
	width := len(world[0])
	i, j, res := 0, 0, 0

	next := newIterator(width, istep, jstep)

	for i < depth-1 {
		if world[i][j] == '#' {
			res++
		}
		i, j = next()
	}

	return res
}

func newIterator(maxWidth, istep, jstep int) func() (int, int) {
	i, j := 0, 0
	return func() (int, int) {
		i = i + istep
		j = (j + jstep) % maxWidth
		return i, j
	}
}

func readFile(path string) []string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(data), "\n")
}
