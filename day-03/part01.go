package main

import (
	"bufio"
	"fmt"
	"os"
)

func priority(b byte) int {
	if 'a' <= b && b <= 'z' {
		return int(b-'a') + 1
	}

	return int(b-'A') + 27
}

func main() {

	score := 0
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		line := scanner.Text()

		items := map[int]struct{}{}

		for i := 0; i < len(line)/2; i++ {
			items[priority(line[i])] = struct{}{}
		}

		for i := len(line) / 2; i < len(line); i++ {
			p := priority(line[i])
			if _, found := items[p]; found {
				score += p
				break
			}
		}
	}

	fmt.Println(score)
}
