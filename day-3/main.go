package main

import (
	"bufio"
	"log"
	"os"
)

const (
	maxVoltage = 9
	asciiShift = '0'
)

// calculateVoltage calculates max voltage of the battery arrange
func calculateVoltage(battery string) int {
	first := battery[0]
	second := battery[1]

	for i := 1; i < len(battery); i++ {
		n := battery[i]
		if first != (maxVoltage+asciiShift) && n > first && i != len(battery)-1 {
			first = n
			second = battery[i+1]
			continue
		}

		if n > second {
			second = battery[i]
		}
	}

	return int((first-asciiShift)*10 + second - asciiShift)
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

	sum := 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		sum += calculateVoltage(sc.Text())
	}

	log.Printf("result is %d", sum)
}
