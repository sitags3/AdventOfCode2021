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
	dataArr := make([]int, len(dataArrString))

	for i := 0; i < len(dataArrString); i++ {
		temp, _ := strconv.Atoi(dataArrString[i])
		dataArr[i] = temp
	}

	for i := 0; i < 80; i++ {
		for j := 0; j < len(dataArr); j++ {
			if dataArr[j] > 0 {
				dataArr[j]--
				continue
			} else {
				dataArr[j] = 6
				dataArr = append(dataArr, 9)
			}
		}
	}

	fmt.Println(len(dataArr))
}
