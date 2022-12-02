package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	top := []int{0, 0, 0}

	scanner := bufio.NewScanner(os.Stdin)
	eof := false

	cur := 0
	for !eof {
		eof = !scanner.Scan()
		text := ""
		if !eof {
			text = scanner.Text()
		}
		eoe := text == ""

		if eoe {
			if cur > top[0] {
				top[0] = cur
			}
			sort.Ints(top)
			cur = 0
			continue
		}

		cal, err := strconv.Atoi(text)
		if err != nil {
			panic("not a number")
		}

		cur += cal
	}

	sum := 0
	for _, cal := range top {
		sum += cal
	}

	fmt.Println(sum)
}
