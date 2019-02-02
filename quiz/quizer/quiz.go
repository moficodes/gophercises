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
)

// quiz is the struct to hold quiz question and answer
type quiz struct {
	question string
	answer   string
}

// StartQuiz controls the quiz game
func StartQuiz(file string, questions int, timelimit int) {
	quizes := parseFile(file)

	scanner := bufio.NewScanner(os.Stdin)

	correct := 0
	fmt.Println("Press [Enter] to begin")

	for i := 0; i < questions; i++ {
		index := getQuestion(&quizes)
		q := quizes[index]
		printQuestion(q)
		quizes = append(quizes[:index], quizes[index+1:]...)
		fmt.Println(len(quizes))
		s := getUserInput(scanner)
		if s == q.answer {
			fmt.Println("Correct")
			correct++
		} else {
			fmt.Println("Incorrect")
		}
	}
}

func getUserInput(scanner *bufio.Scanner) (s string) {
	for scanner.Scan() {
		s = processString(scanner.Text())
		break
	}
	return
}

func processString(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func printQuestion(q quiz) {
	fmt.Printf("What is %s", q.question)
}

// GetQuestion Returns a single quiz to user
// Remove the question that was just asked
func getQuestion(quizes *[]quiz) int {
	return rand.Intn(len(*quizes))
}

// Check if a answer is correct
func isCorrect(answer string, question quiz) bool {
	return answer == question.answer
}

func parseFile(file string) []quiz {
	var quizes []quiz
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
		quizes = append(quizes, quiz{question: processString(record[0]), answer: processString(record[1])})
	}
	return quizes
}
