package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main(){
	fmt.Println("testing hwo to read csv files")
	data, err := readHtmlFile("problems.csv")
	if err != nil {
		log.Fatal("failed to read file")
	}
	n := csv.NewReader(strings.NewReader(data))
	
	for {
		record, err := n.Read()

		if err == io.EOF{
			break
		}

		if err != nil {
			log.Fatal("something is wrong with the file")
		}

		fmt.Println(record)
	}
}

func readHtmlFile(filename string) (string, error){
	bs, err := os.ReadFile(filename)
	return string(bs), err
}