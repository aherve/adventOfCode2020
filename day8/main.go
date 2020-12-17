package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	original := readFile("./input.txt")
	original = original[0 : len(original)-1] // get rid of last empty line

	// Part1
	part1, _ := run(original)
	log.Printf("part 1 => %v", part1)

	// Part2
	log.Printf("Part 2 => %v", fixAndReturn(original))
}

func fixAndReturn(original []string) int {
	for i := range original {
		fixed := make([]string, len(original))
		copy(fixed, original)
		if strings.HasPrefix(original[i], "acc") {
			continue
		} else if strings.HasPrefix(original[i], "jmp") {
			fixed[i] = strings.ReplaceAll(fixed[i], "jmp", "nop")
			if res, err := run(fixed); err == nil {
				return res
			}
		} else if strings.HasPrefix(original[i], "nop") {
			fixed[i] = strings.ReplaceAll(fixed[i], "nop", "jmp")
			if res, err := run(fixed); err == nil {
				return res
			}
		}
	}

	log.Fatalf("sorry, could not find any fix")
	return 0
}

func run(program []string) (int, error) {
	i, acc := 0, 0
	read := map[int]int{}

	for i < len(program) {
		read[i]++
		if read[i] >= 2 {
			return acc, fmt.Errorf("loop at line %v", i)
		}

		args := strings.Split(program[i], " ")
		instr := args[0]
		arg, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal(err)
		}

		switch instr {
		case "acc":
			acc += arg
			i++
		case "jmp":
			i += arg
		case "nop":
			i++
		}

	}
	return acc, nil
}

func readFile(path string) []string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(data), "\n")
}
