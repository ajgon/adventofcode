package day21

import (
	"reflect"
	"strings"
	"testing"
	"year2020/helpers"
)

var testData []string = strings.Split(`mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`, "\n")

func TestMapAllergensCandidates(t *testing.T) {
	got := mapAllergensCandidates(testData)
	want := map[string][][]string{
		"dairy": [][]string{
			[]string{"mxmxvkd", "kfcds", "sqjhc", "nhms"},
			[]string{"trh", "fvjkl", "sbzzf", "mxmxvkd"},
		},
		"fish": [][]string{
			[]string{"mxmxvkd", "kfcds", "sqjhc", "nhms"},
			[]string{"sqjhc", "mxmxvkd", "sbzzf"},
		},
		"soy": [][]string{[]string{"sqjhc", "fvjkl"}},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestNarrowDownAllergens(t *testing.T) {
	candidates := mapAllergensCandidates(testData)

	got := narrowDownAllergens(candidates)
	want := map[string][]string{
		"dairy": []string{"mxmxvkd"},
		"fish":  []string{"mxmxvkd", "sqjhc"},
		"soy":   []string{"sqjhc", "fvjkl"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestGetAllProducts(t *testing.T) {
	got := getAllProducts(testData)
	want := []string{
		"mxmxvkd", "kfcds", "sqjhc", "nhms",
		"trh", "fvjkl", "sbzzf", "mxmxvkd",
		"sqjhc", "fvjkl",
		"sqjhc", "mxmxvkd", "sbzzf",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMapProductToAllergen(t *testing.T) {
	candidates := mapAllergensCandidates(testData)

	got := mapProductToAllergen(candidates)
	want := map[string]string{
		"dairy": "mxmxvkd",
		"fish":  "sqjhc",
		"soy":   "fvjkl",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestSimpleSolution(t *testing.T) {
	want := helpers.Answers.Day21.Simple
	got := SimpleSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestAdvancedSolution(t *testing.T) {
	want := helpers.Answers.Day21.Advanced
	got := AdvancedSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
