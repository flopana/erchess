package commands

import (
	"reflect"
	"testing"
)

func TestRemoveFromSlice(t *testing.T) {
	arr := []string{"position", "fen", "123/321"}
	expected := []string{"fen", "123/321"}
	was := removeElementFromSlice(arr, 0)
	if !reflect.DeepEqual(was, expected) {
		t.Errorf("Wrong slice expected: %s\nwas: %s", expected, was)
	}
}