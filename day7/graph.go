package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func NewGraph() *graphStruct {
	return &graphStruct{
		nodes:            map[string]*node{},
		containedByEdges: map[*node][]containedByEdge{},
		containsEdges:    map[*node][]containsEdge{},
	}
}

func (graph graphStruct) searchAllContainers(n *node) map[*node]bool {
	res := map[*node]bool{}

	for _, container := range graph.containedByEdges[n] {
		// declare node
		res[container.target] = true
		// recursive search
		res = merge(res, graph.searchAllContainers(container.target))
	}

	return res
}

// returns all bags, INCLUDING self (substract 1 to answer the question)
func (graph graphStruct) countContents(n *node) int {
	res := 1

	for _, content := range graph.containsEdges[n] {
		res += content.count * graph.countContents(content.target)
	}

	return res
}

func (graph *graphStruct) ParseLine(line string) error {
	line = strings.TrimSpace(line)
	if line == "" {
		return nil
	}
	reg1 := regexp.MustCompile(`(\w+ \w+) bags contain (no other bags|.*)\.`)
	match1 := reg1.FindStringSubmatch(line)

	if len(match1) != 3 {
		return fmt.Errorf("failed to parse %s => %+v", line, match1)
	}

	container := graph.GetNode(match1[1])

	rules := strings.Split(match1[2], ", ")
	reg2 := regexp.MustCompile(`(\d) (\w+ \w+) bag`)
	regNoOther := regexp.MustCompile(`no other bags`)

	for _, rule := range rules {

		if regNoOther.MatchString(rule) {
			continue
		}

		match2 := reg2.FindStringSubmatch(rule)
		if len(match2) != 3 {
			return fmt.Errorf("failed to parse %s => %+v with lengh %v", rule, match2, len(match2))
		}
		count, err := strconv.Atoi(match2[1])
		if err != nil {
			return err
		}
		target := graph.GetNode(match2[2])

		if _, ok := graph.containsEdges[container]; !ok {
			graph.containsEdges[container] = []containsEdge{}
		}

		if _, ok := graph.containedByEdges[target]; !ok {
			graph.containedByEdges[target] = []containedByEdge{}
		}

		graph.containsEdges[container] = append(graph.containsEdges[container], containsEdge{
			target: target,
			count:  count,
		})

		graph.containedByEdges[target] = append(graph.containedByEdges[target], containedByEdge{
			target: container,
		})

	}

	return nil
}

func (g *graphStruct) GetNode(s string) *node {
	if n, ok := g.nodes[s]; ok {
		return n
	} else {
		n := node(s)
		g.nodes[s] = &n
		return &n
	}
}

func merge(m1, m2 map[*node]bool) map[*node]bool {
	for k := range m2 {
		m1[k] = true
	}
	return m1
}
