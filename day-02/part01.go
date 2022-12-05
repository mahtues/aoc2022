package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const ROCK = 1
const PAPER = 2
const SCISSORS = 3

var ME = map[string]int{
	"X": ROCK,
	"Y": PAPER,
	"Z": SCISSORS,
}

var OTHER = map[string]int{
	"A": ROCK,
	"B": PAPER,
	"C": SCISSORS,
}

var WINS = map[int]int{
	ROCK:     SCISSORS,
	PAPER:    ROCK,
	SCISSORS: PAPER,
}

func calcMyScore(other int, me int) int {
	if other == me {
		return 3 + me
	}

	if WINS[me] == other {
		return 6 + me
	}

	return me
}

func main() {

	score := 0

	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		line := scanner.Text()
		parts := strings.Fields(line)
		other, me := OTHER[parts[0]], ME[parts[1]]

		score += calcMyScore(other, me)
	}

	fmt.Println(score)
}
