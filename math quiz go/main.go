package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	//csv
	csvFilename := flag.String("csv", "problems.csv", "a csv file in question and answer format")
	//timelimit
	timeLimit := flag.Int("limit", 30, "time limit for the quiz in seconds")
	//read flags
	flag.Parse()

	//read csv
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("failed to open csv file: %s", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("failed to read th csv file"))
	}

	problems := parseLines(lines)                                   //read lines in csv
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second) //convert int type to type time.duration to avoid mismatch errors
	correct := 0                                                    //default correct count
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q) //print out the questions
		answerChannel := make(chan string)           //create channel
		//go routine anonymous function
		go func() {
			var answer string          //create variable for the answer to store when entred
			fmt.Scanf("%s\n", &answer) //scan answer
			answerChannel <- answer    //nb: arrow always point to where the data is moving, sending answer to answerChannel
		}()
		select {
		case <-timer.C: //channel wait for problems
			fmt.Printf("\nScored %d out of %d. \n", correct, len(problems)) //compare correct answers with the total number of problems
			return
		case answer := <-answerChannel:
			//check for correct answer
			if answer == p.a {
				correct++ //increment correct count
			}

		}

	}

}

//error exit
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]), //handle invalid csv
		}
	}
	return ret
}

//problem struct
type problem struct {
	q string
	a string
}
