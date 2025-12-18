package main

import (
	"aoc-2015/common"
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Entry struct {
	Subject string
	Value   int
	Target  string
}

func createEntries(data []byte) ([]Entry, error) {

	var res []Entry

	for entry := range bytes.SplitSeq(data, []byte{10}) {
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

func main() {
	filePath := "./inputExample.txt"
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
	fmt.Println(adj)
}
