package main

import (
	"aoc/util"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Part 1")
	part1()

	fmt.Println("\n\nPart2")
	part2()
}

func part1() {
	input := util.NewInput("inputs/input.txt")

	puzzle := [][]string{}
	for _, line := range input.Lines {
		puzzle = append(puzzle, strings.Split(line, ""))
	}

	count := 0
	for y := 0; y < len(puzzle); y++ {
		for x := 0; x < len(puzzle[0]); x++ {
			if puzzle[y][x] == "X" {
				count += checkAll(puzzle, y, x)
			}
		}
	}

	fmt.Println(count)
}

func part2() {
	input := util.NewInput("inputs/input.txt")

	puzzle := [][]string{}
	for _, line := range input.Lines {
		puzzle = append(puzzle, strings.Split(line, ""))
	}

	count := 0
	for y := 1; y < len(puzzle)-1; y++ {
		for x := 1; x < len(puzzle[0])-1; x++ {
			if puzzle[y][x] == "A" {
				if checkX(puzzle, y, x) {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}

func checkX(puzzle [][]string, y int, x int) bool {
	if puzzle[y-1][x-1] == "M" && puzzle[y+1][x+1] == "S" {
		if (puzzle[y-1][x+1] == "M" && puzzle[y+1][x-1] == "S") || (puzzle[y-1][x+1] == "S" && puzzle[y+1][x-1] == "M") {
			return true
		}
	}

	if puzzle[y-1][x-1] == "S" && puzzle[y+1][x+1] == "M" {
		if (puzzle[y-1][x+1] == "M" && puzzle[y+1][x-1] == "S") || (puzzle[y-1][x+1] == "S" && puzzle[y+1][x-1] == "M") {
			return true
		}
	}

	return false
}

func checkAll(puzzle [][]string, y int, x int) int {
	matches := 0

	if checkN(puzzle, y, x) {
		matches++
	}

	if checkS(puzzle, y, x) {
		matches++
	}

	if checkE(puzzle, y, x) {
		matches++
	}

	if checkW(puzzle, y, x) {
		matches++
	}

	if checkNE(puzzle, y, x) {
		matches++
	}

	if checkNW(puzzle, y, x) {
		matches++
	}

	if checkSE(puzzle, y, x) {
		matches++
	}

	if checkSW(puzzle, y, x) {
		matches++
	}

	if matches > 0 {
		fmt.Printf("[%d][%d] FOUND %d\n", y, x, matches)
	}

	return matches
}

func checkN(puzzle [][]string, y int, x int) bool {
	if y < 3 {
		return false
	}

	if puzzle[y-1][x] == "M" && puzzle[y-2][x] == "A" && puzzle[y-3][x] == "S" {
		return true
	}

	return false
}

func checkS(puzzle [][]string, y int, x int) bool {
	if y > len(puzzle)-3 {
		return false
	}

	if puzzle[y+1][x] == "M" && puzzle[y+2][x] == "A" && puzzle[y+3][x] == "S" {
		return true
	}

	return false
}

func checkE(puzzle [][]string, y int, x int) bool {
	if x > len(puzzle[0])-4 {
		return false
	}

	if puzzle[y][x+1] == "M" && puzzle[y][x+2] == "A" && puzzle[y][x+3] == "S" {
		return true
	}

	return false
}

func checkW(puzzle [][]string, y int, x int) bool {
	if x < 3 {
		return false
	}

	if puzzle[y][x-1] == "M" && puzzle[y][x-2] == "A" && puzzle[y][x-3] == "S" {
		return true
	}

	return false
}

func checkNE(puzzle [][]string, y int, x int) bool {
	if y < 3 || x > len(puzzle[0])-4 {
		return false
	}

	if puzzle[y-1][x+1] == "M" && puzzle[y-2][x+2] == "A" && puzzle[y-3][x+3] == "S" {
		return true
	}

	return false
}

func checkNW(puzzle [][]string, y int, x int) bool {
	if y < 3 || x < 3 {
		return false
	}

	if puzzle[y-1][x-1] == "M" && puzzle[y-2][x-2] == "A" && puzzle[y-3][x-3] == "S" {
		return true
	}

	return false
}

func checkSE(puzzle [][]string, y int, x int) bool {
	if y > len(puzzle)-4 || x > len(puzzle[0])-4 {
		return false
	}

	if puzzle[y+1][x+1] == "M" && puzzle[y+2][x+2] == "A" && puzzle[y+3][x+3] == "S" {
		return true
	}

	return false
}

func checkSW(puzzle [][]string, y int, x int) bool {
	if y > len(puzzle)-4 || x < 3 {
		return false
	}

	if puzzle[y+1][x-1] == "M" && puzzle[y+2][x-2] == "A" && puzzle[y+3][x-3] == "S" {
		return true
	}

	return false
}
