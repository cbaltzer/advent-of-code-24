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
	input := util.NewInput("inputs/input1.txt")

	totalSafe := 0
	for _, report := range input.Lines {
		levels := strings.Fields(report)

		safe := checkSafety(levels)

		if safe {
			totalSafe++
		}
	}
	fmt.Println(totalSafe)
}

func part2() {
	input := util.NewInput("inputs/input1.txt")

	totalSafe := 0
	for _, report := range input.Lines {
		allLevels := strings.Fields(report)

		allSafe := checkSafety(allLevels)
		if allSafe {
			totalSafe++
			continue
		} else {

			for i := range allLevels {
				levels := []string{}
				levels = append(levels, allLevels[:i]...)
				levels = append(levels, allLevels[i+1:]...)

				safe := checkSafety(levels)
				if safe {
					totalSafe++
					break
				}
			}

		}

	}
	fmt.Println(totalSafe)
}

func checkSafety(levels []string) bool {

	safe := true
	delta := 0

	for i := 0; i < len(levels)-1; i++ {

		l, _ := strconv.Atoi(levels[i])
		r, _ := strconv.Atoi(levels[i+1])

		diff := max(l, r) - min(l, r)
		if diff < 1 || diff > 3 {
			safe = false
			continue
		}

		if l > r {
			delta -= 1
		} else {
			delta += 1
		}
	}

	if delta < 0 {
		delta *= -1
	}

	if delta != len(levels)-1 {
		safe = false
	}

	return safe
}
