package main

import (
	"bufio"
	"fmt"
	"os"
)

type pos struct {
	x, y int
}

func main() {
	f, _ := os.Open("./input")
	r := bufio.NewReader(f)

	m := make([][]rune, 0)

	for {
		l, _, err := r.ReadLine()
		if err != nil {
			break
		}

		row := make([]rune, 0)

		for _, r := range []rune(string(l)) {
			row = append(row, r)
		}

		m = append(m, row)
	}

	fmt.Println(sol1(m))
	fmt.Println(sol2(m))
}

func sol1(m [][]rune) int {
	height, width := len(m), len(m[0])

	locs := make(map[rune][]pos)

	nodes := make(map[[2]int]struct{})

	inboundsFunc := func(loc1, loc2 pos) (res int) {
		dx, dy := loc1.x-loc2.x, loc1.y-loc2.y

		if px, py := loc1.x+dx, loc1.y+dy; px >= 0 && px < width && py >= 0 && py < height {
			if _, ok := nodes[[2]int{px, py}]; !ok {
				nodes[[2]int{px, py}] = struct{}{}
				res++
			}
		}

		if px, py := loc2.x-dx, loc2.y-dy; px >= 0 && px < width && py >= 0 && py < height {
			if _, ok := nodes[[2]int{px, py}]; !ok {
				nodes[[2]int{px, py}] = struct{}{}
				res++
			}
		}

		return
	}

	res := 0
	for i, row := range m {
		for j, r := range row {
			if r != '.' {
				if loc, ok := locs[r]; ok {
					newLoc := pos{j, i}
					for _, l := range loc {
						res += inboundsFunc(newLoc, l)
					}
					loc = append(loc, newLoc)
					locs[r] = loc
				} else {
					locs[r] = []pos{{j, i}}
				}
			}
		}
	}

	return res
}

func sol2(m [][]rune) int {
	height, width := len(m), len(m[0])

	locs := make(map[rune][]pos)

	nodes := make(map[[2]int]struct{})

	inboundsFunc := func(loc1, loc2 pos) (res int) {
		dx, dy := loc1.x-loc2.x, loc1.y-loc2.y

		for px, py := loc1.x+dx, loc1.y+dy; px >= 0 && px < width && py >= 0 && py < height; px, py = px+dx, py+dy {
			if _, ok := nodes[[2]int{px, py}]; !ok {
				nodes[[2]int{px, py}] = struct{}{}
				res++
			}
		}

		for px, py := loc2.x-dx, loc2.y-dy; px >= 0 && px < width && py >= 0 && py < height; px, py = px-dx, py-dy {
			if _, ok := nodes[[2]int{px, py}]; !ok {
				nodes[[2]int{px, py}] = struct{}{}
				res++
			}
		}

		return
	}

	res := 0
	for i, row := range m {
		for j, r := range row {
			if r != '.' && r != '#' {
				if loc, ok := locs[r]; ok {
					newLoc := pos{j, i}
					for _, l := range loc {
						res += inboundsFunc(newLoc, l)
					}
					loc = append(loc, newLoc)
					locs[r] = loc
				} else {
					locs[r] = []pos{{j, i}}
				}
			}
		}
	}

	for _, loc := range locs {
		if len(loc) > 2 {
			for _, l := range loc {
				if _, ok := nodes[[2]int{l.x, l.y}]; !ok {
					res++
				}
			}
		}
	}

	return res
}
