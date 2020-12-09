package main

import (
    "fmt"
    "strconv"

    "aoc2020/getpuzzle"
)

func parseInt (intStr string) int {
    num, _ := strconv.Atoi(intStr)
    return num
}

func invalidInt (nums []string, num int) bool {
    var ints []int
    for _, num := range nums {
        ints = append(ints, parseInt(num))
    }
    for i, numOne := range ints {
        for _, numTwo := range ints[i+1:] {
            if numOne + numTwo == num {
                return false
            }
        }
    }
    return true
}

func main() {
    intStrings := getpuzzle.GetPuzzleFile("puzzle.txt")
    preambleLen := 25
    for i := preambleLen; i < len(intStrings); i++ {
        num := parseInt(intStrings[i])
        if invalidInt(intStrings[i-preambleLen:i], num) {
            fmt.Println(num)
            break
        }
    }
}
