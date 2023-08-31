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

func numTreesVisibleTop(i, j int, grid Grid) int {
    height := grid[i][j]
    treesVisible := 0

    for i := i - 1; i >= 0; i-- {
        treesVisible++

        if grid[i][j] >= height {
            return treesVisible
        }
    }

    return treesVisible
}

func numTreesVisibleBot(i, j int, grid Grid) int {
    height := grid[i][j]
    treesVisible := 0

    for i := i + 1; i < len(grid); i++ {
        treesVisible++

        if grid[i][j] >= height {
            return treesVisible
        }
    }

    return treesVisible
}

func numTreesVisibleLeft(i, j int, grid Grid) int {
    height := grid[i][j]
    treesVisible := 0

    for j := j - 1; j >= 0; j-- {
        treesVisible++

        if grid[i][j] >= height {
            return treesVisible
        }
    }

    return treesVisible
}

func numTreesVisibleRight(i, j int, grid Grid) int {
    height := grid[i][j]
    treesVisible := 0

    for j := j + 1; j < len(grid[0]); j++ {
        treesVisible++

        if grid[i][j] >= height {
            return treesVisible
        }
    }

    return treesVisible
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

    maxTreesVisible := 0

    for i := 0; i < N; i++ {
        for j := 0; j < M; j++ {
            treesVisible := numTreesVisibleBot(i, j, grid) * numTreesVisibleTop(i, j, grid) * numTreesVisibleLeft(i, j, grid) * numTreesVisibleRight(i, j, grid)
            if treesVisible > maxTreesVisible {
                maxTreesVisible = treesVisible
            }
        }
    }

    fmt.Println(maxTreesVisible)
}
