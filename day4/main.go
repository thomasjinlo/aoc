package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }

    overlap := 0
    for bufScanner := bufio.NewScanner(file); bufScanner.Scan(); {
        line := bufScanner.Text()

        ranges := strings.Split(line, ",")

        firstRange := strings.Split(ranges[0], "-")
        secondRange := strings.Split(ranges[1], "-")

        firstStart, _ := strconv.Atoi(firstRange[0])
        firstEnd, _ := strconv.Atoi(firstRange[1])
        secondStart, _ := strconv.Atoi(secondRange[0])
        secondEnd, _ := strconv.Atoi(secondRange[1])

        if firstStart <= secondStart && firstEnd >= secondEnd {
            overlap ++
            continue
        }

        if secondStart <= firstStart && secondEnd >= firstEnd {
            overlap ++
        }
    }

    fmt.Println(overlap)
}
