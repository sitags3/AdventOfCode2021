package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("data.txt")
	dataArr := strings.Split(string(data), "\n")
	var count int

	for i := 0; i < len(dataArr); i++ {
		entry := strings.Split(dataArr[i], "| ")
		output := strings.Split(entry[1], " ")
		for j := 0; j < len(output); j++ {
			if (len(output[j]) == 2) || (len(output[j]) == 3) || (len(output[j]) == 4) || (len(output[j]) == 7) {
				count++
			}
		}
	}

	fmt.Println(count)
}
