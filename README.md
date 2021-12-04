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
