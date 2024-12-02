package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("./test")
	reader := bufio.NewReader(f)

	lines := make([][]int, 0)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		seg := strings.Split(string(line), " ")
		lines = append(lines, make([]int, len(seg)))

		for i, val := range seg {
			value, _ := strconv.Atoi(val)
			lines[len(lines)-1][i] = value
		}
	}

	fmt.Println(sol1(lines))
	_, res1 := sol2(lines)
	_, res2 := sol3(lines)
	diff(res1, res2)
}

func sol1(lines [][]int) int {
	safe := func(line []int) bool {
		asc := line[0] < line[1]
		curr := line[0]
		for i := 1; i < len(line); i++ {
			if curr == line[i] {
				return false
			}

			if curr < line[i] != asc {
				return false
			}

			if int(math.Abs(float64(curr-line[i]))) > 3 {
				return false
			}
			curr = line[i]
		}
		return true
	}

	safeCount := 0
	for _, line := range lines {
		if safe(line) {
			safeCount++
		}
	}

	return safeCount
}

func sol2(lines [][]int) (int, [][]int) {
	safe := func(line []int) (bool, int) {
		asc := line[0] < line[1]
		curr := line[0]
		for i := 1; i < len(line); i++ {
			if curr == line[i] {
				return false, i
			}

			if curr < line[i] != asc {
				return false, i
			}

			if int(math.Abs(float64(curr-line[i]))) > 3 {
				return false, i
			}
			curr = line[i]
		}
		return true, 0
	}

	res := make([][]int, 0)
	safeCount := 0
	for _, line := range lines {
		s, idx := safe(line)
		if s {
			safeCount++
		} else {
			if idx == 1 {
				if s, _ := safe(line[1:]); !s {
					if s, _ := safe(slices.Concat([]int{line[0]}, line[2:])); !s {
						continue
					}
				}
			} else {
				if s, _ := safe(slices.Concat(line[:idx-1], line[idx:])); !s {
					if s, _ := safe(slices.Concat(line[:idx], line[idx+1:])); !s {
						continue
					}
				}
			}
			res = append(res, line)
			safeCount++
		}
	}
	return safeCount, res
}

func diff(lines1 [][]int, lines2 [][]int) {
	for _, line := range lines2 {
		found := false
		for _, l := range lines1 {
			if slices.Equal(l, line) {
				found = true
				break
			}
		}
		if !found {
			fmt.Println(line)
		}
	}
}

func sol3(lines [][]int) (int, [][]int) {
	safe := func(line []int) (bool, int) {
		asc := line[0] < line[1]
		curr := line[0]
		for i := 1; i < len(line); i++ {
			if curr == line[i] {
				return false, i
			}

			if curr < line[i] != asc {
				return false, i
			}

			if int(math.Abs(float64(curr-line[i]))) > 3 {
				return false, i
			}
			curr = line[i]
		}
		return true, 0
	}

	res := make([][]int, 0)
	safeCount := 0
	for _, line := range lines {
		s, _ := safe(line)
		if s {
			safeCount++
		} else {
			for i := 0; i < len(line); i++ {
				if s, _ := safe(slices.Concat(line[:i], line[i+1:])); s {
					res = append(res, line)
					safeCount++
					break
				}
			}
		}
	}
	return safeCount, res
}
