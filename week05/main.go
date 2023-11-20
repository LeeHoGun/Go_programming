package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) // get the current data and time as an integer
	answer := rand.Intn(100) + 1 // random integer number (1 ~ 100)
	fmt.Println("Guess Number Game~")
	fmt.Println(answer)
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 10; i++ {
		fmt.Println("You have", 10-i, "chances~")
		fmt.Print("Input guess number : ")
		inputGuessString, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		inputGuessString = strings.TrimSpace(inputGuessString)
		inputGuess, err := strconv.Atoi(inputGuessString)
		if err != nil {
			log.Fatal(err)
		}
		if inputGuess < answer {
			if inputGuess == answer {
				fmt.Println("Great! U got number. congratulations~") // Answer is higher~
				break
			} else if inputGuess < answer {
				fmt.Println("Guess number is lower then answer") // Answer is higher~
			} else if inputGuess > answer {
				fmt.Println("Guess number is higher then answer") // Answer is lower
			}
		}
	}
}
