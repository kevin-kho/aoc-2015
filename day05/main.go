package main

import (
	"aoc-2015/common"
	"fmt"
	"log"
	"strings"
)

func isVowel(char string) bool {
	vowels := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
	}
	return vowels[char]
}

func isForbidden(char0, char1 string) bool {
	forbidden := map[string]bool{
		"ab": true,
		"cd": true,
		"pq": true,
		"xy": true,
	}

	return forbidden[char0+char1]

}

func countNiceStrings(stringArr []string) int {
	var count int

	for _, str := range stringArr {

		if len(str) < 3 {
			continue
		}

		seenVowels := make(map[string]int)
		twiceInARow := false
		containsForbidden := false

		// Check the very first character if vowel
		if isVowel(string(str[0])) {
			seenVowels[string(str[0])]++
		}

		for i := 1; i < len(str); i++ {
			char0 := string(str[i-1])
			char1 := string(str[i])

			// Check for vowel
			if isVowel(char1) {
				seenVowels[char1]++
			}

			// Check for twice in a row
			if char0 == char1 {
				twiceInARow = true
			}

			// Check for forbidden
			if isForbidden(char0, char1) {
				containsForbidden = true
			}

		}

		var totalVowels int
		for _, ct := range seenVowels {
			totalVowels += ct
		}

		if totalVowels >= 3 && twiceInARow && !containsForbidden {
			count++
		}

	}

	return count
}

func appearsTwiceNoOverlap(str string) bool {
	// seenPairs := make(map[string][2]int) // [idx, idx]
	var prefixStr string

	for i := 1; i < len(str); i++ {
		char0 := string(str[i-1])
		char1 := string(str[i])

		if len(prefixStr) > 0 && strings.Contains(prefixStr[:i-1], char0+char1) {
			return true
		}

		// Add both chars only for the first loop
		// for i > 2, char0 is an overlap character
		if i == 1 {
			prefixStr += char0
		}
		prefixStr += char1

	}

	return false
}

func charBetween(str string) bool {

	for i := 2; i < len(str); i++ {
		if str[i-2] == str[i] {
			return true
		}
	}

	return false

}

func countNiceStringsNewModel(stringArr []string) int {
	var count int
	for _, str := range stringArr {
		if appearsTwiceNoOverlap(str) && charBetween(str) {
			count++
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

	stringArr := strings.Split(string(data), "\n")
	ct := countNiceStrings(stringArr)
	fmt.Println(ct)

	ctNew := countNiceStringsNewModel(stringArr)
	fmt.Println(ctNew)

}
