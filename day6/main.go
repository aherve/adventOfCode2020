package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	anySum, allSum := 0, 0
	for _, data := range readFile("./input.txt") {
		data = strings.TrimSpace(data)
		anySum += anyScore(data)
		allSum += allScore(data)
	}
	log.Printf("Part 1 => %v", anySum)
	log.Printf("Part 2 => %v", allSum)
}

func anyScore(data string) int {
	data = strings.ReplaceAll(data, "\n", "")
	all := map[rune]bool{}
	for _, c := range data {
		all[c] = true
	}
	return len(all)
}

func allScore(data string) int {
	result := 0
	perPerson := strings.Split(data, "\n")
	breakdown := map[string]int{}
	for _, personData := range perPerson {
		for _, c := range personData {
			breakdown[string(c)]++
			if breakdown[string(c)] == len(perPerson) {
				result++
			}
		}
	}
	return result
}

func readFile(path string) []string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(data), "\n\n")
}
