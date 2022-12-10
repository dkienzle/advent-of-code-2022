package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Stack []byte

func (s Stack) Top() byte {
	if len(s) == 0 {
		return ' '
	}
	return s[len(s)-1]
}

func (s *Stack) Pop() byte {
	b := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return b
}

func (s *Stack) PopN(n int) []byte {
	b := (*s)[len(*s)-n : len(*s)]
	*s = (*s)[:len(*s)-n]
	return b
}

func (s *Stack) PushN(b []byte) {
	*s = append(*s, b...)
}

func main() {

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var stacks []Stack = nil
	for scanner.Scan() {
		t := scanner.Text()
		if t[1] == '1' { // found the divider
			break // done reading the header
		}

		if stacks == nil {
			numcols := (len(t) + 1) / 4
			stacks = make([]Stack, numcols)
		}

		for i := 0; i < (len(t)+1)/4; i++ {
			char := t[i*4+1]
			if char != ' ' {
				stacks[i] = append(stacks[i], char)
			}
		}
	}

	//now reverse all the stacks!
	for i := 0; i < len(stacks); i++ {
		sz := len(stacks[i])
		target := make([]byte, sz)
		for j := 0; j < sz; j++ {
			target[j] = stacks[i][sz-j-1]
		}
		stacks[i] = target
	}

	for i := 0; i < len(stacks); i++ {
		fmt.Println(i, stacks[i])
	}

	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			continue // skip the blank line
		}

		rule := strings.Split(t, " ")
		num, _ := strconv.ParseInt(rule[1], 10, 64)
		from, _ := strconv.ParseInt(rule[3], 10, 64)
		to, _ := strconv.ParseInt(rule[5], 10, 64)

		fmt.Printf("Moving %d from %d to %d\n", num, from-1, to-1)

		stacks[to-1].PushN(stacks[from-1].PopN(int(num)))

		fmt.Println("------------------------------")
		for i := 0; i < len(stacks); i++ {
			fmt.Println(i, stacks[i])
		}

	}

	str := []byte{}
	for i := 0; i < len(stacks); i++ {
		str = append(str, stacks[i].Top())
	}

	fmt.Printf("Total pattern = \"%s\"\n", str)

}
