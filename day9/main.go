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
    next *Position
    snapshots map[string]bool
}

func (p *Position) Snapshot() {
    p.snapshots[fmt.Sprintf("%d,%d", p.x, p.y)] = true
}

func absDiff(x, y int) int {
    return int(math.Abs(float64(x - y)))
}

func (p *Position) IsAdjacent(p2 *Position) bool {
    if p.x == p2.x || p.y == p2.y {
        return absDiff(p.x, p2.x) + absDiff(p.y, p2.y) <= 1
    }

    return absDiff(p.x, p2.x) + absDiff(p.y, p2.y) <= 2
}

func (p *Position) Move(x, y int) {
    p.x, p.y = x, y
    p.Snapshot()
    
    p2 := p.next
    if p2 == nil {
        return
    }

    if p.IsAdjacent(p2) {
        return
    }

    dx, dy := 0, 0
    if p.x == p2.x || p.y == p2.y {
        dx, dy = (p.x - p2.x)/2, (p.y - p2.y)/2
    } else {
        if absDiff(p.x, p2.x) == 2 {
            dx = (p.x - p2.x)/2
        } else {
            dx = p.x - p2.x
        }

        if absDiff(p.y, p2.y) == 2 {
            dy = (p.y - p2.y)/2
        } else {
            dy = p.y - p2.y
        }
    }

    p2.Move(p2.x + dx, p2.y + dy)
}

type Rope struct {
    head *Position
    tail *Position
}

func (r *Rope) Move(direction Direction, distance int) {
    for j := 0; j < distance; j++ {
        dx, dy := r.head.x, r.head.y
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
        r.head.Move(dx, dy)
    }
}

func NewRope(numPos int) *Rope {
    head := &Position{snapshots: make(map[string]bool)}
    prev := head
    for i := 1; i < numPos; i++ {
        prev.next = &Position{snapshots: make(map[string]bool)}
        prev = prev.next
    }
    return &Rope{head: head, tail: prev}
}

func main() {
    inputs := make(chan string)
    go utils.ScanInputs("input.txt", inputs)

    knots := 10
    rope := NewRope(knots)

    for line := range inputs {
        inputs := strings.Split(line, " ")
        direction := Direction(inputs[0])
        distance, _ := strconv.Atoi(inputs[1])

        rope.Move(direction, distance)
    }

    fmt.Println(len(rope.tail.snapshots) + 1)
}
