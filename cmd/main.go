package main

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

func getInput() string {
	input := `3 4
4 3
2 5
1 3
3 9
3 3
`

	return input
}

func main() {
	input := getInput()
	var arr1 []int
	var arr2 []int
	for line := range strings.SplitSeq(input, "\n") {
		numbers := strings.Split(line, " ")
		if len(numbers) == 2 {
			a, err := strconv.Atoi(numbers[0])
			if err != nil {
				log.Fatal("error: invalid input")
			}
			b, err := strconv.Atoi(numbers[1])
			if err != nil {
				log.Fatal("error: invalid input")
			}
			arr1 = append(arr1, a)
			arr2 = append(arr2, b)
		}
	}
	sort.Ints(arr2)
	sort.Ints(arr1)
}
