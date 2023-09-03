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

type Snapshot string

type Position struct {
    x int
    y int
    next *Position
    prev *Position
    snapshots []Snapshot
}

func (p *Position) Snapshot() {
    p.snapshots = append(p.snapshots, Snapshot(fmt.Sprintf("%d,%d", p.x, p.y)))
}

func (p *Position) Move(direction Direction) {
    switch direction {
    case UP:
        p.x--
    case DOWN:
        p.x++
    case RIGHT:
        p.y++
    case LEFT:
        p.y--
    }
    p.Snapshot()
}

func (p *Position) Follow(prev *Position, direction Direction) {
    if math.Abs(float64(p.x - prev.x)) == 2 {
        p.y = prev.y
    }

    if math.Abs(float64(p.y - prev.y)) == 2 {
        p.x = prev.x
    }

    p.Move(direction)
}

func (p *Position) IsAdjacent(p2 *Position) bool {
    // horizontally or vertically adjacent
    if p.x == p2.x || p.y == p2.y {
        return math.Abs(float64(p.x - p2.x)) + math.Abs(float64(p.y - p2.y)) <= 1
    }

    // diagonally adjacent
    return math.Abs(float64(p.x - p2.x)) + math.Abs(float64(p.y - p2.y)) <= 2
}

type Rope struct {
    positions []*Position
}

func NewRope(numPos int) *Rope {
    positions := make([]*Position, numPos)
    for i := range positions {
        positions[i] = &Position{}
    }
    return &Rope{positions: positions}
}

func (r *Rope) Move(direction Direction, distance int) {
    for head, j := r.positions[0], 0; j < distance; j++ {
        prev := head
        head.Move(direction)

        for _, pos := range r.positions[1:] {
            if pos.IsAdjacent(prev) {
                continue
            }

            pos.Follow(prev, direction)
            prev = pos
        }
    }
}

func main() {
    inputs := make(chan string)
    go utils.ScanInputs("input.txt", inputs)

    knots := 2
    rope := NewRope(knots)

    for line := range inputs {
        inputs := strings.Split(line, " ")
        direction := Direction(inputs[0])
        distance, _ := strconv.Atoi(inputs[1])

        rope.Move(direction, distance)
    }

    uniqCount := 1
    snapshotSet := map[Snapshot]bool{}
    for _, snapshot := range rope.positions[knots-1].snapshots {
        if _, ok := snapshotSet[snapshot]; !ok {
            snapshotSet[snapshot] = true
            uniqCount++
        }
    }

    fmt.Println(uniqCount)
}
