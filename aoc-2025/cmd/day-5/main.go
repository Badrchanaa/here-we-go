package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("test.txt")
	if err != nil {
		log.Panicf("error reading input file, err: %v", err)
	}
	lines := strings.Split(string(content), "\n")
	spans := [][2]int{}
	var i int
	var line string
	for i, line = range lines {
		if line == "" {
			break
		}
		rangeValues := strings.Split(line, "-")
		if len(rangeValues) != 2 {
			log.Panicf("invalid range in input file")
		}
		rangeStart, err := strconv.Atoi(rangeValues[0])
		rangeEnd, err2 := strconv.Atoi(rangeValues[1])
		if err != nil || err2 != nil {
			log.Panicf("invalid range in input file")
		}
		spans = append(spans, [2]int{rangeStart, rangeEnd})
	}
	counter := 0
	log.Printf("i: %v spans len: %v", i, len(spans))
	for j := i + 1; j < len(lines); j++ {
		if lines[j] == "" {
			break
		}
		num, err := strconv.Atoi(lines[j])
		if err != nil {
			log.Panicf("invalid number in input file")
		}
		for _, bounds := range spans {
			log.Printf("num: %v | bounds[0]: %v | bounds[1]: %v", num, bounds[0], bounds[1])
			if num >= bounds[0] && num <= bounds[1] {
				counter++
				break
			}
		}
	}
	log.Printf("part1 result is %v", counter)
}
