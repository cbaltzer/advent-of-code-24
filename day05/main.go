package main

import (
	"aoc/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := util.NewInput("inputs/input.txt")

	rules := map[string][]string{}

	p1total := 0
	p2total := 0
	for _, line := range input.Lines {

		ruleSplit := strings.Split(line, "|")
		if len(ruleSplit) == 2 {
			k := ruleSplit[0]
			v := ruleSplit[1]
			rules[k] = append(rules[k], v)
		}

		updateSplit := strings.Split(line, ",")
		if len(updateSplit) > 1 {

			valid := true
			for i, v := range updateSplit {
				for _, n := range rules[v] {
					next := slices.Index(updateSplit, n)
					if next != -1 && next < i {
						valid = false
					}
				}
			}

			if valid {
				midpoint := len(updateSplit) / 2
				mid, _ := strconv.Atoi(updateSplit[midpoint])

				p1total += mid
			} else {
				slices.SortFunc(updateSplit, func(a, b string) int {
					if slices.Contains(rules[a], b) {
						return -1
					}

					return 0
				})

				midpoint := len(updateSplit) / 2
				mid, _ := strconv.Atoi(updateSplit[midpoint])

				p2total += mid
			}

		}

	}

	fmt.Println("Part 1")
	fmt.Println(p1total)

	fmt.Println("\n\nPart 2")
	fmt.Println(p2total)
}
