package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	max, cur := 0, 0

	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		line := scanner.Text()

		if line == "" {
			cur = 0
			continue
		}

		cal, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("not a number")
		}

		cur += cal

		if cur > max {
			max = cur
		}
	}

	fmt.Println(max)
}
