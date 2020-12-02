package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type policyStruct struct {
	min      int
	max      int
	char     string
	password string
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	counter1, counter2 := 0, 0
	for scanner.Scan() {
		p := parsePolicy(scanner.Text())
		if p.part1Validator() {
			counter1++
		}
		if p.part2Validator() {
			counter2++
		}
	}
	fmt.Printf("part 1: %v\n", counter1)
	fmt.Printf("part 2: %v\n", counter2)
}

func (p policyStruct) part1Validator() bool {
	count := strings.Count(p.password, p.char)
	return count >= p.min && count <= p.max
}

func (p policyStruct) part2Validator() bool {
	// A xor B <=> A != B when A and B are booleans :)
	return (p.password[p.min-1] == p.char[0]) != (p.password[p.max-1] == p.char[0])
}

func parsePolicy(line string) policyStruct {
	split := regexp.MustCompile(`(\d+)-(\d+) ([a-zA-Z]): (\S+)`).FindAllStringSubmatch(line, -1)
	strMin, strMax, str, pass := split[0][1], split[0][2], split[0][3], split[0][4]
	min, _ := strconv.Atoi(strMin)
	max, _ := strconv.Atoi(strMax)
	return policyStruct{
		min:      min,
		max:      max,
		char:     str,
		password: pass,
	}
}
