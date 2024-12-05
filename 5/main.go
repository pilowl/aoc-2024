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

	rules := make([][2]int, 0)
	updates := make([][]int, 0)

	second := false
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			break
		}
		line := string(b)
		if line == "" {
			second = true
			continue
		}

		if second {
			newUpdates := make([]int, 0)
			for _, v := range strings.Split(line, ",") {
				num, _ := strconv.Atoi(v)
				newUpdates = append(newUpdates, num)
			}
			updates = append(updates, newUpdates)
		} else {
			split := strings.Split(line, "|")
			f, _ := strconv.Atoi(split[0])
			s, _ := strconv.Atoi(split[1])
			rules = append(rules, [2]int{f, s})
		}
	}

	fmt.Println(sol1(rules, updates))
	fmt.Println(sol2(rules, updates))
}

func sol1(rules [][2]int, updates [][]int) int {
	after := mapAfter(rules)

	correct := make([][]int, 0)

	for _, update := range updates {
		if check(update, after) {
			correct = append(correct, update)
		}
	}

	return sumMed(correct)
}

func sol2(rules [][2]int, updates [][]int) int {
	after := mapAfter(rules)

	correct := make([][]int, 0)

	for _, update := range updates {
		if !check(update, after) {
			correctUpdate(update, after)
			correct = append(correct, update)
		}
	}

	return sumMed(correct)
}

func correctUpdate(update []int, rules map[int]map[int]interface{}) {
	incorrectNum := func(num int) int {
		for i := num + 1; i < len(update); i++ {
			if _, ok := rules[update[num]][update[i]]; !ok {
				return i
			}
		}
		return -1
	}

	for num := 0; num < len(update)-1; num++ {
		if incorrect := incorrectNum(num); incorrect != -1 {
			update[incorrect], update[incorrect-1] = update[incorrect-1], update[incorrect]
			num--
		}
	}
}

func sumMed(updates [][]int) int {
	res := 0
	for _, update := range updates {
		res += update[len(update)/2]
	}
	return res
}

func check(update []int, rules map[int]map[int]interface{}) bool {
	for num := range update {
		for i := num + 1; i < len(update); i++ {
			if _, ok := rules[update[num]][update[i]]; !ok {
				return false
			}
		}
	}
	return true
}

func mapAfter(rules [][2]int) map[int]map[int]interface{} {
	m := make(map[int]map[int]interface{})
	for _, rule := range rules {
		if _, ok := m[rule[0]]; !ok {
			m[rule[0]] = make(map[int]interface{})
		}
		m[rule[0]][rule[1]] = struct{}{}
	}

	return m
}
