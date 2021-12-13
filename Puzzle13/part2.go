package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FoldX(instructions [][]int, foldInstruction int) [][]int {
	for j := 0; j < len(instructions); j++ {
		for i := 1; i < len(instructions[0])-foldInstruction; i++ {
			instructions[j][foldInstruction-i] += instructions[j][foldInstruction+i]
		}
	}
	for i := 0; i < len(instructions); i++ {
		instructions[i] = instructions[i][:foldInstruction]
	}

	return instructions
}

func FoldY(instructions [][]int, foldInstruction int) [][]int {
	for i := 1; i < len(instructions)-foldInstruction; i++ {
		for j := 0; j < len(instructions[i]); j++ {
			instructions[foldInstruction-i][j] += instructions[foldInstruction+i][j]
		}
	}
	instructions = instructions[:foldInstruction]

	return instructions
}

func main() {
	data, _ := os.ReadFile("data.txt")
	dataArrString := strings.Split(string(data), "\n")
	foldInstructions := make([]string, 0)
	dataArr := make([][]int, 0)
	var largestX, largestY int

	for i := 0; i < len(dataArrString); i++ {
		if dataArrString[i] == "" {
			for j := i + 1; j < len(dataArrString); j++ {
				temp := strings.Split(dataArrString[j], " ")
				foldInstructions = append(foldInstructions, temp[2])
			}
			break
		}
		temp := strings.Split(dataArrString[i], ",")
		tempArr := make([]int, 2)
		tempArr[0], _ = strconv.Atoi(temp[0])
		tempArr[1], _ = strconv.Atoi(temp[1])
		dataArr = append(dataArr, tempArr)
		if tempArr[0] > largestX {
			largestX = tempArr[0]
		}
		if tempArr[1] > largestY {
			largestY = tempArr[1]
		}
	}

	instructions := make([][]int, largestY+1)
	for i := 0; i <= largestY; i++ {
		instructions[i] = make([]int, largestX+1)
	}

	for i := 0; i < len(dataArr); i++ {
		instructions[dataArr[i][1]][dataArr[i][0]]++
	}

	for i := 0; i < len(foldInstructions); i++ {
		if int(foldInstructions[i][0]) == 120 {
			temp, _ := strconv.Atoi(foldInstructions[i][2:])
			instructions = FoldX(instructions, temp)
		} else {
			temp, _ := strconv.Atoi(foldInstructions[i][2:])
			instructions = FoldY(instructions, temp)
		}
	}

	for i := 0; i < len(instructions); i++ {
		for j := 0; j < len(instructions[i]); j++ {
			if instructions[i][j] > 0 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
