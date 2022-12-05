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

const WIN = 1
const DRAW = 2
const LOSS = 3

var ME = map[string]int{
	"X": LOSS,
	"Y": DRAW,
	"Z": WIN,
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

var LOSES = map[int]int{
	ROCK:     PAPER,
	PAPER:    SCISSORS,
	SCISSORS: ROCK,
}

func calcMyScore(other int, outcome int) int {
	switch outcome {
	case LOSS:
		return WINS[other]
	case DRAW:
		return 3 + other
	case WIN:
		return 6 + LOSES[other]
	}

	panic("invalid outcome value")
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
