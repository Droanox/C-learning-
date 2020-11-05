package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

var questions int

func timer(timer int) {
	time.Sleep(time.Duration(timer) * time.Second)
	fmt.Println("Oopsie woopsie, you ran out of time retard. Maybe you should have been quicker")
	os.Exit(0)
}

func main() {
	timerFlag := flag.Int("time", 30, "The timer(s)")
	csvFlag := flag.String("quiz", "problems.csv", "Input csv file")
	flag.Parse()

	csvFile, err := os.Open(*csvFlag)
	if err != nil {
		fmt.Println("Error: Could not open csv file, try -help")
		return
	}
	csvReader := csv.NewReader(csvFile)

	line, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	questions = len(line)

	fmt.Println("The quiz: " + string(*csvFlag))
	fmt.Println()
	fmt.Println("The time limit for taking this test is", *timerFlag, "seconds. To begin, please click ENTER.")
	fmt.Scanln()
	fmt.Println()

	go timer(*timerFlag)

	var score int

	incorrectAnswer := []bool{}

	for i := 0; i < questions; i++ {
		fmt.Println(i, ") What's the answer to = "+line[i][0])

		var answer string
		fmt.Scanln(&answer)
		fmt.Println()

		if answer == line[i][1] {
			score++
			incorrectAnswer = append(incorrectAnswer, false)
		} else {
			incorrectAnswer = append(incorrectAnswer, true)

		}
	}
	wrongAnswer(incorrectAnswer, line)
	fmt.Println()
	fmt.Println("You score is: ", score, "/", questions)
	fmt.Println()
	return

}

func wrongAnswer(incorrectAnswer []bool, line [][]string) {
	for i := 0; i < len(incorrectAnswer); i++ {
		if incorrectAnswer[i] == true {
			fmt.Println(i, ") The question you got wrong is: ", line[i][0])
			fmt.Println("The correct answer is: ", line[i][1])
			fmt.Println()
		}
	}
	return
}
