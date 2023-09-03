package utils

import (
    "bufio"
    "log"
    "os"
)

func ScanInputs(filepath string, ch chan<-string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	for bufScanner := bufio.NewScanner(file); bufScanner.Scan(); {
        ch <- bufScanner.Text()
    }

    close(ch)
}
