package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var testRange IPRangeBlacklist

func TestRemoveOverlap(t *testing.T) {
	testRange.removeOverlaps()
	var prevRange IPRange
	for i, v := range testRange {
		if i == 0 {
			prevRange = v
		} else {
			if prevRange.max >= v.min {
				fmt.Printf("Min: %v Max: %v\n", prevRange.min, prevRange.max)
				fmt.Printf("Min: %v Max: %v\n", v.min, v.max)
				t.Error("Expected all minimums to be greater than previous maximums.")
			}
		}
	}
}

func TestMain(m *testing.M) {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	testRange.Read(contents)
	os.Exit(m.Run())
}
