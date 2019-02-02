package quizer

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// problem is the struct to hold problem question and answer
type problem struct {
	question string
	answer   string
}

// StartQuiz controls the problem game
func StartQuiz(file string, questions int, timelimit int) {
	problems := parseFile(file)

	scanner := bufio.NewScanner(os.Stdin)

	correct := 0
	fmt.Println("Press [Enter] to begin")
	for scanner.Scan() {
		break
	}

	answerCh := make(chan string)

	timer := time.NewTimer(time.Duration(timelimit) * time.Second)

	for i := 0; i < questions; i++ {
		index := getQuestion(&problems)
		q := problems[index]
		printQuestion(q)
		problems = append(problems[:index], problems[index+1:]...)
		go getUserInput(scanner, answerCh)
		select {
		case <-timer.C:
			result(correct, questions)
			return
		case answer := <-answerCh:
			if isCorrect(answer, q.answer) {
				correct++
			}
		}
	}

	result(correct, questions)
}

func result(correct int, total int) {
	ratio := (float32(correct) / float32(total)) * 100
	fmt.Printf("\nYou got %d correct out of %d questions. \nYour score is %.2f%% \n", correct, total, ratio)
}

func getUserInput(scanner *bufio.Scanner, ans chan string) {
	var s string
	for scanner.Scan() {
		s = processString(scanner.Text())
		break
	}
	ans <- s
}

func processString(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func printQuestion(q problem) {
	fmt.Printf("What is %s : ", q.question)
}

// GetQuestion Returns a single problem to user
// Remove the question that was just asked
func getQuestion(problems *[]problem) int {
	return rand.Intn(len(*problems))
}

// Check if a answer is correct
func isCorrect(given string, expected string) bool {
	return given == expected
}

func parseFile(file string) []problem {
	var problems []problem
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
			log.Fatalln(err)
		}
		problems = append(problems, problem{question: processString(record[0]), answer: processString(record[1])})
	}
	return problems
}
