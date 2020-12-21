package day21

import (
	"sort"
	"strconv"
	"strings"
	"year2020/helpers"
)

var input, _ = helpers.ProcessInput("day21.input")

func mapAllergensCandidates(data []string) (results map[string][][]string) {
	results = make(map[string][][]string)
	for i := 0; i < len(data); i++ {
		productsAllergens := strings.Split(data[i], " (contains ")
		products := strings.Split(productsAllergens[0], " ")
		allergens := strings.Split(productsAllergens[1][0:len(productsAllergens[1])-1], ", ")

		for j := 0; j < len(allergens); j++ {
			results[allergens[j]] = append(results[allergens[j]], products)
		}
	}

	return
}

func narrowDownAllergens(candidates map[string][][]string) (results map[string][]string) {
	results = make(map[string][]string)

	for allergen, products := range candidates {
		suspectedProducts := products[0]
		for i := 1; i < len(products); i++ {
			suspectedProducts = helpers.ArrayIntersectStr(suspectedProducts, products[i])
		}
		results[allergen] = suspectedProducts
	}

	return
}

func getAllProducts(data []string) (allProducts []string) {
	for i := 0; i < len(data); i++ {
		productsAllergens := strings.Split(data[i], " (contains ")
		products := strings.Split(productsAllergens[0], " ")

		for j := 0; j < len(products); j++ {
			allProducts = append(allProducts, products[j])
		}
	}

	return
}

func mapProductToAllergen(candidates map[string][][]string) (allergenMap map[string]string) {
	allergenMap = make(map[string]string)
	narrowedAllergens := narrowDownAllergens(candidates)

	for {
		done := true
		for allergen, products := range narrowedAllergens {
			if len(products) == 1 {
				done = false
				allergenMap[allergen] = products[0]
				for a, p := range narrowedAllergens {
					narrowedAllergens[a] = helpers.ArrayExceptStr(p, []string{allergenMap[allergen]})
				}
			}
		}
		if done {
			break
		}
	}

	return
}

func SimpleSolution() string {
	allProducts := getAllProducts(input)
	suspectedProductsMap := narrowDownAllergens(mapAllergensCandidates(input))
	suspectedProducts := make([]string, 0)

	for _, products := range suspectedProductsMap {
		for p := 0; p < len(products); p++ {
			if !helpers.ArrayContainsStr(suspectedProducts, products[p]) {
				suspectedProducts = append(suspectedProducts, products[p])
			}
		}
	}

	return strconv.Itoa(len(helpers.ArrayExceptStr(allProducts, suspectedProducts)))
}

func AdvancedSolution() string {
	results := make([]string, 0)
	keys := make([]string, 0)

	candidates := mapAllergensCandidates(input)
	allergenProducts := mapProductToAllergen(candidates)

	for key := range allergenProducts {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for k := 0; k < len(keys); k++ {
		results = append(results, allergenProducts[keys[k]])
	}

	return strings.Join(results, ",")
}
