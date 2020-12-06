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
        for _, v2 := range puzzle[i+1:] {
            v2Int, _ := strconv.Atoi(v2)
            if v1Int + v2Int == 2020 {
                fmt.Println(v1Int * v2Int) //1019571
                found = true
                break
            }
        }
        if found {
            break
        }
    }
}
