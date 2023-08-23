package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

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

        crates := ""
        for _, crate := range stacks[from-1][len(stacks[from-1])-numCrates:] {
            crates = string(crate) + crates
        }
        stacks[from-1] = stacks[from-1][:len(stacks[from-1])-numCrates]
        stacks[to-1] += crates
    }

    result := ""

    for _, stack := range stacks {
        i := len(stack) - 1
        result = result + stack[i:i+1]
    }

    fmt.Println(result)
}
