package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type CrateMover interface {
    Move(stacks *[]string, numCrates, from, to int)
}

type CrateMoverFunc func(stacks *[]string, numCrates, from, to int)

func (f CrateMoverFunc) Move(stacks *[]string, numCrates, from, to int) {
    f(stacks, numCrates, from, to)
}

func Crate9000Mover(stacks *[]string, numCrates, from, to int) {
    derefStacks := (*stacks)
    crates := ""
    for _, crate := range derefStacks[from][len(derefStacks[from])-numCrates:] {
        crates = string(crate) + crates
    }
    derefStacks[from] = derefStacks[from][:len(derefStacks[from])-numCrates]
    derefStacks[to] += crates
}

func Crate9001Mover(stacks *[]string, numCrates, from, to int) {
    derefStacks := (*stacks)
    crates := derefStacks[from][len(derefStacks[from])-numCrates:]
    derefStacks[from] = derefStacks[from][:len(derefStacks[from])-numCrates]
    derefStacks[to] += crates
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }

    crateArrangementLines := []string{}
    isCrateArrangementLine := true
    stacks := []string{}
    var line string

    for bufScanner := bufio.NewScanner(file); bufScanner.Scan(); {
        line = bufScanner.Text()

        if len(line) == 0 {
            isCrateArrangementLine = false
            numStacks := len(crateArrangementLines) - 1
            lineLength := len(crateArrangementLines[0])

            for i := 0; i <= numStacks; i++ {
                stacks = append(stacks, "")
            }

            for _, crateLine := range crateArrangementLines[:numStacks + 1] {
                for i := 1; i < lineLength; i += 4 {
                    char := crateLine[i:i+1]

                    if char == " " {
                        char = ""
                    }

                    stacks[i/4] = char + stacks[i/4]
                }
            }

            continue
        }

        if isCrateArrangementLine {
            crateArrangementLines = append(crateArrangementLines, line)
            continue
        }

        matchDigits := `\d+`
        re := regexp.MustCompile(matchDigits)
        digits := re.FindAllString(line, -1)
        numCrates, _ := strconv.Atoi(digits[0])
        from, _ := strconv.Atoi(digits[1])
        to, _ := strconv.Atoi(digits[2])


        crateMoverFunc := CrateMoverFunc(Crate9001Mover)
        crateMoverFunc.Move(&stacks, numCrates, from-1, to-1)
    }

    result := ""

    for _, stack := range stacks {
        i := len(stack) - 1
        result = result + string(stack[i])
    }

    fmt.Println(result)
}
