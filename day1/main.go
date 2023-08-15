package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
    "sort"
)

type Elf struct {
    id int
    calories int
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

    bufScanner := bufio.NewScanner(input)

    elf := &Elf{id: 0, calories: 0}
    elves := []*Elf{}
    elves = append(elves, elf)

    for bufScanner.Scan() {
        text := bufScanner.Text()

        if text == "" {
            elf = &Elf{id: elf.id + 1, calories: 0}
            elves = append(elves, elf)
            continue
        }

        calories, err := strconv.Atoi(text)
        if err != nil {
            log.Fatal(err)
        }

        elf.calories = elf.calories + calories
    }

    sort.Slice(elves, func(i, j int) bool {
        return elves[i].calories > elves[j].calories
    })

    totalCalories := 0
    for _, elf := range elves[:3] {
        totalCalories = totalCalories + elf.calories
    }
    fmt.Println(totalCalories)
}
