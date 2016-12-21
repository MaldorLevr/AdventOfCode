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

// IPRangeBlacklist holds a list of IP ranges read from a file
type IPRangeBlacklist []IPRange

func (ipRanges IPRangeBlacklist) Len() int {
	return len(ipRanges)
}

func (ipRanges IPRangeBlacklist) Swap(i, j int) {
	ipRanges[i], ipRanges[j] = ipRanges[j], ipRanges[i]
}

func (ipRanges IPRangeBlacklist) Less(i, j int) bool {
	return ipRanges[i].min < ipRanges[j].min
}

func (ipRanges *IPRangeBlacklist) Read(p []byte) (n int, err error) {
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

func (ipRanges *IPRangeBlacklist) removeOverlaps() {
	nextRangeList := *ipRanges
	var filteredRanges IPRangeBlacklist
	sort.Sort(ipRanges)
	for len(nextRangeList) > 0 {
		var origMin, highestMax int64
		for i, singleRange := range nextRangeList {
			if i == 0 {
				nextRangeList = nil
				origMin = singleRange.min
				highestMax = singleRange.max
			}
			if singleRange.min <= highestMax+1 {
				if singleRange.max > highestMax {
					highestMax = singleRange.max
				}
			} else if i != 0 {
				nextRangeList = append(nextRangeList, singleRange)
			}
		}
		filteredRanges = append(filteredRanges, IPRange{
			min: origMin,
			max: highestMax,
		})
	}
	*ipRanges = filteredRanges
}

func (ipRanges IPRangeBlacklist) countPossibleIPs() int64 {
	var prevRange IPRange
	possibleIPCount := int64(1)
	for i, singleRange := range ipRanges {
		if i == 0 {
			possibleIPCount += singleRange.min - 0
		} else if i == len(ipRanges)-1 {
			possibleIPCount += 4294967295 - singleRange.max
		} else {
			possibleIPCount += singleRange.min - prevRange.max - 1
		}
		prevRange = singleRange
	}
	return possibleIPCount
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	var ranges IPRangeBlacklist
	ranges.Read(contents)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	ranges.removeOverlaps()
	// Part 1
	fmt.Printf("Lowest possible IP: %v\n", ranges[0].max+1)
	// Part 2
	fmt.Printf("Possible IPs: %v\n", ranges.countPossibleIPs())
}
