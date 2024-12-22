package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("./input")
	r := bufio.NewReader(f)

	m := make([][]rune, 0)

	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}

		nm := make([]rune, 0)
		for _, r := range []rune(string(line)) {
			nm = append(nm, r)
		}
		m = append(m, nm)
	}

	fmt.Println(sol1(m))
	for k := range vis {
		delete(vis, k)
	}
	fmt.Println(sol2(m))
}

type pos struct {
	i, j int
}

var (
	vis map[pos]struct{} = make(map[pos]struct{})
)

func sol1(m [][]rune) int {
	res := 0

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			if _, ok := vis[pos{i, j}]; !ok {
				area, per := find(i, j, m)
				res += area * per
			}
		}
	}

	return res
}

// 892142 -- too high
func sol2(m [][]rune) int {
	res := 0

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			if _, ok := vis[pos{i, j}]; !ok {
				nodes := make(map[pos]struct{})
				area, _ := find2(i, j, m, nodes)
				sides := calculateSides(nodes, m)
				fmt.Printf("%c has %d sides\n", m[i][j], sides)

				res += area * sides
			}
		}
	}

	return res
}

func calculateSides(nodes map[pos]struct{}, m [][]rune) int {
	hasSide := func(i, j int) bool {
		_, ok := nodes[pos{i, j}]
		return ok
	}
	sides := 0

	for node := range nodes {
		i, j := node.i, node.j
		if !hasSide(i+1, j) && !hasSide(i, j+1) {
			sides++
		}
		if !hasSide(i-1, j) && !hasSide(i, j-1) {
			sides++
		}
		if !hasSide(i+1, j) && !hasSide(i, j-1) {
			sides++
		}
		if !hasSide(i-1, j) && !hasSide(i, j+1) {
			sides++
		}

		if hasSide(i, j+1) && hasSide(i-1, j) && !hasSide(i-1, j+1) {
			sides += 1
		}
		if hasSide(i+1, j) && hasSide(i, j+1) && !hasSide(i+1, j+1) {
			sides += 1
		}
		if hasSide(i-1, j) && hasSide(i, j-1) && !hasSide(i-1, j-1) {
			sides += 1
		}
		if hasSide(i+1, j) && hasSide(i, j-1) && !hasSide(i+1, j-1) {
			sides += 1
		}

		/*
			if hasSide(i-1, j) && hasSide(i, j+1) && !hasSide(i-1, j+1) {
				m[i][j] = 'X'
				sides += 1
			}
			if hasSide(i+1, j) && hasSide(i, j+1) && !hasSide(i+1, j+1) {
				m[i][j] = 'X'
				sides += 1
			}
			if hasSide(i-1, j) && hasSide(i, j-1) && !hasSide(i-1, j-1) {
				m[i][j] = 'X'
				sides += 1
			}
			if hasSide(i+1, j) && hasSide(i, j-1) && !hasSide(i+1, j-1) {
				fmt.Println(i, j)
				sides += 1
				m[i][j] = 'X'
			}

			if hasSide(i-1, j) && hasSide(i, j-1) && !hasSide(i-1, j+1) && !hasSide(i+1, j-1) {
				sides++
				m[i][j] = 'X'
			}
			if hasSide(i+1, j) && hasSide(i, j-1) && !hasSide(i+1, j+1) && !hasSide(i-1, j-1) {
				sides++
				m[i][j] = 'X'
			}
			if hasSide(i-1, j) && hasSide(i, j+1) && !hasSide(i-1, j+1) && !hasSide(i+1, j-1) {
				sides++
				m[i][j] = 'X'
			}
			if hasSide(i+1, j) && hasSide(i, j+1) && !hasSide(i+1, j-1) && !hasSide(i-1, j+1) {
				sides++
				m[i][j] = 'X'
			}
		*/

	}

	// for i := 0; i < len(m); i++ {
	// 	for j := 0; j < len(m[0]); j++ {
	// 		fmt.Printf("%c ", m[i][j])
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println()

	return sides
}

func find2(i, j int, m [][]rune, nodes map[pos]struct{}) (area, per int) {
	r := m[i][j]
	vis[pos{i, j}] = struct{}{}
	nodes[pos{i, j}] = struct{}{}

	// upper
	if i > 0 && m[i-1][j] == r {
		if _, ok := vis[pos{i - 1, j}]; !ok {
			da, dp := find2(i-1, j, m, nodes)
			area += da
			per += dp
		}
	} else {
		per += 1
	}

	// lower
	if i < len(m)-1 && m[i+1][j] == r {
		if _, ok := vis[pos{i + 1, j}]; !ok {
			da, dp := find2(i+1, j, m, nodes)
			area += da
			per += dp
		}
	} else {
		per += 1
	}

	// left
	if j > 0 && m[i][j-1] == r {
		if _, ok := vis[pos{i, j - 1}]; !ok {
			da, dp := find2(i, j-1, m, nodes)
			area += da
			per += dp
		}
	} else {
		per += 1
	}

	if j < len(m[0])-1 && m[i][j+1] == r {
		if _, ok := vis[pos{i, j + 1}]; !ok {
			da, dp := find2(i, j+1, m, nodes)
			area += da
			per += dp
		}
	} else {
		per += 1
	}

	return area + 1, per
}

func find(i, j int, m [][]rune) (area, per int) {
	r := m[i][j]
	vis[pos{i, j}] = struct{}{}

	// upper
	if i > 0 && m[i-1][j] == r {
		if _, ok := vis[pos{i - 1, j}]; !ok {
			da, dp := find(i-1, j, m)
			area += da
			per += dp
		}
	} else {
		per += 1
	}

	// lower
	if i < len(m)-1 && m[i+1][j] == r {
		if _, ok := vis[pos{i + 1, j}]; !ok {
			da, dp := find(i+1, j, m)
			area += da
			per += dp
		}
	} else {
		per += 1
	}

	// left
	if j > 0 && m[i][j-1] == r {
		if _, ok := vis[pos{i, j - 1}]; !ok {
			da, dp := find(i, j-1, m)
			area += da
			per += dp
		}
	} else {
		per += 1
	}

	if j < len(m[0])-1 && m[i][j+1] == r {
		if _, ok := vis[pos{i, j + 1}]; !ok {
			da, dp := find(i, j+1, m)
			area += da
			per += dp
		}
	} else {
		per += 1
	}

	return area + 1, per
}
