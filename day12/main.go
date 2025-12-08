package main

import (
	"aoc-2015/common"
	"fmt"
	"log"
	"unicode"
)

func getNumbers(data []byte) []int {
	var res []int
	var num int
	var negative bool
	for _, char := range string(data) {
		if string(char) == "-" {
			negative = true
		} else if unicode.IsDigit(char) {
			num = num*10 + int(char-48)
		} else if num != 0 {
			if negative {
				num = -num
			}
			res = append(res, num)
			num = 0
			negative = false
		}
	}

	if num != 0 {
		res = append(res, num)
	}

	return res
}

func sumInts(slc []int) int {
	var res int
	for _, val := range slc {
		res += val
	}

	return res

}

func main() {

	filePath := "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	data = common.TrimNewLineSuffix(data)

	nums := getNumbers(data)
	sum := sumInts(nums)

	fmt.Println(sum)

}
