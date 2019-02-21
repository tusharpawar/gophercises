package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	csvFile := flag.String("csv", "problems.csv", "csv file name to read questions from")
	timeLimit := flag.Int("time", 30, "time limit for quiz in seconds")

	flag.Parse()

	//fmt.Println("csvFile", *csvFile)
	fmt.Println("Reading file", *csvFile)

	fileReader, _ := os.Open(*csvFile)
	r := csv.NewReader(bufio.NewReader(fileReader))

	inputReader := bufio.NewReader(os.Stdin)

	correct := 0
	incorrect := 0
	fmt.Println("Press enter to start, your time will end in ", *timeLimit, "seconds")
	input, _ := inputReader.ReadString('\n')

	timeChannel := make(chan string, 1)
	limit := *timeLimit
	go func() {
		time.Sleep(time.Duration(limit) * time.Second)
		timeChannel <- "complete"
	}()

	end := false

	for {

		select {
		case res := <-timeChannel:
			fmt.Println("Time: ", res)
			end = true
		default:
			record, err := r.Read()
			if err == io.EOF {
				end = true
			}
			if err != nil {
				log.Fatal(err)
				end = true
			}
			fmt.Printf("%s ", record[0])
			input, _ = inputReader.ReadString('\n')
			if strings.EqualFold(record[1], strings.TrimSpace(input)) {
				correct++
			} else {
				incorrect++
			}
		}
		if end {
			break
		}

	}
	fmt.Println("Correct", correct, " Incorrect:", incorrect)

}
