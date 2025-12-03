package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Part1() {

	isInvalidId := func(id string) bool {
		length := len(id)
		if length%2 != 0 {
			return false
		}
		return id[0:length/2] == id[length/2:length]
	}

	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Panicf("Error reading iput file %v", err)
	}
	ranges := strings.Split(string(content), ",")
	var result int64 = 0
	for i, currentRange := range ranges {
		log.Printf("%v range is %v", i, currentRange)
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
