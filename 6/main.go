package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var ROWS = 4

func getInput() ([][]int, []rune) {
	numbers := make([][]int, ROWS)
	scn := bufio.NewScanner(os.Stdin)
	for i := range ROWS {
		scn.Scan()
		line := scn.Text()
		for val := range strings.SplitSeq(line, " ") {
			if val != "" {
				num, _ := strconv.Atoi(val)
				numbers[i] = append(numbers[i], num)
			}
		}
	}

	var operators []rune
	scn.Scan()
	line := scn.Text()
	for val := range strings.SplitSeq(line, " ") {
		if val != "" {
			operators = append(operators, rune(val[0]))
		}
	}

	return numbers, operators
}

func partA() int {
	nums, ops := getInput()
	result := 0
	for i := range len(ops) {
		val := 0
		if ops[i] == '*' {
			val = 1
		}
		for j := range ROWS {
			if ops[i] == '*' {
				val *= nums[j][i]
			} else {
				val += nums[j][i]
			}
		}
		result += val
	}

	return result
}

func getInputB() []string {
	data, _ := io.ReadAll(os.Stdin)
	lines := strings.Split(string(data), "\n")
	return lines
}

func readNum(lines []string, col int) int {
	n := 0
	for i := range ROWS {
		r := rune(lines[i][col])
		if r != ' ' {
			n = 10*n + int(r-'0')
		}
	}
	return n
}

func partB() int {
	lines := getInputB()
	cols := len(lines[0])
	result := 0
	op := rune(lines[ROWS][0])

	val := 0
	if op == '*' {
		val = 1
	}

	for i := range cols {
		if i < cols-1 && rune(lines[ROWS][i+1]) != ' ' {
			result += val
			val = 0
			op = rune(lines[ROWS][i+1])
			fmt.Println("OP", op)
			if op == '*' {
				val = 1
			}
			continue
		}
		num := readNum(lines, i)

		if op == '*' {
			val *= num
		} else {
			val += num
		}
	}
	fmt.Println(lines[ROWS])
	result += val
	return result
}

func main() {
	//	fmt.Println(partA())
	fmt.Println(partB())
}
