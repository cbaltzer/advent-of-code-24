package main

import (
	"aoc/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 1")
	part1("inputs/input1.txt")

	fmt.Println("\n\nDay 2")
	part2("inputs/input2.txt")
}

func part1(filename string) {
	input := util.NewInput(filename)

	first := []int{}
	second := []int{}

	for _, line := range input.Lines {
		parts := strings.Fields(line)

		i, _ := strconv.Atoi(parts[0])
		j, _ := strconv.Atoi(parts[1])

		first = append(first, i)
		second = append(second, j)
	}

	slices.Sort(first)
	slices.Sort(second)

	totalDistance := 0
	for i := range first {

		f := first[i]
		s := second[i]

		totalDistance += max(f, s) - min(f, s)
	}

	fmt.Println(totalDistance)
}

func part2(filename string) {
	input := util.NewInput(filename)

	left := []int{}
	right := map[int]int{}

	for _, line := range input.Lines {
		parts := strings.Fields(line)

		i, _ := strconv.Atoi(parts[0])
		j, _ := strconv.Atoi(parts[1])

		left = append(left, i)
		right[j]++
	}

	score := 0
	for _, v := range left {
		count := right[v]
		score += v * count
	}

	fmt.Println(score)
}
