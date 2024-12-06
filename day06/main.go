package main

import (
	"aoc/util"
	"fmt"
	"strings"
)

var grid [][]string

type Guard struct {
	y           int
	x           int
	orientation string
}

func (g *Guard) checkNext() string {
	switch g.orientation {
	case "^":
		if g.y == 0 {
			return "DONE"
		}
		return grid[g.y-1][g.x]
	case ">":
		if g.x == len(grid[0])-1 {
			return "DONE"
		}
		return grid[g.y][g.x+1]
	case "v":
		if g.y == len(grid)-1 {
			return "DONE"
		}
		return grid[g.y+1][g.x]
	case "<":
		if g.x == 0 {
			return "DONE"
		}
		return grid[g.y][g.x-1]
	}
	return "DONE"
}

func (g *Guard) castRight() bool {
	castX := 0
	castY := 0
	switch g.orientation {
	case "^":
		castX = 1
	case ">":
		castY = 1
	case "v":
		castX = -1
	case "<":
		castY = -1
	}

	x := g.x
	y := g.y

	fmt.Printf("\n-----\nGuard: %v\n", g)
	for x > 0 && x < len(grid[0]) && y > 0 && y < len(grid) {

		fmt.Printf("Cast: [%d][%d]\n", y, x)

		check := grid[y][x]
		if check == "R" {
			return true
		}

		x += castX
		y += castY
	}

	return false
}

func (g *Guard) advance() {
	switch g.orientation {
	case "^":
		if g.y != 0 {
			g.y -= 1
		}
	case ">":
		if g.x != len(grid[0])-1 {
			g.x += 1
		}
	case "v":
		if g.y != len(grid)-1 {
			g.y += 1
		}
	case "<":
		if g.x != 0 {
			g.x -= 1
		}
	}
}

func (g *Guard) mark() {
	if grid[g.y][g.x] != "R" {
		grid[g.y][g.x] = "X"
	}
}

func (g *Guard) markTurn() {
	grid[g.y][g.x] = "R"
}

func (g *Guard) turn() {
	switch g.orientation {
	case "^":
		g.orientation = ">"
	case ">":
		g.orientation = "v"
	case "v":
		g.orientation = "<"
	case "<":
		g.orientation = "^"
	}
}

func main() {
	fmt.Println("Part 1")
	part1()

	fmt.Println("\n\nPart 2")
	part2()
}

func part1() {
	input := util.NewInput("inputs/input.txt")

	grid = [][]string{}
	for _, line := range input.Lines {
		grid = append(grid, strings.Split(line, ""))
	}

	var guard Guard
	for y, row := range grid {
		for x, val := range row {

			//fmt.Printf("%s", val)
			if val == "^" || val == ">" || val == "v" || val == "<" {

				guard = Guard{
					y:           y,
					x:           x,
					orientation: val,
				}

			}
		}
		//fmt.Printf("\n")
	}

	for guard.checkNext() != "DONE" {
		guard.mark()
		next := guard.checkNext()

		if next == "#" {
			guard.turn()
		}
		guard.advance()
	}

	marked := 1
	for _, row := range grid {
		//fmt.Println(row)
		for _, v := range row {
			if v == "X" {
				marked += 1
			}
		}
	}

	fmt.Println(marked)
}

func part2() {
	input := util.NewInput("inputs/sample.txt")

	grid = [][]string{}
	for _, line := range input.Lines {
		grid = append(grid, strings.Split(line, ""))
	}

	var guard Guard
	for y, row := range grid {
		for x, val := range row {

			//fmt.Printf("%s", val)
			if val == "^" || val == ">" || val == "v" || val == "<" {

				guard = Guard{
					y:           y,
					x:           x,
					orientation: val,
				}

			}
		}
		//fmt.Printf("\n")
	}

	loopCount := 0
	for {
		guard.mark()
		next := guard.checkNext()

		if guard.castRight() {
			loopCount += 1
		}

		if next == "#" {
			guard.turn()
			guard.markTurn()

		}

		if next != "DONE" {
			guard.advance()
		} else {
			break
		}
	}

	marked := 1
	for _, row := range grid {
		//fmt.Println(row)
		for _, v := range row {
			if v == "X" {
				marked += 1
			}
		}
	}

	fmt.Println(loopCount)
}
