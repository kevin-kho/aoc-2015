package main

import (
	"fmt"
	"testing"
)

func TestAppearsTwiceNoOverlap(t *testing.T) {
	res := appearsTwiceNoOverlap("aaa")
	res2 := appearsTwiceNoOverlap("xyxy")
	res3 := appearsTwiceNoOverlap("aabcdefgaa")
	res4 := appearsTwiceNoOverlap("aaaa")
	fmt.Println(res)
	fmt.Println(res2)
	fmt.Println(res3)
	fmt.Println(res4)

}

func TestCharBetween(t *testing.T) {
	res := charBetween("xyx")
	res2 := charBetween("abcdefeghi")
	res3 := charBetween("aaa")
	res4 := charBetween("xxyxx")
	res5 := charBetween("aaaa")

	fmt.Println(res)
	fmt.Println(res2)
	fmt.Println(res3)
	fmt.Println(res4)
	fmt.Println(res5)

}

func TestCountNiceStringsNewModel(t *testing.T) {
	res := countNiceStringsNewModel([]string{"qjhvhtzxzqqjkmpb"})
	fmt.Println(res)

	res2 := countNiceStringsNewModel([]string{"xxyxx"})
	fmt.Println(res2)

	res3 := countNiceStringsNewModel([]string{"uurcxstgmygtbstg"})
	fmt.Println(res3)

	res4 := countNiceStringsNewModel([]string{"ieodomkazucvgmuy"})
	fmt.Println(res4)

	res5 := countNiceStringsNewModel([]string{"aaaa"})
	fmt.Println(res5)

}
