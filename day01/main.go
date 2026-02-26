package main

import (
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
)

func solve(directions []byte) int {
	var floor int
	// For reference
	// 40 = ( -> go up one floor
	// 41 = ) -> go down one floor

	for _, dir := range directions {

		switch dir {
		case '(':
			floor++
		case ')':
			floor--
		default:
			continue
		}

	}

	return floor
}

func solveBasement(directions []byte) int {

	var floor int
	for i, dir := range directions {

		switch dir {
		case '(':
			floor++
		case ')':
			floor--
		default:
			continue
		}

		if floor == -1 {
			return i + 1
		}

	}

	return -1

}

func main() {

	filePath := "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	finalFloor := solve(data)

	posBasement := solveBasement(data)

	fmt.Println(finalFloor)
	fmt.Println(posBasement)

}
