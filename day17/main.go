package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/kevin-kho/aoc-utilities/common"
)

func getContainers(data []byte) ([]int, error) {
	var res []int

	for entry := range bytes.SplitSeq(data, []byte{'\n'}) {

		size, err := strconv.Atoi(string(entry))
		if err != nil {
			return res, err
		}

		res = append(res, size)

	}

	return res, nil
}

func getCombinations(containers []int) int {
	var res int

	var recurse func(i, curr int)

	recurse = func(i, curr int) {
		// exit condition
		if curr >= 150 {
			if curr == 150 {
				res++
			}
			return
		}

		// exit condition: oob
		if !(i < len(containers)) {
			return
		}

		// take
		recurse(i+1, curr+containers[i])

		// no take
		recurse(i+1, curr)

	}
	recurse(0, 0)

	return res
}

func getMinNumWays(containers []int) int {

	containerWays := make(map[int]int)

	var recurse func(i, curr, ct int)
	recurse = func(i, curr, ct int) {
		// exit condition
		if curr >= 150 {
			if curr == 150 {
				containerWays[ct]++
			}
			return
		}

		// exit condition: oob
		if !(i < len(containers)) {
			return
		}

		// take
		recurse(i+1, curr+containers[i], ct+1)

		// no take
		recurse(i+1, curr, ct)

	}
	recurse(0, 0, 0)

	minContainers := math.MaxInt

	for con, _ := range containerWays {
		minContainers = min(minContainers, con)
	}

	return containerWays[minContainers]

}

func main() {
	filePath := "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	containers, err := getContainers(data)
	if err != nil {
		log.Fatal(err)
	}

	res := getCombinations(containers)
	fmt.Println(res)

	res2 := getMinNumWays(containers)
	fmt.Println(res2)

}
