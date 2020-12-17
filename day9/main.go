package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data := readFile("./input.txt")
	numbers := make([]int, len(data))
	var err error
	for i := range data {
		numbers[i], err = strconv.Atoi(data[i])
		if err != nil {
			log.Fatalf("cannot convert %s to integer", data[i])
		}
	}

	p1, err := part1(numbers, 25)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Part1 => %v", p1)

	log.Printf("Part2 => %v", part2(numbers, p1))
}

func part2(numbers []int, searchFor int) int {
	var min, max int
	for i := 0; i < len(numbers)-1; i++ {
		min, max = numbers[i], numbers[i]
		j := i + 1
		if numbers[j] < min {
			min = numbers[j]
		}
		if numbers[j] > max {
			max = numbers[j]
		}
		sum := numbers[i] + numbers[j]
		for sum < searchFor {
			j++
			sum += numbers[j]
			if numbers[j] < min {
				min = numbers[j]
			}
			if numbers[j] > max {
				max = numbers[j]
			}
			if sum == searchFor {
				return min + max
			}
		}
	}
	return -1
}

func part1(numbers []int, size int) (int, error) {
	for i := size; i < len(numbers); i++ {
		if !isSum(numbers[i], numbers[i-size:i]) {
			return numbers[i], nil
		}
	}
	return -1, fmt.Errorf("sorry, could not find any value")
}

func isSum(n int, base []int) bool {
	for i := 0; i < len(base)-1; i++ {
		if base[i] > n {
			continue
		}
		for j := i + 1; j < len(base); j++ {
			if base[i]+base[j] == n {
				return true
			}
		}
	}
	return false
}

func readFile(path string) []string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	res := strings.Split(string(data), "\n")
	res = res[0 : len(res)-1]
	return res
}
