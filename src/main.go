package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var inputedInt int
	var j string
	tries := 0
	println := true
	var intToGuess int

	fmt.Println("Welcome to my number guessing game")
	fmt.Println("To start please type any key")
	fmt.Scanln(&j)

	for {
		if println {
			fmt.Println("Enter a number between 0 and 9:")
			intToGuess = rand.Intn(10)
		}

		println = false

		tries ++
		fmt.Scanln(&inputedInt)

		if intToGuess == inputedInt {
			fmt.Println("Congratulations! You guessed correctly. ", tries)
			println = true
			fmt.Println("Do you want to continue? (y/n)")
			fmt.Scanln(&j)
			
			for {
				if j == "n" || j == "N" {
					fmt.Println("Thank you for playing!")
					return
				} else if j == "y" || j == "Y" {
					tries = 0
					break
				} else {
					fmt.Println("Please enter a valid answer (y/n):")
					fmt.Scanln(&j)
				}
			}
		} else {
			fmt.Println("Please try again.")
		}
	}
}
