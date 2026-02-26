package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Command struct {
	Value       *int
	Destination string
	Action      *string // AND, OR, LSHIFT, RSHIFT
	Left        *string
	Right       *string
}

func commandAssign(entry []string, wire map[string]uint16) {

	var value uint16
	if v, err := strconv.Atoi(entry[0]); err != nil {
		value = wire[entry[0]]
	} else {
		value = uint16(v)
	}

	dst := entry[len(entry)-1]

	wire[dst] = value

}

func commandNot(entry []string, wire map[string]uint16) {

	var value uint16
	if v, err := strconv.Atoi(entry[1]); err != nil {
		value = wire[entry[1]]
	} else {
		value = uint16(v)
	}

	value = ^value

	dst := entry[len(entry)-1]

	wire[dst] = value

}

func commandAndOr(entry []string, wire map[string]uint16) {

	var left uint16
	if l, err := strconv.Atoi(entry[0]); err != nil {
		left = wire[entry[0]]
	} else {
		left = uint16(l)
	}

	var right uint16
	if r, err := strconv.Atoi(entry[0]); err != nil {
		right = wire[entry[2]]
	} else {
		right = uint16(r)
	}

	dst := entry[len(entry)-1]

	var value uint16
	if entry[1] == "AND" {
		value = left & right
	} else {
		value = left | right
	}

	wire[dst] = value

}

func commandShift(entry []string, wire map[string]uint16) {

	var left uint16
	if l, err := strconv.Atoi(entry[0]); err != nil {
		left = wire[entry[0]]
	} else {
		left = uint16(l)
	}

	var right uint16
	if r, err := strconv.Atoi(entry[2]); err != nil {
		right = wire[entry[0]]
	} else {
		right = uint16(r)
	}

	dst := entry[len(entry)-1]

	var value uint16
	if entry[1] == "LSHIFT" {
		value = left << right
	} else {
		value = left >> right
	}

	wire[dst] = value

}

func parseData(data []byte) map[string]uint16 {
	wire := make(map[string]uint16)
	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {

		entryStrArr := strings.Split(string(entry), " ")
		switch len(entryStrArr) {
		case 3:
			commandAssign(entryStrArr, wire)
		case 4:
			commandNot(entryStrArr, wire)
		case 5:
			// AND, OR, LSHIFT, RSHIFT
			if entryStrArr[1] == "LSHIFT" || entryStrArr[1] == "RSHIFT" {
				commandShift(entryStrArr, wire)
			} else {
				commandAndOr(entryStrArr, wire)
			}
		}

	}

	return wire

}

func main() {

	filePath := "./inputExample.txt"
	filePath = "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	wire := parseData(data)
	fmt.Println(wire)

}
