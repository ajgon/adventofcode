package day16

import (
	"reflect"
	"sort"
	"strings"
	"testing"
	"year2020/helpers"
)

var testSimpleData []string = strings.Split(`class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`, "\n")

var testAdvancedData []string = strings.Split(`class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9`, "\n")

func TestExtractTicketsData(t *testing.T) {
	gotRules, gotTicket, gotNearby := extractTicketsData(testSimpleData)
	wantRules, wantTicket, wantNearby := map[string][]int{
		"class": []int{1, 2, 3, 5, 6, 7},
		"row":   []int{6, 7, 8, 9, 10, 11, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44},
		"seat":  []int{13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 45, 46, 47, 48, 49, 50},
	}, []int{7, 1, 14}, [][]int{[]int{7, 3, 47}, []int{40, 4, 50}, []int{55, 2, 20}, []int{38, 6, 12}}

	if !reflect.DeepEqual(gotRules, wantRules) || !reflect.DeepEqual(gotTicket, wantTicket) || !reflect.DeepEqual(gotNearby, wantNearby) {
		t.Errorf("got %+v, %v, %v - want %+v, %v, %v", gotRules, gotTicket, gotNearby, wantRules, wantTicket, wantNearby)
	}
}

func TestFindInvalidTicketValues(t *testing.T) {
	rules, _, nearby := extractTicketsData(testSimpleData)
	got := findInvalidTicketValues(rules, nearby)
	want := []int{4, 55, 12}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFindValidNearbyTickets(t *testing.T) {
	rules, _, nearby := extractTicketsData(testSimpleData)
	got := findValidNearbyTickets(rules, nearby)
	want := [][]int{[]int{7, 3, 47}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestTranspose(t *testing.T) {
	got := transpose([][]int{[]int{1, 4}, []int{2, 5}, []int{3, 6}})
	want := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDetermineApplicableNamesFromValues01(t *testing.T) {
	rules, _, _ := extractTicketsData(testAdvancedData)
	got := determineApplicableNamesFromValues(rules, []int{3, 15, 5})
	want := []string{"row"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDetermineApplicableNamesFromValues02(t *testing.T) {
	rules, _, _ := extractTicketsData(testAdvancedData)
	got := determineApplicableNamesFromValues(rules, []int{9, 1, 14})
	sort.Strings(got)
	want := []string{"class", "row"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDetermineApplicableNamesFromValues03(t *testing.T) {
	rules, _, _ := extractTicketsData(testAdvancedData)
	got := determineApplicableNamesFromValues(rules, []int{18, 5, 9})
	sort.Strings(got)
	want := []string{"class", "row", "seat"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestNarrowPossibilities(t *testing.T) {
	got := narrowPossibilities([][]string{[]string{"row"}, []string{"class", "row"}, []string{"class", "row", "seat"}})
	want := []string{"row", "class", "seat"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSimpleSolution(t *testing.T) {
	want := helpers.Answers.Day16.Simple
	got := SimpleSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestAdvancedSolution(t *testing.T) {
	want := helpers.Answers.Day16.Advanced
	got := AdvancedSolution()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
