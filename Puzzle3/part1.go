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
	resultArr := make([]int, len(dataArr[0]))

	for i := 0; i < len(dataArr); i++ {
		for j := 0; j < len(resultArr); j++ {
			if dataArr[i][j] == '1' {
				resultArr[j]++
			} else {
				resultArr[j]--
			}
		}
	}

	var gamma, epsilon string
	for i := 0; i < len(resultArr); i++ {
		if resultArr[i] > 0 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)

	fmt.Println(gammaInt * epsilonInt)
}
