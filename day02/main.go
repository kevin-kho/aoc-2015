package main

import (
	"bytes"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type dimensions struct {
	length int
	width  int
	height int
}

func (d dimensions) getTotalSurfaceArea() int {
	return 2*d.length*d.width + 2*d.width*d.height + 2*d.height*d.length
}

func (d dimensions) getExtra() int {
	bob := []int{d.length * d.width, d.width * d.height, d.height * d.length}
	return slices.Min(bob)
}

func (d dimensions) getRibbonFeet() int {

	dims := []int{d.height, d.length, d.width}
	slices.Sort(dims)

	ribbon := dims[0]*2 + dims[1]*2
	bow := d.length * d.height * d.width

	return ribbon + bow
}

func getDimensionStructs(byteArr []byte) ([]dimensions, error) {

	// For reference "x" is represented by 120

	var d []dimensions

	for b := range bytes.SplitSeq(byteArr, []byte{'\n'}) {
		entry := string(b)
		split := strings.Split(entry, "x")
		if len(split) == 3 {
			l, err := strconv.Atoi(split[0])
			w, err := strconv.Atoi(split[1])
			h, err := strconv.Atoi(split[2])

			if err != nil {
				return d, err
			}
			dim := dimensions{
				length: l,
				width:  w,
				height: h,
			}
			d = append(d, dim)
		}
	}

	return d, nil

}

func getTotalSquareFeet(dims []dimensions) int {
	var total int
	for _, d := range dims {
		total += d.getTotalSurfaceArea()
		total += d.getExtra()
	}
	return total
}

func getTotalRibbon(dims []dimensions) int {

	var total int
	for _, d := range dims {
		ribbon := d.getRibbonFeet()
		total += ribbon
	}

	return total

}

func main() {

	filePath := "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	dims, err := getDimensionStructs(data)
	if err != nil {
		log.Fatal(err)
	}

	totalSquareFeet := getTotalSquareFeet(dims)
	fmt.Println(totalSquareFeet)

	totalRibbon := getTotalRibbon(dims)
	fmt.Println(totalRibbon)

	// fmt.Println(data)

}
