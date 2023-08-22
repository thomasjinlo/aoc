package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Range struct {
    Start int
    End int
}

func NewRange(start, end int) Range {
    return Range{
        Start: start,
        End: end,
    }
}

type Overlapper interface {
    Overlap(r1, r2 Range) bool
}

type OverlapperFunc func(r1, r2 Range) bool

func (f OverlapperFunc) Overlap(r1, r2 Range) bool {
    return f(r1, r2)
}

func FullOverlap(r1, r2 Range) bool {
    return ((r1.Start <= r2.Start && r1.End >= r2.End) ||
            (r2.Start <= r1.Start && r2.End >= r1.End ))
}

func PartialOverlap(r1, r2 Range) bool {
    return (r1.Start <= r2.End && r2.Start <= r1.End) ||
           (r2.Start <= r1.End && r1.Start <= r2.End)
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }

    overlap := 0
    for bufScanner := bufio.NewScanner(file); bufScanner.Scan(); {
        line := bufScanner.Text()

        matchDigits := `\d+`
        re := regexp.MustCompile(matchDigits)
        digits := re.FindAllString(line, -1)
        numbers := []int{}

        for _, digit := range digits {
            if number, err := strconv.Atoi(digit); err == nil {
                numbers = append(numbers, number)
            }
        }

        range1 := NewRange(numbers[0], numbers[1])
        range2 := NewRange(numbers[2], numbers[3])

        doesOverlap := OverlapperFunc(PartialOverlap)
        if doesOverlap(range1, range2) {
            overlap ++
        }
    }

    fmt.Println(overlap)
}
