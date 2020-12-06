package main

import (
    "fmt"
    "strconv"
    "regexp"

    "aoc2020/getpuzzle"
)

// func
func main() {
    puzzle := getpuzzle.GetPuzzleFile("puzzle.txt")

    goodPasswords := 0

    for _, puzzleLine := range puzzle {
        r := regexp.MustCompile(`(\d+)-(\d+) ([a-z]): ([[:alpha:]]+)`)
        matches := r.FindStringSubmatch(puzzleLine)

        min, _ := strconv.Atoi(matches[1])
        max, _ := strconv.Atoi(matches[2])

        matchCnt := 0
        if (string(matches[4][min-1]) == matches[3]) {
            matchCnt++
        }
        if (string(matches[4][max-1]) == matches[3]) {
            matchCnt++
        }
        if matchCnt == 1 {
            goodPasswords++
        }
    }
    fmt.Println(goodPasswords) // 346
}
