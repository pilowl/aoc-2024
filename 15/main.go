package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Move rune

const (
	up    Move = '^'
	down  Move = 'v'
	left  Move = '<'
	right Move = '>'
)

func main() {
	f, _ := os.Open("./input")
	r := bufio.NewReader(f)

	mp := make([][]rune, 0)
	mp2 := make([][]rune, 0)
	for {
		b, _, _ := r.ReadLine()
		if len(b) == 0 {
			break
		}

		newLine := make([]rune, 0)
		for _, r := range []rune(string(b)) {
			newLine = append(newLine, r)
		}

		mp = append(mp, slices.Clone(newLine))
		mp2 = append(mp2, slices.Clone(newLine))
	}

	moves := make([]Move, 0)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			break
		}

		for _, r := range []rune(string(b)) {
			moves = append(moves, Move(r))
		}
	}

	// fmt.Println(sol1(mp, moves))
	fmt.Println(sol2(mp2, moves))
}

var moveDiff map[Move]struct{ dx, dy int } = map[Move]struct {
	dx int
	dy int
}{
	up:    {0, -1},
	down:  {0, 1},
	left:  {-1, 0},
	right: {1, 0},
}

func sol1(mp [][]rune, moves []Move) int {
	var px, py int = 0, 0

	for i := 0; i < len(mp); i++ {
		for j := 0; j < len(mp[i]); j++ {
			if mp[i][j] == '@' {
				px, py = j, i
				mp[i][j] = '.'
				break
			}
		}
		if px != 0 || py != 0 {
			break
		}
	}

	moveFunc := func(move Move) {
		diff := moveDiff[move]
		sx, sy := px+diff.dx, py+diff.dy

		for sy < len(mp)-1 && sx < len(mp[sy])-1 && sx > 0 && sy > 0 {
			if mp[sy][sx] == '#' {
				break
			}
			if mp[sy][sx] == '.' {
				for sx != px || sy != py {
					mp[sy][sx] = mp[sy-diff.dy][sx-diff.dx]
					sx -= diff.dx
					sy -= diff.dy
				}
				px += diff.dx
				py += diff.dy
				break
			}
			sx += diff.dx
			sy += diff.dy
		}
	}

	for _, move := range moves {
		moveFunc(move)
		mp[py][px] = '@'
		print(mp)
		mp[py][px] = '.'
		bufio.NewReader(os.Stdin).ReadString('\n')
	}

	return boxScore(mp, 'O')
}

func sol2(mp [][]rune, moves []Move) int {
	mp = remap(mp)

	var px, py int = 0, 0

	for i := 0; i < len(mp); i++ {
		for j := 0; j < len(mp[i]); j++ {
			if mp[i][j] == '@' {
				px, py = j, i
				mp[i][j] = '.'
				break
			}
		}
		if px != 0 || py != 0 {
			break
		}
	}

	for _, move := range moves {
		nx, ny := rmove(px, py, mp, move, true)
		if nx != px || ny != py {
			rmove(px, py, mp, move, false)
			px, py = nx, ny
		}
		mp[py][px] = '.'
	}
	mp[py][px] = '@'
	print(mp)
	bufio.NewReader(os.Stdin).ReadString('\n')

	return boxScore(mp, '[')
}

func rmove(px, py int, mp [][]rune, move Move, check bool) (newPx, newPy int) {
	sx, sy := px+moveDiff[move].dx, py+moveDiff[move].dy

	if mp[sy][sx] == '.' {
		return sx, sy
	}
	if mp[sy][sx] == '#' {
		return px, py
	}
	if mp[sy][sx] == '[' && (move == up || move == down) {
		mx1, my1 := rmove(sx, sy, mp, move, check)
		mx2, my2 := rmove(sx+1, sy, mp, move, check)
		if (mx1 == sx && my1 == sy) || (mx2 == sx+1 && my2 == sy) {
			return px, py
		}
		if !check {
			mp[my1][mx1] = mp[sy][sx]
			mp[my2][mx2] = mp[sy][sx+1]
			mp[sy][sx+1] = '.'
		}
		return sx, sy
	}

	if mp[sy][sx] == ']' && (move == up || move == down) {
		mx1, my1 := rmove(sx, sy, mp, move, check)
		mx2, my2 := rmove(sx-1, sy, mp, move, check)
		if (mx1 == sx && my1 == sy) || (mx2 == sx-1 && my2 == sy) {
			return px, py
		}
		if !check {
			mp[my1][mx1] = mp[sy][sx]
			mp[my2][mx2] = mp[sy][sx-1]
			mp[sy][sx-1] = '.'
		}
		return sx, sy
	}

	nx, ny := rmove(sx, sy, mp, move, check)
	if nx != sx || ny != sy {
		if !check {
			mp[ny][nx] = mp[sy][sx]
		}
		return sx, sy
	}
	return px, py
}

func remap(mp [][]rune) [][]rune {
	newMp := make([][]rune, 0)

	for i := 0; i < len(mp); i++ {
		newRow := make([]rune, 0)
		for j := 0; j < len(mp[i]); j++ {
			switch mp[i][j] {
			case '#':
				newRow = append(newRow, '#', '#')
			case 'O':
				newRow = append(newRow, '[', ']')
			case '.':
				newRow = append(newRow, '.', '.')
			case '@':
				newRow = append(newRow, '@', '.')
			}
		}
		newMp = append(newMp, newRow)
	}

	return newMp
}

func boxScore(m [][]rune, identifier rune) int {
	score := 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			if m[i][j] == identifier {
				score += i*100 + j
			}
		}
	}

	return score
}

func print(m [][]rune) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			fmt.Printf("%c", m[i][j])
		}
		fmt.Println()
	}
}
