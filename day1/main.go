package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// open file
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln("could not open input", err)
	}
	defer file.Close()

	// read as csv
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("could not parse file", err)
	}

	// parse ints
	numbers := make([]int64, len(data))
	for i := range data {
		numbers[i], err = strconv.ParseInt(data[i][0], 10, 64)
		if err != nil {
			log.Fatalf("could not convert %s to an int64", data[i][0])
		}
	}

	// run part1
	n1, n2, err := part1(numbers)
	if err != nil {
		log.Fatalln("woops", err)
	}
	fmt.Printf("PART1: %v * %v = %v\n", n1, n2, n1*n2)

	// run part2
	n1, n2, n3, err := part2(numbers)
	if err != nil {
		log.Fatalln("woops", err)
	}
	fmt.Printf("PART2: %v * %v * %v = %v\n", n1, n2, n3, n1*n2*n3)

}

func part1(numbers []int64) (int64, int64, error) {
	// Search for answer
	for i := 0; i < len(numbers)-1; i++ {
		for j := i; j < len(numbers); j++ {
			if a, b := numbers[i], numbers[j]; a+b == 2020 {
				return a, b, nil
			}
		}
	}

	return 0, 0, fmt.Errorf("sorry, I could not find any solution")
}

func part2(numbers []int64) (int64, int64, int64, error) {

	// Search for answer
	for i := 0; i < len(numbers)-2; i++ {
		for j := i + 1; j < len(numbers)-1; j++ {
			for k := j + 1; k < len(numbers); k++ {
				if a, b, c := numbers[i], numbers[j], numbers[k]; a+b+c == 2020 {
					return a, b, c, nil
				}
			}
		}
	}

	return 0, 0, 0, fmt.Errorf("sorry, I could not find any solution")
}
