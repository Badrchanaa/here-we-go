package main

import (
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		return
	}
	lines := strings.Split(string(content), "\n")
	result := 0
	for i, line := range lines {
		maxNum := '0'
		secondMax := '0'
		for j, c := range line {
			if !unicode.IsDigit(c) {
				log.Panicf("invalid digit in line %v", line)
			}
			if c > maxNum && j != len(line)-1 {
				secondMax = '0'
				maxNum = c
			} else if c > secondMax {
				secondMax = c
			}
		}
		log.Printf("line %v | number is %v-%v", i, string(maxNum), string(secondMax))
		result += int((maxNum-'0'))*10 + int((secondMax - '0'))
	}
	log.Printf("result: %v", result)
}
