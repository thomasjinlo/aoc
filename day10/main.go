package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

type CpuInstruction string

const (
    ADDX CpuInstruction = "addx"
    NOOP CpuInstruction = "noop"
)

type CpuSignal struct {
    instruction CpuInstruction
    xValue int
}

func ExtractSignal(signal string) CpuSignal {
    instructions := strings.Split(signal, " ")
    cpuInstruction := CpuInstruction(instructions[0])
    xValue := 0

    if len(instructions) == 2 {
        xValue, _ = strconv.Atoi(instructions[1])
    }

    return CpuSignal{instruction: cpuInstruction, xValue: xValue}
}

type VideoDevice struct {
    cycleClock int
    registerX int
    image []string
    registerXCycles []int
}

func (v *VideoDevice) processCycle() {
    v.registerXCycles = append(v.registerXCycles, v.registerX)

    if len(v.image[len(v.image)-1]) == 40 {
        v.image = append(v.image, "")
    }

    pixelIdx := v.cycleClock - ((len(v.image) - 1) * 40)
    if pixelIdx >= v.registerX - 1 && pixelIdx <= v.registerX + 1 {
        v.image[len(v.image)-1] += "#"
    } else {
        v.image[len(v.image)-1] += "."
    }

    v.cycleClock++
}

func (v *VideoDevice) DrawImage() {
    for _, imageRow := range v.image {
        fmt.Println(imageRow)
    }
}

func (v *VideoDevice) ProcessCpuSignal(signal CpuSignal) {
    switch signal.instruction {
    case NOOP:
        v.processCycle()
    case ADDX:
        v.processCycle()
        v.processCycle()
        v.registerX += signal.xValue
    }
}

func (v *VideoDevice) SignalStrength() int {
    return ((v.getRegisterX(20) * 20) +
            (v.getRegisterX(60) * 60) +
            (v.getRegisterX(100) * 100) +
            (v.getRegisterX(140) * 140) +
            (v.getRegisterX(180) * 180) +
            (v.getRegisterX(220) * 220))
}

func (v *VideoDevice) getRegisterX(cycle int) int {
    return v.registerXCycles[cycle]
}

func NewVideoDevice() *VideoDevice {
    return &VideoDevice{
        registerX: 1,
        registerXCycles: []int{1},
        image: []string{""},
    }
}

func partOne() {
    inputs := make(chan string)
    go utils.ScanInputs("input.txt", inputs)

    videoDevice := NewVideoDevice()

    for line := range inputs {
        cpuSignal := ExtractSignal(line)
        videoDevice.ProcessCpuSignal(cpuSignal)
    }

    fmt.Println(videoDevice.SignalStrength())
}

func partTwo() {
    inputs := make(chan string)
    go utils.ScanInputs("input.txt", inputs)

    videoDevice := NewVideoDevice()

    for line := range inputs {
        cpuSignal := ExtractSignal(line)
        videoDevice.ProcessCpuSignal(cpuSignal)
    }

    videoDevice.DrawImage()
}

func main() {
    //partOne()
    partTwo()
}
