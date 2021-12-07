package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("data.txt")
	dataArrString := strings.Split(string(data), ",")
	dataArr := make([]int, len(dataArrString))
	for i := 0; i < len(dataArr); i++ {
		dataArr[i], _ = strconv.Atoi(dataArrString[i])
	}

	sort.Ints(dataArr)
	median := dataArr[int(len(dataArr)/2)]

	var fuel int

	for i := 0; i < len(dataArr); i++ {
		if dataArr[i] > median {
			fuel += dataArr[i] - median
		} else {
			fuel += median - dataArr[i]
		}
	}

	fmt.Println(fuel)
}
