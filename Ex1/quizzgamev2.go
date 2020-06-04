package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)
var correct int

func main()  {
	csvfilename := flag.String("file", "problems.csv", "The filename of the csv file")
	flag.Parse()


	file, err:= os.Open(*csvfilename)
	if err!= nil {
		exit(fmt.Sprintf("An error occured opening file %s",*csvfilename))
		
	}
	r := csv.NewReader(file)
	//Parse the read files
	lines,err := r.ReadAll()
	if err!= nil {
		exit("Error Parsing csv file")
	}
	go func() {
		timer:= time.NewTimer(5 * time.Second)
		<-timer.C
		fmt.Printf("\nYou got %d out of %d", correct, len(lines))
		exit("\nSorry you have run out of time.")
	}()

	parsequizz(lines,quizzMaster)  //pardon the complexity, I was testing a new concept.
	

	
}

//problems models the quizz structure.
type problems struct {
	question string
	answer string
}

//exit prints a message and makes a call to Exit with staus code 1.
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

//parsequizz transforms the 2D array into an array of problem structs and passes the result
// to another function for further processing.
func parsequizz(lines [][]string, next func(q []problems)) {
	test := make([]problems, len(lines))  // make an array of type problem and capacity of lines(no of rows in the the 2D array)
	for i , quizz := range lines{
		test[i] = problems{
			question: quizz[0],
			answer: strings.TrimSpace(quizz[1]), 

		}
	}
	next(test)
}

//quizzMaster takes in an array of problems and presents them to the user.
func quizzMaster(q []problems) {
	
	for i, problem:= range q{
		fmt.Printf("Question #%d %s \n=", i+1, problem.question)
		var answer string
		fmt.Scanf("%s\n",&answer)
		if answer == problem.answer{
			correct++
		}
	}
	
}