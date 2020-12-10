package helpers

import (
	"reflect"
	"testing"
)

func TestProcessInput(t *testing.T) {
	got, _ := ProcessInput("helpers.input")
	want := []string{"lorem", "ipsum", "dolor", "sit", "amet"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestArrayContainsStr(t *testing.T) {
	got1, got2 := ArrayContainsStr([]string{"lorem", "ipsum"}, "ipsum"), ArrayContainsStr([]string{"lorem", "ipsum"}, "x")
	want1, want2 := true, false

	if got1 != want1 || got2 != want2 {
		t.Errorf("got %t, %t - want %t, %t", got1, got2, want1, want2)
	}
}

func TestArrayContainsInt(t *testing.T) {
	got1, got2 := ArrayContainsInt([]int{4, 2}, 2), ArrayContainsInt([]int{4, 2}, 7)
	want1, want2 := true, false

	if got1 != want1 || got2 != want2 {
		t.Errorf("got %t, %t - want %t, %t", got1, got2, want1, want2)
	}
}

func TestConvertStringsToNumbers(t *testing.T) {
	got := ConvertStringsToNumbers([]string{"1", "24", "-13", "7"})
	want := []int{1, 24, -13, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
