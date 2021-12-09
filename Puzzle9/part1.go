package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("data.txt")
	dataArr := strings.Split(string(data), "\n")
	field := make([][]int, len(dataArr))

	for i := 0; i < len(dataArr); i++ {
		tempArr := make([]int, len(dataArr[i]))
		for j := 0; j < len(dataArr[i]); j++ {
			tempArr[j], _ = strconv.Atoi(string(dataArr[i][j]))
		}
		field[i] = tempArr
	}

	var riskLevel int

	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if i > 0 {
				if field[i][j] >= field[i-1][j] {
					continue
				}
			}
			if i < len(field)-1 {
				if field[i][j] >= field[i+1][j] {
					continue
				}
			}
			if j > 0 {
				if field[i][j] >= field[i][j-1] {
					continue
				}
			}
			if j < len(field[i])-1 {
				if field[i][j] >= field[i][j+1] {
					continue
				}
			}
			riskLevel += field[i][j] + 1
		}
	}

	fmt.Println(riskLevel)
}
