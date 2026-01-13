package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Instruction struct {
	Command  string
	Register string
	Offset   int
}

func CreateInstructions(data []byte) ([]Instruction, error) {
	var res []Instruction
	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {
		entryStrArr := strings.Split(string(entry), " ")

		cmd := entryStrArr[0]
		var reg string
		var offset int
		var err error

		switch cmd {
		case "hlf":
			reg = entryStrArr[1]
		case "tpl":
			reg = entryStrArr[1]
		case "inc":
			reg = entryStrArr[1]
		case "jmp":
			offset, err = strconv.Atoi(entryStrArr[1])
		case "jie":
			reg = strings.TrimSuffix(entryStrArr[1], ",")
			offset, err = strconv.Atoi(entryStrArr[2])
		case "jio":
			reg = strings.TrimSuffix(entryStrArr[1], ",")
			offset, err = strconv.Atoi(entryStrArr[2])
		}

		if err != nil {
			return res, err
		}

		res = append(res, Instruction{
			Command:  cmd,
			Register: reg,
			Offset:   offset,
		})

	}

	return res, nil
}

func SolvePartOne(instructions []Instruction) int {
	reg := make(map[string]int)
	var i int
	for i < len(instructions) {
		in := instructions[i]
		switch in.Command {

		case "hlf":
			reg[in.Register] /= 2
		case "tpl":
			reg[in.Register] *= 3
		case "inc":
			reg[in.Register]++
		case "jmp":
			i += in.Offset
			continue
		case "jie":
			if reg[in.Register]%2 == 0 {
				i += in.Offset
				continue
			}
		case "jio":
			if reg[in.Register] == 1 {
				i += in.Offset
				continue
			}

		}
		i++
	}

	return reg["b"]
}

func main() {
	filePath := "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)
	in, err := CreateInstructions(data)
	if err != nil {
		log.Fatal(err)
	}
	res := SolvePartOne(in)
	fmt.Println(res)
}
