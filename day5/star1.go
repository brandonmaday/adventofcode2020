package main

import (
    "fmt"

    "aoc2020/getpuzzle"
)

func seatId(row, col int) int {
    return (row * 8) + col
}

func halfIt(left bool, min, max int) (int, int) {
    h := max-((max-min+1)/2)
    if left {
        return min, h
    }
    return h+1, max
}

func walkToSeat(directions []bool, min, max int) (int) {
    for _, r := range directions {
        min, max = halfIt(r, min, max)
    }
    return min
}

func main() {
    puzzle := getpuzzle.GetPuzzleFile("puzzle.txt")
    highestId := 0
    for _, passport := range puzzle {
        lefts := make([]bool, 10)
        for i, letter := range passport {
            char := string(letter)
            if char == "F" || char == "L" {
                lefts[i] = true
            }
        }
        row := walkToSeat(lefts[:7], 0, 127)
        col := walkToSeat(lefts[7:], 0, 7)
        seat := seatId(row, col)
        if seat > highestId {
            highestId = seat
        }
    }
    fmt.Println(highestId) //855
}
