package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"flag"
)

const(
	WelcomeMessage= "Please provide file with questions"
)

type ScoreManager struct {
	QuestionCount int
	wrongAnswer int
	rightAnswer int
}
var scoreData ScoreManager
func main(){
	
	fmt.Println("testing how to read csv files")
	
	input :=dataInput(WelcomeMessage)

	if input == ""{
		fmt.Println("Defaulting to existing test...")
	}
	data, err := scoreData.ParseCsv(input)
	if err != nil {
		log.Fatal("could not parse csv file.\nReason: ",err)
	}

	for _, i := range data {

		if scoreData.wrongAnswer + scoreData.rightAnswer == scoreData.QuestionCount{
			break
		}
		question := dataInput("Calculate " +  i[0])
		if question == i[1]{
			scoreData.rightAnswer+=1
		} else {
			scoreData.wrongAnswer +=1
		}
	}
	fmt.Println(scoreData)
}




func readFile(filename string) (string, error){
	if filename == ""{
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

func dataInput(str string) (data string){
	fmt.Println(str)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	data = strings.TrimSpace(input)
	return
}