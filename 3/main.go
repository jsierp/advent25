package main

import (
	"bufio"
	"fmt"
	"os"
)

func findMax(start int, stop int, line string) (int, int) {
	mval := -1
	mpos := -1
	for i := start; i < stop; i++ {
		val := int(line[i] - '0')
		if val > mval {
			mval = val
			mpos = i
		}
		start++
	}
	return mval, mpos
}

func findNum(line string, digits int) int {
	result, val, pos := 0, 0, 0
	for i := range digits {
		val, pos = findMax(pos, len(line)-(digits-i-1), line)
		pos++
		result = result*10 + val
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	sum2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		sum += findNum(line, 2)
		sum2 += findNum(line, 12)
	}
	fmt.Println(sum, sum2)
}
