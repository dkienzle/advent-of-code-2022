package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Elf struct {
	start int
	end   int
}

func parse(s string) Elf {
	strs := strings.Split(s, "-")
	start, _ := strconv.ParseInt(strs[0], 10, 64)
	end, _ := strconv.ParseInt(strs[1], 10, 64)

	return Elf{start: int(start), end: int(end)}
}

func (e1 Elf) contains(e2 Elf) bool {
	return e1.start <= e2.start && e1.end >= e2.end
}

func (e1 Elf) overlapping(e2 Elf) bool {
	return !(e1.end < e2.start || e1.start > e2.end)
}

func main() {
	contains := 0
	overlaps := 0

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			continue
		}

		assignments := strings.Split(t, ",")
		elf1 := parse(assignments[0])
		elf2 := parse(assignments[1])

		if elf1.contains(elf2) || elf2.contains(elf1) {
			contains++
		}

		if elf1.overlapping(elf2) {
			overlaps++
		}

	}
	fmt.Println("Total contains =", contains)
	fmt.Println("Total overlap =", overlaps)
}
