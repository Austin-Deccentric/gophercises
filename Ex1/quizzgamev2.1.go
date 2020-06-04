package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var (
	csvfile string
	timeLimit int
	correct int
	//quizzStruct []problem
)

func init() {
	
	const (
		defaultFile = "problem.csv"
		usage = "The file to use for the problem"
		defaultTime = 5
		timeUsage = "This is the time duration of the quizz"
	)

	flag.StringVar(&csvfile,"file",defaultFile, usage)
	flag.StringVar(&csvfile,"f", defaultFile, usage)
	flag.IntVar(&timeLimit,"time", defaultTime, timeUsage)
	flag.IntVar(&timeLimit,"t", defaultTime, timeUsage)
}

type problem struct {
	q string
	a string
}

func readCsv(file string) ([][]string, error ) {
	f, err := os.Open(file); if err!= nil {
		return nil, err
	}
	defer f.Close()
	var quizz [][]string
	reader := csv.NewReader(f)
	quizz, err = reader.ReadAll()
	return quizz, err
}

func makeQuizz(question [][]string) ([]problem) {
	quizz := make([]problem, len(question))
	for i, line := range question{
		quizz[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return quizz
}

func askQuizz(problems []problem) {

	for i, question := range problems{
		timer := counter(len(problems))
		defer timer.Stop()
		fmt.Printf("Question #%d: %s \n",i+1, question.q)
		var answer string
		fmt.Scanf("%s\n",&answer)
		if answer == question.a{
			correct++
			//fmt.Println(correct)
		}
	}
}

func counter(n int) *time.Timer{
	return  time.AfterFunc(time.Duration(timeLimit)* time.Second, func ()  {
		fmt.Println("\n Your time has elasped")
		fmt.Printf("You scored %d out of a possible %d",correct, n)
		os.Exit(0)	
	})
}

func main()  {
	flag.Parse()
	quizz, err := readCsv(csvfile); if err!= nil{
		log.Fatal(err)
	}
	// fmt.Println(quizz)
	
	quizzStruct:= makeQuizz(quizz)
	
	askQuizz(quizzStruct)
	fmt.Printf("You scored %d out of a possible %d",correct, len(quizzStruct))

}

