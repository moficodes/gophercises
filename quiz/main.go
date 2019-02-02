package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

var quizTime int
var numQuestions int

type quiz struct {
	question string
	answer   string
}

func init() {
	const (
		timeValue = 30
		timeUsage = "Duration of quiz in seconds"
		numValue  = 10
		numUsage  = "Number of questions"
	)
	flag.IntVar(&quizTime, "time", timeValue, timeUsage)
	flag.IntVar(&quizTime, "t", timeValue, timeUsage+" (shorthand)")
	flag.IntVar(&numQuestions, "quiz", numValue, numUsage)
	flag.IntVar(&numQuestions, "q", numValue, numUsage+" (shorthand)")
}

func parseFile(file string) []quiz {
	quizes := make([]quiz, 0)
	data, err := ioutil.ReadFile("problems.csv")
	if err != nil {
		log.Panic("Problem Reading FILE")
	}
	r := csv.NewReader(bytes.NewReader(data))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panic(err)
		}
		quizes = append(quizes, quiz{question: record[0], answer: record[1]})
	}
	return quizes
}

func getQuestion(quizes []quiz) quiz {
	index := rand.Intn(len(quizes))
	return quizes[index]
}

func isCorrect(answer string, question quiz) bool {
	return answer == question.answer
}

func main() {
	rand.Seed(time.Now().UnixNano())
	flag.Parse()
	quizes := parseFile("problems.csv")
	fmt.Println(getQuestion(quizes))
}
