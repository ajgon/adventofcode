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
