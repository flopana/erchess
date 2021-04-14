package commands

import (
	"erchess/src/commands/position"
	"fmt"
)

func ReceiveArgumentsAndStartAccordingly(arguments []string){
	switch arguments[0] {
	case "uci" :
		fmt.Println("uciok")
	case "isready":
		fmt.Println("readyok")
	case "ucinewgame":
		fmt.Println("isready")
	case "position":
		position.Fen(removeElementFromSlice(arguments, 0))
	default:
		fmt.Println("Unknown command "+arguments[0])

	}
}

func removeElementFromSlice(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
