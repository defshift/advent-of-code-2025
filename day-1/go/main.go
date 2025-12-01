package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	leftLabel  = 'L'
	rightLabel = 'R'

	circleMin = 0
	circleMax = 99

	startArrowPosition = 50
)

func shift(state int, count int) int {
	res := state + count

	overflow := res > circleMax || res < circleMin
	// no overflow, return as is
	if !overflow {
		return res
	}

	// check sign and correct the number after overflow
	sign := 1
	overflowPositive := res > 0
	if overflowPositive {
		sign = -1
	}

	return res + sign*(circleMax+1)
}

// parseLine converts string representation of the dial rotation command to signed integer within range -99..99
// Example:
// L1 -> -1
// R12 -> 12
func parseLine(line string) (int, error) {
	var dirRaw byte
	var countRaw int

	// validate string
	if _, err := fmt.Sscanf(line, "%c%d", &dirRaw, &countRaw); err != nil {
		return 0, fmt.Errorf("cannot parse line: %w", err)
	}

	var sign int
	switch dirRaw {
	case leftLabel:
		sign = -1
	case rightLabel:
		sign = 1
	default:
		// strong requirements for case
		return 0, fmt.Errorf("unrecognized direction %c", dirRaw)
	}

	// validate absolute amount of steps
	// zero or negative steps values are not accepted
	if countRaw <= 0 {
		return 0, fmt.Errorf("invalid steps %d", countRaw)
	}

	// if count > 100, it means we made a full circle at least one time
	// then modulo gives us the final destination
	if countRaw > circleMax {
		countRaw = countRaw % (circleMax + 1)
	}

	return countRaw * sign, nil
}

func main() {
	// get the logfile from the command line arguments
	if len(os.Args) < 2 {
		log.Fatalf("usage: %s <logfile>", os.Args[0])
	}
	srcFile := os.Args[1]

	f, err := os.Open(srcFile)
	if err != nil {
		log.Fatalf("open file %s: %v", srcFile, err)
	}
	defer f.Close()

	arrowState := startArrowPosition
	amountOfZeros := 0

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		count, err := parseLine(line)
		if err != nil {
			log.Fatalf("parse line error: %v", err)
		}

		arrowState = shift(arrowState, count)
		if arrowState == 0 {
			amountOfZeros++
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
	}

	log.Printf("password is %d", amountOfZeros)
}
