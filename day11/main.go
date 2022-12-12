package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const ADD = 1
const MUL = 2
const SQUARE = 3

type Monkey struct {
	id          int
	items       []uint64
	op          int
	operand     int
	divisor     int
	trueDest    int
	falseDest   int
	inspections int
}

func parseList(buf string) []int {
	textList := ""
	num, err := fmt.Sscanf(buf, "Starting items: %s", &textList)
	if num != 1 || err != nil {
		return nil
	}

	var items []int
	for _, str := range strings.Split(textList, ",") {
		num64, err := strconv.ParseInt(strings.TrimSpace(str), 10, 64)
		if err != nil {
			fmt.Println(err)
			items = append(items, int(num64))
		}
	}
	return items
}

func parseOperation(s string) (int, int) {
	var mnemonic rune
	var target string

	num, err := fmt.Sscanf(s, "Operation: new = old %c %s\n", &mnemonic, &target)
	if num != 2 || err != nil {
		return 0, 0
	}

	if mnemonic == '*' && target == "old" {
		return SQUARE, 0
	}

	num64, err := strconv.ParseInt(target, 10, 64)
	if err != nil {
		return 0, 0
	}
	if mnemonic == '*' {
		return MUL, int(num64)
	}
	if mnemonic == '+' {
		return ADD, int(num64)
	}
	return 0, 0
}

func readMonkey(scan *bufio.Scanner) *Monkey {
	m := new(Monkey)

	buf := scan.Text()
	num, err := fmt.Sscanf(buf, "Monkey %d:", &m.id)
	if num != 1 || err != nil {
		return nil
	}

	if !scan.Scan() {
		return nil
	}
	buf = strings.TrimSpace(scan.Text())
	if !strings.HasPrefix(buf, "Starting items:") {
		fmt.Println(buf)
		return nil
	}

	for _, b := range strings.Split(buf[15:], ",") {
		n, err := strconv.ParseInt(strings.TrimSpace(b), 10, 64)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		m.items = append(m.items, uint64(n))
	}

	if !scan.Scan() {
		println(4)
		return nil
	}
	buf = strings.TrimSpace(scan.Text())
	m.op, m.operand = parseOperation(buf)
	if m.op == 0 {
		println(5)
		return nil
	}

	if !scan.Scan() {
		println(6)
		return nil
	}
	buf = strings.TrimSpace(scan.Text())
	num, err = fmt.Sscanf(buf, "Test: divisible by %d", &m.divisor)
	if num != 1 || err != nil {
		println(7)
		return nil
	}

	if !scan.Scan() {
		println(8)
		return nil
	}
	buf = strings.TrimSpace(scan.Text())
	num, err = fmt.Sscanf(buf, "If true: throw to monkey %d", &m.trueDest)
	if num != 1 || err != nil {
		println(9)
		return nil
	}

	if !scan.Scan() {
		println(10)
		return nil
	}
	buf = strings.TrimSpace(scan.Text())
	num, err = fmt.Sscanf(buf, "If false: throw to monkey %d", &m.falseDest)
	if num != 1 || err != nil {
		println(11)
		return nil
	}

	return m
}

func readMonkeys(filename string) []*Monkey {
	var monkeys []*Monkey

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open file", err)
	}

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)
	for scan.Scan() {
		t := scan.Text()
		if strings.TrimSpace(t) == "" {
			//println("skipping blank line")
			continue
		}
		//fmt.Println("New monkey starts with", t)
		m := readMonkey(scan)
		if m != nil {
			//fmt.Println("Got one!")
			monkeys = append(monkeys, m)
		} else {
			println("NIL")
		}
	}
	return monkeys
}

const ROUNDS = 10000

func main() {
	monkeys := readMonkeys(os.Args[1])

	// this is a horrible kludge that recognizes that since the worry
	// is only increasing, we can reduce it recognizing that the divisors
	// are all primes.  So reducing the worry by this amount won't
	// impact future divisibility.
	FUDGE := uint64(1)
	for _, m := range monkeys {
		FUDGE *= uint64(m.divisor)
	}
	fmt.Println("Fudge factor", FUDGE)

	for round := 1; round <= ROUNDS; round++ {
		for _, m := range monkeys {
			olditems := m.items
			m.items = make([]uint64, 0)
			for _, item := range olditems {
				m.inspections++

				worry := uint64(item)
				switch m.op {
				case SQUARE:
					worry = worry * worry
				case ADD:
					worry = worry + uint64(m.operand)
				case MUL:
					worry = worry * uint64(m.operand)
				}

				worry %= FUDGE

				target := monkeys[m.falseDest]
				if worry%uint64(m.divisor) == 0 {
					target = monkeys[m.trueDest]
				}
				target.items = append(target.items, worry)
			}
		}

		if round == 1 || round == 20 || round%1000 == 0 {
			fmt.Println("After round", round)
			for i, m := range monkeys {
				fmt.Println(i, m.inspections)
			}
		}

	}

	sort.Slice(monkeys, func(i, j int) bool { return monkeys[i].inspections > monkeys[j].inspections })

	fmt.Printf("Largest (%d) * Second (%d) = %d\n", monkeys[0].inspections, monkeys[1].inspections, monkeys[0].inspections*monkeys[1].inspections)
}
