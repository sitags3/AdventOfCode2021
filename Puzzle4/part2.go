package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve(boards [][][]string, draw []string) ([][]string, int) {
	boardCount := len(boards)
	winningBoards := make([]int, boardCount)
	var count int
	//number draw
	for i := 0; i < len(draw); i++ {
		for j := 0; j < len(boards); j++ {
			for g := 0; g < 5; g++ {
				for h := 0; h < 5; h++ {
					if boards[j][g][h] == draw[i] {
						boards[j][g][h] = "*"
						break
					}
				}
			}
			//check if board wins
			for g := 0; g < 5; g++ {
				for h := 0; h < 5; h++ {
					if boards[j][g][h] != "*" {
						break
					}
					if h == len(boards[j][g])-1 {
						//winner
						if winningBoards[j] == 0 {
							count++
							winningBoards[j] = 1
						}
						if boardCount == count {
							winningNumber, _ := strconv.Atoi(draw[i])
							return boards[j], winningNumber
						}
					}
				}
			}

			for g := 0; g < 5; g++ {
				for h := 0; h < 5; h++ {
					if boards[j][h][g] != "*" {
						break
					}
					if h == len(boards[j][g])-1 {
						//winner
						if winningBoards[j] == 0 {
							count++
							winningBoards[j] = 1
						}
						if boardCount == count {
							winningNumber, _ := strconv.Atoi(draw[i])
							return boards[j], winningNumber
						}
					}
				}
			}
		}
	}
	return boards[0], 1
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func RemoveEmptyElements(row []string) []string {
	for j := 0; j < len(row); j++ {
		if j >= len(row) {
			break
		}
		if row[j] == "" {
			row = remove(row, j)
			j--
		}
	}
	return row
}

func main() {
	data, _ := os.ReadFile("data.txt")
	dataArr := strings.Split(string(data), "\n")
	draw := strings.Split(dataArr[0], ",")
	boards := make([][][]string, 0)

	//fill boards
	for i := 2; i < len(dataArr); i += 6 {
		board := make([][]string, 0)
		row := RemoveEmptyElements(strings.Split(dataArr[i], " "))
		board = append(board, row)
		row = RemoveEmptyElements(strings.Split(dataArr[i+1], " "))
		board = append(board, row)
		row = RemoveEmptyElements(strings.Split(dataArr[i+2], " "))
		board = append(board, row)
		row = RemoveEmptyElements(strings.Split(dataArr[i+3], " "))
		board = append(board, row)
		row = RemoveEmptyElements(strings.Split(dataArr[i+4], " "))
		board = append(board, row)
		boards = append(boards, board)
	}

	winningBoard, winningNumber := Solve(boards, draw)
	var sum int

	for i := 0; i < len(winningBoard); i++ {
		for j := 0; j < len(winningBoard[i]); j++ {
			if winningBoard[i][j] != "*" {
				temp, _ := strconv.Atoi(winningBoard[i][j])
				sum += temp
			}
		}
	}

	fmt.Println(sum * winningNumber)
}
