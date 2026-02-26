package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Pos struct {
	X int
	Y int
}

type Command struct {
	Start     Pos
	End       Pos
	Operation string
}

func parseInstructions(byteArr []byte) ([]Command, error) {

	// remove trailing whitespace at end of file
	byteArr, _ = bytes.CutSuffix(byteArr, []byte{'\n'})

	lines := strings.SplitSeq(string(byteArr), "\n")

	var commands []Command
	for line := range lines {
		cmd, err := parseLine(line)
		if err != nil {
			return commands, err
		}
		commands = append(commands, *cmd)
	}

	return commands, nil
}

func parseLine(line string) (*Command, error) {

	var start []string
	var end []string
	var operation string
	splitLine := strings.Split(line, " ")

	// turn on/off
	if len(splitLine) == 5 {
		start = strings.Split(splitLine[2], ",")
		end = strings.Split(splitLine[4], ",")
		operation = strings.ToUpper(splitLine[1])

	}

	// toggle
	if len(splitLine) == 4 {
		start = strings.Split(splitLine[1], ",")
		end = strings.Split(splitLine[3], ",")
		operation = strings.ToUpper(splitLine[0])
	}
	sx, err := strconv.Atoi(start[0])
	sy, err := strconv.Atoi(start[1])
	ex, err := strconv.Atoi(end[0])
	ey, err := strconv.Atoi(end[1])
	if err != nil {
		return nil, err
	}

	cmd := Command{
		Start: Pos{
			X: sx,
			Y: sy,
		},
		End: Pos{
			X: ex,
			Y: ey,
		},
		Operation: operation,
	}

	return &cmd, nil
}

func lightStrategyPartOne(cmd Command, originalValue int) int {
	var newValue int
	switch cmd.Operation {
	case "ON":
		newValue = 1
	case "OFF":
		newValue = 0
	default:
		newValue = originalValue ^ 1
	}

	return newValue
}

func lightStrategyPartTwo(cmd Command, originalValue int) int {
	var newValue int = originalValue
	switch cmd.Operation {
	case "ON":
		newValue++
	case "OFF":
		newValue--
		newValue = max(0, newValue)
	default:
		newValue += 2
	}
	return newValue
}

func operateLights(commands []Command, grid [][]int, strategy func(cmd Command, originalValue int) int) {
	for _, cmd := range commands {

		// For all Commands, they follow these rules
		// (cmd.Start.X <= cmd.End.X) && (cmd.Start.Y <= cmd.End.Y)
		// Build a rectangle

		for x := cmd.Start.X; x < cmd.End.X+1; x++ {
			for y := cmd.Start.Y; y < cmd.End.Y+1; y++ {
				val := grid[y][x]
				newValue := strategy(cmd, val)
				grid[y][x] = newValue
			}
		}
	}
}

func countLights(grid [][]int) int {
	var count int
	for _, row := range grid {
		for _, light := range row {
			count += light
		}
	}

	return count
}

func constructGrid() [][]int {
	var grid [][]int
	for range 1000 {
		row := make([]int, 1000)
		grid = append(grid, row)
	}

	return grid

}

func main() {

	partOneGrid := constructGrid()
	partTwoGrid := constructGrid()

	filePath := "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	cmds, err := parseInstructions(data)
	if err != nil {
		log.Fatal(err)
	}

	operateLights(cmds, partOneGrid, lightStrategyPartOne)
	litLightPartOne := countLights(partOneGrid)
	fmt.Println(litLightPartOne)

	operateLights(cmds, partTwoGrid, lightStrategyPartTwo)
	litLightsPartTwo := countLights(partTwoGrid)
	fmt.Println(litLightsPartTwo)

}
