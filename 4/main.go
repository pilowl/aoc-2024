package main

import (
	"bufio"
	"fmt"
	"os"
)

const XMAS = "XMAS"

func main() {
	f, _ := os.Open("./input")

	reader := bufio.NewReader(f)

	data := make([][]rune, 0)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		data = append(data, []rune(string(line)))
	}

	fmt.Println(sol1(data))
	fmt.Println(sol2(data))
}

func sol1(data [][]rune) int {
	countX := func(i, j int) (cnt int) {
		// i, j -> A pos

		runes := []rune(XMAS)
		lenI := len(data)
		lenJ := len(data[i])
		for c := 0; c < len(runes) && i+c < lenI; c++ {
			if data[(i+c)%lenI][j] != runes[c] {
				break
			}
			if c == len(XMAS)-1 {
				cnt++
			}
		}
		for c := 0; c < len(runes) && i-c >= 0; c++ {
			if data[i-c][j] != runes[c] {
				break
			}
			if c == len(XMAS)-1 {
				cnt++
			}
		}
		for c := 0; c < len(runes) && j+c < lenJ; c++ {
			if data[i][j+c] != runes[c] {
				break
			}
			if c == len(XMAS)-1 {
				cnt++
			}
		}
		for c := 0; c < len(runes) && j-c >= 0; c++ {
			if data[i][j-c] != runes[c] {
				break
			}
			if c == len(XMAS)-1 {
				cnt++
			}
		}

		for c := 0; c < len(runes) && i+c < lenI && j+c < lenJ; c++ {
			if data[i+c][j+c] != runes[c] {
				break
			}
			if c == len(XMAS)-1 {
				cnt++
			}
		}

		for c := 0; c < len(runes) && i-c >= 0 && j-c >= 0; c++ {
			if data[i-c][j-c] != runes[c] {
				break
			}
			if c == len(XMAS)-1 {
				cnt++
			}
		}

		for c := 0; c < len(runes) && i+c < lenI && j-c >= 0; c++ {
			if data[i+c][j-c] != runes[c] {
				break
			}
			if c == len(XMAS)-1 {
				cnt++
			}
		}

		for c := 0; c < len(runes) && i-c >= 0 && j+c < len(data[i]); c++ {
			if data[i-c][j+c] != runes[c] {
				break
			}
			if c == len(XMAS)-1 {
				cnt++
			}
		}

		return cnt
	}

	occ := 0
	for i, _ := range data {
		for j, _ := range data[i] {
			if data[i][j] == 'X' {
				cnt := countX(i, j)
				occ += cnt
			}
		}
	}

	return occ
}

func sol2(data [][]rune) int {
	isMas := func(i, j int) bool {
		if !((data[i-1][j-1] == 'M' && data[i+1][j+1] == 'S') || (data[i-1][j-1] == 'S' && data[i+1][j+1] == 'M')) {
			return false
		}

		if !((data[i-1][j+1] == 'M' && data[i+1][j-1] == 'S') || (data[i-1][j+1] == 'S' && data[i+1][j-1] == 'M')) {
			return false
		}

		return true
	}

	occ := 0
	for i, _ := range data {
		for j, _ := range data[i] {
			if data[i][j] == 'A' && i > 0 && i < len(data)-1 && j > 0 && j < len(data[i])-1 {
				if isMas(i, j) {
					occ++
				}
			}
		}
	}

	return occ
}
