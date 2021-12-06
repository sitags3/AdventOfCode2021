package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println("Incorrect input")
		os.Exit(0)
	}
	dataArr := strings.Split(string(f), "\n")
	var result int
	dataIntArr := make([]int, len(dataArr))

	for i := 0; i < len(dataArr); i++ {
		temp, _ := strconv.Atoi(dataArr[i])
		dataIntArr[i] = temp
	}

	for i := 0; i < len(dataArr)-3; i++ {
		if dataIntArr[i+1]+dataIntArr[i+2]+dataIntArr[i+3] > dataIntArr[i]+dataIntArr[i+1]+dataIntArr[i+2] {
			result++
		}
	}
	fmt.Println(result)
}
