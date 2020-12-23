package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

type readerV2 struct {
	mask maskStruct
	data map[int64]int64
}

func newReaderV2() readerV2 {
	return readerV2{
		mask: maskStruct{
			floatingPowers: []int64{},
		},
		data: map[int64]int64{},
	}
}

func (r *readerV2) readLine(line string) {
	var ok bool
	if r.updateMask(line) {
		return
	}

	var assign assignment
	if assign, ok = getAssign(line); ok {
		keys := r.applyMask(assign.address, true)
		for _, k := range keys {
			r.data[k] = assign.value
		}
		return
	}

	log.Fatalf("can't parse line %s", line)
}

func (r readerV2) applyMask(input int64, applyStatic bool) []int64 {
	if applyStatic {
		input = (input | r.mask.orMask) & r.mask.andMask
	}
	if len(r.mask.floatingPowers) == 0 {
		return []int64{input}
	}
	res := []int64{}
	sub := r
	pow := r.mask.floatingPowers[0]
	sub.mask.floatingPowers = r.mask.floatingPowers[1:]

	res = append(res, sub.applyMask(input, false)...)
	res = append(res, sub.applyMask(input+intPow(2, pow), false)...)
	return res
}

func (r *readerV2) updateMask(line string) bool {
	reg := regexp.MustCompile(`mask = ([01X]+)`)
	m := reg.FindStringSubmatch(line)
	if len(m) != 2 {
		return false
	}

	input := m[1]
	andMaskString := strings.ReplaceAll(input, "0", "1")
	andMaskString = strings.ReplaceAll(andMaskString, "X", "0") // remove Xs
	orMaskString := strings.ReplaceAll(input, "X", "0")         // write ones

	var err error
	if r.mask.andMask, err = strconv.ParseInt(andMaskString, 2, 64); err != nil {
		log.Fatal(err)
	}
	if r.mask.orMask, err = strconv.ParseInt(orMaskString, 2, 64); err != nil {
		log.Fatal(err)
	}

	newPowers := []int64{}
	for i, c := range input {
		if c == 'X' {
			newPowers = append(newPowers, int64(len(input)-i-1))
		}
	}
	r.mask.floatingPowers = newPowers
	return true
}

// power in integers rather than floats
func intPow(base, exp int64) int64 {
	if exp > 0 {
		return base * intPow(base, exp-1)
	}
	return 1
}
