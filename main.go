package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const (
	WelcomeMessage = "Please provide file with questions"
)

type (
	test struct {
		Time int
		csv  string
	}
	ScoreManager struct {
		QuestionCount int
		wrongAnswer   int
		rightAnswer   int
	}
)

const (
	DefaultTestFile = "problems.csv"
	DefaultTestTime = 15
)

func main() {
	var scoreData ScoreManager
	var initializer test

	flag.IntVar(&initializer.Time, "time", DefaultTestTime, "Set time for test")
	flag.StringVar(&initializer.csv, "file", DefaultTestFile, "tells the program where the test questoins are")
	flag.Parse()

	

	if initializer.csv == DefaultTestFile {
		fmt.Println("Defaulting to existing test...")
	}

	data, err := scoreData.ParseCsv(initializer.csv)

	if err != nil {
		log.Fatal("could not parse csv file.\nReason: ", err)
	}

	// Start timer goroutine
	timer := time.NewTimer(time.Duration(initializer.Time) * time.Second)

	go func() {
		<-timer.C
	}()

	for _, i := range data {

		if scoreData.wrongAnswer+scoreData.rightAnswer == scoreData.QuestionCount {
			break
		}

		answerch := make(chan string)

		go func() {
			answerch <- dataInput("Calculate " + i[0])
		}()

		select {
		case <-timer.C:
			fmt.Println("\nTime's up!")

			fmt.Printf("Your score from the test taken is %d out of %d", scoreData.rightAnswer, scoreData.QuestionCount)
			os.Exit(0)
		case answer := <-answerch:
			if answer == i[1] {
				scoreData.rightAnswer++
			} else {
				scoreData.wrongAnswer++
			}
		}
	}
	fmt.Printf("Your score from the test taken is %d out of %d\n", scoreData.rightAnswer, scoreData.QuestionCount)
}

func readFile(filename string) (string, error) {
	if filename == "" {
		filename = "problems.csv"
	}
	bs, err := os.ReadFile(filename)
	return string(bs), err
}

func (s *ScoreManager) ParseCsv(file string) ([][]string, error) {
	data, err := readFile(file)
	if err != nil {
		log.Fatal("failed to read file")
	}
	n := csv.NewReader(strings.NewReader(data))

	d, err := n.ReadAll()
	s.QuestionCount = len(d)

	return d, err
}

func dataInput(str string) (data string) {
	fmt.Println(str)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	data = strings.TrimSpace(input)
	return
}
