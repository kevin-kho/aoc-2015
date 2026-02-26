package main

import (
	"fmt"
	"log"

	"github.com/kevin-kho/aoc-utilities/common"
)

func getDelta(b byte) [2]int {
	var delta [2]int
	switch b {
	case '^': // ^
		delta = [2]int{0, -1}
	case 'v': // v
		delta = [2]int{0, 1}
	case '>': // >
		delta = [2]int{1, 0}
	case '<': // <
		delta = [2]int{-1, 0}
	default:
		delta = [2]int{0, 0}
	}
	return delta
}

func countUniqueHouses(dirs []byte) int {
	var count int
	pos := [2]int{0, 0}
	seen := make(map[[2]int]bool)
	seen[pos] = true
	count++

	for _, d := range dirs {
		delta := getDelta(d)
		pos = [2]int{pos[0] + delta[0], pos[1] + delta[1]}

		if !seen[pos] {
			count++
		}
		seen[pos] = true
	}

	return count

}

func countUniqueHousesTwoSantas(dirs []byte) int {
	var count int
	santa0 := [2]int{0, 0}
	santa1 := [2]int{0, 0}
	seen := make(map[[2]int]bool)
	seen[santa0] = true
	seen[santa1] = true
	count++

	for i, d := range dirs {
		delta := getDelta(d)
		if i%2 == 0 {
			santa0 = [2]int{santa0[0] + delta[0], santa0[1] + delta[1]}
			if !seen[santa0] {
				count++
			}
			seen[santa0] = true

		} else {
			santa1 = [2]int{santa1[0] + delta[0], santa1[1] + delta[1]}
			if !seen[santa1] {
				count++
			}
			seen[santa1] = true
		}

	}

	return count

}

func main() {

	filePath := "./input.txt"

	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	uniqueHouses := countUniqueHouses(data)
	fmt.Println(uniqueHouses)

	uniqueHousesTwoSantas := countUniqueHousesTwoSantas(data)
	fmt.Println(uniqueHousesTwoSantas)

}
