package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var mulRegex = regexp.MustCompile("mul\\(\\d+,\\d+\\)")

func main() {
	bytes, _ := os.ReadFile("./input")
	data := string(bytes)

	fmt.Println(sol1(data))
	fmt.Println(sol2(data))
}

func sol1(data string) int {
	founds := mulRegex.FindAllString(data, 1000)
	res := 0

	for _, found := range founds {
		found = strings.TrimRight(strings.TrimLeft(found, "mul("), ")")
		nums := strings.Split(found, ",")
		n1, _ := strconv.Atoi(nums[0])
		n2, _ := strconv.Atoi(nums[1])
		res += n1 * n2
	}

	return res
}

func sol2(data string) int {

	founds := mulRegex.FindAllString(data, 1000)
	idxs := mulRegex.FindAllStringIndex(data, 1000)
	dos := regexp.MustCompile("do\\(\\)").FindAllStringIndex(data, 2000)
	donts := regexp.MustCompile("don't\\(\\)").FindAllStringIndex(data, 2000)

	lastLeft := func(idx int) (_ int, ddo bool) {
		do := -1
		dont := -1

		for _, d := range dos {
			if d[0] > idx {
				break
			}
			do = d[0]
		}

		for _, d := range donts {
			if d[0] > idx {
				break
			}
			dont = d[0]
		}

		if do > dont || do == dont {
			return do, true
		} else {
			return dont, false
		}
	}

	res := 0
	for i, found := range founds {
		_, do := lastLeft(idxs[i][0])
		if !do {
			continue
		}

		found = strings.TrimRight(strings.TrimLeft(found, "mul("), ")")
		nums := strings.Split(found, ",")
		n1, _ := strconv.Atoi(nums[0])
		n2, _ := strconv.Atoi(nums[1])
		res += n1 * n2
	}

	return res
}
