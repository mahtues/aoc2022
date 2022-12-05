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

	result := 0

	scanner := bufio.NewScanner(os.Stdin)
	for {
		items := map[int]struct{}{}
		for p := 1; p < 53; p++ {
			items[p] = struct{}{}
		}

		for e := 0; e < 3; e++ {
			if !scanner.Scan() {
				if e != 0 {
					panic("invalid number of lines")
				}
				fmt.Println(result)
				return
			}

			line := scanner.Text()

			nitems := map[int]struct{}{}
			for i := 0; i < len(line); i++ {
				p := priority(line[i])
				if _, found := items[p]; found {
					nitems[p] = struct{}{}
				}
			}
			items = nitems
		}

		if len(items) != 1 {
			fmt.Println(items)
			panic("bad file or implementation")
		}

		for p, _ := range items {
			result += p
		}
	}
}
