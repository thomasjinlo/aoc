package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
    input, err := os.Open("input.txt")
    if err != nil {
        log.Fatal("input.txt")
    }

    const (
        elfRock = 65
        elfPaper = 66
        elfScissor = 67
        myRock   = 88
        myPaper = 89
        myScissor = 90
    )

    totalScore := 0
    calculateScore := map[uint8]map[uint8]int{
        elfRock: {
            myRock: 3,
            myPaper: 6,
            myScissor: 0,
        },
        elfPaper: {
            myRock: 0,
            myPaper: 3,
            myScissor: 6,
        },
        elfScissor: {
            myRock: 6,
            myPaper: 0,
            myScissor: 3,
        },
    }
    calculateFreeScore := map[uint8]int{
        myRock: 1,
        myPaper: 2,
        myScissor: 3,
    }

    for bufScanner := bufio.NewScanner(input); bufScanner.Scan(); {
        moves := bufScanner.Text()
        elfMove, myMove := moves[0], moves[2]

        totalScore += calculateFreeScore[myMove]
        totalScore += calculateScore[elfMove][myMove]
    }

    fmt.Println(totalScore)
}
