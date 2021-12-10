package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	data, _ := os.ReadFile("data.txt")
	dataArr := strings.Split(string(data), "\n")
	score := make([]int, 0)

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
						corrupted = true
					}
				}
			case ']':
				if len(stack) > 0 {
					if stack[len(stack)-1] == 2 {
						stack = stack[:len(stack)-1]
					} else {
						corrupted = true
					}
				}
			case '}':
				if len(stack) > 0 {
					if stack[len(stack)-1] == 3 {
						stack = stack[:len(stack)-1]
					} else {
						corrupted = true
					}
				}
			case '>':
				if len(stack) > 0 {
					if stack[len(stack)-1] == 4 {
						stack = stack[:len(stack)-1]
					} else {
						corrupted = true
					}
				}
			}
		}

		if corrupted == false {
			score = append(score, 0)
			for j := len(stack) - 1; j >= 0; j-- {
				score[len(score)-1] = (score[len(score)-1] * 5) + stack[j]
			}
		}

	}

	sort.Ints(score)
	fmt.Println(score[(len(score)-1)/2])
}
