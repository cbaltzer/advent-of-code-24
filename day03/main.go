package main

import (
	"aoc/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1")
	part1()

	fmt.Printf("\n\nPart 2\n")
	part2()
}

func part1() {
	input := util.NewInput("inputs/input.txt")

	total := 0
	for _, line := range input.Lines {
		total += executeCode(line)
	}

	fmt.Println(total)
}

func part2() {
	input := util.NewInput("inputs/input.txt")

	total := 0
	allCode := strings.Join(input.Lines, "")

	//for _, line := range input.Lines {
	enabledCode := getEnabledCode(allCode)
	for _, code := range enabledCode {
		total += executeCode(code)
	}
	//}

	fmt.Println(total)
}

func getEnabledCode(line string) []string {
	enabled := []string{}

	doSplit := strings.Split(line, "do()")
	for _, section := range doSplit {
		dontSplit := strings.Split(section, "don't()")
		enabled = append(enabled, dontSplit[0])
	}

	fmt.Printf("%q\n\n", enabled)

	return enabled
}

func executeCode(line string) int {
	total := 0

	commandPattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := commandPattern.FindAllStringSubmatch(line, -1)

	for _, match := range matches {
		l, _ := strconv.Atoi(match[1])
		r, _ := strconv.Atoi(match[2])
		fmt.Printf("%q | mul(%d, %d) = %d\n", match, l, r, l*r)

		total += l * r
	}

	return total
}
