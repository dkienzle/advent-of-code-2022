package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func oldmain() {

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	biggestNum := 0
	biggestTotal := 0

	elfnum := 1
	total := 0
	for scanner.Scan() {
		str := scanner.Text()
		//fmt.Println(str)
		if str == "" {
			fmt.Println("Elf", elfnum, "has", total)
			if total > biggestTotal {
				biggestTotal = total
				biggestNum = elfnum
			}
			total = 0
			elfnum++
		} else {
			cals, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				log.Fatal("Parse error", err)
			}
			total += int(cals)

		}
	}
	fmt.Printf("Elf %d has %d calories\n", biggestNum, biggestTotal)
}

func main() {

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	totals := make([]int, 0)

	total := 0
	for scanner.Scan() {
		str := scanner.Text()
		if str == "" {
			fmt.Println("Elf has", total)
			totals = append(totals, total)
			total = 0
		} else {
			cals, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				log.Fatal("Parse error", err)
			}
			total += int(cals)

		}
	}

	sort.IntSlice(totals).Sort()

	i := len(totals) - 3
	fmt.Println("Top three are:", totals[i:])
	fmt.Println("Total = ", totals[i]+totals[i+1]+totals[i+2])
}
