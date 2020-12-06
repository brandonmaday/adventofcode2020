package main

import (
    "fmt"

    "aoc2020/getpuzzle"
)

func main() {
    puzzle := getpuzzle.GetPuzzleFile("puzzle.txt")
    totalCount := 0
    group := make(map[rune]bool)
    puzzleLen := len(puzzle)
    for i, line := range puzzle {
        for _, letter := range line {
            group[letter] = true
        }
        if line == "" || i+1 == puzzleLen {
            totalCount = totalCount + len(group)
            group = make(map[rune]bool)
            continue
        }
    }

    fmt.Println(totalCount)//6249
}
