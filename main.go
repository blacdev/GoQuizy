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

	
	problem struct {
		q, a string
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

	data := scoreData.ParseCsv(initializer.csv, initializer.randomize)

	

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
			answerch <- dataInput("Calculate " + i.q + "\nAnswer: ")
		}()

		select {

		case <-timer.C:
			fmt.Println("\nTime's up!")
			break loop

		case answer := <-answerch:

			if answer == i.a {
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

func (s *ScoreManager) ParseCsv(file string, randomize string) (ret []problem) {
	fileData, err := readFile(file)
	
	if err != nil {
		log.Fatal("failed to read file")
	}
	
	nData := csv.NewReader(strings.NewReader(fileData))
	records, err := nData.ReadAll()

	if err != nil {
		log.Fatal("failed to read CSV data")
	}
	
	ret = make([]problem, len(records))

	for i, j := range records {
		ret[i] = problem{
			q: j[0],
			a: j[1],
		}
	}

	if randomize == "true" {

		newSource := rand.NewSource(seedValue)
		rand.New(newSource)
		rand.Shuffle(len(ret), func(i, j int) {
			ret[i], ret[j] = ret[j], ret[i]
		})
	}
	s.QuestionCount = len(ret)

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
