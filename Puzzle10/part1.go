package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("data.txt")
	dataArr := strings.Split(string(data), "\n")
	var score int

	for i := 0; i < len(dataArr); i++ {
		stack := make([]int, 0)
		corrupted := false
		for j := 0; j < len(dataArr[i]); j++ {
			if corrupted == true {
				break
			}
			switch dataArr[i][j] {
			case '(':
				stack = append(stack, 1)
			case '[':
				stack = append(stack, 2)
			case '{':
				stack = append(stack, 3)
			case '<':
				stack = append(stack, 4)
			case ')':
				if len(stack) > 0 {
					if stack[len(stack)-1] == 1 {
						stack = stack[:len(stack)-1]
					} else {
						score += 3
						corrupted = true
						break
					}
				}
			case ']':
				if len(stack) > 0 {
					if stack[len(stack)-1] == 2 {
						stack = stack[:len(stack)-1]
					} else {
						score += 57
						corrupted = true
						break
					}
				}
			case '}':
				if len(stack) > 0 {
					if stack[len(stack)-1] == 3 {
						stack = stack[:len(stack)-1]
					} else {
						score += 1197
						corrupted = true
						break
					}
				}
			case '>':
				if len(stack) > 0 {
					if stack[len(stack)-1] == 4 {
						stack = stack[:len(stack)-1]
					} else {
						score += 25137
						corrupted = true
						break
					}
				}
			}
		}
	}

	fmt.Println(score)
}
