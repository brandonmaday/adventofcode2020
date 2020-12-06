package main

import (
    "fmt"
    "strconv"

    "aoc2020/getpuzzle"
)

func between(min int, max int, val int) bool {
    return min <= val && val <= max
}

type validCk func(string) bool

func validByr(val: string) bool {
    v, _ := strconv.Atoi(val)
    return between(1920, 2020, v)
}

func validIyr(val: string) bool {
    v, _ := strconv.Atoi(val)
    return between(2010, 2020, v)
}

func validEyr(val: string) bool {
    v, _ := strconv.Atoi(val)
    return between(2010, 2020, v)
}

func main() {
    puzzle := getpuzzle.GetPuzzleFile("puzzle.txt")
    lineLen := len(puzzle)
    reqs := map[string]validCk{
        "byr": validByr,
        "iyr": validIyr,
        "eyr": validEyr,
        "hgt": validHgt,
        "hcl": validHcl,
        "ecl": validEcl,
        "pid": validPid,
    }
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
    fmt.Println(validPassports)
}
