package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	defaultFile = "problems.csv"
	usage       = "a csv file in the format of 'question,answer'"
)

func main() {
	csvFile := flag.String("csv", defaultFile, usage)
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	score := 0
	total := len(lines)
	for index, line := range lines {
		question := line[0]
		answer := line[1]

		fmt.Printf("Problem #%d: %s = ", index+1, question)

		reader := bufio.NewReader(os.Stdin)

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSuffix(input, "\n")

		// log.Printf("input=%q, answer=%q", input, answer)
		if input == answer {
			score++
		}

	}

	fmt.Printf("You scored %d of %d.\n", score, total)
}
