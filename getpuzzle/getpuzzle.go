package getpuzzle

import (
    "io/ioutil"
    "log"
    "strings"
)

func GetPuzzleFile(filename string) []string {
    puzzleData, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }
    puzzle := strings.Split(string(puzzleData), "\n")
    puzzle = puzzle[:len(puzzle)-1]
    return puzzle
}
