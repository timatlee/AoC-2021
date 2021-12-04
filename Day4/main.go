package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"./bingo"
)

func main() {
	filecontent := readfile("day4input.txt")
	draws, bingos := parse_input_data(filecontent)

	fmt.Println("Numbers drawn:", draws)
	//fmt.Println("Bingo cards:", bingos)

	var hasWin bool = false
	var winningNumber int = 0
	var winningBoardindex int = -1
	for d := 0; d < len(draws); d++ {
		draw_number, _ := strconv.ParseInt(draws[d], 10, 64)
		if !hasWin {
			for i := 0; i < len(bingos); i++ {
				bingos[i].PlayNumber(int(draw_number))

				if bingos[i].IsWon {
					fmt.Println("Found winning number", draw_number, "board", i)
					winningNumber = int(draw_number)
					winningBoardindex = i
					//hasWin = true
					//break
				}

				// Check if all boards are in the winningBoards list
				if check_all_boards_played(bingos) {
					hasWin = true
					break
				}
			}
		}

	}

	fmt.Println("Winning board:", winningBoardindex, "Winning number", winningNumber, "Score", bingos[winningBoardindex].GetScore(), "final score", winningNumber*bingos[winningBoardindex].GetScore())

	//fmt.Println("Bingo cards:\n", bingos)

}

func check_all_boards_played(s []bingo.Bingo) bool {
	for _, board := range s {
		if !board.IsWon {
			return false
		}
	}
	return true
}

func parse_input_data(filecontent []string) ([]string, []bingo.Bingo) {
	// This is gross, but works. First element of filecontent is our plays.
	var number_draws = strings.Split(filecontent[0], ",")

	var bingos []bingo.Bingo
	var rawrows []string
	// Skip the first row, since it's the list of number draws.
	for i := 1; i < len(filecontent); i++ {
		if len(filecontent[i]) > 0 {
			// Append the current row to the list of rawrows
			rawrows = append(rawrows, filecontent[i])

			// Once that list of rawrows is 5, we're done with this card - create a new Bingo, give it the rows, start over.
			if len(rawrows) == 5 {
				b := bingo.New(rawrows)
				bingos = append(bingos, b)
				rawrows = nil
			}
		}
	}
	return number_draws, bingos
}

func readfile(filename string) []string {
	dir, _ := os.Getwd()

	file, err := os.Open(dir + "/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var commands []string

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		commands = append(commands, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return commands
}
