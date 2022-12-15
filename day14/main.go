package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type Pair struct {
	x int
	y int
}

func parseLine(s string) []Pair {
	coords := strings.Split(s, "->")

	ret := make([]Pair, len(coords))
	for i := 0; i < len(coords); i++ {
		p := Pair{}
		n, err := fmt.Sscanf(coords[i], "%d,%d", &p.x, &p.y)
		if n != 2 || err != nil {
			log.Fatal(n, err)
		}
		ret[i] = p
	}
	return ret
}

// in the spirit of the advent of code, I should just skip this and
// use numbers that I got from eyeballing the only data file it needs
// to work with.  but I can't do that.
func scanFile(filename string) (Pair, Pair) {
	min := Pair{x: math.MaxInt, y: 0}
	max := Pair{x: 0, y: 0}
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		t := scan.Text()
		for _, pair := range parseLine(t) {
			if pair.x < min.x {
				min.x = pair.x
			} else if pair.x > max.x {
				max.x = pair.x
			}
			if pair.y < min.y {
				min.y = pair.y
			} else if pair.y > max.y {
				max.y = pair.y
			}
		}
	}
	return min, max
}

func fillLine(grid [][]int8, start Pair, from Pair, to Pair) {
	fmt.Printf("Draw from (%d,%d) to (%d,%d)\n", from.x, from.y, to.x, to.y)
	// draw a vertical line
	if from.x == to.x {
		dir := 1
		if to.y < from.y {
			dir = -1
		}
		for y := from.y; y != to.y+dir; y += dir {
			fmt.Printf("Drawing point at (%d,%d)\n", from.x, y)
			grid[y-start.y][from.x-start.x] = '#'
		}
		return
	}
	// draw a horizontal line
	dir := 1
	if to.x < from.x {
		dir = -1
	}
	for x := from.x; x != to.x+dir; x += dir {
		fmt.Printf("Drawing point at (%d,%d)\n", x, from.y)
		grid[from.y-start.y][x-start.x] = '#'
	}
}

func readFile(filename string, grid [][]int8, start Pair, extent Pair) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		t := scan.Text()
		pairs := parseLine(t)
		cursor := pairs[0]
		for i := 1; i < len(pairs); i++ {
			fillLine(grid, start, cursor, pairs[i])
			cursor = pairs[i]
		}
	}
}

func main() {
	min, max := scanFile(os.Args[1])

	start := Pair{x: min.x - 1, y: min.y}
	extent := Pair{x: max.x - min.x + 3, y: max.y + 1}

	grid := make([][]int8, extent.y)
	for i := 0; i < extent.y; i++ {
		grid[i] = make([]int8, extent.x)
	}

	fmt.Println(start, extent)

	readFile(os.Args[1], grid, start, extent)

	count := dumpSand(grid, start)

	drawGrid(grid, start)

	fmt.Println("Count =", count)
}

func drawGrid(grid [][]int8, start Pair) {
	offset500 := 500 - start.x
	fmt.Println(strings.Repeat(" ", offset500) + "5")
	fmt.Println(strings.Repeat(" ", offset500) + "0")
	fmt.Println(strings.Repeat(" ", offset500) + "0")

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 0 {
				fmt.Print(".")
			} else {
				fmt.Printf("%c", grid[y][x])
			}
		}
		fmt.Println()
	}
}

func dumpSand(grid [][]int8, start Pair) int {
	count := 0
	origin := Pair{x: 500 - start.x, y: 0}
	for {
		dest, ok := dropSand(grid, origin)
		if !ok {
			return count
		}
		count++
		grid[dest.y][dest.x] = 'o'
		//drawGrid(grid, start)
	}
}

// finally a function where we just use the Pair as the
// index into the array without worrying about the offset.
func dropSand(grid [][]int8, start Pair) (Pair, bool) {
	loc := start
	for {
		if loc.y == len(grid)-1 {
			return Pair{}, false
		}
		if grid[loc.y+1][loc.x] == 0 {
			loc.y++
		} else if grid[loc.y+1][loc.x-1] == 0 {
			loc.y++
			loc.x--
		} else if grid[loc.y+1][loc.x+1] == 0 {
			loc.y++
			loc.x++
		} else {
			return loc, true
		}
	}
}
