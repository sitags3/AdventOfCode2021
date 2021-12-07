package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func FuelCalculation(average int, dataArr []int) int {
	var fuel int
	for i := 0; i < len(dataArr); i++ {
		moves := int(math.Abs(float64(dataArr[i]) - float64(average)))
		fuel += moves * (moves + 1) / 2
	}
	return fuel
}

func main() {
	data, _ := os.ReadFile("data.txt")
	dataArrString := strings.Split(string(data), ",")
	dataArr := make([]int, len(dataArrString))
	for i := 0; i < len(dataArr); i++ {
		dataArr[i], _ = strconv.Atoi(dataArrString[i])
	}

	var average int
	for i := 0; i < len(dataArr); i++ {
		average += dataArr[i]
	}
	average = int(math.Floor(float64(average) / float64(len(dataArr))))

	fuel := FuelCalculation(average, dataArr)
	temp := FuelCalculation(average+1, dataArr)
	if fuel > temp {
		fuel = temp
	}

	fmt.Println(fuel)
}
