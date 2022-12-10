package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func findCommon(l, r []byte) byte {
	sort.Slice(l, func(i, j int) bool {
		return l[i] < l[j]
	})

	for _, v := range r {
		i := sort.Search(len(l), func(i int) bool { return l[i] >= v })
		if i < len(l) && l[i] == v {
			return v
		}
	}

	return 0
}

func contains(a []byte, c byte) bool {
	i := sort.Search(len(a), func(i int) bool { return a[i] >= c })
	return i < len(a) && a[i] == c

}

func findCommon3(x, y, z []byte) byte {
	sort.Slice(x, func(i, j int) bool {
		return x[i] < x[j]
	})
	sort.Slice(y, func(i, j int) bool {
		return y[i] < y[j]
	})

	for _, c := range z {
		if contains(x, c) && contains(y, c) {
			return c
		}
	}

	return 0
}

func getWeight(c byte) byte {
	if c >= 'a' {
		return c - byte('a') + 1
	}
	return c - byte('A') + 27
}

func main2() {
	totalWeight := 0

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

		sl := []byte(str)
		sz := len(sl) / 2
		left := sl[:sz]
		right := sl[sz:]

		//fmt.Printf("%s\n", left)
		//fmt.Printf("%s\n", right)

		dup := findCommon(left, right)
		weight := getWeight(dup)

		fmt.Printf("%c=>%d\n", dup, weight)

		totalWeight += int(weight)
	}
	fmt.Println("Total score =", totalWeight)
}

func main() {
	totalWeight := 0

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	set := make([]string, 0, 3)

	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			continue
		}
		set = append(set, t)
		if len(set) == 3 {
			s1 := []byte(set[0])
			s2 := []byte(set[1])
			s3 := []byte(set[2])

			dup := findCommon3(s1, s2, s3)
			weight := getWeight(dup)

			fmt.Printf("%c=>%d\n", dup, weight)

			totalWeight += int(weight)
			set = set[:0]
		}
	}
	fmt.Println("Total score =", totalWeight)
}
