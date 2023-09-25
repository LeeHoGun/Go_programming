package main

import (
	"fmt"
)

func main() {
	fmt.Print("Guess number game!")
	answer := rand.intn(100) + 1 // 1 ~ 100
	fmt.Println(answer)
}
