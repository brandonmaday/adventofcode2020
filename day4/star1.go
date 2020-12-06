package main

import (
    "fmt"
    "strings"

    "aoc2020/getpuzzle"
)

func main() {
    puzzle := getpuzzle.GetPuzzleFile("puzzle.txt")
    lineLen := len(puzzle)
    reqs := [7]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
    validPassports := 0
    passportCheck := 0
    for i, line := range puzzle {
        data := strings.Split(line, " ")
        for _, attr := range data {
            field := strings.Split(attr, ":")
            for _, req := range reqs {
                if field[0] == req {
                    passportCheck++
                    break
                }
            }
        }
        if i + 1 == lineLen || len(line) == 0 {
            if passportCheck == 7 {
                validPassports++
            }
            passportCheck = 0
        }
    }
    fmt.Println(validPassports) //219
}
