package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
    A = 65
    B = 66
    C = 67
    X   = 88
    Y = 89
    Z = 90
)

type rpsScorer interface {
    calculateWinningScore(elfMove, myMove uint8) int
    calculateFreeScore (move uint8) int
}

type naiveScorer struct {}

func (s *naiveScorer) calculateWinningScore(elfMove, myMove uint8) int {
    calculateScore := map[uint8]map[uint8]int{
        A: {
            X: 3,
            Y: 6,
            Z: 0,
        },
        B: {
            X: 0,
            Y: 3,
            Z: 6,
        },
        C: {
            X: 6,
            Y: 0,
            Z: 3,
        },
    }

    return calculateScore[elfMove][myMove]
}

func (s *naiveScorer) calculateFreeScore(move uint8) int {
    calculateFreeScore := map[uint8]int{
        X: 1,
        Y: 2,
        Z: 3,
    }

    return calculateFreeScore[move]
}

type smartScorer struct {}

func (s *smartScorer) calculateWinningScore(elfMove, myMove uint8) int {
    calculateScore := map[uint8]map[uint8]int{
        A: {
            X: 3,
            Y: 1,
            Z: 2,
        },
        B: {
            X: 1,
            Y: 2,
            Z: 3,
        },
        C: {
            X: 2,
            Y: 3,
            Z: 1,
        },
    }

    return calculateScore[elfMove][myMove]
}

func (s *smartScorer) calculateFreeScore(move uint8) int {
    calculateFreeScore := map[uint8]int{
        X: 0,
        Y: 3,
        Z: 6,
    }

    return calculateFreeScore[move]
}

func main() {
    input, err := os.Open("input.txt")
    if err != nil {
        log.Fatal("input.txt")
    }

    totalScore := 0
    scorer := smartScorer{}

    for bufScanner := bufio.NewScanner(input); bufScanner.Scan(); {
        moves := bufScanner.Text()
        elfMove, myMove := moves[0], moves[2]

        totalScore += scorer.calculateFreeScore(myMove)
        totalScore += scorer.calculateWinningScore(elfMove, myMove)
    }

    fmt.Println(totalScore)
}
