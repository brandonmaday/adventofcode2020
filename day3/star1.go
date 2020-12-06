package main

import (
    "fmt"

    "aoc2020/getpuzzle"
)

func main() {
    puzzle := getpuzzle.GetPuzzleFile("puzzle.txt")
    x, trees := 0, 0
    for _, line := range puzzle[1:] {
        l := len(line) - 1
        x = x + 3
        if x > l {
            x = x-l-1
        }
        if string(line[x]) == "#" {
            trees++
        }
    }
    fmt.Println(trees) // 242
}
