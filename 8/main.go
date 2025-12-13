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

type Node struct {
	X int
	Y int
	Z int
}

type Dist struct {
	Dist int
	A    int
	B    int
}

func dist2(a Node, b Node) int {
	return (a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y) + (a.Z-b.Z)*(a.Z-b.Z)
}

func distances(nodes []Node) []Dist {
	var dists []Dist
	for a := range len(nodes) {
		for b := a + 1; b < len(nodes); b++ {
			dists = append(dists, Dist{dist2(nodes[a], nodes[b]), a, b})
		}
	}
	return dists
}

func strToInt(num string) int {
	res, _ := strconv.Atoi(num)
	return res
}

func getInput() []Node {
	scn := bufio.NewScanner(os.Stdin)
	var nodes []Node

	for scn.Scan() {
		line := scn.Text()
		vals := strings.Split(line, ",")
		nodes = append(nodes, Node{strToInt(vals[0]), strToInt(vals[1]), strToInt(vals[2])})
	}
	return nodes
}

func main() {
	nodes := getInput()
	dists := distances(nodes)
	slices.SortFunc(dists, func(a Dist, b Dist) int { return cmp.Compare(a.Dist, b.Dist) })

	circuits := make([]map[int]struct{}, len(nodes))
	nodeInCircuit := make([]int, len(nodes))
	for i := range len(nodes) {
		circuits[i] = map[int]struct{}{}
		circuits[i][i] = struct{}{}
		nodeInCircuit[i] = i
	}

	for i := range len(dists) {
		dist := dists[i]
		circuitA := nodeInCircuit[dist.A]
		circuitB := nodeInCircuit[dist.B]
		if circuitA == circuitB {
			continue
		}

		for node := range circuits[circuitB] {
			circuits[circuitA][node] = struct{}{}
			nodeInCircuit[node] = circuitA
		}
		circuits[circuitB] = nil

		if len(circuits[circuitA]) == len(nodes) {
			fmt.Println("PartB:", nodes[dist.A].X*nodes[dist.B].X)
			break
		}
	}

	// partA
	// slices.SortFunc(circuits, func(a map[int]struct{}, b map[int]struct{}) int { return -cmp.Compare(len(a), len(b)) })
	// fmt.Println(len(circuits[0]) * len(circuits[1]) * len(circuits[2]))
}
