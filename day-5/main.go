package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// isInRange checks if the id is in range
func isInRange(r [2]uint64, id uint64) bool {
	return r[0] <= id && id <= r[1]
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

	sc := bufio.NewScanner(f)

	ranges := make([][2]uint64, 0)
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			break
		}
		var left uint64
		var right uint64
		n, err := fmt.Sscanf(line, "%d-%d", &left, &right)
		if err != nil || n != 2 {
			log.Fatalf("failed to parse range: %s", line)
		}

		ranges = append(ranges, [2]uint64{left, right})
	}

	count := 0
	for sc.Scan() {
		line := sc.Text()
		var id uint64

		n, err := fmt.Sscanf(line, "%d", &id)
		if err != nil || n != 1 {
			log.Fatalf("failed to parse ID: %s", line)
		}

		for _, r := range ranges {
			if isInRange(r, id) {
				count++
				break
			}
		}
	}

	if err := sc.Err(); err != nil {
		log.Fatalf("reading file: %v", err)
	}

	log.Printf("fresh products count: %d", count)
}
