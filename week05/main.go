package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Input score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScore := reader.ReadString('\n') // 1 variable but reader.ReadString returns 2 values
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(inputScore)
	if inputScore >= 90 {
		grade := "A grade!"
	} else {
		grade := "under A grade..."
	}
}
