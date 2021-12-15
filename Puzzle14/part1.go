package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	data, _ := os.ReadFile("data.txt")
	dataArrString := strings.Split(string(data), "\n")
	template := dataArrString[0]

	for i := 0; i < 10; i++ {
		var newTemplate string
		for j := 1; j < len(template); j++ {
			newTemplate += string(template[j-1])
			pair := string(template[j-1]) + string(template[j])
			for g := 2; g < len(dataArrString); g++ {
				if pair == dataArrString[g][:2] {
					newTemplate += dataArrString[g][6:]
					break
				}
			}
		}
		newTemplate += string(template[len(template)-1])
		template = newTemplate
	}

	count := make([]int, 26)
	for i := 0; i < len(template); i++ {
		count[int(template[i])-65]++
	}

	sort.Ints(count)
	for i := 0; i < 26; i++ {
		if count[i] != 0 {
			fmt.Println(count[len(count)-1] - count[i])
			break
		}
	}
}
