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
	lastNumber, _ := strconv.Atoi(dataArr[0])
	for i := 1; i < len(dataArr); i++ {
		temp, _ := strconv.Atoi(dataArr[i])
		if temp > lastNumber {
			result++
		}
		lastNumber = temp
	}
	fmt.Println(result)
}
