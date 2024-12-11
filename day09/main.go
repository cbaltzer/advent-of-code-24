package main

import (
	"aoc/util"
	"fmt"
	"slices"
	"strconv"
)

func main() {
	fmt.Println("Part 1")
	part1()

	fmt.Println("\n\nPart 2")
	part2()
}

func part1() {
	input := util.NewInput("inputs/sample.txt")
	inputDisk := input.Lines[0]

	unpacked := unpackRespresentation(inputDisk)
	fmt.Println(unpacked)

	defragged := defrag(unpacked)
	fmt.Println(defragged)
	fmt.Println(diskToString(defragged))

	fmt.Println(checksum(defragged))
}

func part2() {
	input := util.NewInput("inputs/sample.txt")
	inputDisk := input.Lines[0]

	unpacked := unpackRespresentation(inputDisk)
	defragged := compress(unpacked)

	fmt.Println(diskToString(defragged))
	fmt.Println(checksum(defragged))
}

func unpackRespresentation(encoded string) []int {
	disk := []int{}

	fileID := 0

	for i, r := range encoded {
		s := string(r)
		count, _ := strconv.Atoi(s)

		for range count {
			if i%2 == 0 {
				disk = append(disk, fileID)
			} else {
				disk = append(disk, -1)
			}
		}

		if i%2 == 0 {
			fileID++
		}
	}

	return disk
}

func defrag(disk []int) []int {
	sectors := disk

	head := 0
	tail := len(sectors) - 1

	for range sectors {
		if head >= tail {
			break
		}
		if sectors[head] == -1 && sectors[tail] != -1 {
			tmp := sectors[head]
			sectors[head] = sectors[tail]
			sectors[tail] = tmp

			head++
			tail--
		} else {
			if sectors[head] != -1 {
				head++
			}
			if sectors[tail] == -1 {
				tail--
			}
		}

		//fmt.Printf("%v\nh[%d] = %d, t[%d] = %d\n\n", sectors, head, sectors[head], tail, sectors[tail])
	}

	return sectors
}

func compress(disk []int) []int {
	sectors := disk

	files := map[int]int{}

	for _, s := range sectors {
		if s != -1 {
			files[s] += 1
		}
	}

	ids := make([]int, len(files))
	i := 0
	for k := range files {
		ids[i] = k
		i += 1
	}

	head := 0
	for f := len(ids) - 1; f > 0; f-- {
		fid := ids[f]
		fileLen := files[fid]
		fileLoc := slices.Index(sectors, fid)

		if head > fileLoc {
			break
		}

		for i := head; i < len(sectors); i++ {
			if sectors[i] == -1 {
				contiguous := 0
				for j := i; j < fileLen; j++ {
					if sectors[j] == -1 {
						contiguous += 1
					}
				}
				if contiguous >= fileLen {

				}
			}
		}

	}

	fmt.Println(ids)

	return sectors
}

func checksum(disk []int) int {
	sum := 0
	for i, v := range disk {
		if v != -1 {
			sum += i * v
		}
	}
	return sum
}

func diskToString(disk []int) string {
	output := ""
	for _, v := range disk {
		if v == -1 {
			output += "."
		} else {
			s := strconv.Itoa(v)
			output += s
		}
	}
	return output
}
