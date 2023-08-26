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

type MarkerFinder interface {
    FindMarker(line string) (int, error)
}

type MarkerFinderFunc func(line string) (int, error)

func (f MarkerFinderFunc) FindMarker(line string) (int, error) {
    return f(line)
}

func findContiguousUniqChars(line string, uniqCount int) (int, error) {
    charSet := map[byte]bool {}

    for i, j := 0, 0; i < len(line); i++ {
        _, ok := charSet[line[i]]
        for ok {
            delete(charSet, line[j])

            j ++
            _, ok = charSet[line[i]]
        }

        charSet[line[i]] = true

        if (i - j) == uniqCount {
            return i + 1, nil
        }
    }

    return 0, &MarkerNotFoundError{}
}

func PacketMarkerFinder(line string) (int, error) {
    return findContiguousUniqChars(line, 3)
}

func MessageMarkerFinder(line string) (int, error) {
    return findContiguousUniqChars(line, 13)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

    markerFinder := MarkerFinderFunc(MessageMarkerFinder)

	for bufScanner := bufio.NewScanner(file); bufScanner.Scan(); {
		line := bufScanner.Text()

        marker, err := markerFinder.FindMarker(line)
        if err == nil {
            fmt.Println(marker)
        }
	}
}
