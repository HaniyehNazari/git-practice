package main

import (
	"fmt"
	"math/rand"
)

func main() {
	showWelcome()

	for {
		playGame()

		if !playAgain() {
			fmt.Println("Thanks for playing!")
			break
		}
	}
}

func showWelcome() {
	fmt.Println("Guess the Number")
	fmt.Println("Guess a number between 1 and 100!")
	fmt.Println("You have 5 attempts.")
}

func playGame() {
	secret := rand.Intn(100) + 1
	attempts := 0
	maxAttempts := 5

	for attempts < maxAttempts {
		guess, err := getGuess()

		if err != nil {
			fmt.Println("Please enter a number!")
			continue
		}

		attempts++

		if guess < secret {
			fmt.Println("Try a higher number!")
		} else if guess > secret {
			fmt.Println("Try a lower number!")
		} else {
			fmt.Println("Correct!")
			fmt.Println("Attempts:", attempts)
			return
		}
	}

	fmt.Println("No more attempts!")
	fmt.Println("The correct number was:", secret)
}

func getGuess() (int, error) {
	var guess int
	fmt.Print("Guess: ")

	_, err := fmt.Scan(&guess)

	return guess, err
}

func playAgain() bool {
	var answer string
	fmt.Print("Play again? (y/n): ")
	fmt.Scan(&answer)

	return answer == "y"
}
