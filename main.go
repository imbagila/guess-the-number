package main

import (
	"fmt"
	"math/rand"
	"time"
)

var guess int
var answer int

func main() {
	rand.Seed(time.Now().UnixNano())
	play()
}

func play() {
	randomNumber()
	inputNumber()
}

func randomNumber() {
	answer = rand.Intn(100)
}

func inputNumber() {
	fmt.Print("Guess the number: ")
	_, err := fmt.Scanln(&guess)
	if err != nil {
		inputNumber()
	}
	check()
}

func check() {
	switch {
	case guess > answer:
		fmt.Println("Too big !")
		inputNumber()
	case guess < answer:
		fmt.Println("Too small !")
		inputNumber()
	default:
		fmt.Println("Correct !")
	}
}
