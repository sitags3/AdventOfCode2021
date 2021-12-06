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
	grid := make([][1000]int, 1000)

	for i := 0; i < len(dataArr); i++ {
		tempArr := strings.Split(dataArr[i], ",")
		tempArr2 := strings.Split(tempArr[1], "-")
		x1, _ := strconv.Atoi(tempArr[0])
		y1, _ := strconv.Atoi(tempArr2[0][:len(tempArr2[0])-1])
		x2, _ := strconv.Atoi(tempArr2[1][2:])
		y2, _ := strconv.Atoi(tempArr[2])

		if (x1 == x2) || (y1 == y2) {
			if x1 == x2 {
				if y1 > y2 {
					y1, y2 = y2, y1
				}
				for z := y1; z <= y2; z++ {
					grid[z][x1]++
				}
			} else {
				if x1 > x2 {
					x1, x2 = x2, x1
				}
				for z := x1; z <= x2; z++ {
					grid[y1][z]++
				}
			}
		}
	}

	var count int
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if grid[i][j] > 1 {
				count++
			}
		}
	}

	fmt.Println(count)
}
