package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func getLists() ([]string, []string) {
	left := make([]string, 0)
	right := make([]string, 0)
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	ranges := strings.Split(strings.TrimSpace(string(input)), ",")
	for _, r := range ranges {
		pair := strings.Split(r, "-")
		left = append(left, pair[0])
		right = append(right, pair[1])

	}

	return left, right
}

func countA(left string, right string) int {
	start, stop := 0, 0
	size := len(right) / 2

	if len(left)%2 == 0 {
		ll, _ := strconv.Atoi(left[:size])
		lr, _ := strconv.Atoi(left[size:])
		start = ll
		if lr > ll {
			start += 1
		}

		stop = int(math.Pow(10, float64(size))) - 1
		if len(right)%2 == 0 {
			rl, _ := strconv.Atoi(right[:size])
			rr, _ := strconv.Atoi(right[size:])
			stop = min(stop, rl)
			if rl > rr {
				stop -= 1
			}
		}
	} else if len(right)%2 == 0 {
		rl, _ := strconv.Atoi(right[:size])
		rr, _ := strconv.Atoi(right[size:])
		start = int(math.Pow(10, float64(size)-1))
		stop = rl
		if rl > rr {
			stop -= 1
		}
	}

	if start != 0 && stop != 0 && start <= stop {
		fmt.Println("Count A range", start, stop)
		toAdd := (stop + start) * (stop - start + 1) / 2

		return toAdd + toAdd*int(math.Pow(10, float64(size)))
	}
	return 0
}

func toN(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func canStart(start int, num string, size int) bool {
	for i := range len(num)/size - 1 {
		n := toN(num[(i+1)*size : (i+2)*size])
		if n < start {
			return true
		} else if n > start {
			return false
		}
	}
	return true
}

func canStop(start int, num string, size int) bool {
	for i := range len(num)/size - 1 {
		n := toN(num[(i+1)*size : (i+2)*size])
		if n < start {
			return false
		} else if n > start {
			return true
		}
	}
	return true
}

func pow10(n int) int {
	return int(math.Pow(10, float64(n)))
}

func countB(left string, right string) int {
	count := 0
	leftCounted, rightCounted := false, false
	size4 := false

	for _, size := range []int{5, 4, 3, 2, 1} {
		start, stop := 0, 0
		if size > 1 && size >= len(left) {
			continue
		}
		if size == 1 {
			fmt.Println("---------------------", rightCounted)
			// theoretically I should add right here, but there is only one
			// example where it counts and it's 11

			if rightCounted {
				fmt.Println("RRRRRRR", size)
			}
			if leftCounted {
				// some size was repeated, so size 1 is already covered
				break
			}
		}
		if size == 2 && size4 {
			// size4 was included, so size 2 is already covered
			continue
		}

		if len(left)%size == 0 && len(left) > 1 {
			start = toN(left[:size])
			if !canStart(start, left, size) {
				start += 1
			}
			stop = pow10(size) - 1
			if len(right) == len(left) {
				stop = toN(right[:size])
				if !canStop(stop, right, size) {
					stop -= 1
				}
			}
		} else if len(right)%size == 0 && len(right) > 1 {
			start = pow10(size - 1)
			stop = toN(right[:size])
			if !canStop(stop, right, size) {
				stop -= 1
			}
		}
		fmt.Println("Count B range", size, start, stop)

		if start != 0 && stop != 0 && start <= stop {
			if len(left)%size == 0 {
				leftCounted = true
			} else {
				rightCounted = true
			}
			if size == 4 {
				size4 = true
			}
			toAdd := (stop + start) * (stop - start + 1) / 2
			fmt.Println("ADDING ######################", toAdd, len(right)/size-1)
			count += toAdd
			for i := range len(right)/size - 1 {
				count += toAdd * pow10(size*(i+1))
			}
		}
	}
	return count
}

func calcB(left string, right string) int {
	fmt.Println("CALC", left, right)
	N := len(left)
	result := 0

	sizesToConsider := []int{5, 4, 3}
	if N == 2 || N == 3 || N == 5 || N == 7 || N == 11 {
		sizesToConsider = []int{1}
	} else if N%2 == 0 && N != 8 && N != 12 {
		sizesToConsider = []int{5, 4, 3, 2}
	}

	for _, size := range sizesToConsider {
		if N%size == 0 && N > size {
			start := toN(left[:size])
			if !canStart(start, left, size) {
				start += 1
			}
			stop := toN(right[:size])
			if !canStop(stop, right, size) {
				stop -= 1
			}

			if start <= stop {
				fmt.Println("HIT", start, stop)
				toAdd := (stop + start) * (stop - start + 1) / 2
				result += toAdd
				for i := range N/size - 1 {
					result += toAdd * pow10(size*(i+1))
				}
			}
		}
	}
	fmt.Println("RESULT", result)

	return result
}

func partB2() int {
	result := 0
	input, _ := io.ReadAll(os.Stdin)

	for line := range strings.SplitSeq(string(input), ",") {
		limits := strings.Split(strings.TrimSpace(line), "-")
		left := limits[0]
		right := limits[1]
		fmt.Println("|", left, "---", right, "|")
		if len(left) == len(right) {
			if len(left) > 1 {
				result += calcB(left, right)
			}
		} else {
			if len(left) > 1 {
				result += calcB(left, strings.Repeat("9", len(left)))
			}
			if len(right) > 1 {
				result += calcB("1"+strings.Repeat("0", len(right)-1), right)
			}
		}
		fmt.Println()
	}

	return result
}

func main() {
	// left, right := getLists()
	// counta, countb := 0, 0

	// for i := range len(left) {
	// 	fmt.Println(left[i], "-", right[i], len(left[i]), len(right[i]))
	// 	counta += countA(left[i], right[i])
	// 	countb += countB(left[i], right[i])
	//
	// }
	// fmt.Println("Part A", counta)
	// fmt.Println("Part B", countb)
	fmt.Println("Part B2", partB2())
}
