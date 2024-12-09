package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	f, _ := os.Open("./input")

	r := bufio.NewReader(f)

	runes := make([]rune, 0)

	for {
		rn, _, err := r.ReadRune()
		if err != nil || rn == '\n' || rn == 0 {
			break
		}

		runes = append(runes, rn)
	}

	fmt.Println(sol1(runes))
	fmt.Println(sol2(runes))
}

func sol1(runes []rune) int {
	// 0 == [48]
	res := make([]int, 0)
	for i, r := range runes {
		if r == 48 {
			continue
		}

		if i%2 == 0 {
			for j := 0; j < int(r)-48; j++ {
				res = append(res, i/2)
			}
		} else {
			for j := 0; j < int(r)-48; j++ {
				res = append(res, -1)
			}
		}
	}

	hash := 0
	for i := range res {
		for i < len(res) && res[i] == -1 {
			res[i] = res[len(res)-1]
			res = res[:len(res)-1]
		}
		if i < len(res) {
			hash += res[i] * i
		}
	}

	for res[len(res)-1] == -1 {
		res = res[:len(res)-1]
	}

	return hash
}

type block struct {
	id    int
	count int
	free  bool
}

func sol2(runes []rune) int {
	blocks := make([]block, 0)
	for i, r := range runes {
		if i%2 == 0 {
			blocks = append(blocks, block{i / 2, int(r) - 48, false})
		} else {
			blocks = append(blocks, block{-1, int(r) - 48, true})
		}
	}

	for i := len(blocks) - 1; i >= 0; i-- {
		b := blocks[i]
		if !b.free {
			for j := 0; j < i; j++ {
				if freeBlock := blocks[j]; freeBlock.free && freeBlock.count >= b.count {
					if freeBlock.count == b.count {
						blocks[j] = b
					} else {
						blocks[j].count -= blocks[i].count
					}
					if i == len(blocks)-1 {
						blocks[i].free = true
						blocks[i].id = -1
					} else {
						if blocks[i-1].free && i+1 < len(blocks) && blocks[i+1].free {
							blocks[i-1].count += blocks[i+1].count + b.count
							blocks = slices.Concat(blocks[:i], blocks[i+2:])
						} else if !blocks[i-1].free && i+1 < len(blocks) && blocks[i+1].free {
							blocks[i+1].count += b.count
							blocks = slices.Concat(blocks[:i], blocks[i+1:])
						} else if blocks[i-1].free && i+1 < len(blocks) && !blocks[i+1].free {
							blocks[i-1].count += b.count
							blocks = slices.Concat(blocks[:i], blocks[i+1:])
						} else {
							blocks[i].free = true
							blocks[i].id = -1
						}
					}
					if freeBlock.count != b.count {
						blocks = append(blocks[:j], append([]block{b}, blocks[j:]...)...)
					}

					break
				}
			}
		}

	}

	hash := 0
	pos := 0
	for _, block := range blocks {
		for i := 0; i < block.count; i++ {
			if !block.free {
				hash += pos * block.id
			}
			pos++
		}
	}

	return hash
}
