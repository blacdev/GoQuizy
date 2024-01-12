package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	WelcomeMessage   = "Please provide file with questions"
	DefaultTestFile  = "problems.csv"
	DefaultTestTime  = 30
	seedValue = 4
	Defaultrandomize = ""
)

type (
	test struct {
		Time      int
		csv       string
		randomize string
	}
	ScoreManager struct {
		QuestionCount int
		wrongAnswer   int
		rightAnswer   int
	}
)


func main() {
	var scoreData ScoreManager
	var initializer test

	
	flag.IntVar(&initializer.Time, "time", DefaultTestTime, "Set time for test")
	flag.StringVar(&initializer.csv, "file", DefaultTestFile, "tells the program where the test questoins are")
	flag.StringVar(&initializer.randomize,"random", Defaultrandomize, "rearranges the set of questions each time the test is taken.")
	flag.Parse()
	fmt.Println(initializer.randomize)

	
	 
	if initializer.csv == DefaultTestFile {
		
		fmt.Println("Defaulting to existing test...")
	}

	data, err := scoreData.ParseCsv(initializer.csv, initializer.randomize)

	if err != nil {
		log.Fatal("could not parse csv file.\nReason: ", err)
	}

	// Start timer goroutine
	dataInput("Press 'enter' to start test...")
	timer := time.NewTimer(time.Duration(initializer.Time) * time.Second)

	//goroutine to send a message to timer channel to stop test
	go func() {
		<-timer.C
	}()

	var questionNumber int

loop:
	for _, i := range data {

		questionNumber++

		if scoreData.wrongAnswer+scoreData.rightAnswer == scoreData.QuestionCount {
			break
		}

		answerch := make(chan string)

		fmt.Printf("Question %d:\n", questionNumber)

		go func() {
			answerch <- dataInput("Calculate " + i[0] + "\nAnswer: ")
		}()

		select {

		case <-timer.C:
			fmt.Println("\nTime's up!")
			break loop

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

func (s *ScoreManager) ParseCsv(file string, randomize string) (data [][]string, err error) {

	fileData, err := readFile(file)

	if err != nil {
		log.Fatal("failed to read file")
	}

	nData := csv.NewReader(strings.NewReader(fileData))
	data, err = nData.ReadAll()

	if randomize == "true" {

		newSource := rand.NewSource(seedValue)
		rand.New(newSource)
		rand.Shuffle(len(data), func(i, j int) {
			data[i], data[j] = data[j], data[i]
		})
	} else {

	}
	s.QuestionCount = len(data)

	return
}

func dataInput(str string) (data string) {

	fmt.Print(str)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	data = strings.TrimSpace(input)
	return
}
