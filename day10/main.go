package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Instruction struct {
	opcode  int
	operand int
}

const NOP = 0
const ADDX = 1

func getInstruction(s string) Instruction {
	i := Instruction{}
	mnemonic := ""
	operand := 0
	fmt.Sscanf(s, "%s %d", &mnemonic, &operand)

	switch mnemonic {
	case "noop":
		i.opcode = NOP
	case "addx":
		{
			i.opcode = ADDX
			i.operand = operand
		}
	}
	return i
}

func drawPixel(clock int, sprite int) {
	xpos := (clock - 1) % 40
	if xpos == sprite || xpos+1 == sprite || xpos-1 == sprite {
		print("#")
	} else {
		print(" ")
	}
	if xpos == 39 {
		println()
	}
}

func main() {

	totalStrength := 0

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	state := 0 // 0 - ready for next instruction, 1 - mid-add instruction
	x := 1
	var inst Instruction

	for clock := 1; clock <= 420; clock++ {
		if state == 0 {
			if !scanner.Scan() {
				break
			}
			inst = getInstruction(scanner.Text())
		}

		drawPixel(clock, x)

		if clock%40 == 20 {
			totalStrength += (x * clock)
		}

		if inst.opcode == NOP {
			state = 0
		}

		if inst.opcode == ADDX {
			if state == 0 {
				state = 1
			} else {
				state = 0
				x += inst.operand
			}
		}
	}

	println("\ntotal signal strength:", totalStrength)

}
