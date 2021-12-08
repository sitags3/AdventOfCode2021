package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func Decode(output string, decodedPatern string) bool {
	if len(output) != len(decodedPatern) {
		return false
	}
	for i := 0; i < len(output); i++ {
		if strings.Contains(output, string(decodedPatern[i])) == false {
			return false
		}
	}
	return true
}

func main() {
	data, _ := os.ReadFile("data.txt")
	dataArr := strings.Split(string(data), "\n")
	var result int

	for i := 0; i < len(dataArr); i++ {
		entry := strings.Split(dataArr[i], " | ")
		paterns := strings.Split(entry[0], " ")
		output := strings.Split(entry[1], " ")
		decodedPatern := make([]string, 10)
		//find patern for 1
		for j := 0; j < len(paterns); j++ {
			if len(paterns[j]) == 2 {
				decodedPatern[1] = paterns[j]
				paterns = remove(paterns, j)
				break
			}
		}
		//find patern for 7
		for j := 0; j < len(paterns); j++ {
			if len(paterns[j]) == 3 {
				decodedPatern[7] = paterns[j]
				paterns = remove(paterns, j)
				break
			}
		}
		//find patern for 4
		for j := 0; j < len(paterns); j++ {
			if len(paterns[j]) == 4 {
				decodedPatern[4] = paterns[j]
				paterns = remove(paterns, j)
				break
			}
		}
		//find patern for 8
		for j := 0; j < len(paterns); j++ {
			if len(paterns[j]) == 7 {
				decodedPatern[8] = paterns[j]
				paterns = remove(paterns, j)
				break
			}
		}
		//find patern for 3
		for j := 0; j < len(paterns); j++ {
			if len(paterns[j]) == 5 {
				for g := 0; g < len(decodedPatern[7]); g++ {
					if strings.Contains(paterns[j], string(decodedPatern[7][g])) == false {
						break
					}
					if g == len(decodedPatern[7])-1 {
						decodedPatern[3] = paterns[j]
						paterns = remove(paterns, j)
					}
				}
				if decodedPatern[3] != "" {
					break
				}
			}
		}
		//find patern for 9
		for j := 0; j < len(paterns); j++ {
			if len(paterns[j]) == 6 {
				for g := 0; g < len(decodedPatern[3]); g++ {
					if strings.Contains(paterns[j], string(decodedPatern[3][g])) == false {
						break
					}
					if g == len(decodedPatern[3])-1 {
						decodedPatern[9] = paterns[j]
						paterns = remove(paterns, j)
					}
				}
				if decodedPatern[9] != "" {
					break
				}
			}
		}
		//find patern for 0
		for j := 0; j < len(paterns); j++ {
			if len(paterns[j]) == 6 {
				for g := 0; g < len(decodedPatern[1]); g++ {
					if strings.Contains(paterns[j], string(decodedPatern[1][g])) == false {
						break
					}
					if g == len(decodedPatern[1])-1 {
						decodedPatern[0] = paterns[j]
						paterns = remove(paterns, j)
					}
				}
				if decodedPatern[0] != "" {
					break
				}
			}
		}
		//find patern for 6
		for j := 0; j < len(paterns); j++ {
			if len(paterns[j]) == 6 {
				decodedPatern[6] = paterns[j]
				paterns = remove(paterns, j)
				break
			}
		}
		//find patern for 5 and 2
		contains := false
		if strings.Contains(decodedPatern[6], string(decodedPatern[1][0])) {
			contains = true
		}

		if strings.Contains(paterns[0], string(decodedPatern[1][0])) == true {
			if contains == true {
				decodedPatern[5] = paterns[0]
				decodedPatern[2] = paterns[1]
			} else {
				decodedPatern[5] = paterns[1]
				decodedPatern[2] = paterns[0]
			}
		} else {
			if contains == false {
				decodedPatern[5] = paterns[0]
				decodedPatern[2] = paterns[1]
			} else {
				decodedPatern[5] = paterns[1]
				decodedPatern[2] = paterns[0]
			}
		}
		//decode output
		var decodedAnswer string
		for j := 0; j < len(output); j++ {
			for g := 0; g < len(decodedPatern); g++ {
				if Decode(output[j], decodedPatern[g]) == true {
					decodedAnswer += strconv.Itoa(g)
					break
				}
			}
		}
		temp, _ := strconv.Atoi(decodedAnswer)
		result += temp
	}

	fmt.Println(result)
}
