package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Cell struct {
	loc    Loc
	height int
	cost   int
	goal   bool
	seen   bool
}

type Loc struct {
	x int
	y int
}

func readMap(filename string) ([][]*Cell, Loc) {
	var rows [][]*Cell

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open file", err)
	}

	var start Loc
	rowNum := 0

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)
	for scan.Scan() {
		t := scan.Bytes()
		if len(t) == 0 {
			continue
		}
		var row []*Cell
		for i, height := range t {
			here := Loc{x: i, y: rowNum}
			c := Cell{loc: here}
			if height == 'E' {
				c.goal = true
				c.height = 26
			} else if height == 'S' {
				start = here
				c.height = 0
			} else {
				c.height = int(height - 'a' + 1)
			}
			row = append(row, &c)
		}
		rows = append(rows, row)
		rowNum++
	}
	return rows, start
}

func main() {
	land, start := readMap(os.Args[1])
	maxY := len(land)
	maxX := len(land[0])

	frontier := make([]*Cell, 0, 100)
	land[start.y][start.x].seen = true
	frontier = append(frontier, land[start.y][start.x])

	var here *Cell

	// This for loop is only needed for part 2.  Skip it for part 1.
	for i := 0; i < len(frontier); i++ {
		here = frontier[i]
		for _, dir := range neighbors(here.loc, maxX, maxY) {
			there := land[dir.y][dir.x]
			if there.height > 1 {
				continue // it's not in our start set
			}

			if there.seen {
				continue // it's already been visited or queued
			}
			there.seen = true
			there.cost = here.cost // 0
			frontier = append(frontier, there)
		}
	}
	fmt.Println("Start set contains ", len(frontier))

outer:
	for {
		if len(frontier) == 0 {
			fmt.Println("Failed")
			break
		}
		here, frontier = frontier[0], frontier[1:]
		for _, dir := range neighbors(here.loc, maxX, maxY) {
			there := land[dir.y][dir.x]
			if there.height > here.height+1 {
				continue // it's not reachable
			}

			if there.seen {
				continue // it's already been visited or queued
			}
			there.seen = true
			there.cost = here.cost + 1

			if there.goal {
				fmt.Println("found the goal in ", there.cost)
				break outer
			}
			frontier = append(frontier, there)

		}

	}

	fmt.Println()
}

func neighbors(here Loc, maxX int, maxY int) []Loc {
	dirs := make([]Loc, 0, 4)
	up := Loc{here.x, here.y - 1}
	if up.y >= 0 {
		dirs = append(dirs, up)
	}
	down := Loc{here.x, here.y + 1}
	if down.y < maxY {
		dirs = append(dirs, down)
	}
	left := Loc{here.x - 1, here.y}
	if left.x >= 0 {
		dirs = append(dirs, left)
	}
	right := Loc{here.x + 1, here.y}
	if right.x < maxX {
		dirs = append(dirs, right)
	}
	return dirs
}
