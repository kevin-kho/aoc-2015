package main

import (
	"aoc-2015/common"
	"bytes"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

type Ingredient struct {
	Name       string
	Capacity   int
	Durability int
	Flavor     int
	Texture    int
	Calories   int
}

func parseIngredients(data []byte) ([]Ingredient, error) {

	var res []Ingredient

	for entry := range bytes.SplitSeq(data, []byte{10}) {

		entryStr := strings.Split(string(entry), " ")

		name := entryStr[0]
		name, _ = strings.CutSuffix(name, ":")

		capacity := entryStr[2]
		capacity, _ = strings.CutSuffix(capacity, ",")
		capacityInt, err := strconv.Atoi(capacity)
		if err != nil {
			return res, err
		}

		durability := entryStr[4]
		durability, _ = strings.CutSuffix(durability, ",")
		durabilityInt, err := strconv.Atoi(durability)
		if err != nil {
			return res, err
		}

		flavor := entryStr[6]
		flavor, _ = strings.CutSuffix(flavor, ",")
		flavorInt, err := strconv.Atoi(flavor)
		if err != nil {
			return res, err
		}

		texture := entryStr[8]
		texture, _ = strings.CutSuffix(texture, ",")
		textureInt, err := strconv.Atoi(texture)
		if err != nil {
			return res, err
		}

		calories := entryStr[10]
		caloriesInt, err := strconv.Atoi(calories)
		if err != nil {
			return res, err
		}

		res = append(res, Ingredient{
			Name:       name,
			Capacity:   capacityInt,
			Durability: durabilityInt,
			Flavor:     flavorInt,
			Texture:    textureInt,
			Calories:   caloriesInt,
		})

	}

	return res, nil

}

func calcIngredientValue(in []Ingredient) int {
	var capacity int
	var durability int
	var flavor int
	var texture int

	// Ignore calories for now

	for _, i := range in {
		capacity += i.Capacity
		durability += i.Durability
		flavor += i.Flavor
		texture += i.Texture
	}

	capacity = max(0, capacity)
	durability = max(0, durability)
	flavor = max(0, flavor)
	texture = max(0, texture)

	return capacity * durability * flavor * texture

}

func calculateCalories(ingred []Ingredient) int {
	var calories int
	for _, in := range ingred {
		calories += in.Calories
	}
	return calories
}

func solvePartOne(ingreds []Ingredient) int {

	var res int

	var recurse func(i int, curr []Ingredient)
	recurse = func(i int, curr []Ingredient) {

		// case: we have 100 ingredients
		if len(curr) == 100 {
			res = max(res, calcIngredientValue(curr))
			return
		}

		// case: out of bounds
		if !(i < len(ingreds)) {
			return
		}

		// take
		recurse(i, append(curr, ingreds[i]))

		// no take
		recurse(i+1, slices.Clone(curr))
	}
	recurse(0, []Ingredient{})
	return res

}

func solvePartTwo(ingreds []Ingredient) int {
	var res int
	var recurse func(i int, curr []Ingredient)
	recurse = func(i int, curr []Ingredient) {
		// case: we have 100 ingredients
		if len(curr) == 100 {
			if calculateCalories(curr) == 500 {
				res = max(res, calcIngredientValue(curr))
			}
			return
		}

		// case: out of bounds
		if !(i < len(ingreds)) {
			return
		}

		// take
		recurse(i, append(curr, ingreds[i]))

		// no take
		recurse(i+1, slices.Clone(curr))
	}
	recurse(0, []Ingredient{})

	return res
}

func main() {
	filePath := "./inputExample.txt"
	filePath = "./input.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	data = common.TrimNewLineSuffix(data)

	ingredients, err := parseIngredients(data)
	if err != nil {
		log.Fatal(err)
	}

	res := solvePartOne(ingredients)
	fmt.Println(res)

	res2 := solvePartTwo(ingredients)
	fmt.Println(res2)

}
