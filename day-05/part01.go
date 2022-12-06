package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	Value byte
	Next  *Node
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	lines := []string{}
	for {
		if !scanner.Scan() {
			panic("invalid file")
		}
		line := scanner.Text()

		if line == "" {
			break
		}

		lines = append(lines, line)
	}

	index_to_stack := map[int]int{}
	{
		line := lines[len(lines)-1]
		for i := 0; i < len(line); i++ {
			b := line[i]
			if '1' <= b && b <= '9' {
				index_to_stack[i] = int(b - '0')
			}
		}
	}

	stacks := make([]*Node, len(index_to_stack)+1)

	for l := len(lines) - 2; l >= 0; l-- {
		line := lines[l]

		for i := 0; i < len(line); i++ {
			b := line[i]

			if 'A' <= b && b <= 'Z' {
				stack := index_to_stack[i]

				stacks[stack] = &Node{
					Value: b,
					Next:  stacks[stack],
				}
			}
		}
	}

	for scanner.Scan() {
		var q int
		var s int
		var t int

		fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &q, &s, &t)

		for i := 0; i < q; i++ {
			top := stacks[s]
			stacks[s] = stacks[s].Next

			top.Next = stacks[t]
			stacks[t] = top
		}
	}

	tops := make([]byte, 0, len(stacks)-1)
	for i := 1; i < len(stacks); i++ {
		tops = append(tops, stacks[i].Value)
	}

	fmt.Println(string(tops))
}
