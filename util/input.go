package util

import (
	"bufio"
	"log"
	"os"
)

type Input struct {
	filename string
	Lines    []string
}

func NewInput(filename string) Input {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := Input{
		filename: filename,
		Lines:    make([]string, 0),
	}

	for scanner.Scan() {
		line := scanner.Text()

		input.Lines = append(input.Lines, line)
	}

	return input
}
