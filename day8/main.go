package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Grid [][]int

type Set map[int]map[int]bool

func isTreeVisibleVert(i, j int, grid Grid) bool {
    isVisibleTop, isVisibleBot := true, true
    height := grid[i][j]

    for i := i - 1; i >= 0; i-- {
        if grid[i][j] >= height {
            isVisibleTop = false
        }
    }

    for i := i + 1; i < len(grid); i++ {
        if grid[i][j] >= height {
            isVisibleBot = false
        }
    }

    return isVisibleTop || isVisibleBot
}

func isTreeVisibleHoriz(i, j int, grid Grid) bool {
    isVisibleLeft, isVisibleRight := true, true
    height := grid[i][j]

    for j := j - 1; j >= 0; j-- {
        if grid[i][j] >= height {
            isVisibleLeft = false
        }
    }

    for j := j + 1; j < len(grid[0]); j++ {
        if grid[i][j] >= height {
            isVisibleRight = false
        }
    }

    return isVisibleLeft || isVisibleRight
}

func isTreeVisible(i, j int, grid Grid) bool {
    return isTreeVisibleVert(i, j, grid) || isTreeVisibleHoriz(i, j, grid)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

    grid := Grid{}

	for bufScanner := bufio.NewScanner(file); bufScanner.Scan(); {
        row := []int{}

        for _, digit := range bufScanner.Text() {
            height, _ := strconv.Atoi(string(digit))
            row = append(row, height)
        }

        grid = append(grid, row)
    }

    N, M := len(grid), len(grid[0])
    treesVisible := N * 2 + M * 2 - 4

    for i := 1; i < N - 1; i++ {
        for j := 1; j < M - 1; j++ {
            if isTreeVisible(i, j, grid) {
                treesVisible ++
            }
        }
    }

    fmt.Println(treesVisible)
}
