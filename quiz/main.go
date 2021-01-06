package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	csvFile := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
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
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

problemloop:
	for index, line := range lines {
		p := problem{question: line[0], answer: line[1]}
		fmt.Printf("Problem #%d: %s = ", index+1, p.question)

		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-answerCh:
			if answer == p.answer {
				score++
			}
		}
	}

	fmt.Printf("You scored %d of %d.\n", score, total)
}

type problem struct {
	question string
	answer   string
}
