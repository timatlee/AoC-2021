package bingo

import (
	"strconv"
	"strings"
)

type Bingo struct {
	board   [5][5]int
	marked  [5][5]bool
	rawRows []string
}

func New(rawrows []string) Bingo {
	bingo := Bingo{
		rawRows: rawrows,
	}

	bingo = bingo.process_raw_rows()

	return bingo
}

func (b Bingo) process_raw_rows() Bingo {
	for i := 0; i < len(b.rawRows); i++ {
		numbers := strings.Fields(b.rawRows[i])
		for j := 0; j < len(numbers); j++ {
			if len(numbers[j]) > 0 {
				number, _ := strconv.ParseInt(numbers[j], 10, 64)
				b.board[i][j] = int(number)
				b.marked[i][j] = false
			}
		}
	}
	return b
}
