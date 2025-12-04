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

func main() {
	left, right := getLists()
	count := 0
	all := 0

	for i := range len(left) {
		fmt.Println(left[i], "-", right[i], len(left[i]), len(right[i]))
		lc, _ := strconv.Atoi(left[i])
		rc, _ := strconv.Atoi(right[i])
		all += rc - lc

		start, stop := 0, 0
		size := len(right[i]) / 2

		if len(left[i])%2 == 0 {
			ll, _ := strconv.Atoi(left[i][:size])
			lr, _ := strconv.Atoi(left[i][size:])
			start = ll
			if lr > ll {
				start += 1
			}

			stop = int(math.Pow(10, float64(size))) - 1
			if len(right[i])%2 == 0 {
				rl, _ := strconv.Atoi(right[i][:size])
				rr, _ := strconv.Atoi(right[i][size:])
				stop = min(stop, rl)
				if rl > rr {
					stop -= 1
				}
			}
		} else if len(right[i])%2 == 0 {
			rl, _ := strconv.Atoi(right[i][:size])
			rr, _ := strconv.Atoi(right[i][size:])
			start = int(math.Pow(10, float64(size)-1))
			stop = rl
			if rl > rr {
				stop -= 1
			}
		}

		if start != 0 && stop != 0 && start <= stop {
			fmt.Println("Counting", start, stop)
			toAdd := (stop + start) * (stop - start + 1) / 2

			count += toAdd + toAdd*int(math.Pow(10, float64(size)))
		}
	}
	fmt.Println("final", count)
	fmt.Println("all", all)
}
