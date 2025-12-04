package main

import (
	"bufio"
	"log"
	"os"
)

// countIsolatedAt counts the number of isolated '@' characters in the grid
func countIsolatedAt(grid [][]rune) int {
	m := len(grid)
	n := len(grid[0])

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '@' {
				adjacentCount := 0
				for di := -1; di <= 1; di++ {
					for dj := -1; dj <= 1; dj++ {
						i2 := i + di
						j2 := j + dj
						if i2 >= 0 && i2 < m && j2 >= 0 && j2 < n && grid[i2][j2] == '@' {
							adjacentCount++
						}
					}
				}
				if adjacentCount <= 4 {
					count++
				}
			}
		}
	}
	return count
}

// countRemovableAt counts total '@' that can be removed by repeatedly
// removing those with <= 4 adjacent '@' neighbors
func countRemovableAt(grid [][]rune) int {
	m := len(grid)
	n := len(grid[0])

	// Make a copy to avoid mutating original
	g := make([][]rune, m)
	for i := range grid {
		g[i] = make([]rune, n)
		copy(g[i], grid[i])
	}

	total := 0
	for {
		var toRemove [][2]int
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if g[i][j] == '@' {
					adjacentCount := 0
					for di := -1; di <= 1; di++ {
						for dj := -1; dj <= 1; dj++ {
							i2 := i + di
							j2 := j + dj
							if i2 >= 0 && i2 < m && j2 >= 0 && j2 < n && g[i2][j2] == '@' {
								adjacentCount++
							}
						}
					}
					if adjacentCount <= 4 {
						toRemove = append(toRemove, [2]int{i, j})
					}
				}
			}
		}
		if len(toRemove) == 0 {
			break
		}
		for _, pos := range toRemove {
			g[pos[0]][pos[1]] = 'x'
		}
		total += len(toRemove)
	}
	return total
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: %s <inputfile>", os.Args[0])
	}
	srcFile := os.Args[1]

	f, err := os.Open(srcFile)
	if err != nil {
		log.Fatalf("open file %s: %v", srcFile, err)
	}
	defer f.Close()

	var grid [][]rune
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		grid = append(grid, []rune(sc.Text()))
	}

	log.Printf("part1: %d", countIsolatedAt(grid))
	log.Printf("part2: %d", countRemovableAt(grid))
}
