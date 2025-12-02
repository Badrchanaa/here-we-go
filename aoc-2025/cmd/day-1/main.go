package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Printf("Error reading input file: %v\n", err)
	}
	lines := strings.Split(string(content), "\n")
	password := 0
	dial := 50
	for _, v := range lines {
		if len(v) == 0 {
			continue
		}
		val, err := strconv.Atoi(v[1:])
		if err != nil {
			log.Fatalf("invalid input line %v\n", v)
		}
		RotationDirection := v[0]
		// log.Printf("dial: %v direction %v value %v", dial, string(RotationDirection), val)
		switch RotationDirection {
		case 'R':
			password += (dial + val) / 100
			// log.Printf("password + %v", (dial+val)/100)
			dial = (dial + val) % 100
		case 'L':
			password += val / 100
			// log.Printf("password + %v", (val)/100)
			if dial != 0 && val%100 >= dial {
				// log.Print("password + 1")
				password++
			}
			dial = (dial - val + ((val/100)+1)*100) % 100
		default:
			log.Fatalf("invalid input line %v\n", v)
		}
		// log.Printf("new dial %v\n", dial)
	}
	log.Printf("password is: %v\n", password)
}
