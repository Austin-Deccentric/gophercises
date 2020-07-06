package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

//var err error

func main() {
	csvfile := flag.String("file", "problem.csv", "Specifies the csv file with the problem set")
	timeLimit := flag.Int("time", 10, "The duration of the quizz.")
	quizzLength := flag.Int("n", 14, "Specifies the number of questions in the quizz.")
	flag.Parse()

	file, err := os.Open(*csvfile)
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(file)
	numofQuestion := 0
	correct := 0
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for {
		quizz, err := r.Read()
		if err == io.EOF {
			return
		}
		if err != nil {
			panic(err)
		}

		answerchan := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerchan <- answer
		}()

		numofQuestion++
		fmt.Printf("Question %d: %s = ", numofQuestion, quizz[0])
		select {
		case <-timer.C:
			fmt.Printf("\nScored %d out of %d", correct, *quizzLength)
			return
		case answer := <-answerchan:
			if answer == quizz[1] {
				correct++
			}
		}

		if numofQuestion == *quizzLength {
			fmt.Printf("\nScored %d out of %d", correct, *quizzLength)
			return
		}
	}
}
