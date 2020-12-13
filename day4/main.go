package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	Data map[string]string
}

func main() {
	lines := readFile("./input.txt")
	passports := make([]passport, len(lines))
	for i, line := range lines {
		passports[i] = parse(line)
	}

	validCntV1 := 0
	validCntV2 := 0
	for _, p := range passports {
		if p.validv1() {
			validCntV1++
		}
		if p.validv2() {
			validCntV2++
		}
	}

	log.Println(validCntV1)
	log.Println(validCntV2)
}

func (p passport) validv2() bool {
	res := map[string]bool{}
	for k, v := range p.Data {
		switch k {
		case "byr":
			if between(v, 1920, 2020) {
				res[k] = true
			}
		case "iyr":
			if between(v, 2010, 2020) {
				res[k] = true
			}
		case "eyr":
			if between(v, 2020, 2030) {
				res[k] = true
			}
		case "hgt":
			matches := regexp.MustCompile(`^(\d+)(cm|in)$`).FindStringSubmatch(v)
			if len(matches) == 3 {
				h, _ := strconv.Atoi(matches[1])
				if matches[2] == "cm" && h >= 150 && h <= 193 {
					res[k] = true
				} else if matches[2] == "in" && h >= 59 && h <= 76 {
					res[k] = true
				}
			}
		case "hcl":
			if regexp.MustCompile(`^#([a-f0-9]{6})$`).MatchString(v) {
				res[k] = true
			}
		case "ecl":
			switch v {
			case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
				res[k] = true
			}
		case "pid":
			if regexp.MustCompile(`^(\d{9})$`).MatchString(v) {
				res[k] = true
			}
		}
	}
	return len(res) == 7
}

func between(str string, min, max int) bool {
	num, err := strconv.Atoi(str)
	return err == nil && num >= min && num <= max
}

func (p passport) validv1() bool {
	res := map[string]bool{}
	for k := range p.Data {
		switch k {
		case "byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid":
			res[k] = true
		}
	}
	return len(res) == 7
}

func parse(line string) passport {
	pass := passport{
		Data: map[string]string{},
	}
	line = strings.ReplaceAll(line, "\n", " ")
	words := strings.Split(line, " ")
	for _, word := range words {
		if word != "" {
			s := strings.Split(word, ":")
			if len(s) != 2 {
				log.Fatalf("cannot parse %s", word)
			}
			pass.Data[s[0]] = s[1]
		}
	}
	return pass
}

func readFile(path string) []string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(data), "\n\n")
}
