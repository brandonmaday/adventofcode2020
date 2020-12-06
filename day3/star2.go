package main

import (
    "fmt"

    "aoc2020/getpuzzle"
)

func main() {
    puzzle := getpuzzle.GetPuzzleFile("puzzle.txt")
    trees := [5]int{0,0,0,0,0}
    rights := [5]int{1,3,5,7,1}
    downs := [5]int{1,1,1,1,2}
    for i, right := range rights {
        x := 0
        for z, line := range puzzle[downs[i]:] {
            if z % downs[i] != 0 {
                continue
            }
            l := len(line) - 1
            x = x + right
            if x > l {
                x = x-l-1
            }
            if string(line[x]) == "#" {
                trees[i]++
            }
        }
    }
    answer := trees[0]
    for _, tree := range trees[1:] {
        answer = answer * tree
    }
    fmt.Println(answer) // 2265549792
}
