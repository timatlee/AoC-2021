package bingo

import (
	"fmt"
	"strconv"
	"strings"
)

type Bingo struct {
	board   [5][5]int
	marked  [5][5]bool
	rawRows []string
	IsWon   bool
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

func (b Bingo) String() string {
	var output string
	// non_marked := color.New(color.FgBlue).SprintFunc()
	//marked := color.new(color.FgCyan).SprintFunc()

	for i := 0; i < len(b.board); i++ {
		for j := 0; j < len(b.board[i]); j++ {
			if b.marked[i][j] {
				output += fmt.Sprintf("m:%2d ", b.board[i][j])
			} else {
				output += fmt.Sprintf("%4d ", b.board[i][j])
			}
		}
		output += fmt.Sprintln()
	}
	return output
}

func (b *Bingo) PlayNumber(number int) bool {
	for i := 0; i < len(b.board); i++ {
		for j := 0; j < len(b.board[i]); j++ {
			if b.board[i][j] == number {
				b.marked[i][j] = true
			}
		}
	}

	return b.checkBoard()
}

func (b *Bingo) GetScore() int {
	var unmarkedSum int = 0
	for i := 0; i < len(b.board); i++ {
		for j := 0; j < len(b.board[i]); j++ {
			if !b.marked[i][j] {
				unmarkedSum += b.board[i][j]
			}
		}
	}
	return unmarkedSum

}

func (b *Bingo) checkBoard() bool {
	// Check if there are any rows that are winning.  I is Y axis, J is X.
	var foundRow bool = false
	for i := 0; i < len(b.marked); i++ {
		var rowTrue bool = true
		for j := 0; j < len(b.marked[i]); j++ {
			if rowTrue && !b.marked[i][j] {
				rowTrue = false
			}
		}
		if rowTrue {
			foundRow = true
			break
		}
	}

	if foundRow {
		b.IsWon = true
		return true
	}

	// Check if there are any rows that are winning.  I is Y axis, J is X.
	var foundCol bool = false
	for j := 0; j < 5; j++ {
		var colTrue bool = true
		for i := 0; i < 5; i++ {
			if colTrue && !b.marked[i][j] {
				colTrue = false
			}
		}
		if colTrue {
			foundCol = true
			break
		}
	}

	if foundCol {
		b.IsWon = true
		return true
	}

	b.IsWon = false
	return false
}
