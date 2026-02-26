package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Reindeer struct {
	Name         string
	Speed        int
	MoveDuration int
	RestDuration int
}

func createReindeers(data []byte) ([]Reindeer, error) {

	var res []Reindeer

	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {
		entryStr := strings.Split(string(entry), " ")

		name := entryStr[0]
		speed, err := strconv.Atoi(entryStr[3])
		if err != nil {
			return res, err
		}
		moveDuration, err := strconv.Atoi(entryStr[6])
		if err != nil {
			return res, err
		}
		restDuration, err := strconv.Atoi(entryStr[len(entryStr)-2])
		if err != nil {
			return res, err
		}

		res = append(res, Reindeer{
			Name:         name,
			Speed:        speed,
			MoveDuration: moveDuration,
			RestDuration: restDuration,
		})

	}

	return res, nil

}

func moveReindeer(r Reindeer, seconds int) []int {
	position := 0
	move := r.MoveDuration
	rest := 0
	var posArr []int
	for i := 1; i < seconds+1; i++ {

		// case: movable
		if move > 0 && rest == 0 {
			position += r.Speed
			posArr = append(posArr, position)
			move--
			if move == 0 { // finish moving
				rest = r.RestDuration
				continue
			}
		}

		// case resting
		if rest > 0 {
			posArr = append(posArr, position)
			rest--
			if rest == 0 {
				move = r.MoveDuration
			}
		}

	}

	return posArr

}

func getReindeerFinalPosition(r Reindeer, seconds int) int {
	posArr := moveReindeer(r, seconds)
	return posArr[len(posArr)-1]

}

func solvePartOne(reindeers []Reindeer, seconds int) int {
	var maxDistance int
	for _, r := range reindeers {
		maxDistance = max(maxDistance, getReindeerFinalPosition(r, seconds))
	}

	return maxDistance

}

func solvePartTwo(reindeers []Reindeer, seconds int) int {
	reindeerPos := map[string][]int{}
	for _, r := range reindeers {
		reindeerPos[r.Name] = moveReindeer(r, seconds)
	}

	maxDistanceArr := getMaxDistanceArr(reindeerPos)
	scoreboard := make(map[string]int)

	for i, maxDistance := range maxDistanceArr {
		for name, pos := range reindeerPos {
			if pos[i] == maxDistance {
				scoreboard[name]++
			}
		}
	}

	var mostPoints int
	for _, points := range scoreboard {
		mostPoints = max(mostPoints, points)
	}

	return mostPoints

}

func getMaxDistanceArr(reindeerPos map[string][]int) []int {
	var distance []int

	var size int
	for _, v := range reindeerPos {
		size = len(v)
	}

	for i := range size {
		var maxDistance int
		for k := range reindeerPos {
			maxDistance = max(maxDistance, reindeerPos[k][i])
		}
		distance = append(distance, maxDistance)

	}

	return distance

}

func main() {
	filePath := "inputExample.txt"
	filePath = "input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	data = common.TrimNewLineSuffix(data)

	reindeers, err := createReindeers(data)
	if err != nil {
		log.Fatal(err)
	}

	res := solvePartOne(reindeers, 2503)
	fmt.Println(res)

	res2 := solvePartTwo(reindeers, 2503)
	fmt.Println(res2)

}
