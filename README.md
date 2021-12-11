# Tim's Advent of Code 2021

Some random thoughts as I am working through Advent Of Code 2021.

## Day 1

Chris pointed out that if you're checking `b + c + d > a + b + c`, then you might as well check that `d > a` since the other factors are equal. Funny enough, this works :P

## Day 3

In part 2, I did not understand the assignment. I kept using the full list of commands, not the subset of commands after the initial filtering.

I wasn't able to leverage the work I did in the first part - which depended on the information from the WHOLE collection.

## Day 4

These bingo boards sound like the should be classes, which I am woefully weak with in Go.

Some points to remember:
- member names in the `struct` that are lower case are effectively private.
- You "hang" methods off the struct. There is no class definition.
- `range` seemed to be hanging me up when calling a method on a class.  I would call the method, observe that the object changed, but when iterating over a list of objects with `range`, it seems that `range` was returning a copy of the object - not the reference to the object.  The fix is to iterate over the slice with more conventional menas:

```go
	for d := 0; d < len(draws); d++ {
		draw_number, _ := strconv.ParseInt(draws[d], 10, 64)
			for i := 0; i < len(bingos); i++ {
				bingos[i].PlayNumber(int(draw_number))
```
- Also, "class members" (there are no classes, just `struct`s with functions that hang off them) work best when referenced by a pointer.  Otherwise..  you get a copy. The behaviour is to copy-by-value, which isn't what we want when trying to change internal properties of an object.

Lessons in OOP in Golang.

## Day 5

I spent quite a bit of time getting an understanding of where files go in Go.

I'm using a container for my Go stuff, so there's some "magic" that seems to be happening.  In the end, my `devcontainer.json` file these settings:

```
	"settings": {
		"go.inferGopath": true,
		"go.toolsManagement.checkForUpdates": "local",
		"go.useLanguageServer": true,
		// "go.gopath": "/go",
		"go.goroot": "/usr/local/go"
	},
```

This still has `/go` set as GOPATH, but also `/workspaces` as a default path (defined by `go.inferGopath`).  Again, magic.

I then moved all my stuff to `/src`, because Go is expecting to find packages under `$GOPATH/src/`.  Final directory tree looks like:
```
vscode ➜ /workspaces/AoC-2021 (main ✗) $ tree ./
./
├── README.md
└── src
    ├── Day4
    │   ├── bingo
    │   │   └── bingo.go
    │   ├── day4input.txt
    │   ├── main.go
    │   └── testinput.txt
    └── Day5
        ├── line
        │   └── line.go
        ├── main.go
        ├── testinput.txt
        └── vertex
            └── vertex.go
```

Package imports are then called like:

```go
# main.go
package main

import (
	"Day5/line"
	"bufio"
	"fmt"
	"log"
	"os"
)
```

Packages that are in a subdirectory follow the same pattern:
```go
# line/line.go

package line

import (
	"Day5/vertex"
	"fmt"
	"regexp"
)

type Line struct {
	start vertex.Vertex
	end   vertex.Vertex
}
```

I also moved around Day 4 to follow this setup.

https://stackoverflow.com/questions/36017724/can-i-have-multiple-gopath-directories

### Declaring 2d arrays
I need to make a fixed length 2d array based on the maximum size of the X and Y coordinates of the lines.  This means I can't statically declare an array (go doesn't like that..).

Found https://stackoverflow.com/questions/39804861/what-is-a-concise-way-to-create-a-2d-slice-in-go/39806983 , but the memory-contiguous way I couldn't quite wrap my skull around, so I went for the "easy" way.

### Unit Testing

Started playing around with this.  VSCode ships with some Table Driven Tests.  Detailed at https://dave.cheney.net/2019/05/07/prefer-table-driven-tests.

### Learned about fmt.Sscanf

From https://skarlso.github.io/2021/12/05/aoc-day5/.

Replaced a bunch of Regex:
```go
		var compiledRegex = regexp.MustCompile(`^(?P<x1>\d+),(?P<y1>\d+) -> (?P<x2>\d+),(?P<y2>\d+)$`)
		matches := compiledRegex.FindStringSubmatch(rowInfo)
		x1 := matches[compiledRegex.SubexpIndex("x1")]
		y1 := matches[compiledRegex.SubexpIndex("y1")]
		x2 := matches[compiledRegex.SubexpIndex("x2")]
		y2 := matches[compiledRegex.SubexpIndex("y2")]
```

with one function call: `fmt.Sscanf(rowInfo, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)`

Will have to try to remember this in the future. This would make parsing a lot of these inputs considerably quicker.

## Day 6

I looked at ths before I went to bed, and thought "ok, I'll make a Fish object, then that object will keep track of it's own age".  Easy. I would need to figure out how to "tick" each fish to increment their day count, then figure out how use a member method as an object factory ..  all new to me, but stuff to learn.

The next morning, I saw the memes on the AoC Subreddit, people talking about exhausting compute resources, etc.  And someone mentioned **being able to solve it in compile time**.

The thing is..  you don't need an object to represent a fish. It doesn't get us anyhting, other than an exponentially growing set of objects. A fish can only be around for 9 days at absolute maximum, and we just need to change counters of an age index on a counter.

So..  basically a map, or an array, of ages and fish counts.

## Day 8

Had a real hard time understanding this one for a bit.

Followed this "logic on paper" solution:
https://www.reddit.com/r/adventofcode/comments/rbvpui/2021_day_8_part_2_my_logic_on_paper_i_used_python/

## Day 9

Fun with recusion. I forgot to set a flag so that the recursive function wouldn't just "bounce" back and forth between X-1 and X+1. This was fixed by setting the elevation to 9 once we counted it as part of the basin.

## Day 10

Hard time with this one, probably because of the day I had.

Wound up treating the list like a stack, and popping elements off it as they were "allowed".  There's a lot of copy-pasta here, and opportunities to improve the code considerably.
