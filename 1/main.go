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
	f, _ := os.Open("./input")

	reader := bufio.NewReader(f)

	l1 := make([]int, 0)
	l2 := make([]int, 0)

	for {
		str, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		seg := strings.Split(string(str), "   ")
		i1, _ := strconv.Atoi(seg[0])
		i2, _ := strconv.Atoi(seg[1])

		l1 = append(l1, i1)
		l2 = append(l2, i2)

	}

	fmt.Println(sol1(l1, l2))

	fmt.Println(sol2(l1, l2))
}

func sol1(l1, l2 []int) int {
	slices.Sort(l1)
	slices.Sort(l2)

	dist := 0

	for i := range l1 {
		dist += int(math.Abs(float64(l2[i] - l1[i])))
	}

	return dist
}

func sol2(l1, l2 []int) int {
	score := 0

	count := func(sl []int, el int) (count int) {
		for _, v := range sl {
			if v == el {
				count++
			}
		}

		return
	}

	for _, v := range l1 {
		score += v * count(l2, v)
	}

	return score
}
