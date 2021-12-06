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
	var depth int
	var distance int
	for i := 0; i < len(dataArr); i++ {
		temp := strings.Split(dataArr[i], " ")
		x, _ := strconv.Atoi(temp[1])

		switch temp[0] {
		case "forward":
			distance += x
		case "up":
			depth -= x
		case "down":
			depth += x
		}
	}
	fmt.Println(depth * distance)
}
