package helpers_test

import (
	"reflect"
	"testing"

	"year2021/helpers"
)

func TestProcessInput(t *testing.T) {
	t.Parallel()

	got, _ := helpers.ProcessInput("helpers.input")
	want := []string{"lorem", "ipsum", "dolor", "sit", "amet"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestConvertStringsToNumbers(t *testing.T) {
	t.Parallel()

	got := helpers.ConvertStringsToNumbers([]string{"1", "24", "-13", "7"})

	if want := []int{1, 24, -13, 7}; !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
