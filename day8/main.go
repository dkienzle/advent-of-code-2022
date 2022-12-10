package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	trees := make([]string, 0)

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	rows := 0
	cols := 0

	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			continue // skip the blank line
		}

		if cols == 0 {
			cols = len(t)
		} else {
			if cols != len(t) {
				println("Warning got string of length", len(t), "expected", cols)
			}
		}

		trees = append(trees, t)
	}
	rows = len(trees)

	visible := 0

	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			if !isHidden(i, j, trees) {
				visible++
			}
		}
	}
	println(visible, "trees are visible")

	maxScore := 0
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			score := calcScore(i, j, trees)
			if score > maxScore {
				maxScore = score
			}
		}
	}

	println("maximum scenic score is", maxScore)

}

func isHidden(x int, y int, trees []string) bool {

	return blockedFromLeft(x, trees[y]) &&
		blockedFromRight(x, trees[y]) &&
		blockedFromTop(x, y, trees) &&
		blockedFromBottom(x, y, trees)
}

func blockedFromLeft(x int, treeRow string) bool {
	for i := 0; i < x; i++ {
		if treeRow[i] >= treeRow[x] {
			return true
		}
	}
	return false
}

func blockedFromRight(x int, treeRow string) bool {
	for i := x + 1; i < len(treeRow); i++ {
		if treeRow[i] >= treeRow[x] {
			return true
		}
	}
	return false
}

func blockedFromTop(x int, y int, trees []string) bool {
	for j := 0; j < y; j++ {
		if trees[j][x] >= trees[y][x] {
			return true
		}
	}
	return false
}

func blockedFromBottom(x int, y int, trees []string) bool {
	for j := y + 1; j < len(trees); j++ {
		if trees[j][x] >= trees[y][x] {
			return true
		}
	}
	return false
}

func calcScore(x int, y int, trees []string) int {

	return visibleFromLeft(x, trees[y]) *
		visibleFromRight(x, trees[y]) *
		visibleFromTop(x, y, trees) *
		visibleFromBottom(x, y, trees)
}

func visibleFromRight(x int, treeRow string) int {
	count := 0
	for i := x + 1; i < len(treeRow); i++ {
		count++
		if treeRow[i] >= treeRow[x] {
			break
		}
	}
	return count
}

func visibleFromLeft(x int, treeRow string) int {
	count := 0
	for i := x - 1; i >= 0; i-- {
		count++
		if treeRow[i] >= treeRow[x] {
			break
		}
	}
	return count
}

func visibleFromTop(x int, y int, trees []string) int {
	count := 0
	for j := y - 1; j >= 0; j-- {
		count++
		if trees[j][x] >= trees[y][x] {
			break
		}
	}
	return count
}

func visibleFromBottom(x int, y int, trees []string) int {
	count := 0
	for j := y + 1; j < len(trees); j++ {
		count++
		if trees[j][x] >= trees[y][x] {
			break
		}
	}
	return count
}
