package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

const INORDER = 1
const OUTOFORDER = -1
const KEEPCHECKING = 0

type Node struct {
	val  int
	list []*Node
	leaf bool // technically this is redundant (equiv to list == nil)
}

func parseInteger(s []byte) (*Node, []byte) {
	value := 0
	for s[0] >= '0' && s[0] <= '9' {
		value = value*10 + int(s[0]-'0')
		s = s[1:]
	}
	n := &Node{leaf: true, val: value}
	return n, s
}

func parseList(s []byte) (*Node, []byte) {
	s = s[1:] // consume the starting [
	n := &Node{leaf: false, list: []*Node{}}
	for s[0] != ']' {
		if s[0] == '[' {
			var child *Node
			child, s = parseList(s)
			n.list = append(n.list, child)
		}
		if s[0] >= '0' && s[0] <= '9' {
			var child *Node
			child, s = parseInteger(s)
			n.list = append(n.list, child)
		}
		if s[0] == ',' {
			s = s[1:]
		}
	}
	s = s[1:] // consume the closing ]
	return n, s
}

func inOrder(l *Node, r *Node) int {
	if l.leaf && r.leaf {
		if l.val < r.val {
			return INORDER
		}
		if l.val > r.val {
			return OUTOFORDER
		}
		return KEEPCHECKING
	}

	if l.leaf {
		newl := &Node{leaf: false}
		newl.list = append(newl.list, l)
		l = newl
	}

	if r.leaf {
		newr := &Node{leaf: false}
		newr.list = append(newr.list, r)
		r = newr
	}

	// they are both lists now.
	for i := 0; i < len(l.list) && i < len(r.list); i++ {
		retval := inOrder(l.list[i], r.list[i])
		if retval != KEEPCHECKING {
			return retval
		}
	}
	if len(l.list) == len(r.list) {
		return KEEPCHECKING
	}
	if len(l.list) < len(r.list) {
		return INORDER
	}
	return OUTOFORDER

}

func main() {

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Failed to open file", err)
	}

	total := 0
	index := 0

	allPackets := make([]*Node, 0)
	s1 := []byte("[[2]]")
	p1, s1 := parseList(s1)
	s2 := []byte("[[6]]")
	p2, s2 := parseList(s2)
	allPackets = append(allPackets, p1, p2)

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)
	for scan.Scan() {
		// skip any blank lines
		for scan.Text() == "" {
			scan.Scan()
		}

		index++
		// for some reason my kludgey parser was randomly failing and I
		// assumed it had something to do with slices moving while I still
		// had pointers to them.  So I allocated a large buffer for the slices
		// to make sure this wouldn't happen and voila it stopped failing.
		// one good kludge deserves another!
		l1 := scan.Bytes()
		left := make([]byte, len(l1), 2048)
		copy(left, l1)
		lstr := scan.Text()
		scan.Scan()
		r1 := scan.Bytes()
		right := make([]byte, len(r1), 2048)
		copy(right, r1)
		rstr := scan.Text()
		fmt.Println("Index =", index)
		//fmt.Println("\tleft =", lstr)
		//fmt.Println("\tright=", rstr)

		leftTree, left := parseList(left)
		if len(left) != 0 {
			fmt.Println("leftovers", left)
			fmt.Println(lstr)
			PrintTree(leftTree)
			fmt.Println()
		}
		rightTree, right := parseList(right)
		if len(right) != 0 {
			fmt.Println("leftovers", right)
			fmt.Println(rstr)
			PrintTree(rightTree)
			fmt.Println()
		}

		allPackets = append(allPackets, leftTree, rightTree)

		retval := inOrder(leftTree, rightTree)
		//fmt.Println(index, retval)
		if retval == INORDER {
			total += index
		}

	}
	fmt.Println("Total = ", total)

	sort.Slice(allPackets, func(i, j int) bool { return inOrder(allPackets[i], allPackets[j]) == INORDER })

	div1 := 0
	div2 := 0
	for i := 0; i < len(allPackets); i++ {
		if p1 == allPackets[i] {
			div1 = i + 1
		}
		if p2 == allPackets[i] {
			div2 = i + 1
		}
	}
	fmt.Printf("Dividers %d and %d == %d\n", div1, div2, div1*div2)
}

func PrintTree(n *Node) {
	if n.leaf {
		fmt.Printf("%d", n.val)
	} else {
		fmt.Print("[")
		for i := 0; i < len(n.list); i++ {
			if i != 0 {
				fmt.Print(",")
			}
			PrintTree(n.list[i])
		}
		fmt.Print("]")
	}
	return
}
