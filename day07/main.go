package main

import (
	"aoc/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1")
	part1()

	fmt.Println("\n\nPart 2")
	part2()
}

func part1() {
	input := util.NewInput("inputs/input.txt")

	total := 0
	for _, line := range input.Lines {

		parts := strings.Split(line, ":")
		solution, _ := strconv.Atoi(parts[0])
		components := []int{}
		for _, c := range strings.Fields(parts[1]) {
			i, _ := strconv.Atoi(c)
			components = append(components, i)
		}

		operationsList := allPermutations(len(components)-1, "+*")
		for _, list := range operationsList {
			operations := strings.Split(list, "")

			i := runPostfix(components, operations)
			if i == solution {
				fmt.Printf("%d == %v %v\n", solution, components, operations)
				total += solution
				break
			}
		}

	}

	fmt.Println(total)
}

func part2() {
	input := util.NewInput("inputs/input.txt")

	total := 0
	for _, line := range input.Lines {

		parts := strings.Split(line, ":")
		solution, _ := strconv.Atoi(parts[0])
		components := []int{}
		for _, c := range strings.Fields(parts[1]) {
			i, _ := strconv.Atoi(c)
			components = append(components, i)
		}

		operationsList := allPermutations(len(components)-1, "+*&") // using & instead of || to fit in single char
		for _, list := range operationsList {
			operations := strings.Split(list, "")

			i := runPostfix(components, operations)

			if i == solution {
				fmt.Printf("%d == %v %v\n", i, components, operations)
				total += solution
				break
			}
		}

	}

	fmt.Println(total)
}

func runPostfix(values []int, operations []string) int {

	if len(values) == 1 {
		return values[0]
	} else {
		n := 0
		if operations[0] == "+" {
			n = values[0] + values[1]
		} else if operations[0] == "*" {
			n = values[0] * values[1]
		} else {
			nStr := strconv.Itoa(values[0]) + strconv.Itoa(values[1])
			n, _ = strconv.Atoi(nStr)
		}

		newValues := []int{
			n,
		}

		newValues = append(newValues, values[2:]...)

		return runPostfix(newValues, operations[1:])
	}

}

func allPermutations(length int, commands string) []string {
	operations := []string{}
	perm := make([]string, length)
	permute(perm, 0, commands, &operations)

	return operations
}

func permute(perm []string, pos int, values string, operations *[]string) {
	if pos == len(perm) {
		*operations = append(*operations, strings.Join(perm, ""))
	} else {
		for i := 0; i < len(values); i++ {
			perm[pos] = string(values[i])
			permute(perm, pos+1, values, operations)
		}
	}
}
