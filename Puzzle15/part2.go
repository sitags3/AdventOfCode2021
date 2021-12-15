package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FindSmallest(shortestDistance [][]int, visited [][]bool) (int, int) {
	smallest := 2147483647
	var x, y int
	for i := 0; i < len(visited); i++ {
		for j := 0; j < len(visited[i]); j++ {
			if visited[i][j] == true {
				continue
			}
			if shortestDistance[i][j] < smallest {
				smallest = shortestDistance[i][j]
				x, y = j, i
			}
		}
	}

	return x, y
}

func main() {
	data, _ := os.ReadFile("data.txt")
	dataArrString := strings.Split(string(data), "\n")
	grid := make([][]int, len(dataArrString)*5)
	for z := 0; z < 5; z++ {
		for i := 0; i < len(dataArrString); i++ {
			row := make([]int, len(dataArrString[i])*5)
			for g := 0; g < 5; g++ {
				for j := 0; j < len(dataArrString[i]); j++ {
					row[j+g*len(dataArrString[i])], _ = strconv.Atoi(string(dataArrString[i][j]))
					row[j+g*len(dataArrString[i])] += g + z
					if row[j+g*len(dataArrString[i])] > 9 {
						row[j+g*len(dataArrString[i])] -= 9
					}
				}
			}

			grid[i+z*len(dataArrString)] = row
		}
	}

	visited := make([][]bool, len(grid))
	for i := 0; i < len(visited); i++ {
		row := make([]bool, len(grid[i]))
		visited[i] = row
	}

	shortestDistane := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		row := make([]int, len(grid[i]))
		for j := 0; j < len(row); j++ {
			row[j] = 2147483647
		}
		shortestDistane[i] = row
	}

	shortestDistane[0][0] = 0
	var x, y int
	for visited[len(visited[0])-1][len(visited)-1] == false {
		x, y = FindSmallest(shortestDistane, visited)

		if x < len(grid[0])-1 {
			if visited[y][x+1] == false {
				if shortestDistane[y][x+1] > (shortestDistane[y][x] + grid[y][x+1]) {
					shortestDistane[y][x+1] = shortestDistane[y][x] + grid[y][x+1]
				}
			}
		}
		if x > 0 {
			if visited[y][x-1] == false {
				if shortestDistane[y][x-1] > (shortestDistane[y][x] + grid[y][x-1]) {
					shortestDistane[y][x-1] = shortestDistane[y][x] + grid[y][x-1]
				}
			}
		}
		if y < len(grid)-1 {
			if visited[y+1][x] == false {
				if shortestDistane[y+1][x] > (shortestDistane[y][x] + grid[y+1][x]) {
					shortestDistane[y+1][x] = shortestDistane[y][x] + grid[y+1][x]
				}
			}
		}
		if y > 0 {
			if visited[y-1][x] == false {
				if shortestDistane[y-1][x] > (shortestDistane[y][x] + grid[y-1][x]) {
					shortestDistane[y-1][x] = shortestDistane[y][x] + grid[y-1][x]
				}
			}
		}
		visited[y][x] = true
	}

	fmt.Println(shortestDistane[y][x])
}
