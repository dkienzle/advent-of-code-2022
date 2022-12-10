package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	file     bool
	size     int64
	parent   *Node
	children map[string]*Node
}

func main() {

	root := new(Node)
	root.parent = root
	root.children = make(map[string]*Node)

	cwd := root

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

		if t[0] != '$' {
			ln := strings.Split(t, " ")
			n := new(Node)
			n.parent = cwd
			cwd.children[ln[1]] = n
			if ln[0] == "dir" {
				n.file = false
				n.children = make(map[string]*Node)
			} else {
				n.file = true
				n.size, _ = strconv.ParseInt(ln[0], 10, 64)
			}
			continue
		}

		cmd := strings.Split(t[2:], " ")
		if cmd[0] == "ls" {
			continue
		}

		if cmd[1] == ".." {
			cwd = cwd.parent
			continue
		}

		if cmd[1] == "/" {
			cwd = root
			continue
		}

		cwd = cwd.children[cmd[1]]

	}

	// now do a depth-first traversal.  Do we need to store the dir size?

	var total int64
	summarize("/", root, &total)
	println("the number you seek is", total)

	totalfree := 70000000 - root.size
	needtofind := 30000000 - totalfree

	println("best candidate is ", bestfit(root, needtofind, root.size))
}

func bestfit(node *Node, needtofind int64, bestthusfar int64) int64 {
	if node.file {
		return bestthusfar
	}
	fit := bestthusfar
	for _, n := range node.children {
		fit = bestfit(n, needtofind, fit)
	}
	if node.size >= needtofind && node.size < fit {
		fit = node.size
	}
	return fit
}

func summarize(name string, node *Node, total *int64) int64 {
	if node.file {
		return node.size
	}
	var dirsize int64
	for name, n := range node.children {
		dirsize += summarize(name, n, total)
	}
	if dirsize <= 100000 {
		*total += dirsize
	}
	node.size = dirsize
	println(name, dirsize, *total)
	return dirsize
}
