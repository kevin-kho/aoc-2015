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

func commandAssign(entry []string, wire map[string]int) {

	value, err := strconv.Atoi(entry[0])
	if err != nil {
		value = wire[entry[0]] // case: it's referring a wire
	}
	dst := entry[len(entry)-1]

	wire[dst] = value

}

func commandNot(entry []string, wire map[string]int) {

	value, err := strconv.Atoi(entry[1])
	if err != nil {
		value = wire[entry[1]] // case: it's referring a wire
	}
	value = ^value

	dst := entry[len(entry)-1]

	wire[dst] = value

}

func commandAndOr(entry []string, wire map[string]int) {
	left, err := strconv.Atoi(entry[0])
	if err != nil {
		left = wire[entry[0]]
	}

	right, err := strconv.Atoi(entry[2])
	if err != nil {
		right = wire[entry[2]]
	}
	dst := entry[len(entry)-1]

	var value int
	if entry[1] == "AND" {
		value = left & right
	} else {
		value = left | right
	}

	wire[dst] = value

}

func commandShift(entry []string, wire map[string]int) {

	left, err := strconv.Atoi(entry[0])
	if err != nil {
		left = wire[entry[0]]
	}
	right, err := strconv.Atoi(entry[2])
	if err != nil {
		right = wire[entry[2]]
	}
	dst := entry[len(entry)-1]

	var value int
	if entry[1] == "LSHIFT" {
		value = left << right
	} else {
		value = left >> right
	}

	wire[dst] = value

}

func parseData(data []byte) map[string]int {
	wire := make(map[string]int)
	for entry := range bytes.SplitSeq(data, []byte{10}) {

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
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	wire := parseData(data)
	fmt.Println(wire)

}
