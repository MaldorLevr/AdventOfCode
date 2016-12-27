package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type CodeLock struct {
	keypad       [][]string
	instructions []string
}

func (lock *CodeLock) findCode() []string {
	x := 1
	y := 1
	var retArray []string
	for _, instruction := range lock.instructions {
		for _, letter := range []rune(instruction) {
			switch letter {
			case 'U':
				if y > 0 && lock.keypad[x][y-1] != " " {
					y--
				}
			case 'R':
				if x < len(lock.keypad)-1 && lock.keypad[x+1][y] != " " {
					x++
				}
			case 'D':
				if y < len(lock.keypad[x])-1 && lock.keypad[x][y+1] != " " {
					y++
				}
			case 'L':
				if x > 0 && lock.keypad[x-1][y] != " " {
					x--
				}
			}
		}
		retArray = append(retArray, lock.keypad[x][y])
	}
	return retArray
}

func (lock *CodeLock) read(p []byte) {
	lock.instructions = strings.Split(string(p), "\n")
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")

	// Part 1
	codeLock := CodeLock{
		keypad: [][]string{
			[]string{"1", "4", "7"},
			[]string{"2", "5", "8"},
			[]string{"3", "6", "9"}},
		instructions: []string{""},
	}
	codeLock.read(contents)
	fmt.Printf("Code: %v\n", codeLock.findCode())

	// Part 2
	codeLock = CodeLock{
		keypad: [][]string{
			[]string{" ", " ", "5", " ", " "},
			[]string{" ", "2", "6", "A", " "},
			[]string{"1", "3", "7", "B", "D"},
			[]string{" ", "4", "8", "C", " "},
			[]string{" ", " ", "9", " ", " "}},
		instructions: []string{""},
	}
	codeLock.read(contents)
	fmt.Printf("Code w/ expanded keypad: %v\n", codeLock.findCode())

	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
	}
}
