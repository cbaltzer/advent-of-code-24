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

	topography := make([][]int, len(input.Lines))

	for y, line := range input.Lines {
		chars := strings.Split(line, "")
		topography[y] = make([]int, len(chars))

		for x, c := range chars {
			i, _ := strconv.Atoi(c)
			topography[y][x] = i
		}
	}

	totalScore := 0
	for y, row := range topography {
		fmt.Println(row)
		for x, s := range row {

			if s == 0 {
				score := trailSearch(topography, y, x)
				totalScore += score
				//fmt.Printf("Score(%d,%d) = %d\n\n", y, x, score)
			}

		}
	}

	fmt.Println(totalScore)
}

func part2() {
	input := util.NewInput("inputs/input.txt")

	topography := make([][]int, len(input.Lines))

	for y, line := range input.Lines {
		chars := strings.Split(line, "")
		topography[y] = make([]int, len(chars))

		for x, c := range chars {
			i, _ := strconv.Atoi(c)
			topography[y][x] = i
		}
	}

	totalScore := 0
	for y, row := range topography {
		fmt.Println(row)
		for x, s := range row {

			if s == 0 {
				score := trailSearchAll(topography, y, x)
				totalScore += score
			}

		}
	}

	fmt.Println(totalScore)
}

func trailSearch(grid [][]int, y, x int) int {

	reachable := map[string]int{}
	search(grid, y, x, &reachable)

	return len(reachable)
}

func trailSearchAll(grid [][]int, y, x int) int {

	reachable := map[string]int{}
	search(grid, y, x, &reachable)

	total := 0
	for _, v := range reachable {
		total += v
	}

	return total
}

func search(grid [][]int, y, x int, visited *map[string]int) {

	current := grid[y][x]

	if current == 9 {
		coord := fmt.Sprintf("%dx%d", y, x)
		(*visited)[coord]++
	}

	//fmt.Printf("Search(%d,%d)[%d] := %d\n", y, x, current, len(*visited))

	// up
	if y-1 >= 0 && grid[y-1][x] == current+1 {
		search(grid, y-1, x, visited)
	}

	// down
	if y+1 < len(grid) && grid[y+1][x] == current+1 {
		search(grid, y+1, x, visited)
	}

	// left
	if x-1 >= 0 && grid[y][x-1] == current+1 {
		search(grid, y, x-1, visited)
	}

	// left
	if x+1 < len(grid[0]) && grid[y][x+1] == current+1 {
		search(grid, y, x+1, visited)
	}
}
