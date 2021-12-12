package main

import (
	"fmt"
	"os"
	"strings"
)

func ConnectionPos(rooms []string, roomNames []string) (int, int) {
	var value1 int
	for j := 0; j < len(rooms); j++ {
		for g := 0; g < len(roomNames); g++ {
			if rooms[j] == roomNames[g] {
				if j == 0 {
					value1 = g
				} else {
					return value1, g
				}
			}
		}
	}
	return 0, 0
}

func Solve(pathCount *int, path []int, connections [][]int, roomNames []string, end int) {
	if path[len(path)-1] != end {
		for i := 0; i < len(connections[path[len(path)-1]]); i++ {
			if roomNames[connections[path[len(path)-1]][i]] == "start" {
				continue
			}
			if int(roomNames[connections[path[len(path)-1]][i]][0]) > 96 {
				allowedToRepeat := true
				for j := 0; j < len(path)-1; j++ {
					if int(roomNames[path[j]][0]) < 97 {
						continue
					}
					for g := j + 1; g < len(path); g++ {
						if roomNames[path[j]] == roomNames[path[g]] {
							allowedToRepeat = false
							break
						}
					}
				}

				repeat := false
				for j := 0; j < len(path); j++ {
					if path[j] == connections[path[len(path)-1]][i] {
						repeat = true
					}
				}
				if repeat == true && allowedToRepeat == false {
					continue
				}
			}
			path = append(path, connections[path[len(path)-1]][i])
			Solve(pathCount, path, connections, roomNames, end)
			path = path[:len(path)-1]
		}
	} else {
		*pathCount++
		return
	}
	return
}

func main() {
	data, _ := os.ReadFile("data.txt")
	dataArrString := strings.Split(string(data), "\n")
	roomNames := make([]string, 0)
	var start, end int

	for i := 0; i < len(dataArrString); i++ {
		rooms := strings.Split(dataArrString[i], "-")
		for j := 0; j < len(rooms); j++ {
			roomExits := false
			for g := 0; g < len(roomNames); g++ {
				if rooms[j] == roomNames[g] {
					roomExits = true
					break
				}
			}
			if roomExits == false {
				if rooms[j] == "start" {
					start = len(roomNames)
				}
				if rooms[j] == "end" {
					end = len(roomNames)
				}
				roomNames = append(roomNames, rooms[j])
			}
		}
	}

	connections := make([][]int, len(roomNames))
	for i := 0; i < len(dataArrString); i++ {
		rooms := strings.Split(dataArrString[i], "-")
		value1, value2 := ConnectionPos(rooms, roomNames)
		connections[value1] = append(connections[value1], value2)
		connections[value2] = append(connections[value2], value1)
	}

	var pathCount int
	var path []int
	path = append(path, start)
	Solve(&pathCount, path, connections, roomNames, end)

	fmt.Println(pathCount)
}
