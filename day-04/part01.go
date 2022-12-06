package main

import (
	"bufio"
	"fmt"
	"os"
)

type Assignment struct {
	Start int
	End   int
}

func main() {

	r := 0
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		line := scanner.Text()

		var a Assignment
		var b Assignment

		fmt.Sscanf(line, "%d-%d,%d-%d", &a.Start, &a.End, &b.Start, &b.End)

		// longest length on a
		if a.End-a.Start < b.End-b.Start {
			a, b = b, a
		}

		if a.Start <= b.Start && b.End <= a.End {
			r += 1
		}
	}

	fmt.Println(r)
}
