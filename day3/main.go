package main

import (
	"bufio"
	"fmt"
    "log"
	"os"
)

func calculatePriority(chr rune) int {
    if ord := int(chr); ord <= 90 {
        return int(chr) - int('A') + 27
    } else {
        return int(chr) - int('a') + 1
    }
}

func main() {
    input, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }

    totalPriority := 0
    groupLine := 1
    firstSet := map[rune]bool{}
    secondSet := map[rune]bool{}
    thirdSet := map[rune]bool{}

    for bufScanner := bufio.NewScanner(input); bufScanner.Scan(); {
        items := bufScanner.Text()

        for _, item := range items {
            switch groupLine {
            case 1:
                firstSet[item] = true
            case 2:
                secondSet[item] = true
            case 3:
                thirdSet[item] = true
            }
        }

        if groupLine == 3 {
            for item := range firstSet {
                _, okSecond := secondSet[item]
                _, okThird := thirdSet[item]

                if okSecond && okThird {
                    totalPriority += calculatePriority(item)
                    break
                }
            }

            firstSet = map[rune]bool{}
            secondSet = map[rune]bool{}
            thirdSet = map[rune]bool{}
            groupLine = 1
        } else {
            groupLine ++
        }

    }

    fmt.Println(totalPriority)
}

