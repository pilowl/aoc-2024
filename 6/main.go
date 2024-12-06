package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

		runes := []rune(string(line))
		lineRunes := make([]rune, len(runes))
		for i, r := range runes {
			lineRunes[i] = r
		}
		m = append(m, lineRunes)
	}

	loc := sol1(m)
	fmt.Println(len(loc))
	fmt.Println(sol2(m, loc))
}

func sol1(m [][]rune) map[string]interface{} {
	x, y := getCurrentXY(m)

	dx, dy := 0, -1
	loc := make(map[string]interface{})

	for y >= 0 && y < len(m) && x >= 0 && x < len(m[y]) {
		if m[y][x] == '#' {
			x, y = x-dx, y-dy
			dx, dy = changeDirection(dx, dy)
			x, y = x+dx, y+dy
			continue
		}
		loc[fmt.Sprintf("%d %d", x, y)] = struct{}{}
		x, y = x+dx, y+dy
	}

	return loc
}

func sol2(m [][]rune, locs map[string]interface{}) int {
	visitedFunc := func(x, y int, dx, dy int) string {
		return fmt.Sprintf("%d %d %d %d", x, y, dx, dy)
	}

	loopFunc := func(x, y, dx, dy int) bool {
		visited := make(map[string]interface{})

		for y >= 0 && y < len(m) && x >= 0 && x < len(m[y]) {
			if m[y][x] == '#' {
				x, y = x-dx, y-dy
				dx, dy = changeDirection(dx, dy)

				continue
			}

			if _, ok := visited[visitedFunc(x, y, dx, dy)]; ok {
				return true
			}
			visited[visitedFunc(x, y, dx, dy)] = struct{}{}

			x, y = x+dx, y+dy
		}

		return false
	}

	sx, sy := getCurrentXY(m)

	res := 0
	for l := range locs {
		sl := strings.Split(l, " ")
		x, _ := strconv.Atoi(sl[0])
		y, _ := strconv.Atoi(sl[1])
		m[y][x] = '#'
		if loopFunc(sx, sy, 0, -1) {
			res++
		}
		m[y][x] = '.'
	}

	return res
}

func print(m [][]rune) {
	for i := range m {
		for j := range m[i] {
			fmt.Printf("%c", m[i][j])
		}
		fmt.Println()
	}
}

func changeDirection(dx, dy int) (int, int) {
	switch {
	case dx == 0 && dy == -1:
		return 1, 0
	case dx == 1 && dy == 0:
		return 0, 1
	case dx == 0 && dy == 1:
		return -1, 0
	default:
		return 0, -1
	}
}

func getCurrentXY(m [][]rune) (int, int) {
	for i := range m {
		for j := range m[i] {
			if m[i][j] == '^' {
				return j, i
			}
		}
	}
	return -1, -1
}
