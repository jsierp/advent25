package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	low  int
	high int
}

func toInt(num string) int {
	i, _ := strconv.Atoi(num)
	return i
}

func toRange(line string) Range {
	pair := strings.Split(line, "-")
	return Range{toInt(pair[0]), toInt(pair[1])}
}

func getInput() ([]Range, []int) {
	scn := bufio.NewScanner(os.Stdin)

	var ranges []Range
	var ids []int

	for scn.Scan() {
		line := scn.Text()
		if line == "" {
			break
		}
		ranges = append(ranges, toRange(line))
	}

	for scn.Scan() {
		line := scn.Text()
		ids = append(ids, toInt(line))
	}
	return ranges, ids
}

func inAnyRange(ranges []Range, id int) bool {
	for _, r := range ranges {
		if r.low <= id && id <= r.high {
			return true
		}
	}
	return false
}

func partA(ranges []Range, ids []int) int {
	result := 0
	for _, id := range ids {
		if inAnyRange(ranges, id) {
			result++
		}
	}
	return result
}

func partB(ranges []Range) int {
	result := 0
	slices.SortFunc(ranges, func(a, b Range) int {
		return cmp.Compare(a.low, b.low)
	})
	fmt.Println(ranges)
	fmt.Println(len(ranges))
	start, stop := 0, -1
	for _, r := range ranges {
		if r.low <= stop {
			stop = max(r.high, stop)
		} else {
			fmt.Println("ADDING", start, stop, stop-start+1)
			result += stop - start + 1
			start, stop = r.low, r.high
		}
	}
	result += stop - start + 1
	return result
}

func main() {
	ranges, ids := getInput()
	fmt.Println(partA(ranges, ids))
	fmt.Println(partB(ranges))
}
