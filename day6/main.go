package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func analyze1(s string) {
	for i := 3; i < len(s); i++ {
		if s[i] != s[i-1] &&
			s[i] != s[i-2] &&
			s[i] != s[i-3] &&
			s[i-1] != s[i-2] &&
			s[i-1] != s[i-3] &&
			s[i-2] != s[i-3] {
			println("Offset", i+1)
			return
		}
	}
}

func distinct(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}

func distinct2(s string) bool {
	//println(len(s))
	b := []byte(s[:])
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	for i := 1; i < len(b); i++ {
		if b[i] == b[i-1] {
			return false
		}
	}
	return true
}

func analyze2(s string) {
	for i := 14; i < len(s); i++ {
		if distinct2(s[i-14 : i]) {
			fmt.Println("Offset", i)
			return
		}
	}
}

func main() {

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
			continue // skip the blank line
		}
		analyze2(t)

	}

}
