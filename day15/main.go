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

// Sensor at x=655450, y=2013424: closest beacon is at x=967194, y=2000000
func parseLine(s string) (Pair, Pair) {
	sensor := Pair{}
	beacon := Pair{}
	t := strings.Split(s, " ")
	t[2] = t[2][2 : len(t[2])-1]

	_, err := fmt.Sscanf(t[2], "%d", &sensor.x)
	if err != nil {
		fmt.Println('A', t[2])
		log.Fatal(err)
	}

	t[3] = t[3][2 : len(t[3])-1]
	_, err = fmt.Sscanf(t[3], "%d", &sensor.y)
	if err != nil {
		fmt.Println('B', t[3])
		log.Fatal(err)
	}

	t[8] = t[8][2 : len(t[8])-1]
	_, err = fmt.Sscanf(t[8], "%d", &beacon.x)
	if err != nil {
		fmt.Println('C', t[8])
		log.Fatal(err)
	}

	t[9] = t[9][2:len(t[9])]
	_, err = fmt.Sscanf(t[9], "%d", &beacon.y)
	if err != nil {
		fmt.Println('D', t[9])
		log.Fatal(err)
	}

	return sensor, beacon
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func manhattan(a, b Pair) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func scanFile(filename string) (Pair, Pair) {

	min := Pair{x: math.MaxInt, y: math.MaxInt}
	max := Pair{x: math.MinInt, y: math.MinInt}

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		t := scan.Text()
		if strings.TrimSpace(t) == "" {
			continue
		}
		s, b := parseLine(t)

		dist := manhattan(s, b)
		if min.x > s.x-dist {
			min.x = s.x - dist
		}
		if min.y > s.y-dist {
			min.y = s.y - dist
		}
		if max.x < s.x+dist {
			max.x = s.x + dist
		}
		if max.y < s.y+dist {
			max.y = s.y + dist
		}
	}
	return min, max
}

func fillGrid(grid [][]int8, min Pair, sensor Pair, beacon Pair) {
	grid[sensor.y-min.y][sensor.x-min.x] = 'S'
	grid[beacon.y-min.y][beacon.x-min.x] = 'B'

	dist := manhattan(sensor, beacon)
	for y_delta := -dist; y_delta <= dist; y_delta++ {
		diff := dist - abs(y_delta)
		for x_delta := -diff; x_delta <= diff; x_delta++ {
			if grid[y_delta+sensor.y-min.y][x_delta+sensor.x-min.x] == 0 {
				grid[y_delta+sensor.y-min.y][x_delta+sensor.x-min.x] = '#'
			}
		}
	}

}

func readFile(filename string, grid [][]int8, min Pair) []Pair {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	sensors := []Pair{}
	for scan.Scan() {
		t := scan.Text()
		if strings.TrimSpace(t) == "" {
			continue
		}
		s, b := parseLine(t)

		sensors = append(sensors, s)

		fillGrid(grid, min, s, b)
	}
	return sensors
}

func main() {
	min, max := scanFile(os.Args[1])
	fmt.Printf("min = (%d,%d)\n", min.x, min.y)
	fmt.Printf("max = (%d,%d)\n", max.x, max.y)

	var grid [][]int8
	for i := min.y; i <= max.y; i++ {
		grid = append(grid, make([]int8, max.x-min.x+1))
	}

	readFile(os.Args[1], grid, min)

	//printGrid(grid, min)

	row := 10
	fmt.Printf("Number of non-beacons for row %d = %d\n", row, nonbeacons(grid, min, row))
	row = 2000000
	fmt.Printf("Number of non-beacons for row %d = %d\n", row, nonbeacons(grid, min, row))

}

func printGrid(grid [][]int8, min Pair) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%c", grid[i][j])
			}
		}
		fmt.Println()
	}

}

func nonbeacons(grid [][]int8, min Pair, row int) int {
	count := 0
	for i := 0; i < len(grid[0]); i++ {
		cell := grid[row-min.y][i]
		if cell == '#' || cell == 'S' {
			count++
		}
	}
	return count
}
