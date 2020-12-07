package main

import (
    "fmt"
    "regexp"
    "strings"
    "strconv"

    "aoc2020/getpuzzle"
)

type Holds struct {
    num int
    name string
}

func main() {
    puzzle := getpuzzle.GetPuzzleFile("puzzle.txt")

    // parse puzzel
    bags := make(map[string][]Holds)
    puzzleLineRe := regexp.MustCompile(`^(.+?) bags contain (.+)\.$`)
    holdsLineRe := regexp.MustCompile(`^(\d+) (.+?) bag(s?)`)
    for _, puzzleLine := range puzzle {
        matches := puzzleLineRe.FindStringSubmatch(puzzleLine)
        var holds []Holds
        if matches[2] != "no other bags" {
            holdLines := strings.Split(matches[2], ", ")
            for _, holdLine := range holdLines {
                otherBag := holdsLineRe.FindStringSubmatch(holdLine)
                cnt, _ := strconv.Atoi(otherBag[1])
                holds = append(holds, Holds{cnt, otherBag[2]})
            }
        }
        bags[matches[1]] = holds
    }

    // check rules

    bagCount := -1
    searchList := []string{"shiny gold"}

    for (len(searchList) > 0) {
        var found []string
        for _, searchFor := range searchList {
            bagCount++
            for _, hold := range bags[searchFor] {
                for i := 0; i < hold.num; i++ {
                    found = append(found, hold.name)
                }
            }
        }
        searchList = found
    }

    fmt.Println(bagCount) // 29547
}
