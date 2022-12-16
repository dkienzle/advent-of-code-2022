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

type Span struct {
	low  int
	high int
}

type Sensor struct {
	Pair
	beacon Pair
	dist   int
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

func readFile(filename string) []Sensor {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	sensors := []Sensor{}
	for scan.Scan() {
		t := scan.Text()
		if strings.TrimSpace(t) == "" {
			continue
		}
		s, b := parseLine(t)

		sensor := Sensor{Pair: s, beacon: b, dist: manhattan(s, b)}

		sensors = append(sensors, sensor)
	}
	return sensors
}

const rownum = 2000000

func main() {

	sensors := readFile(os.Args[1])

	//printGrid(grid, min)

	spans := []Span{}

	for _, s := range sensors {
		sp, ok := getSpan(s, rownum)
		if ok {
			spans = append(spans, sp)
		}
	}
	min := math.MaxInt
	max := math.MinInt
	for _, s := range spans {
		if s.low < min {
			min = s.low
		}
		if s.high > max {
			max = s.high
		}
	}
	row := make([]byte, max-min+1)
	for _, s := range spans {
		for i := s.low; i <= s.high; i++ {
			row[i-min] = '#'
		}
	}
	for _, s := range sensors {
		if s.beacon.y == rownum {
			row[s.beacon.x-min] = 'B'
			fmt.Printf("Beacon at %d,%d\n", s.beacon.x, s.beacon.y)
		}
		if s.y == rownum {
			row[s.x-min] = 'S'
			fmt.Printf("Sensor at %d,%d\n", s.x, s.y)
		}
	}
	count := 0
	for _, c := range row {
		if c == '#' || c == 'S' {
			count++
		}
	}
	fmt.Printf("Row %d has %d non-beacons\n", rownum, count)
}

func getSpan(s Sensor, row int) (Span, bool) {
	delta := s.dist - abs(s.y-row)
	if delta < 0 {
		return Span{}, false
	}
	return Span{low: s.x - delta, high: s.x + delta}, true
}
