package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type tuple struct {
	x int
	y int
}

func (t *tuple) move(direction rune) {
	switch direction {
	case 'U':
		t.y++
	case 'D':
		t.y--
	case 'L':
		t.x--
	case 'R':
		t.x++
	default:
		log.Fatal("illegal direction:", direction)
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (tail *tuple) follow(head tuple) {
	if tail.x == head.x && tail.y-2 == head.y {
		tail.y--
		return
	}
	if tail.x == head.x && tail.y+2 == head.y {
		tail.y++
		return
	}
	if tail.y == head.y && tail.x-2 == head.x {
		tail.x--
		return
	}
	if tail.y == head.y && tail.x+2 == head.x {
		tail.x++
		return
	}

	if abs(tail.x-head.x)+abs(tail.y-head.y) <= 2 {
		return
	}

	if tail.x < head.x {
		tail.x++
	} else {
		tail.x--
	}

	if tail.y < head.y {
		tail.y++
	} else {
		tail.y--
	}

}

func (t *tuple) visit(visited map[string]bool) {
	key := fmt.Sprintf("%d,%d", t.x, t.y)
	visited[key] = true
	//fmt.Println("Visited", t.x, ",", t.y)
}

const LEN = 10

func main() {

	visited := make(map[string]bool)

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	rope := make([]tuple, LEN) // should default to 0,0

	rope[LEN-1].visit(visited)
	for scanner.Scan() {

		t := scanner.Text()
		if t == "" {
			continue // skip the blank line
		}

		dir := ' '
		distance := 0
		num, _ := fmt.Sscanf(t, "%c %d", &dir, &distance)
		if num != 2 {
			log.Fatal("Unabled to parse", t)
		}

		for i := 0; i < distance; i++ {
			rope[0].move(dir)
			for j := 1; j < LEN; j++ {
				rope[j].follow(rope[j-1])
			}
			rope[LEN-1].visit(visited)
		}
		fmt.Printf("%s\t(%d,%d) (%d,%d)\n", t, rope[0].x, rope[0].y, rope[LEN-1].x, rope[LEN-1].y)

	}

	println("locations visited:", len(visited))

}
