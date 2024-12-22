package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("./input")

	r := bufio.NewReader(f)

	line, _, err := r.ReadLine()
	if err != nil {
		return
	}
	num := make([]int, 0)

	for _, token := range strings.Split(string(line), " ") {
		n, _ := strconv.Atoi(token)
		num = append(num, n)
	}

	numsCopy := slices.Clone(num)
	fmt.Println(len(sol1(num)))
	fmt.Println(sol2(numsCopy))
}

func sol1(nums []int) []int {
	for n := 0; n < 25; n++ {
		for i := 0; i < len(nums); i++ {
			if nums[i] == 0 {
				nums[i] = 1
			} else if c := count(nums[i]); c%2 == 0 {
				p := pow10(c / 2)
				num1 := nums[i] / p
				num2 := nums[i] % p
				nums = append(nums[:i], append([]int{num1, num2}, nums[i+1:]...)...)
				i++
			} else {
				v := nums[i] * 2024
				nums[i] = v
			}
		}
	}
	return nums
}

const MAX_DEPTH = 75

var cache = make(map[string]int)

func sol2(nums []int) int {
	res := 0

	for i := 0; i < len(nums); i++ {
		res += conquer(nums[i], MAX_DEPTH)
	}

	return res
}

func conquer(num, depth int) int {
	if depth == 0 {
		return 1
	}

	if num == 0 {
		if v, ok := cache[fmt.Sprintf("%d %d", 1, depth-1)]; ok {
			return v
		} else {
			v = conquer(1, depth-1)
			cache[fmt.Sprintf("%d %d", 1, depth-1)] = v
			return v
		}
	} else if c := count(num); c%2 == 0 {
		p := pow10(c / 2)
		num1 := num / p
		num2 := num % p

		v1, ok := cache[fmt.Sprintf("%d %d", num1, depth-1)]
		if !ok {
			v1 = conquer(num1, depth-1)
			cache[fmt.Sprintf("%d %d", num1, depth-1)] = v1
		}
		v2, ok := cache[fmt.Sprintf("%d %d", num2, depth-1)]
		if !ok {
			v2 = conquer(num2, depth-1)
			cache[fmt.Sprintf("%d %d", num2, depth-1)] = v2
		}

		return v1 + v2
	} else {
		if v, ok := cache[fmt.Sprintf("%d %d", num*2024, depth-1)]; ok {
			return v
		} else {
			v = conquer(num*2024, depth-1)
			cache[fmt.Sprintf("%d %d", num*2024, depth-1)] = v
			return v
		}
	}
}

func pow10(power int) int {
	res := 1
	for i := 0; i < power; i++ {
		res *= 10
	}
	return res
}

func count(n int) int {
	num := 0
	for {
		n /= 10
		num++
		if n == 0 {
			return num
		}
	}
}
