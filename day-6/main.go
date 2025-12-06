package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	multOp = "*"
	addOp  = "+"
	subOp  = "-"
)

func Op(op string, first int64, second int64) int64 {
	if op == multOp {
		return first * second
	}
	if op == addOp {
		return first + second
	}
	if op == subOp {
		return first - second
	}
	return 0
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
	tokens := make([][]string, 0)
	for sc.Scan() {
		t := strings.Fields(sc.Text())
		tokens = append(tokens, t)
	}

	var total int64
	for j := 0; j < len(tokens[0]); j++ {
		op := tokens[len(tokens)-1][j]
		var res int64
		for i := 0; i < len(tokens)-1; i++ {
			var num int64
			if _, err := fmt.Sscanf(tokens[i][j], "%d", &num); err != nil {
				log.Fatalf("incorrect token: %s", tokens[i][j])
			}
			if i == 0 {
				res = num
			} else {
				res = Op(op, res, num)
			}
		}
		total += res
	}

	if err := sc.Err(); err != nil {
		log.Fatalf("reading file: %v", err)
	}
	log.Printf("result sum is %d", total)
}
