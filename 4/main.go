package main

import (
	"bufio"
	"fmt"
	"os"
)

var N = 138

func getArr() [][]bool {
	scanner := bufio.NewScanner(os.Stdin)

	arr := make([][]bool, N)
	for i := range N {
		arr[i] = make([]bool, N)
	}

	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		for j, c := range line {
			if c == '@' {
				arr[i][j] = true
			} else {
				arr[i][j] = false
			}
		}
		i++
	}
	return arr
}

func canAccess(arr [][]bool, i int, j int) bool {
	if !arr[i][j] {
		return false
	}
	count := 0
	for ii := i - 1; ii < i+2; ii++ {
		for jj := j - 1; jj < j+2; jj++ {
			if 0 <= ii && ii < N && 0 <= jj && jj < N && arr[ii][jj] {
				count++
			}
		}
	}
	return count < 4+1
}

func ifAccessRemove(arr [][]bool, i int, j int) bool {
	if !arr[i][j] {
		return false
	}
	count := 0
	for ii := i - 1; ii < i+2; ii++ {
		for jj := j - 1; jj < j+2; jj++ {
			if 0 <= ii && ii < N && 0 <= jj && jj < N && arr[ii][jj] {
				count++
			}
		}
	}
	if count < 4+1 {
		arr[i][j] = false
		return true
	}
	return false
}

type Pair struct {
	i int
	j int
}

func addNeighbours(arr [][]bool, i int, j int, q *[]Pair) {
	for ii := i - 1; ii < i+2; ii++ {
		for jj := j - 1; jj < j+2; jj++ {
			if 0 <= ii && ii < N && 0 <= jj && jj < N && arr[ii][jj] {
				*q = append(*q, Pair{ii, jj})
			}
		}
	}
}

func partB(arr [][]bool) int {
	result := 0
	var q []Pair

	for i := range N {
		for j := range N {
			if ifAccessRemove(arr, i, j) {
				result++
				addNeighbours(arr, i, j, &q)
			}
		}
	}

	for k := 0; k < len(q); k++ {
		i := q[k].i
		j := q[k].j

		if ifAccessRemove(arr, i, j) {
			result++
			addNeighbours(arr, i, j, &q)
		}
	}
	return result
}

func partA(arr [][]bool) int {
	result := 0
	for i := range N {
		for j := range N {
			if canAccess(arr, i, j) {
				result++
			}
		}
	}
	return result
}

func main() {
	arr := getArr()
	fmt.Println(partA(arr))
	fmt.Println(partB(arr))
}
