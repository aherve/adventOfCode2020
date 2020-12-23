package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

type readerV1 struct {
	mask maskStruct
	data map[int64]int64
}

func newReaderV1() readerV1 {
	return readerV1{
		mask: maskStruct{},
		data: map[int64]int64{},
	}
}

func (r *readerV1) readLine(line string) {

	var ok bool
	if r.updateMask(line) {
		return
	}

	var assign assignment
	if assign, ok = getAssign(line); ok {
		r.data[assign.address] = r.applyMask(assign.value)
		return
	}
	log.Fatalf("can't parse line %s", line)
}

func (r readerV1) applyMask(input int64) int64 {
	return (input | r.mask.orMask) & r.mask.andMask
}

func (r *readerV1) updateMask(line string) bool {

	reg := regexp.MustCompile(`mask = ([01X]+)`)
	m := reg.FindStringSubmatch(line)
	if len(m) != 2 {
		return false
	}

	input := m[1]
	andMaskString := strings.ReplaceAll(input, "X", "1")
	orMaskString := strings.ReplaceAll(input, "X", "0")

	var err error
	if r.mask.andMask, err = strconv.ParseInt(andMaskString, 2, 64); err != nil {
		log.Fatal(err)
	}
	if r.mask.orMask, err = strconv.ParseInt(orMaskString, 2, 64); err != nil {
		log.Fatal(err)
	}

	return true
}
