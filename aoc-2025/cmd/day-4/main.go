package main

import (
	"log"
	"os"
	"strings"
)

func solve(grid [][]rune, fn func([][]rune, int, int)) int {
	counter := 0
	neighbors := [][]int{
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left

		{-1, -1}, // up left
		{-1, 1},  // up right
		{1, 1},   // down right
		{1, -1},  // down left
	}

	for i := range grid {
		width := len(grid[i])
		for j, c := range grid[i] {
			if c != '@' {
				continue
			}
			rollNeighbors := 0
			for _, neighbor := range neighbors {
				x, y := neighbor[0], neighbor[1]
				if x+i >= len(grid) || x+i < 0 || y+j >= width || y+j < 0 {
					continue
				}
				if grid[x+i][y+j] == '@' {
					rollNeighbors++
				}
			}
			if rollNeighbors < 4 {
				counter++
				if fn != nil {
					fn(grid, i, j)
				}
			}
		}
	}
	return counter
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Panicf("error reading input file: %v", err)
	}
	lines := strings.Split(string(content), "\n")
	grid := [][]rune{}
	for i, line := range lines {
		if line == "" {
			break
		}
		grid = append(grid, []rune{})
		for _, c := range line {
			grid[i] = append(grid[i], c)
		}
	}
	part1Result := solve(grid, nil)
	log.Printf("Part1: total rolls of paper that can be accessed %v", part1Result)
	log.Print("------------------------")
	part2Result := 0
	part2fn := func(grid [][]rune, x int, y int) {
		grid[x][y] = 'x'
	}
	for result := solve(grid, part2fn); result != 0; result = solve(grid, part2fn) {
		part2Result += result
	}
	log.Printf("Part2: total rolls of paper that can be accessed %v", part2Result)
}
