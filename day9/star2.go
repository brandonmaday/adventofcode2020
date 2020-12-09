package main

import (
    "fmt"
    "strconv"

    "aoc2020/getpuzzle"
)

func invalidInt (ints []int, num int) bool {
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
    var ints []int
    for _, num := range intStrings {
        i, _ := strconv.Atoi(num)
        ints = append(ints, i)
    }
    preambleLen := 25
    invalidNum := 0
    for i := preambleLen; i < len(ints); i++ {
        if invalidInt(ints[i-preambleLen:i], ints[i]) {
            invalidNum = ints[i]
            break
        }
    }
    for i :=0; i < len(ints); i++ {
        for j := i+2; j < len(ints); j++ {
            min := 0
            max := 0
            total := 0
            for k, n := range ints[i:j] {
                if k == 0 {
                    min = n
                    max = n
                } else {
                    if n < min {
                        min = n
                    } else if n > max {
                        max = n
                    }
                }
                total += n
            }
            if invalidNum == total {
                fmt.Println(min + max)
            }
        }
    }
}
