package utils

import "strconv"

func MustInt(digit string) int {
    number, err := strconv.Atoi(digit)
    if err != nil {
        panic("could not convert to int")
    }

    return number
}
