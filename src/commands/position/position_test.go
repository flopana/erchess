package position

import (
	"reflect"
	"testing"
)

func TestMakeBasicBoard(t *testing.T) {
	was := makeBasicBoard()
	expected := Board{
		Field:      [][]Piece{
			{
				{},{},{},{},{},{},{},{},
			},{
				{},{},{},{},{},{},{},{},
			},{
				{},{},{},{},{},{},{},{},
			},{
				{},{},{},{},{},{},{},{},
			},{
				{},{},{},{},{},{},{},{},
			},{
				{},{},{},{},{},{},{},{},
			},{
				{},{},{},{},{},{},{},{},
			},{
				{},{},{},{},{},{},{},{},
			},
		},
		EnpassantX: 0,
		EnpassantY: 0,
		ToMove:     0,
	}

	if !reflect.DeepEqual(was, expected) {
		t.Errorf("Arrays didn't equal for board.Field")
	}
}