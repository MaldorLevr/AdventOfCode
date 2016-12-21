package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func scramble(password string, commands CommandList) string {
	for _, command := range commands {
		words := strings.Split(command, " ")
		if words[0] == "swap" {
			var x, y int64
			if words[1] == "position" {
				err := error(nil)
				x, err = strconv.ParseInt(words[2], 10, 0)
				y, err = strconv.ParseInt(words[5], 10, 0)
				if err != nil {
					fmt.Printf("Error: %v", err.Error())
				}
			} else {
				x = int64(strings.Index(password, words[2]))
				y = int64(strings.Index(password, words[5]))
			}
			password = swapLetters(password, int(x), int(y))
		} else if words[0] == "rotate" {
			if words[1] == "based" {
				x := strings.Index(password, words[6])
				amount := x + 1
				if x >= 4 {
					amount++
				}
				password = rotateRight(password, amount)
			} else {
				x, err := strconv.ParseInt(words[2], 10, 0)
				if err != nil {
					fmt.Printf("err")
				}
				if words[1] == "right" {
					password = rotateRight(password, int(x))
				} else {
					password = rotateLeft(password, int(x))
				}
			}
		} else if words[0] == "reverse" {
			x, err := strconv.ParseInt(words[2], 10, 0)
			y, err := strconv.ParseInt(words[4], 10, 0)
			if err != nil {
				fmt.Printf("error")
			}
			password = reverseStringRange(password, int(x), int(y))
		} else if words[0] == "move" {
			x, err := strconv.ParseInt(words[2], 10, 0)
			y, err := strconv.ParseInt(words[5], 10, 0)
			if err != nil {
				fmt.Printf("error")
			}
			password = moveLetter(password, int(x), int(y))
		}
	}
	return password
}

func unscramble(password string, commands CommandList) string {
	commands = commands.reverse()
	for _, command := range commands {
		words := strings.Split(command, " ")
		if words[0] == "swap" {
			var x, y int64
			if words[1] == "position" {
				err := error(nil)
				x, err = strconv.ParseInt(words[2], 10, 0)
				y, err = strconv.ParseInt(words[5], 10, 0)
				if err != nil {
					fmt.Printf("Error: %v", err.Error())
				}
			} else {
				x = int64(strings.Index(password, words[2]))
				y = int64(strings.Index(password, words[5]))
			}
			password = swapLetters(password, int(y), int(x))
		} else if words[0] == "rotate" {
			if words[1] == "based" {
				x := strings.Index(password, words[6])
				var origX int
				switch x {
				case 0:
					origX = 7
				case 1:
					origX = 0
				case 2:
					origX = 4
				case 3:
					origX = 1
				case 4:
					origX = 5
				case 5:
					origX = 2
				case 6:
					origX = 6
				case 7:
					origX = 3
				}
				amount := x - origX
				if amount < 0 {
					password = rotateRight(password, -1*amount)
				} else {
					password = rotateLeft(password, amount)
				}
			} else {
				x, err := strconv.ParseInt(words[2], 10, 0)
				if err != nil {
					fmt.Printf("Error: %v", err.Error())
				}
				if words[1] == "right" {
					password = rotateLeft(password, int(x))
				} else {
					password = rotateRight(password, int(x))
				}
			}
		} else if words[0] == "reverse" {
			x, err := strconv.ParseInt(words[2], 10, 0)
			y, err := strconv.ParseInt(words[4], 10, 0)
			if err != nil {
				fmt.Printf("Error: %v", err.Error())
			}
			password = reverseStringRange(password, int(x), int(y))
		} else if words[0] == "move" {
			x, err := strconv.ParseInt(words[2], 10, 0)
			y, err := strconv.ParseInt(words[5], 10, 0)
			if err != nil {
				fmt.Printf("Error: %v", err.Error())
			}
			password = moveLetter(password, int(y), int(x))
		}
	}
	return password
}

func moveLetter(input string, x, y int) string {
	runeInput := []rune(input)
	xLetter := runeInput[x]
	for i := x; i < len(input)-1; i++ {
		runeInput[i] = runeInput[i+1]
	}
	first := true
	var previousLetter rune
	for i := y; i < len(runeInput) || first; i++ {
		v := runeInput[i]
		if first {
			runeInput[i] = xLetter
			first = false
		} else if i != len(input) {
			runeInput[i] = previousLetter
		}
		previousLetter = v
	}
	return string(runeInput)
}

func swapLetters(input string, x, y int) string {
	runeInput := []rune(input)
	runeInput[x], runeInput[y] = runeInput[y], runeInput[x]
	return string(runeInput)
}

func reverseStringRange(input string, x, y int) string {
	runeInput := []rune(input)
	reversedRange := reverseRunes(runeInput[x : y+1])
	for i := 0; x < y; x, i = x+1, i+1 {
		runeInput[x] = reversedRange[i]
	}
	return string(runeInput)
}

func reverseRunes(input []rune) []rune {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
	return input
}

func (input CommandList) reverse() CommandList {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
	return input
}

func rotateRight(input string, amount int) string {
	for i := 0; i < amount; i++ {
		runeString := []rune(input)
		var previousLetter rune
		for j, v := range runeString {
			if j == 0 {
				runeString[j] = runeString[len(runeString)-1]
			} else {
				runeString[j] = previousLetter
			}
			previousLetter = v
		}
		input = string(runeString)
	}
	return input
}

func rotateLeft(input string, amount int) string {
	for i := 0; i < amount; i++ {
		runeString := []rune(input)
		var firstLetter rune
		for j, v := range runeString {
			if j == len(runeString)-1 {
				runeString[j] = firstLetter
			} else if j == 0 {
				firstLetter = v
				runeString[j] = runeString[j+1]
			} else {
				runeString[j] = runeString[j+1]
			}
		}
		input = string(runeString)
	}
	return input
}

type CommandList []string

func (commands *CommandList) Read(p []byte) (n int, err error) {
	lines := strings.Split(string(p), "\n")
	for _, line := range lines {
		*commands = append(*commands, line)
	}
	return 0, nil
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("err")
	}
	var list CommandList
	list.Read(contents)
	// Part 2
	scrambledPass := scramble("abcdefgh", list)
	fmt.Printf("Scrambled   abcdefgh: %v", scrambledPass)
	// Part 2
	unscrambledPass := unscramble("fbgdceah", list)
	fmt.Printf("Unscrambled fbgdceah: %v", unscrambledPass)
}
