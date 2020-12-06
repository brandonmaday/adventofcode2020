package main

import (
    "fmt"
    "strconv"

    "aoc2020/getpuzzle"
)

func main() {
    puzzle := getpuzzle.GetPuzzleFile("puzzle.txt")

    found := false
    for i, v1 := range puzzle {
        v1Int, _ := strconv.Atoi(v1)
        for x, v2 := range puzzle[i+1:] {
            v2Int, _ := strconv.Atoi(v2)
            for _, v3 := range puzzle[x+1:] {
                v3Int, _ := strconv.Atoi(v3)
                if v1Int + v2Int + v3Int == 2020 {
                    fmt.Print(v1Int * v2Int * v3Int) // 100655544
                    found = true
                    break
                }
            }
            if found {
                break
            }
        }
        if found {
            break
        }
    }
}
