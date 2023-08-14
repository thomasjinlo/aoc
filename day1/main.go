package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

    bufScanner := bufio.NewScanner(input)

    var (
        maxCalories int
        numCalories int
    )

    for bufScanner.Scan() {
        text := bufScanner.Text()

        if text == "" {
            numCalories = 0
            continue
        }

        intCalories, err := strconv.Atoi(text)
        if err != nil {
            log.Fatal(err)
        }

        numCalories = numCalories + intCalories

        if numCalories > maxCalories {
            maxCalories = numCalories 
        }
    }

    fmt.Println(maxCalories)
}
