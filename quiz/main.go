package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/moficodes/gophercises/quiz/quizer"
)

var quizTime int
var numQuestions int
var fileName string

func init() {
	const (
		timeValue = 30
		timeUsage = "Duration of quiz in seconds"
		numValue  = 10
		numUsage  = "Number of questions"
		fileValue = "problems.csv"
		fileUsage = "File to get the questions from"
	)
	flag.IntVar(&quizTime, "time", timeValue, timeUsage)
	flag.IntVar(&quizTime, "t", timeValue, timeUsage+" (shorthand)")
	flag.IntVar(&numQuestions, "quiz", numValue, numUsage)
	flag.IntVar(&numQuestions, "q", numValue, numUsage+" (shorthand)")
	flag.StringVar(&fileName, "file", fileValue, fileUsage)
	flag.StringVar(&fileName, "f", fileValue, fileUsage+" (shorthand)")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	flag.Parse()
	quizer.StartQuiz(fileName, numQuestions, quizTime)
}
