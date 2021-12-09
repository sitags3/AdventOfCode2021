package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Combine(field [][]int, basins []int, value1 int, value2 int) ([][]int, []int) {
	basins[(value1*-1)-1] += basins[(value2*-1)-1]
	basins[(value2*-1)-1] = value1

	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if field[i][j] == value2 {
				field[i][j] = value1
			}
		}
	}

	return field, basins
}

func main() {
	data, _ := os.ReadFile("data.txt")
	dataArr := strings.Split(string(data), "\n")
	field := make([][]int, len(dataArr))

	for i := 0; i < len(dataArr); i++ {
		tempArr := make([]int, len(dataArr[i]))
		for j := 0; j < len(dataArr[i]); j++ {
			tempArr[j], _ = strconv.Atoi(string(dataArr[i][j]))
		}
		field[i] = tempArr
	}

	basins := make([]int, 0)
	basinCount := 0

	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if field[i][j] != 9 {
				if j > 0 {
					if field[i][j-1] != 9 {
						basins[(field[i][j-1]*-1)-1]++
						field[i][j] = field[i][j-1]
					} else {
						basinCount++
						basins = append(basins, 0)
						basins[basinCount-1]++
						field[i][j] = basinCount * -1
					}
				} else {
					basinCount++
					basins = append(basins, 0)
					basins[basinCount-1]++
					field[i][j] = basinCount * -1
				}
			}
		}
	}

	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if field[i][j] != 9 {
				if i > 0 {
					if (field[i][j] != field[i-1][j]) && (field[i-1][j] != 9) {
						field, basins = Combine(field, basins, field[i][j], field[i-1][j])
					}
				}
				if i < len(field)-1 {
					if (field[i][j] != field[i+1][j]) && (field[i+1][j] != 9) {
						field, basins = Combine(field, basins, field[i][j], field[i+1][j])
					}
				}
				if j > 0 {
					if (field[i][j] != field[i][j-1]) && (field[i][j-1] != 9) {
						field, basins = Combine(field, basins, field[i][j], field[i][j-1])
					}
				}
				if j < len(field[i])-1 {
					if (field[i][j] != field[i][j+1]) && (field[i][j+1] != 9) {
						field, basins = Combine(field, basins, field[i][j], field[i][j+1])
					}
				}
			}
		}
	}

	sort.Ints(basins)
	fmt.Println(basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3])
}
