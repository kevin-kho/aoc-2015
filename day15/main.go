package main

import (
	"aoc-2015/common"
	"bytes"
	"fmt"
	"log"
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

func main() {
	filePath := "./inputExample.txt"
	data, err := common.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}

	data = common.TrimNewLineSuffix(data)

	ingredients, err := parseIngredients(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ingredients)

}
