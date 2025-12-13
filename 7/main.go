package main

import (
	"bufio"
	"fmt"
	"os"
)

func partA() int {
	scn := bufio.NewScanner(os.Stdin)
	scn.Scan()
	line := []byte(scn.Text())
	splits := 0

	for scn.Scan() {
		nextLine := []byte(scn.Text())
		for i := range len(nextLine) {
			if line[i] == byte('S') {
				if nextLine[i] == byte('^') {
					splits += 1
					if i > 0 && nextLine[i-1] == byte('.') {
						nextLine[i-1] = byte('S')
					}
					if i < len(nextLine)-1 && nextLine[i+1] == byte('.') {
						nextLine[i+1] = byte('S')
					}
				} else {
					nextLine[i] = byte('S')
				}
			}
		}
		line = nextLine
	}
	return splits
}

func partB() int {
	scn := bufio.NewScanner(os.Stdin)
	scn.Scan()
	line := scn.Text()
	n := len(line)
	rays := make([]int, n)

	for i, c := range line {
		if c == 'S' {
			rays[i] = 1
		}
	}

	for scn.Scan() {
		line := scn.Text()
		nextRays := make([]int, n)
		for i, c := range line {
			if rays[i] > 0 {
				if c == '.' {
					nextRays[i] += rays[i]
				} else if c == '^' {
					nextRays[i-1] += rays[i]
					nextRays[i+1] += rays[i]
				}
			}
		}
		fmt.Println(rays)
		rays = nextRays
	}
	result := 0
	for _, r := range rays {
		result += r
	}
	return result
}

func main() {
	// fmt.Println(partA())
	fmt.Println(partB())
}
