package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"./bingo"
)

func main() {
	filecontent := readfile("testinput.txt")

	// This is gross, but works. First element of filecontent is our plays.
	var number_draws = strings.Split(filecontent[0], ",")
	fmt.Println(number_draws)

	var boards_rows []string
	// Skip the first row, since it's the list of number draws.
	for i := 1; i < len(filecontent); i++ {
		if len(filecontent[i]) > 0 {
			boards_rows = append(boards_rows, filecontent[i])
		}
	}

	// Loop over our board_rows, and add 5 rows to a new Bingo.
	var bingos []bingo.Bingo
	var rawrows []string
	for i := 0; i < len(boards_rows); i++ {
		// Append the current row to the list of rawrows
		rawrows = append(rawrows, boards_rows[i])

		// Once that list of rawrows is 5, we're done with this card - create a new Bingo, give it the rows, start over.
		if len(rawrows) == 5 {
			b := bingo.New(rawrows)
			bingos = append(bingos, b)
			rawrows = nil
		}
	}
	// boards_rows has all valid rows of each board. Now how many boards do we have?
	// board_count := len(boards_rows)/5

	fmt.Println(bingos)
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
