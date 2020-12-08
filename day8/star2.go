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

func nop(acc int, i int) (int, int) {
    return acc, i + 1
}

func jmp(cmd Cmd, acc int, i int) (int, int) {
    return acc, chgVal(cmd.dir, i, cmd.offset)
}

func process(cmds []Cmd, chgI int) (acc int, i int) {
    acc = 0
    i = 0
    for x :=0; x < len(cmds); x++ {
        cmds[x].ran = false
    }
    for i < len(cmds) {
        if cmds[i].ran == true {
            break
        }
        cmds[i].ran = true

        switch cmds[i].cmd {
        case "nop":
            if i == chgI {
                acc, i = jmp(cmds[i], acc, i)
            } else {
                acc, i = nop(acc, i)
            }
        case "acc":
            acc = chgVal(cmds[i].dir, acc, cmds[i].offset)
            i++
        case "jmp":
            if i == chgI {
                acc, i = nop(acc, i)
            } else {
                acc, i = jmp(cmds[i], acc, i)
            }
        }
    }
    return
}

func main() {
    puzzle := getpuzzle.GetPuzzleFile("puzzle.txt")
    puzzleLineRe := regexp.MustCompile(`^(.+) ([+|-])(.+)$`)

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
    // goalI := len(cmds)-1
    i := 0
    acc := 9
    cmdsLen := len(cmds)
    for chgI := 0; chgI < cmdsLen; chgI++ {
        i = 0
        acc = 0
        if cmds[chgI].cmd == "nop" || cmds[chgI].cmd == "jmp" {
            acc, i = process(cmds, chgI)
            if i >= cmdsLen-1 {
                break
            }
        }
    }
    fmt.Println(acc) //703
}
