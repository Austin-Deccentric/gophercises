package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"flag"
)

func readCsv(filename string) error{
	f, err := os.Open(filename)
	if err != nil {
		return errors.New("An error occured opening the file: " + err.Error())
	}
	// parse the file
	r := csv.NewReader(f)
	var answer string
	score := 0
	numOfQuestions := 0
	for {
		question, err := r.Read()
		if err == io.EOF{
			break
		}
		// Handle errors gracefully
		if err != nil{
			return errors.New(err.Error())
		}
		// Display the question and collect user input
		numOfQuestions++
		for {
			fmt.Println(numOfQuestions,question[0])
			_,err = fmt.Scan(&answer)
			if err !=nil{
				fmt.Println("Input only one number",err)
				continue	// Continues the loop if an error is encountered
			}
			break
		}
		// Checks the answer and score the question.
		if answer == question[1] {
			score++
		}
		
	}
	_,err = fmt.Printf("You got %d out of %d ",score,numOfQuestions)
	if err != nil {
		return errors.New(err.Error())
	}
	
	defer f.Close()
	return nil
}

//Declare a flag `file`
var filename string
func init() {
	const (
		defaultFile = "problems.csv"
		usage = "The filename of the csv file"
	)
	
	flag.StringVar(&filename,"file",defaultFile,usage)
	flag.StringVar(&filename,"f",defaultFile,usage+"(shorthand)")
	flag.Parse()  // Execute command-line parsing of flags
}
func main() {
	//wrdPtr := flag.String("file","problems.csv","This is the the csv filename")
	err := readCsv(filename)
	if err != nil {
		log.Fatal(err)
	}
}

