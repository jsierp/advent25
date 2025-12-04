package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getList() []string {
	list := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		list = append(list, line)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return list
}

func main() {
	list := getList()
	dial := 50
	countA := 0
	countB := 0

	for _, rot := range list {
		jump, err := strconv.Atoi(rot[1:])
		if err != nil {
			panic(err)
		}
		if rot[0] == 'L' {
			jump = -jump
		}

		dial += jump
		if dial >= 100 {
			countB += dial / 100
		} else if dial <= 0 {
			countB += -dial / 100
			if dial > jump {
				// dial was positive before, so it went through 0
				countB += 1
			}
		}
		dial = (dial%100 + 100) % 100
		if dial == 0 {
			countA += 1
		}
	}
	fmt.Println("Part A:", countA)
	fmt.Println("Part B:", countB)
}
