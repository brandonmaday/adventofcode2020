package main

import (
    "fmt"
    "strconv"
    "sort"

    "aoc2020/getpuzzle"
)


func main() {
    puzzleLines := getpuzzle.GetPuzzleFile("puzzle.txt")
    var adapters []int
    for _, n := range puzzleLines {
        i, _ := strconv.Atoi(n)
        adapters = append(adapters, i)
    }
    sort.Ints(adapters)
    adapters = append(adapters, adapters[len(adapters)-1]+3)
    currJolt := 0
    jumps := make(map[int]int)
    for _, adapter := range adapters {
        jump := adapter-currJolt
        if jump <= 3 {
            jumps[jump]++
            currJolt = adapter
        }
    }
    fmt.Println(jumps[1]*jumps[3])
}
