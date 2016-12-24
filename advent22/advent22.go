package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	size, used, available, usePercent int
}

type NodeList [][]Node

func (nodes *NodeList) DoubleMake(x, y int) {
	*nodes = make([][]Node, x)
	for i := range *nodes {
		(*nodes)[i] = make([]Node, y)
	}
}

func (nodes *NodeList) Read(p []byte) (n int, err error) {
	lines := strings.Split(string(p), "\n")
	for _, line := range lines {
		var node Node
		var nodeRegex = regexp.MustCompile(`\/dev\/grid\/node-x(\d+)-y(\d+)`)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		matches := nodeRegex.FindAllStringSubmatch(parts[0], -1)
		x, err := strconv.ParseInt(matches[0][1], 10, 0)
		y, err := strconv.ParseInt(matches[0][2], 10, 0)

		var tempSize int64
		tempSize, err = strconv.ParseInt(strings.TrimSuffix(parts[1], "T"), 10, 0)
		node.size = int(tempSize)

		var tempUsed int64
		tempUsed, err = strconv.ParseInt(strings.TrimSuffix(parts[2], "T"), 10, 0)
		node.used = int(tempUsed)

		var tempAvailable int64
		tempAvailable, err = strconv.ParseInt(strings.TrimSuffix(parts[3], "T"), 10, 0)
		node.available = int(tempAvailable)

		var tempUsePercent int64
		tempUsePercent, err = strconv.ParseInt(strings.TrimSuffix(parts[4], "%"), 10, 0)
		node.usePercent = int(tempUsePercent)

		if err != nil {
			fmt.Printf("Error: %v", err.Error())
		}
		(*nodes)[x][y] = node
	}
	return 0, nil
}

func isPairViable(nodeA, nodeB Node, xA, yA, xB, yB int) bool {
	return nodeA.used != 0 && !(xA == xB && yA == yB) && nodeA.used <= nodeB.available
}

func (nodes *NodeList) CountViablePairs() int {
	var count int
	for xA := range *nodes {
		for yA, nodeA := range (*nodes)[xA] {
			for xB := range *nodes {
				for yB, nodeB := range (*nodes)[xB] {
					if isPairViable(nodeA, nodeB, xA, yA, xB, yB) {
						count++
					}
				}
			}
		}
	}
	return count
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
	}

	var nodes NodeList
	nodes.DoubleMake(32, 31)
	nodes.Read(contents)

	// Part 1
	fmt.Printf("Viable Pairs: %v\n", nodes.CountViablePairs())

	// Part 2
	// A general solution is very difficult so I just printed and solved by hand
	// E represents an empty node
	// - represents a wall (large node which data cant be moved in our out of)
	// * represents the goal zone (node which we need to move our data to)
	// . represents a node with data we can move
	// D represents the data we're moving
	// Also, the printed x and y axis are switched
	for x := range nodes {
		for y, node := range nodes[x] {
			if node.used == 0 {
				fmt.Printf("E")
			} else if node.size > 200 {
				fmt.Printf("-")
			} else if x == 0 && y == 0 {
				fmt.Printf("*")
			} else if x == 31 && y == 0 {
				fmt.Printf("D")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}
