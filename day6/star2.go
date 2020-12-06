package main

import (
    "fmt"
    "strings"

    "aoc2020/getpuzzle"
)

func main() {
    puzzle := getpuzzle.GetPuzzleFile("puzzle.txt")
    totalCount := 0
    group := make(map[string]bool)
    groupIdx := 0
    puzzleLen := len(puzzle)
    for i, line := range puzzle {
        if line != "" {
            for letter := range group {
                if strings.Contains(line, letter) == false {
                    delete(group, letter)
                }
            }
        }
        if groupIdx == 0 {
            for _, letter := range line {
                group[string(letter)] = true
            }
        }
        if line == "" || i+1 == puzzleLen {
            totalCount = totalCount + len(group)
            group = make(map[string]bool)
            groupIdx = 0
            continue
        }
        groupIdx++
    }

    fmt.Println(totalCount) // 3103
}
