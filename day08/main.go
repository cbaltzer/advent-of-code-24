package main

import (
	"aoc/util"
	"fmt"
	"strings"
)

type Point struct {
	y, x int
}

func main() {
	fmt.Println("Part 1")
	part1()

	fmt.Println("\n\nPart 2")
	part2()
}

func part1() {
	input := util.NewInput("inputs/input.txt")

	antennas := map[string][]Point{}
	antinodeMap := [][]string{}
	antinodesList := map[string]int{}

	for y, line := range input.Lines {
		mapLine := []string{}

		for x, freq := range strings.Split(line, "") {
			mapLine = append(mapLine, ".")

			if freq != "." {
				p := Point{y: y, x: x}
				antennas[freq] = append(antennas[freq], p)
			}
		}

		antinodeMap = append(antinodeMap, mapLine)
	}

	maxX := len(antinodeMap[0]) - 1
	maxY := len(antinodeMap) - 1

	for _, points := range antennas {

		for i, p1 := range points {
			for j, p2 := range points {

				if i != j {
					dY := p2.y - p1.y
					dX := p2.x - p1.x

					a1 := Point{
						y: p2.y + dY,
						x: p2.x + dX,
					}

					a2 := Point{
						y: p1.y - dY,
						x: p1.x - dX,
					}

					//fmt.Printf("%s: %v ++ %v => %v | %v\n", f, p1, p2, a1, a2)

					if a1.y >= 0 && a1.y <= maxY && a1.x >= 0 && a1.x <= maxX {
						pos := fmt.Sprintf("%d,%d", a1.y, a1.x)
						antinodesList[pos]++
						antinodeMap[a1.y][a1.x] = "#"
					}

					if a2.y >= 0 && a2.y <= maxY && a2.x >= 0 && a2.x <= maxX {
						pos := fmt.Sprintf("%d,%d", a2.y, a2.x)
						antinodesList[pos]++
						antinodeMap[a2.y][a2.x] = "#"
					}

				}

			}
		}

	}

	for _, line := range antinodeMap {
		fmt.Println(line)
	}

	fmt.Println(len(antinodesList))
}

func part2() {
	input := util.NewInput("inputs/input.txt")

	antennas := map[string][]Point{}
	antinodeMap := [][]string{}
	antinodesList := map[string]int{}

	for y, line := range input.Lines {
		mapLine := []string{}

		for x, freq := range strings.Split(line, "") {
			mapLine = append(mapLine, ".")

			if freq != "." {
				p := Point{y: y, x: x}
				antennas[freq] = append(antennas[freq], p)
			}
		}

		antinodeMap = append(antinodeMap, mapLine)
	}

	maxX := len(antinodeMap[0]) - 1
	maxY := len(antinodeMap) - 1

	for f, points := range antennas {

		for i, p1 := range points {
			for j, p2 := range points {

				if i != j {
					dY := p2.y - p1.y
					dX := p2.x - p1.x

					a1y := p2.y + dY
					a1x := p2.x + dX

					for a1y >= 0 && a1y <= maxY && a1x >= 0 && a1x <= maxX {
						a := Point{
							y: a1y,
							x: a1x,
						}

						fmt.Printf("%s: %v ++ %v => %v \n", f, p1, p2, a)

						pos := fmt.Sprintf("%d,%d", a.y, a.x)
						antinodesList[pos]++
						antinodeMap[a.y][a.x] = "#"

						a1y += dY
						a1x += dX
					}

					a2y := p1.y - dY
					a2x := p1.x - dX

					for a2y >= 0 && a2y <= maxY && a2x >= 0 && a2x <= maxX {
						a := Point{
							y: a2y,
							x: a2x,
						}

						//fmt.Printf("%s: %v -- %v => %v \n", f, p1, p2, a)

						pos := fmt.Sprintf("%d,%d", a.y, a.x)
						antinodesList[pos]++
						antinodeMap[a.y][a.x] = "#"

						a2y -= dY
						a2x -= dX
					}

				} else {
					if len(points) > 1 {
						a := Point{
							y: p1.y,
							x: p1.x,
						}

						//fmt.Printf("%s: %v -- %v => %v \n", f, p1, p2, a)

						pos := fmt.Sprintf("%d,%d", a.y, a.x)
						antinodesList[pos]++
						antinodeMap[a.y][a.x] = "#"
					}
				}

			}
		}

	}

	for _, line := range antinodeMap {
		fmt.Println(line)
	}

	fmt.Println(len(antinodesList))
}
