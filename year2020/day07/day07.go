package day07

import (
	"strconv"
	"strings"
	"year2020/helpers"
)

var input, _ = helpers.ProcessInput("day07.input")

type Bags struct {
	count int
	color string
}

func bagsContain(bags []Bags, bagColor string) bool {
	for _, bag := range bags {
		if bag.color == bagColor {
			return true
		}
	}

	return false
}

func getParentBags(data []string, bags []Bags) []Bags {
	results := make([]Bags, 0)

	for _, row := range data {
		for _, bag := range bags {
			if strings.Contains(row, bag.color) {
				bagColor := strings.Join(strings.Split(row, " ")[0:2], " ")
				if bagColor != bag.color && !bagsContain(results, bagColor) {
					results = append(results, Bags{count: 1, color: bagColor})
				}
			}
		}
	}

	return results
}

func getChildrenBags(data []string, bag Bags) []Bags {
	results := make([]Bags, 0)

	for _, row := range data {
		if strings.Index(row, bag.color) == 0 {
			bagStrings := strings.Split(strings.Split(row, " contain ")[1], ", ")
			for _, bagString := range bagStrings {
				bagData := strings.Split(bagString, " ")
				count, _ := strconv.Atoi(bagData[0])
				bag := Bags{count: count, color: strings.Join(bagData[1:3], " ")}
				if !bagsContain(results, bag.color) && count != 0 {
					results = append(results, bag)
				}
			}
		}
	}

	return results
}

func countChildrenBags(data []string, bag Bags) int {
	childrenBags := getChildrenBags(data, bag)

	if len(childrenBags) == 0 {
		return bag.count
	}

	var count int

	for _, childBag := range childrenBags {
		count += bag.count * countChildrenBags(data, childBag)
	}

	return count + bag.count
}

func SimpleSolution() string {
	var bags []Bags = make([]Bags, 0)
	var parentBags []Bags = getParentBags(input, []Bags{Bags{count: 1, color: "shiny gold"}})
	for {
		for _, bag := range parentBags {
			if !bagsContain(bags, bag.color) {
				bags = append(bags, bag)
			}
		}

		if len(parentBags) == 0 {
			break
		}

		parentBags = getParentBags(input, parentBags)
	}

	return strconv.Itoa(len(bags))
}

func AdvancedSolution() string {
	return strconv.Itoa(countChildrenBags(input, Bags{count: 1, color: "shiny gold"}) - 1)
}
