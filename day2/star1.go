package main

import (
    "fmt"
    "strconv"
    "regexp"

    "aoc2020/getpuzzle"
)

func main() {
    puzzle := getpuzzle.GetPuzzleFile("puzzle.txt")

    goodPasswords := 0

    for _, puzzleLine := range puzzle {
        r := regexp.MustCompile(`(\d+)-(\d+) ([a-z]): ([[:alpha:]]+)`)
        matches := r.FindStringSubmatch(puzzleLine)

        min, _ := strconv.Atoi(matches[1])
        max, _ := strconv.Atoi(matches[2])

        letterCnt := 0
        for _, l := range matches[4] {
            if string(l) == matches[3] {
                letterCnt++
            }
        }
        if letterCnt >= min && letterCnt <= max {
            goodPasswords++
        }
    }
    fmt.Println(goodPasswords) // 569
}
