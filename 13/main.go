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

	machines := make([]Machine, 0)
	for {
		b, _, err := r.ReadLine()
		str := strings.TrimLeft(string(b), "Button A: ")
		spl := strings.Split(str, ", ")
		x, _ := strconv.Atoi(strings.TrimLeft(spl[0], "X+"))
		y, _ := strconv.Atoi(strings.TrimLeft(spl[1], "Y+"))
		A := Button{x, y}

		b, _, err = r.ReadLine()
		str = strings.TrimLeft(string(b), "Button B: ")
		spl = strings.Split(str, ", ")
		x, _ = strconv.Atoi(strings.TrimLeft(spl[0], "X+"))
		y, _ = strconv.Atoi(strings.TrimLeft(spl[1], "Y+"))
		B := Button{x, y}

		b, _, err = r.ReadLine()
		str = strings.TrimLeft(string(b), "Prize: ")
		spl = strings.Split(str, ", ")
		x, _ = strconv.Atoi(strings.TrimLeft(spl[0], "X="))
		y, _ = strconv.Atoi(strings.TrimLeft(spl[1], "Y="))

		machines = append(machines, Machine{
			a:     A,
			b:     B,
			prize: Prize{x, y},
		})

		_, _, err = r.ReadLine()
		if err != nil {
			break
		}
	}

	fmt.Println(sol1(machines))
	for i := range machines {
		machines[i].prize.x += 10000000000000
		machines[i].prize.y += 10000000000000
	}
	fmt.Println(sol2(machines))
}

func sol1(machines []Machine) int {
	tokens := 0

	for _, machine := range machines {
		for i := 0; machine.prize.x > machine.a.dx*i && machine.prize.y > machine.a.dy*i; i++ {
			if (machine.prize.x-i*machine.a.dx)%machine.b.dx == 0 && (machine.prize.y-i*machine.a.dy)%machine.b.dy == 0 &&
				(machine.prize.x-i*machine.a.dx)/machine.b.dx == (machine.prize.y-i*machine.a.dy)/machine.b.dy {
				tokens += 3*i + (machine.prize.x-i*machine.a.dx)/machine.b.dx
				break
			}
		}
	}

	return tokens
}

// 69*A + 27*B = PRIZE X
// 23*A + 71*B = PRIZE Y
// ---------------------
// PRIZE X | B.x
// PRIZE Y | B.y
// ---------------------
// A.x     | PRIZE X
// A.y     | PRIZE Y

func sol2(machines []Machine) int {
	tokens := 0

	for _, machine := range machines {
		d := machine.a.dx*machine.b.dy - machine.b.dx*machine.a.dy
		if d == 0 {
			continue
		}
		da := machine.prize.x*machine.b.dy - machine.prize.y*machine.b.dx
		af := float64(da) / float64(d)
		if af != float64(int(af)) {
			continue
		}

		db := machine.a.dx*machine.prize.y - machine.a.dy*machine.prize.x
		bf := float64(db) / float64(d)
		if bf != float64(int(bf)) {
			continue
		}

		tokens += int(af)*3 + int(bf)
	}

	return tokens
}

type Button struct {
	dx int
	dy int
}

type Prize struct {
	x int
	y int
}

type Machine struct {
	a     Button
	b     Button
	prize Prize
}
