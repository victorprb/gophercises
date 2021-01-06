package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type problem struct {
	question string
	answer   string
}

func main() {
	csvFile := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Duration("limit", 30*time.Second, "the time limit for the quiz in seconds")
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
		p := problem{question: line[0], answer: line[1]}

		fmt.Printf("Problem #%d: %s = ", index+1, p.question)

		select {
		case <-time.After(*timeLimit):
			fmt.Printf("You scored %d of %d.\n", score, total)
			return
		case answer := <-receiveAnswer():
			if checkAnswer(p.answer, answer) {
				score++
			}
		}
	}

	fmt.Printf("You scored %d of %d.\n", score, total)
}

func receiveAnswer() chan string {
	ch := make(chan string)
	var answer string

	go func() {
		fmt.Scanf("%s\n", &answer)
		ch <- answer
	}()

	return ch
}

func checkAnswer(want, got string) bool {
	// debug
	log.Printf("got=%q, want=%q", got, want)
	return got == want
}
