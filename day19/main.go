package main

import (
	"bytes"
	"fmt"
	"log"
	"slices"
	"strings"
	"unicode"

	"github.com/kevin-kho/aoc-utilities/common"
)

func parseData(data []byte) (map[string][]string, string) {

	entries := bytes.SplitSeq(data, []byte{'\n'})

	// Build substition map
	substitions := make(map[string][]string)
	for entry := range entries {
		if len(entry) == 0 {
			break
		}
		entryStr := strings.Split(string(entry), " ")
		element := entryStr[0]
		sub := entryStr[len(entryStr)-1]

		substitions[element] = append(substitions[element], sub)
	}

	var molecule string
	for entry := range entries {
		if len(entry) == 0 {
			continue
		}
		molecule = string(entry)
	}

	return substitions, molecule

}

func createElementList(mol string) []string {
	var res []string

	for i, char := range mol {

		if unicode.IsLower(char) {
			continue
		}

		if i == len(mol)-1 {
			res = append(res, string(char))
			break
		}

		if unicode.IsLower(rune(mol[i+1])) {
			res = append(res, string(mol[i:i+2]))

		} else {
			res = append(res, string(char))
		}

	}

	return res

}

func solvePartOne(subs map[string][]string, eleList []string) int {

	seen := make(map[string]bool)
	for i, ele := range eleList {
		e := slices.Clone(eleList)
		for _, sub := range subs[ele] {
			e[i] = sub
			seen[strings.Join(e, "")] = true
		}
	}

	return len(seen)
}

func main() {

	filePath := "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	subs, mol := parseData(data)

	eleList := createElementList(mol)

	res := solvePartOne(subs, eleList)
	fmt.Println(res)

}
