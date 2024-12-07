package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Eq struct {
	res int
	val []int
}

func main() {
	f, _ := os.Open("./input")
	r := bufio.NewReader(f)

	eqs := make([]Eq, 0)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}

		sp := strings.Split(string(line), ": ")
		res, _ := strconv.Atoi(sp[0])
		sp = strings.Split(sp[1], " ")
		vals := make([]int, len(sp))
		for i := range sp {
			vals[i], _ = strconv.Atoi(sp[i])
		}

		eqs = append(eqs, Eq{
			res: res,
			val: vals,
		})
	}

	fmt.Println(sol1(eqs))
	fmt.Println(sol2(eqs))
}

func sol1(eqs []Eq) int {
	res := 0
	for _, eq := range eqs {
		if isItPossibleToGetThisValueFromThatSlice(eq.val[1:], eq.val[0], eq.res) {
			res += eq.res
		}
	}

	return res
}

const (
	PLUS int = iota
	MUL
)

func sol2(eqs []Eq) int {
	res := 0
	for _, eq := range eqs {
		if isItPossibleToGetThisValueFromThatSliceWithWellHiddenOperator(eq.val[1:], eq.val[0], eq.res) {
			res += eq.res
		}
	}

	return res
}
func isItPossibleToGetThisValueFromThatSliceWithWellHiddenOperator(vals []int, comulative int, res int) bool {
	if len(vals) == 0 {
		return comulative == res
	}

	concatNumber, _ := strconv.Atoi(fmt.Sprintf("%d%d", comulative, vals[0]))

	return isItPossibleToGetThisValueFromThatSliceWithWellHiddenOperator(vals[1:], comulative+vals[0], res) ||
		isItPossibleToGetThisValueFromThatSliceWithWellHiddenOperator(vals[1:], comulative*vals[0], res) ||
		isItPossibleToGetThisValueFromThatSliceWithWellHiddenOperator(vals[1:], concatNumber, res)
}

func isItPossibleToGetThisValueFromThatSlice(vals []int, comulative int, res int) bool {
	if len(vals) == 0 {
		return comulative == res
	}

	return isItPossibleToGetThisValueFromThatSlice(vals[1:], comulative+vals[0], res) ||
		isItPossibleToGetThisValueFromThatSlice(vals[1:], comulative*vals[0], res)
}
