package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 83286528 -- too low
func main() {
	f, _ := os.Open("./input")

	r := bufio.NewReader(f)

	robots := make([]Robot, 0)

	for {
		l, _, err := r.ReadLine()
		if err != nil {
			break
		}

		sp := strings.Split(string(l), " ")
		psp := strings.Split(strings.TrimLeft(sp[0], "p="), ",")
		vsp := strings.Split(strings.TrimLeft(sp[1], "v="), ",")

		px, _ := strconv.Atoi(psp[0])
		py, _ := strconv.Atoi(psp[1])

		vx, _ := strconv.Atoi(vsp[0])
		vy, _ := strconv.Atoi(vsp[1])

		robots = append(robots, Robot{px, py, vx, vy})
	}

	fmt.Println(sol1(robots))
	fmt.Println(sol2(robots))
}

const width, height = 101, 103

func sol1(bots []Robot) int {
	q1, q2, q3, q4 := 0, 0, 0, 0

	for _, bot := range bots {
		x := (bot.x + bot.vx*100) % width
		y := (bot.y + bot.vy*100) % height
		if x < 0 {
			x += width
		}
		if y < 0 {
			y += height
		}

		switch {
		case x < width/2 && y < height/2:
			q1++
		case x > width/2 && y < height/2:
			q2++
		case x < width/2 && y > height/2:
			q3++
		case x > width/2 && y > height/2:
			q4++
		}
	}

	return q1 * q2 * q3 * q4
}

func giveMeTree(bots []Robot) {
	var m [height][width]rune

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			m[i][j] = '.'
		}
	}

	mp := make(map[string]int)
	for _, bot := range bots {
		if _, ok := mp[fmt.Sprintf("%d %d", bot.x, bot.y)]; !ok {
			mp[fmt.Sprintf("%d %d", bot.x, bot.y)] = 1
		} else {
			mp[fmt.Sprintf("%d %d", bot.x, bot.y)]++
		}
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if v, ok := mp[fmt.Sprintf("%d %d", j, i)]; ok {
				fmt.Printf("%d", v)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

// 50, 95, 153, 196
func sol2(bots []Robot) int {
	q1, q2, q3, q4 := 0, 0, 0, 0

	for nr := 95; nr < 10000000000; nr += 101 {
		toDrawBots := make([]Robot, 0)
		for _, bot := range bots {
			x := (bot.x + bot.vx*nr) % width
			y := (bot.y + bot.vy*nr) % height
			if x < 0 {
				x += width
			}
			if y < 0 {
				y += height
			}
			bot.x, bot.y = x, y
			toDrawBots = append(toDrawBots, bot)
		}

		fmt.Printf("%d seconds elapsed\n", nr)
		giveMeTree(toDrawBots)
		bufio.NewReader(os.Stdin).ReadString('\n')
	}

	return q1 * q2 * q3 * q4
}

type Robot struct {
	x  int
	y  int
	vx int
	vy int
}
