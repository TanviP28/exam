package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 10
	target := rand.Intn(10) + 1

	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I've picked a number between 1 and 10. Try to guess it!")

	var guess int
	for {
		fmt.Print("Your guess: ")
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		if guess < 1 || guess > 10 {
			fmt.Println("Please enter a number between 1 and 10.")
			continue
		}

		if guess < target {
			fmt.Println("Too low! Try again.")
		} else if guess > target {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Println("Congratulations! You've guessed it right!")
			break
		}
	}
}
