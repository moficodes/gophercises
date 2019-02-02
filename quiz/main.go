package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/moficodes/quiz/quizer"
)

var quizTime int
var numQuestions int

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

func main() {
	rand.Seed(time.Now().UnixNano())
	flag.Parse()
	file := "problems.csv"
	quizer.StartQuiz(file, numQuestions, quizTime)
}
