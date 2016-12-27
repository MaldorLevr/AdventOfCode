package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Location struct {
	x, y int
}

type Path struct {
	// North: 0, East: 1, South: 2, West: 3
	direction    int
	instructions []string
}

func containsLocation(slice []Location, location Location) bool {
	for _, v := range slice {
		if v.x == location.x && v.y == location.y {
			return true
		}
	}
	return false
}

func findDistance(x, y int) int {
	if x < 0 {
		x *= -1
	}
	if y < 0 {
		y *= -1
	}
	return x + y
}

func (path *Path) firstLocationVisitedTwice() int {
	var x, y int
	path.direction = 0
	var locationsVisited []Location
	for _, instruction := range path.instructions {
		turnDirection := string([]rune(instruction)[0])
		tempDistance, err := strconv.ParseInt(string([]rune(instruction)[1:len(instruction)]), 10, 0)
		if err != nil {
			fmt.Printf("Error: %v\n", err.Error())
		}
		distance := int(tempDistance)
		if turnDirection == "L" {
			path.direction--
		} else {
			path.direction++
		}
		if path.direction == 4 {
			path.direction = 0
		} else if path.direction == -1 {
			path.direction = 3
		}

		var sign int
		switch path.direction {
		case 0:
			fallthrough
		case 1:
			sign = 1
		case 2:
			fallthrough
		case 3:
			sign = -1
		}

		switch path.direction {
		case 0:
			fallthrough
		case 2:
			origY := y
			for y = y + distance*sign; y != origY; y += sign * -1 {
				if containsLocation(locationsVisited, Location{x: x, y: y}) {
					return findDistance(x, y)
				}
				locationsVisited = append(locationsVisited, Location{x: x, y: y})
			}
			y = y + distance*sign
		case 1:
			fallthrough
		case 3:
			origX := x
			for x = x + distance*sign; x != origX; x += sign * -1 {
				if containsLocation(locationsVisited, Location{x: x, y: y}) {
					return findDistance(x, y)
				}
				locationsVisited = append(locationsVisited, Location{x: x, y: y})
			}
			x = x + distance*sign
		}

		locationsVisited = append(locationsVisited, Location{x: x, y: y})
	}
	return 0
}

func (path *Path) findShortestDistance() int {
	var x, y int
	path.direction = 0
	for _, instruction := range path.instructions {
		turnDirection := string([]rune(instruction)[0])
		tempDistance, err := strconv.ParseInt(string([]rune(instruction)[1:len(instruction)]), 10, 0)
		if err != nil {
			fmt.Printf("Error: %v\n", err.Error())
		}
		distance := int(tempDistance)
		if turnDirection == "L" {
			path.direction--
		} else {
			path.direction++
		}
		if path.direction == 4 {
			path.direction = 0
		} else if path.direction == -1 {
			path.direction = 3
		}

		switch path.direction {
		case 0:
			y += distance
		case 1:
			x += distance
		case 2:
			y -= distance
		case 3:
			x -= distance
		}
	}

	return findDistance(x, y)
}

func (path *Path) read(p []byte) {
	path.instructions = strings.Fields(strings.Replace(string(p), ",", " ", -1))
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	var path Path
	path.read(contents)
	// Part 1
	fmt.Printf("Shortest Distance: %v\n", path.findShortestDistance())
	// Part 2
	fmt.Printf("Distance to first Location found twice: %v\n", path.firstLocationVisitedTwice())

	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
	}
}
