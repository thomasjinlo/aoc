package main

import (
	"aoc/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Direction string

const (
    UP Direction = "U"
    DOWN Direction = "D"
    RIGHT Direction = "R"
    LEFT Direction = "L"
)

type Position struct {
    x int
    y int
}

func (p Position) Copy() Position {
    return Position{x: p.x, y: p.y}
}

func (p Position) IsAdjacent(p2 Position) bool {
    // horizontally or vertically adjacent
    if p.x == p2.x || p.y == p2.y {
        return math.Abs(float64(p.x - p2.x)) + math.Abs(float64(p.y - p2.y)) <= 1
    }

    // diagonally adjacent
    return math.Abs(float64(p.x - p2.x)) + math.Abs(float64(p.y - p2.y)) <= 2
}

type Rope struct {
    headPositions []Position
    tailPositions []Position
}

func NewRope() *Rope {
    head, tail := Position{}, Position{}
    return &Rope{
        headPositions: []Position{head},
        tailPositions: []Position{tail},
    }
}

func (r *Rope) Move(direction Direction, distance int) {
    for i := 0; i < distance; i++ {
        head := r.headPositions[len(r.headPositions) - 1]
        tail := r.tailPositions[len(r.tailPositions) - 1]

        dx, dy := 0, 0
        switch direction {
        case UP:
            dx--
        case DOWN:
            dx++
        case RIGHT:
            dy++
        case LEFT:
            dy--
        }

        newHead := Position{x: head.x + dx, y: head.y + dy}
        r.headPositions = append(r.headPositions, newHead)

        if !tail.IsAdjacent(newHead) {
            //fmt.Println(head, newHead, tail)
            newTail := head.Copy()
            r.tailPositions = append(r.tailPositions, newTail)
        }
    }
}

func main() {
    inputs := make(chan string)
    go utils.ScanInputs("input.txt", inputs)

    rope := NewRope()

    for line := range inputs {
        inputs := strings.Split(line, " ")
        direction := Direction(inputs[0])
        distance, _ := strconv.Atoi(inputs[1])

        rope.Move(direction, distance)
    }

    visited := map[string]bool{}
    uniqPos := 0
    for _, tail := range rope.tailPositions {
        pos := fmt.Sprintf("%d,%d", tail.x, tail.y)

        if _, ok := visited[pos]; !ok {
            visited[pos] = true
            uniqPos++
        }
    }
    fmt.Println(uniqPos)
}
