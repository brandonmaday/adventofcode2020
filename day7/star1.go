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

    searchList := []string{"shiny gold"}
    validBags := make(map[string]bool)

    for (len(searchList) > 0) {
        var found []string
        for _, searchFor := range searchList {
            for bag, holds := range bags {
                for _, hold := range holds {
                    if hold.name == searchFor {
                        found = append(found, bag)
                        validBags[bag] = true
                    }
                }
            }
        }
        searchList = found
    }

    fmt.Println(len(validBags)) //370
}
