package main

import (
	"aoc/utils"
	"fmt"
	"regexp"
	"strings"
	// "strconv"
	// "strings"
)

type Monkey struct {
    id int
    inspectCount int
    inspectOperation string
    inspectNumber int
    divisibleNumber int
    monkeyIdFalse int
    monkeyIdTrue int
    startingItems []int
    useItemReflection bool
}

func (m *Monkey) InspectItems (monkeys map[int]*Monkey) {
    for _, item := range m.startingItems {
        m.inspectCount += 1
        newItem := m.inspect(item)

        if passTest := m.test(newItem); passTest {
            monkeys[m.monkeyIdTrue].AddItem(newItem)
        } else {
            monkeys[m.monkeyIdFalse].AddItem(newItem)
        }
    }

    m.startingItems = []int{}
}

func (m *Monkey) inspect(item int) int {
    inspectNumber := m.inspectNumber
    if m.useItemReflection {
        inspectNumber = item
    }


    newItem := 0
    switch string(m.inspectOperation) {
    case "+":
        newItem = item + inspectNumber
    case "*":
        newItem = item * inspectNumber
    }

    return newItem / 3
}

func (m *Monkey) test(item int) bool {
    return item % m.divisibleNumber == 0
}

func (m *Monkey) AddItem(item int) {
    m.startingItems = append(m.startingItems, item)
}

func (m *Monkey) AddInspectOperation(operation string) {
    m.inspectOperation = operation
}

func (m *Monkey) AddInspectNumber(number int) {
    m.inspectNumber = number
}

func (m *Monkey) UseItemReflection() {
    m.useItemReflection = true
}

func (m *Monkey) AddDivisibleNumber(number int) {
    m.divisibleNumber = number
}

func (m *Monkey) AddMonkeyIdTrue(id int) {
    m.monkeyIdTrue = id
}

func (m *Monkey) AddMonkeyIdFalse(id int) {
    m.monkeyIdFalse = id
}

func partOne() {
    inputs := make(chan string)
    go utils.ScanInputs("input.txt", inputs)

    monkeys := map[int]*Monkey{}
    monkeyIdRegex := regexp.MustCompile(`Monkey (\d+):`)
    startingItemsRegex := regexp.MustCompile(`Starting items: (.+)`)
    operationRegex := regexp.MustCompile(`Operation: new = old (.+)`)
    testDivisionRegex := regexp.MustCompile(`Test: divisible by (\d+)`)
    monkeyIdTrueRegex := regexp.MustCompile(`If true: throw to monkey (\d+)`)
    monkeyIdFalseRegex := regexp.MustCompile(`If false: throw to monkey (\d+)`)

    var monkey *Monkey
    for line := range inputs {
        if len(line) == 0 {
            continue
        }

        if match := monkeyIdRegex.FindStringSubmatch(line); len(match) > 1 {
            monkey = &Monkey{id: utils.MustInt(match[1])}
            monkeys[monkey.id] = monkey
        }

        if match := startingItemsRegex.FindStringSubmatch(line); len(match) > 1 {
            startingItems := strings.Split(match[1], ", ")
            for _, item := range startingItems {
                monkey.AddItem(utils.MustInt(item))
            }
        }

        if match := operationRegex.FindStringSubmatch(line); len(match) > 1 {
            operations := strings.Split(match[1], " ")

            monkey.AddInspectOperation(operations[0])

            if operations[1] == "old" {
                monkey.UseItemReflection()
            } else {
                monkey.AddInspectNumber(utils.MustInt(operations[1]))
            }
        }

        if match := testDivisionRegex.FindStringSubmatch(line); len(match) > 1 {
            monkey.AddDivisibleNumber(utils.MustInt(match[1]))
        }

        if match := monkeyIdTrueRegex.FindStringSubmatch(line); len(match) > 1 {
            monkey.AddMonkeyIdTrue(utils.MustInt(match[1]))
        }

        if match := monkeyIdFalseRegex.FindStringSubmatch(line); len(match) > 1 {
            monkey.AddMonkeyIdFalse(utils.MustInt(match[1]))
        }
    }

    for round := 0; round < 20; round++ {
        for i := 0; i < len(monkeys); i++ {
            monkey := monkeys[i]
            monkey.InspectItems(monkeys)
        }
    }

    maxInspect, nextMaxInspect := 0, 0
    for i := 0; i < len(monkeys); i++ {
        monkey := monkeys[i]
        if monkey.inspectCount > maxInspect {
            nextMaxInspect = maxInspect
            maxInspect = monkey.inspectCount
        }
    }

    fmt.Println("multiplying", maxInspect, nextMaxInspect)
    fmt.Println(maxInspect * nextMaxInspect)
}

func main() {
    partOne()
}
