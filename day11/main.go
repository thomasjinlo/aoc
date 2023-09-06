package main

import (
	"aoc/utils"
	"fmt"
//  "regexp"
//	"strconv"
//	"strings"
)

type Monkey struct {
    id int
    inspectCount int
    startingItems []int
    operation func(int) int
    test func(int) bool
    testTrue int
    testFalse int
}

func (m *Monkey) InspectItems (monkeys map[int]*Monkey) {
    for _, item := range m.startingItems {
        m.inspectCount += 1
        newItem := m.operation(item)
        newItem = newItem / 3
        passTest := m.test(newItem)

        if passTest {
            monkeys[m.testTrue].startingItems = append(monkeys[m.testTrue].startingItems, newItem)
        } else {
            monkeys[m.testFalse].startingItems = append(monkeys[m.testFalse].startingItems, newItem)
        }
    }

    m.startingItems = []int{}
}

func partOne() {
    inputs := make(chan string)
    go utils.ScanInputs("sampleInput.txt", inputs)

    monkey0 := &Monkey{
        id: 0,
        startingItems: []int{79, 98},
        operation: func(item int) int { return item * 19 },
        test: func(item int) bool { return item % 23 == 0 },
        testTrue: 2,
        testFalse: 3,
    }
    monkey1 := &Monkey{
        id: 1,
        startingItems: []int{54, 65, 75, 74},
        operation: func(item int) int { return item + 6 },
        test: func(item int) bool { return item % 19 == 0 },
        testTrue: 2,
        testFalse: 0,
    }
    monkey2 := &Monkey{
        id: 2,
        startingItems: []int{79, 60, 97},
        operation: func(item int) int { return item * item },
        test: func(item int) bool { return item % 13 == 0 },
        testTrue: 1,
        testFalse: 3,
    }
    monkey3 := &Monkey{
        id: 3,
        startingItems: []int{74},
        operation: func(item int) int { return item + 3 },
        test: func(item int) bool { return item % 17 == 0 },
        testTrue: 0,
        testFalse: 1,
    }
    monkeys := map[int]*Monkey{
        0: monkey0,
        1: monkey1,
        2: monkey2,
        3: monkey3,
    }

    for round := 0; round < 20; round++ {
        for i := 0; i < len(monkeys); i++ {
            monkey := monkeys[i]
            monkey.InspectItems(monkeys)
        }
    }

    for i := 0; i < len(monkeys); i++ {
        monkey := monkeys[i]
        fmt.Println(monkey.inspectCount)
    }
}

func main() {
    partOne()
}
