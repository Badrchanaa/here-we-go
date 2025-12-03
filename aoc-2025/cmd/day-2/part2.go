package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Part2() {

	isInvalidId := func(id string) bool {
		pattern := id[0:1]
		currentLength := 0
		for i, c := range id {
			currentLength++
			idx := (currentLength - 1) % len(pattern)
			if pattern[idx] != byte(c) {
				if pattern[0] == byte(c) {
					pattern = id[0:i]
					currentLength = 1
				} else {
					pattern = id[0 : i+1]
					currentLength = 0
				}
			}
			if len(pattern) > len(id)/2 {
				return false
			}
		}
		if len(id)%len(pattern) != 0 {
			return false
		}
		return len(pattern) <= len(id)/2
	}

	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Panicf("Error reading iput file %v", err)
	}
	ranges := strings.Split(string(content), ",")
	var result int64 = 0
	for _, currentRange := range ranges {
		//log.Printf("%v range is %v", i, currentRange)
		rangeParts := strings.Split(currentRange, "-")
		start, err := strconv.Atoi(rangeParts[0])
		if err != nil {
			log.Panicf("error range part, err: %v", err)
		}
		end, err := strconv.Atoi(rangeParts[1])
		if err != nil {
			log.Panicf("error range part, err: %v", err)
		}
		for i := start; i <= end; i++ {
			if isInvalidId(strconv.FormatInt(int64(i), 10)) {
				log.Printf("%v is invalid", i)
				result += int64(i)
			}
		}
	}
	log.Printf("result is: %v", result)
}
