package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Flash(dataArr [][]int, i int, j int) [][]int {
	if (i > 0) && (j > 0) {
		dataArr[i-1][j-1]++
	}
	if i > 0 {
		dataArr[i-1][j]++
	}
	if (i > 0) && (j < len(dataArr[i])-1) {
		dataArr[i-1][j+1]++
	}
	if j < len(dataArr[i])-1 {
		dataArr[i][j+1]++
	}
	if (i < len(dataArr)-1) && (j < len(dataArr[i])-1) {
		dataArr[i+1][j+1]++
	}
	if i < len(dataArr)-1 {
		dataArr[i+1][j]++
	}
	if (i < len(dataArr)-1) && (j > 0) {
		dataArr[i+1][j-1]++
	}
	if j > 0 {
		dataArr[i][j-1]++
	}

	return dataArr
}

func CheckFlashes(dataArr [][]int, flashesCount int) ([][]int, int) {
	endTurn := false
	for endTurn == false {
		endTurn = true
		for i := 0; i < len(dataArr); i++ {
			for j := 0; j < len(dataArr[i]); j++ {
				if (dataArr[i][j] > 9) && (dataArr[i][j] < 100) {
					dataArr = Flash(dataArr, i, j)
					dataArr[i][j] = 101
					endTurn = false
					flashesCount++
				}
			}
		}
	}

	for i := 0; i < len(dataArr); i++ {
		for j := 0; j < len(dataArr[i]); j++ {
			if dataArr[i][j] > 100 {
				dataArr[i][j] = 0
			}
		}
	}

	return dataArr, flashesCount
}

func main() {
	data, _ := os.ReadFile("data.txt")
	dataArrString := strings.Split(string(data), "\n")
	dataArr := make([][]int, len(dataArrString))
	for i := 0; i < len(dataArrString); i++ {
		row := make([]int, len(dataArrString[i]))
		for j := 0; j < len(dataArrString[i]); j++ {
			row[j], _ = strconv.Atoi(string(dataArrString[i][j]))
		}
		dataArr[i] = row
	}

	var flashesCount int
	for i := 0; i < 100; i++ {
		for j := 0; j < len(dataArr); j++ {
			for g := 0; g < len(dataArr[j]); g++ {
				dataArr[j][g]++
			}
		}
		dataArr, flashesCount = CheckFlashes(dataArr, flashesCount)
	}

	fmt.Println(flashesCount)
}
