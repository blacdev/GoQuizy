package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main(){
	fmt.Println("testing hwo to read csv files")
	data, err := ParseCsv("")
	if err != nil {
		log.Fatal("could not parse csv file")
	}
	fmt.Println(data)
}




func readFile(filename string) (string, error){
	if filename == ""{
		filename = "problems.csv"
	}
	bs, err := os.ReadFile(filename)
	return string(bs), err
}


func ParseCsv(file string) ([][]string, error) {
	data, err := readFile(file)
	if err != nil {
		log.Fatal("failed to read file")
	}
	n := csv.NewReader(strings.NewReader(data))
	
	d, err := n.ReadAll() 
	return d, err
}