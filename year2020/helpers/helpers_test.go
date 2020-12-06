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

func TestArrayContainsSuccess(t *testing.T) {
	got1, got2 := ArrayContains([]string{"lorem", "ipsum"}, "ipsum"), ArrayContains([]string{"lorem", "ipsum"}, "x")
	want1, want2 := true, false

	if got1 != want1 || got2 != want2 {
		t.Errorf("got %t, %t - want %t, %t", got1, got2, want1, want2)
	}
}
