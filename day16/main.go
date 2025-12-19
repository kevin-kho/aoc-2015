package main

import (
	"aoc-2015/common"
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Sue struct {
	Number      int
	Children    *int
	Cats        *int
	Samoyeds    *int
	Pomeranians *int
	Akitas      *int
	Vizslas     *int
	Goldfish    *int
	Trees       *int
	Cars        *int
	Perfumes    *int
}

func createSues(data []byte) ([]Sue, error) {

	var sues []Sue

	for entry := range bytes.SplitSeq(data, []byte{10}) {

		var number int
		var children *int
		var cats *int
		var samoyeds *int
		var pomeranians *int
		var akitas *int
		var vizslas *int
		var goldfish *int
		var trees *int
		var cars *int
		var perfumes *int

		entryStr := strings.Split(string(entry), " ")
		numberStr, _ := strings.CutSuffix(entryStr[1], ":")
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			return sues, err
		}

		for i := 3; i < len(entryStr); i += 2 {
			obj, _ := strings.CutSuffix(entryStr[i-1], ":")

			val, _ := strings.CutSuffix(entryStr[i], ",")
			valInt, err := strconv.Atoi(val)
			if err != nil {
				return sues, err
			}

			switch obj {
			case "children":
				children = &valInt
			case "cats":
				cats = &valInt
			case "samoyeds":
				samoyeds = &valInt
			case "pomeranians":
				pomeranians = &valInt
			case "akitas":
				akitas = &valInt
			case "vizslas":
				vizslas = &valInt
			case "goldfish":
				goldfish = &valInt
			case "trees":
				trees = &valInt
			case "cars":
				cars = &valInt
			case "perfumes":
				perfumes = &valInt
			}

		}

		sue := Sue{
			Number:      number,
			Children:    children,
			Cats:        cats,
			Samoyeds:    samoyeds,
			Pomeranians: pomeranians,
			Akitas:      akitas,
			Vizslas:     vizslas,
			Goldfish:    goldfish,
			Trees:       trees,
			Cars:        cars,
			Perfumes:    perfumes,
		}

		sues = append(sues, sue)

	}

	return sues, nil

}

func findSue(sues []Sue) *int {
	var res *int
	for _, s := range sues {
		// Check the value
		if s.Children != nil && *s.Children != 3 {
			continue
		}

		if s.Cats != nil && *s.Cats != 7 {
			continue
		}

		if s.Samoyeds != nil && *s.Samoyeds != 2 {
			continue
		}

		if s.Pomeranians != nil && *s.Pomeranians != 3 {
			continue
		}

		if s.Akitas != nil && *s.Akitas != 0 {
			continue
		}

		if s.Vizslas != nil && *s.Vizslas != 0 {
			continue
		}

		if s.Goldfish != nil && *s.Goldfish != 5 {
			continue
		}

		if s.Trees != nil && *s.Trees != 3 {
			continue
		}

		if s.Cars != nil && *s.Cars != 2 {
			continue
		}

		if s.Perfumes != nil && *s.Perfumes != 1 {
			continue
		}

		res = &s.Number

	}

	return res

}

func main() {
	filePath := "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	sues, err := createSues(data)
	if err != nil {
		log.Fatal(err)
	}

	res := findSue(sues)
	fmt.Println(*res)

}
