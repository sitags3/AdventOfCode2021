package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("data.txt")
	dataArrString := strings.Split(string(data), ",")
	dataArr := make([]int, 9)
	for i := 0; i < len(dataArrString); i++ {
		temp, _ := strconv.Atoi(dataArrString[i])
		dataArr[temp]++
	}

	for i := 0; i < 256; i++ {
		dataArr[8], dataArr[7], dataArr[6], dataArr[5], dataArr[4], dataArr[3], dataArr[2], dataArr[1], dataArr[0] = dataArr[0], dataArr[8], dataArr[7]+dataArr[0], dataArr[6], dataArr[5], dataArr[4], dataArr[3], dataArr[2], dataArr[1]
	}

	var result int
	for i := 0; i < len(dataArr); i++ {
		result += dataArr[i]
	}

	fmt.Println(result)
}
