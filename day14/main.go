package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type (
	maskStruct struct {
		andMask        int64
		orMask         int64
		floatingPowers []int64
	}

	assignment struct {
		address int64
		value   int64
	}
)

func main() {

	r1 := newReaderV1()
	r2 := newReaderV2()
	for _, line := range slurpFile("./input.txt") {
		r1.readLine(line)
		r2.readLine(line)
	}

	log.Printf("part 1 => %v", reduce(r1.data))
	log.Printf("part 2 => %v", reduce(r2.data))
}

func reduce(data map[int64]int64) (res int64) {
	for _, v := range data {
		res += v
	}
	return res
}

func getAssign(line string) (assignment, bool) {
	res := assignment{}
	reg := regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)
	m := reg.FindStringSubmatch(line)
	if len(m) != 3 {
		return res, false
	}

	if address, err := strconv.Atoi(m[1]); err != nil {
		log.Fatal(err)
	} else {
		res.address = int64(address)
	}

	if value, err := strconv.Atoi(m[2]); err != nil {
		log.Fatal(err)
	} else {
		res.value = int64(value)
	}

	return res, true
}

func slurpFile(path string) []string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	lines = lines[0 : len(lines)-1]
	return lines
}
