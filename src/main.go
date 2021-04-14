package main

import (
	"bufio"
	"erchess/src/commands"
	"fmt"
	"os"
	"strings"
)

func main() {
	printLogo()
	for {
		input := readInput()
		go commands.ReceiveArgumentsAndStartAccordingly(parseInput(input))
	}
}

func printLogo(){
	fmt.Println("                _                   \n   ___ _ __ ___| |__   ___  ___ ___ \n  / _ \\ '__/ __| '_ \\ / _ \\/ __/ __|\n |  __/ | | (__| | | |  __/\\__ \\__ \\\n  \\___|_|  \\___|_| |_|\\___||___/___/ erchess v0\n ")
}

func readInput() string{
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	return input
}

func parseInput(input string) []string {
	res := strings.Split(input, " ")
	return res
}