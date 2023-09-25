package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Input score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScore := reader.ReadString('\n') // 1 variable but reader.ReadString returns 2 values
	if err != nil {
		log.Fatal(err)
	}
	inputScoreString = strings.TrimSpace(inputScoreString)
	inputScore, err := strconv.ParseFloat(inputScoreString, 32)
	if err != nil {
		log.Fatal(err)
	}
	var grade string
	if inputScore >= 90 {
		grade = "A grade!"
	} else {
		grade = "under A grade..."
	}
	fmt.Println("You got", grade)
}
