package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const Rock = 0
const Paper = 1
const Scissors = 2

var xlate = map[byte]int8{'A': 0, 'B': 1, 'C': 2, 'X': 0, 'Y': 1, 'Z': 2}

var names = []string{"Rock    ", "Paper   ", "Scissors"}

func getChoices(s string) (int8, int8) {

	if len(s) < 3 {
		return -1, -5
	}
	return xlate[s[0]], xlate[s[2]]
}

func getScore(me, them int8) int8 {
	if me == them {
		return 3 + me + 1
	}
	if me == (them+1)%3 {
		return 6 + me + 1
	}
	if them == (me+1)%3 {
		return me + 1
	}
	log.Fatal("Impossible result")
	return 0
}

func getScores(p1, p2 int8) (int8, int8) {
	return getScore(p1, p2), getScore(p2, p1)

}

func main() {
	totalScore := 0

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		str := scanner.Text()
		if str == "" {
			continue
		}
		p1, p2 := getChoices(str)
		_, score := getScores(p1, p2)

		fmt.Printf("%s\t%s\t%s\t%d\n", str, names[p1], names[p2], score)

		totalScore += int(score)
	}
	fmt.Println("Total score =", totalScore)
}
