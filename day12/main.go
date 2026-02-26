package main

import (
	"fmt"
	"log"
	"strings"
	"unicode"

	"github.com/kevin-kho/aoc-utilities/common"
)

func getNumbers(data []byte) []int {
	var res []int
	var num int
	var negative bool
	for _, char := range string(data) {
		if string(char) == "-" {
			negative = true
		} else if unicode.IsDigit(char) {
			num = num*10 + int(char-'0')
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

func getObjects(data []byte) [][]byte {
	var res [][]byte
	var stack []int // index where an open brace occurs

	for i, char := range string(data) {
		if string(char) == "{" {
			stack = append(stack, i)
		} else if string(char) == "}" {

			if len(stack) == 1 { // Ensures no repeated objects
				start := stack[len(stack)-1]
				res = append(res, data[start:i+1])
			}

			stack = stack[:len(stack)-1]

		}

	}

	return res

}

func removeNestedObject(obj []byte) []byte {
	inner := string(obj[1 : len(obj)-1])
	if strings.Contains(inner, "{") && strings.Contains(inner, "}") {
		start := strings.Index(inner, "{")
		end := strings.LastIndex(inner, "}")
		left := obj[:start]
		right := obj[end+1:]

		res := append(left, right...)
		return res
	}

	return obj

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

	objs := getObjects(data)
	for _, obj := range objs {
		fmt.Println(string(obj))
		fmt.Println("---")
	}

	for _, o := range objs {
		fmt.Println(string(o))
		fmt.Println("---")
	}

}
