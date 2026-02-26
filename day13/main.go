package main

import (
	"bytes"
	"cmp"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Entry struct {
	Subject string
	Value   int
	Target  string
}

func createEntries(data []byte) ([]Entry, error) {

	var res []Entry

	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {
		entryStr := strings.Split(string(entry), " ")

		subject := entryStr[0]
		target := entryStr[len(entryStr)-1]
		target = strings.TrimSuffix(target, ".")
		value, err := strconv.Atoi(entryStr[3])
		if err != nil {
			return res, err
		}

		if entryStr[2] == "lose" {
			value = -value
		}

		res = append(res, Entry{
			Subject: subject,
			Value:   value,
			Target:  target,
		})

	}

	return res, nil

}

func createAdjMatrix(entries []Entry) map[string]map[string]int {

	adj := make(map[string]map[string]int)

	for _, entry := range entries {
		if _, ok := adj[entry.Subject]; !ok {
			adj[entry.Subject] = map[string]int{}
		}
		adj[entry.Subject][entry.Target] = entry.Value
	}

	return adj

}

func solvePartOne(entries []Entry, adj map[string]map[string]int) int {

	// Greedily select people to sit next to
	// Doesn't work.
	slices.SortFunc(entries, func(a, b Entry) int {
		return cmp.Compare(a.Value, b.Value)
	})
	slices.Reverse(entries)

	seated := map[string][]string{}
	for _, e := range entries {
		if len(seated[e.Subject]) < 2 && len(seated[e.Target]) < 2 {
			seated[e.Subject] = append(seated[e.Subject], e.Target)
			seated[e.Target] = append(seated[e.Target], e.Subject)
		}
	}

	totalHappiness := calculateTotalHappiness(seated, adj)

	return totalHappiness

}

func calculateTotalHappiness(seated map[string][]string, adj map[string]map[string]int) int {
	var total int
	for subject, neighbors := range seated {
		for _, n := range neighbors {
			total += adj[subject][n]
		}
	}

	return total

}

func main() {
	filePath := "./inputExample.txt"
	filePath = "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	data = common.TrimNewLineSuffix(data)

	entries, err := createEntries(data)
	if err != nil {
		log.Fatal(err)
	}

	adj := createAdjMatrix(entries)

	res := solvePartOne(entries, adj)
	fmt.Println(res)
}
