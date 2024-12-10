package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("./input")

	r := bufio.NewReader(f)

	m := make([][]int, 0)

	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}
		mn := make([]int, 0)

		for _, r := range []rune(string(line)) {
			mn = append(mn, int(r)-48)
		}
		m = append(m, mn)
	}

	fmt.Println(sol1(m))
	fmt.Println(sol2(m))
}

func sol1(m [][]int) int {
	cache := make(map[string]map[string]struct{})

	posFn := func(i, j int) string {
		return fmt.Sprintf("%d %d", i, j)
	}

	mergeNeighbors := func(i, j int, n int) {
		neighbors := make(map[string]struct{})

		merge := func(toMerge map[string]struct{}) {
			for v := range toMerge {
				neighbors[v] = struct{}{}
			}
		}

		if i > 0 && m[i-1][j] == n+1 {
			merge(cache[posFn(i-1, j)])
		}
		if i < len(m)-1 && m[i+1][j] == n+1 {
			merge(cache[posFn(i+1, j)])
		}
		if j > 0 && m[i][j-1] == n+1 {
			merge(cache[posFn(i, j-1)])
		}
		if j < len(m[i])-1 && m[i][j+1] == n+1 {
			merge(cache[posFn(i, j+1)])
		}

		cache[posFn(i, j)] = neighbors
	}

	res := 0
	for n := 9; n >= 0; n-- {
		for i := 0; i < len(m); i++ {
			for j := 0; j < len(m[i]); j++ {
				if n == 9 && m[i][j] == 9 {
					cache[posFn(i, j)] = make(map[string]struct{})
					cache[posFn(i, j)][posFn(i, j)] = struct{}{}
				} else if m[i][j] == n {
					mergeNeighbors(i, j, n)
					if n == 0 {
						res += len(cache[posFn(i, j)])
					}
				}

			}
		}
	}

	return res
}

func rc(i, j int, m [][]int) int {
	n := m[i][j]
	if n == 9 {
		return 1
	}

	sum := 0
	if i > 0 && m[i-1][j] == n+1 {
		sum += rc(i-1, j, m)
	}
	if i < len(m)-1 && m[i+1][j] == n+1 {
		sum += rc(i+1, j, m)
	}
	if j > 0 && m[i][j-1] == n+1 {
		sum += rc(i, j-1, m)
	}
	if j < len(m[i])-1 && m[i][j+1] == n+1 {
		sum += rc(i, j+1, m)
	}

	return sum
}

func sol2(m [][]int) int {
	res := 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			if m[i][j] == 0 {
				res += rc(i, j, m)
			}
		}
	}

	return res
}
