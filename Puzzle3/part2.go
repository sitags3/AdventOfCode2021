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
	oxygenArr := make([]int, len(dataArr[0]))
	scrubberArr := make([]int, len(dataArr[0]))
	var oxygenRating, scrubberRating string

	for j := 0; j < len(oxygenArr); j++ {
		var count int
		for i := 0; i < len(dataArr); i++ {
			if (dataArr[i][j] == '1') && (strings.HasPrefix(string(dataArr[i]), scrubberRating) == true) {
				scrubberArr[j]--
				count++
			} else if strings.HasPrefix(string(dataArr[i]), scrubberRating) == true {
				scrubberArr[j]++
				count++
			}
		}
		for i := 0; i < len(dataArr); i++ {
			if (dataArr[i][j] == '1') && (strings.HasPrefix(string(dataArr[i]), oxygenRating) == true) {
				oxygenArr[j]++
			} else if strings.HasPrefix(string(dataArr[i]), oxygenRating) == true {
				oxygenArr[j]--
			}
		}
		if scrubberArr[j] > 0 {
			if count == 1 {
				scrubberRating += "0"
			} else {
				scrubberRating += "1"
			}
		} else {
			if count == 1 {
				scrubberRating += "1"
			} else {
				scrubberRating += "0"
			}
		}
		if oxygenArr[j] >= 0 {
			oxygenRating += "1"
		} else {
			oxygenRating += "0"
		}
	}

	oxygenInt, _ := strconv.ParseInt(oxygenRating, 2, 64)
	scrubberInt, _ := strconv.ParseInt(scrubberRating, 2, 64)

	fmt.Println(oxygenInt * scrubberInt)
}
