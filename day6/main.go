package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type MarkerNotFoundError struct {}

func (e *MarkerNotFoundError) Error() string {
    return "Unable to find marker"
}

func findMarker(line string) (int, error) {
    charSet := map[byte]bool {}

    for i, j := 0, 0; i < len(line); i++ {
        _, ok := charSet[line[i]]
        for ok {
            delete(charSet, line[j])

            j ++
            _, ok = charSet[line[i]]
        }

        charSet[line[i]] = true

        if (i - j) == 3 {
            return i + 1, nil
        }
    }

    return 0, &MarkerNotFoundError{}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	for bufScanner := bufio.NewScanner(file); bufScanner.Scan(); {
		line := bufScanner.Text()

        marker, err := findMarker(line)
        if err == nil {
            fmt.Println(marker)
        }
	}
}
