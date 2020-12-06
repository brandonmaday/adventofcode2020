package main

import (
    "log"
    "fmt"
    "strings"
    "strconv"
    "regexp"

    "aoc2020/getpuzzle"
)

func between(min int, max int, val int) bool {
    return (min <= val && val <= max)
}

func validYr(min int, max int, val string) bool {
    if len(val) != 4 {
        return false
    }
    v, err := strconv.Atoi(val)
    if err != nil {
        return false
    }
    return between(min, max, v)
}

func validHgt(val string) bool {
    r := regexp.MustCompile(`(.+)(cm|in)`)
    matches := r.FindStringSubmatch(val)
    if len(matches) != 3 {
        return false
    }
    h, err := strconv.Atoi(matches[1])
    if err != nil {
        return false
    }
    if matches[2] == "cm" {
        return between(150, 193, h)
    }
    return between(59, 76, h)
}

func validHcl(val string) bool {
    r := regexp.MustCompile(`^(#)([0-9|a-f]{6})$`)
    matches := r.FindStringSubmatch(val)
    return len(matches) == 3
}

func validEcl(val string) bool {
    r := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
    return r.MatchString(val)
}

func validPid(val string) bool {
    r := regexp.MustCompile(`^\d{9}$`)
    return r.MatchString(val)
}

type passport struct {
    byr bool
    iyr bool
    eyr bool
    hgt bool
    hcl bool
    ecl bool
    pid bool
}



func main() {
    puzzle := getpuzzle.GetPuzzleFile("puzzle.txt")
    lineLen := len(puzzle)
    validPassports := 0
    aPassport := passport{}
    for i, line := range puzzle {
        fmt.Println(line)
        if i + 1 == lineLen || len(line) == 0 {
            if aPassport.byr && aPassport.iyr && aPassport.eyr && aPassport.hgt && aPassport.hcl && aPassport.ecl && aPassport.pid {
                validPassports++
            }
            fmt.Println(aPassport)
            aPassport = passport{}
            continue
        }
        data := strings.Split(line, " ")
        for _, attr := range data {
            field := strings.Split(attr, ":")
            switch field[0] {
            case "byr":
                aPassport.byr = validYr(1920, 2002, field[1])
            case "iyr":
                aPassport.iyr = validYr(2010, 2020, field[1])
            case "eyr":
                aPassport.eyr = validYr(2020, 2030, field[1])
            case "hgt":
                aPassport.hgt = validHgt(field[1])
            case "hcl":
                aPassport.hcl = validHcl(field[1])
            case "ecl":
                aPassport.ecl = validEcl(field[1])
            case "pid":
                aPassport.pid = validPid(field[1])
            default:
                if field[0] != "cid" {
                    fmt.Println(attr)
                    fmt.Println(field)
                    log.Fatal("No idea what passport attr is")
                }
            }
        }
    }
    fmt.Println(validPassports) //127
}
