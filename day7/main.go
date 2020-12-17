package main

import (
	"io/ioutil"
	"log"
	"strings"
)

type (
	node string

	containsEdge struct {
		target *node
		count  int
	}
	containedByEdge struct {
		target *node
	}

	graphStruct struct {
		nodes            map[string]*node
		containedByEdges map[*node][]containedByEdge
		containsEdges    map[*node][]containsEdge
	}
)

func main() {
	g := NewGraph()
	for _, line := range readFile("./input.txt") {
		if err := g.ParseLine(line); err != nil {
			log.Fatal(err)
		}
	}

	n := g.GetNode("shiny gold")
	containers := g.searchAllContainers(n)
	log.Printf("Part1 => %v", len(containers))

	log.Printf("Part2 => %v", g.countContents(n)-1)

}

func readFile(path string) []string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(data), "\n")
}
