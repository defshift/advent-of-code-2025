package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// compareHalves compares halves of the number
func compareHalves(number string) bool {
	middle := len(number) / 2
	first := number[:middle]
	second := number[middle:]
	return first == second
}

// sumBadIdsFromRange finds all bad ids and calculates sum of it
func sumBadIdsFromRange(start int, end int) int {
	sum := 0
	for i := start; i < end; i++ {
		num := strconv.Itoa(i)
		if len(num)%2 != 0 {
			continue
		}
		if compareHalves(num) {
			sum += i
		}
	}
	return sum
}

// parses sources string and returns ranges
func parse(input string) [][2]int {
	rangeStrings := strings.Split(input, ",")
	ranges := make([][2]int, 0, len(rangeStrings))
	for _, r := range rangeStrings {
		parts := strings.Split(r, "-")
		if len(parts) != 2 {
			log.Fatalf("invalid range format: %s", r)
		}
		rng := [2]int{0, 0}
		left, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("cannot parse range start: %v", err)
		}
		rng[0] = left
		right, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("cannot parse range end: %v", err)
		}
		rng[1] = right
		ranges = append(ranges, rng)
	}

	return ranges
}

func main() {
	// get the inputfile from the command line arguments
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
	var line string
	for sc.Scan() {
		line = sc.Text()
	}

	ranges := parse(line)

	sum := 0
	for _, r := range ranges {
		sum += sumBadIdsFromRange(r[0], r[1])
	}

	log.Printf("result is %d", sum)
}
