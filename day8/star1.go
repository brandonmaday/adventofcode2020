package main

import (
    "fmt"
    "regexp"
    "strconv"

    "aoc2020/getpuzzle"
)

type Cmd struct {
    cmd string
    dir string
    offset int
    ran bool
}

func chgVal(dir string, val, chg int) int {
    if dir == "+" {
        return val + chg
    } else {
        return val - chg
    }
}

func main() {
    puzzle := getpuzzle.GetPuzzleFile("puzzle.txt")
    puzzleLineRe := regexp.MustCompile(`^(.+) ([+|-])(.+)$`)

    acc := 0
    var cmds []Cmd
    // Convert puzzle input in to command list
    for _, puzzleLine := range puzzle {
        matches := puzzleLineRe.FindStringSubmatch(puzzleLine)
        offset, _ := strconv.Atoi(matches[3])
        cmds = append(cmds, Cmd{
            matches[1],
            matches[2],
            offset,
            false,
        })
    }
    for i := 0; i < len(cmds); {
        if cmds[i].ran == true {
            break
        }
        cmds[i].ran = true

        switch cmds[i].cmd {
        case "nop":
            i++
        case "acc":
            acc = chgVal(cmds[i].dir, acc, cmds[i].offset)
            i++
        case "jmp":
            i = chgVal(cmds[i].dir, i, cmds[i].offset)
        }
    }
    fmt.Println(acc) //1586
}
