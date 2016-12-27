package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

type Program struct {
	memory       map[string]int
	instructions []string
}

func (program *Program) run() {
	var toggles []int
	for i := 0; i < len(program.instructions); i++ {
		instructionParts := strings.Fields(program.instructions[i])
		var x interface{}
		var y interface{}
		if !isAlpha(instructionParts[1]) {
			tempX, err := strconv.ParseInt(instructionParts[1], 10, 0)
			if err != nil {
				fmt.Printf("Error: %v\n", err.Error())
			}
			x = int(tempX)
		} else {
			x = instructionParts[1]
		}

		if len(instructionParts) > 2 && !isAlpha(instructionParts[2]) {
			tempY, err := strconv.ParseInt(instructionParts[2], 10, 0)
			if err != nil {
				fmt.Printf("Error: %v\n", err.Error())
			}
			y = int(tempY)
		} else if len(instructionParts) > 2 {
			y = instructionParts[2]
		}

		switch instructionParts[0] {
		case "cpy":
			if sliceContains(toggles, i) {
				i += program.jnz(x, y) - 1
			} else {
				y := y.(string)
				program.copy(x, y)
			}
		case "inc":
			if sliceContains(toggles, i) {
				x := x.(string)
				program.dec(x)
			} else {
				x := x.(string)
				program.inc(x)
			}
		case "dec":
			if sliceContains(toggles, i) {
				x := x.(string)
				program.inc(x)
			} else {
				x := x.(string)
				program.dec(x)
			}
		case "jnz":
			if sliceContains(toggles, i) {
				y := y.(string)
				program.copy(x, y)
			} else {
				i += program.jnz(x, y) - 1
			}
		case "tgl":
			if sliceContains(toggles, i) {
				x := x.(string)
				program.inc(x)
			} else {
				toggles = append(toggles, i+program.tgl(x))
			}
		}
	}
}

func isAlpha(s string) bool {
	runes := []rune(s)
	for _, r := range runes {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func sliceContains(slice []int, x int) bool {
	for _, v := range slice {
		if v == x {
			return true
		}
	}
	return false
}

func (program *Program) copy(x interface{}, y string) {
	switch t := x.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t)
	case int:
		x := x.(int)
		program.memory[y] = x
	case string:
		x := x.(string)
		program.memory[y] = program.memory[x]
	}
}

func (program *Program) inc(x string) {
	program.memory[x]++
}

func (program *Program) dec(x string) {
	program.memory[x]--
}

func (program *Program) clearMemory() {
	for i := 97; i <= 122; i++ {
		program.memory[string(i)] = 0
	}
}

func (program *Program) jnz(x interface{}, y interface{}) int {
	var newX int

	switch t := x.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t)
	case int:
		x := x.(int)
		newX = x
	case string:
		x := x.(string)
		newX = program.memory[x]
	}
	if newX != 0 {
		switch t := y.(type) {
		default:
			fmt.Printf("unexpected type %T\n", t)
		case int:
			y := y.(int)
			return y
		case string:
			y := y.(string)
			return program.memory[y]
		}
	} else {
		return 1
	}
	return 1
}

func (program *Program) tgl(x interface{}) int {
	switch t := x.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t)
	case int:
		return x.(int)
	case string:
		return program.memory[x.(string)]
	}
	return 0
}

func (program *Program) read(p []byte) {
	instructions := strings.Split(string(p), "\n")
	program.memory = make(map[string]int)
	for _, instruction := range instructions {
		program.instructions = append(program.instructions, instruction)
	}
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")

	var program Program
	program.read(contents)
	// Part 1
	program.memory["a"] = 7
	program.run()
	fmt.Printf("Value \"A\" with a initialized to 7: %v\n", program.memory["a"])
	// Part 2
	// Works but takes a long time (around 20 mins on my machine)
	program.clearMemory()
	program.memory["a"] = 12
	program.run()
	fmt.Printf("Value \"A\" with a initialized to 12: %v\n", program.memory["a"])

	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
	}
}
