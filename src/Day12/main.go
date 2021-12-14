package main

import (
	"fmt"
	"strings"
	"unicode"

	aocutils "github.com/timatlee/AoC-Common"
)

func main() {
	content := aocutils.Readfile("input.txt")

	exits := map[string][]string{}
	visits := map[string]int{}

	for _, l := range content {
		caves := strings.Split(l, "-")

		exits[string(caves[0])] = append(exits[string(caves[0])], string(caves[1]))
		exits[string(caves[1])] = append(exits[string(caves[1])], string(caves[0]))
	}

	part1 := allPossiblePathsVisitOnce(exits, "start", visits)
	fmt.Printf("part1: %v\n", part1)

	exits = map[string][]string{}
	visits = map[string]int{}

	for _, l := range content {
		caves := strings.Split(l, "-")

		exits[string(caves[0])] = append(exits[string(caves[0])], string(caves[1]))
		exits[string(caves[1])] = append(exits[string(caves[1])], string(caves[0]))
	}

	part2 := allPossiblePathsOneDoubleVisit(exits, "start", visits, false)

	fmt.Printf("part2: %v\n", part2)
}

func allPossiblePathsOneDoubleVisit(exits map[string][]string, currentCave string, visits map[string]int, hasVisitedTwice bool) int {
	if visits[currentCave] > 0 && unicode.IsLower(rune(currentCave[0])) {
		hasVisitedTwice = true
	}
	visits[currentCave] += 1
	sum := 0
	for _, c := range exits[currentCave] {
		if c == "end" {
			sum++
			continue
		}
		if c != "start" && (!unicode.IsLower(rune(c[0])) || visits[c] < 1 || !hasVisitedTwice) {
			nextVisits := visits
			if unicode.IsLower(rune(c[0])) {
				nextVisits = copyMap(visits)
			}
			sum += allPossiblePathsOneDoubleVisit(exits, c, nextVisits, hasVisitedTwice)
		}
	}
	return sum
}

func allPossiblePathsVisitOnce(exits map[string][]string, currentCave string, visits map[string]int) int {
	visits[currentCave] += 1
	sum := 0
	for _, c := range exits[currentCave] {
		if c == "end" {
			sum++
			continue
		}
		if !unicode.IsLower(rune(c[0])) || visits[c] == 0 {
			nextVisits := visits
			if unicode.IsLower(rune(c[0])) {
				nextVisits = copyMap(visits)
			}
			sum += allPossiblePathsVisitOnce(exits, c, nextVisits)
		}
	}
	return sum
}

func copyMap(visits map[string]int) map[string]int {
	res := make(map[string]int, len(visits))
	for k, v := range visits {
		res[k] = v
	}
	return res
}
