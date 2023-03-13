package main

import (
	"fmt"
	"strings"
)

func getInput() string {
	return `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
}

func parseInput() {
	parts := strings.Split(getInput(), "\n\n")
	drawing := parts[0]
	drawing_lines := strings.Split(drawing, "\n")
	stacks := make([][]string, 9)

	for _, stack := range stacks {
		stack = append(stack, "sss")
		fmt.Println(stack)
		fmt.Println(stacks)
	}

	for _, line := range drawing_lines {
		// fmt.Println(len(line))
		// fmt.Println(crates)
		for i := 1; i < len(line); i += 4 {
			if string(line[i]) != " " {
				// stacks[i] = append(stacks[i], string(line[i]))
				fmt.Println(string(line[i]))
			}
		}
	}
	fmt.Println(stacks)
}

func main() {
	parseInput()
}
