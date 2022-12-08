package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var guess int64
var answer int64
var max *int64
var isMe *bool

func main() {
	rand.Seed(time.Now().UnixNano())

	max = flag.Int64("max", 100, "maximum random number")
	isMe = flag.Bool("isMe", true, "true (default) if you wants to play, false if you wants computer to solve")
	flag.Parse()

	play()
}

func play() {
	randomNumber()
	if *isMe {
		inputNumber()
	}
}

func randomNumber() {
	answer = rand.Int63n(*max)
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
