package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T)  {
	res := parseInput("position fen 123/321")
	expected := []string{"position", "fen", "123/321"}
	if !reflect.DeepEqual(res, expected){
		t.Errorf("Arrays didn't equal expected: %s\nWas: %s", res, expected)
	}
}