package main

import (
	"log"
	"math"
	"os"
	"strings"
	"unicode"
)

func clear_nums(nums *[12]rune, x int) {
	for i := x; i < 12; i++ {
		nums[i] = '0'
	}
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		return
	}
	lines := strings.Split(string(content), "\n")
	var result int64 = 0
	for i, line := range lines {
		if line == "" {
			continue
		}
		var nums [12]rune
		for j, c := range line {
			if !unicode.IsDigit(c) {
				log.Panicf("invalid digit in line %v", line)
			}
			for x := 0; x < 12; x++ {
				if c > nums[x] && len(line)-j >= 12-x {
					clear_nums(&nums, x+1)
					nums[x] = c
					break
				}
			}
		}
		lineResult := 0
		for i, v := range nums {
			lineResult += int(v-'0') * int(math.Pow(10, float64(12-(i+1))))
		}
		result += int64(lineResult)
		log.Printf("line %v | number is %v", i, lineResult)
	}
	log.Printf("result: %v", result)
}
