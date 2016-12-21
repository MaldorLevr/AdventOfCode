package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

// IPRange represents a range of ips between min and max
type IPRange struct {
	min, max int64
}

// IPRangeReader holds a list of IP ranges read from a file
type IPRangeReader []IPRange

func (ipRanges IPRangeReader) Len() int {
	return len(ipRanges)
}

func (ipRanges IPRangeReader) Swap(i, j int) {
	ipRanges[i], ipRanges[j] = ipRanges[j], ipRanges[i]
}

func (ipRanges IPRangeReader) Less(i, j int) bool {
	return ipRanges[i].min < ipRanges[j].min
}

func (ipRanges *IPRangeReader) Read(p []byte) (n int, err error) {
	lines := strings.Split(string(p), "\n")
	for _, line := range lines {
		rangeArray := strings.Split(line, "-")
		min, err := strconv.ParseInt(rangeArray[0], 10, 0)
		if err != nil {
			return 0, err
		}
		max, err := strconv.ParseInt(rangeArray[1], 10, 0)
		if err != nil {
			return 0, err
		}
		thisRange := IPRange{
			min: min, // min value
			max: max, // max value
		}
		*ipRanges = append(*ipRanges, thisRange)
	}
	return 0, nil
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	var ranges IPRangeReader
	ranges.Read(contents)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	// gets the lowest range with overlaps filtered out
	var lowestRange IPRange
	sort.Sort(ranges)
	highestMax := int64(0)
	for i, v := range ranges {
		fmt.Printf("Min: %v  Max: %v\n", v.min, v.max)
		if i == 0 {
			highestMax = v.max
		}
		if (v.min <= highestMax+1) && (v.max >= highestMax) {
			highestMax = v.max
		}
	}

	lowestRange.max = highestMax
	lowestRange.min = int64(0)
	fmt.Printf("Lowest Possible IP: %v\n", lowestRange.max+1)
}
