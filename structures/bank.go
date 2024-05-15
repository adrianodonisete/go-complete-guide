package main

import (
	"fmt"
)

var accountBalance float64 = 1000

func main() {
	// for i := 0; i < 100; i++ {
	for {
		if !run() {
			break
		}
	}
}

func run() bool {
	var choice int

	fmt.Println("Welcome to Bank!")
	fmt.Println("What do you want?")

	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit")
	fmt.Println("4. Exit")

	fmt.Print(`Your choice: `)
	fmt.Scan(&choice)

	var keepRunning bool
	if choice == 1 {
		keepRunning = true
		fmt.Println("Your Balance is: ", accountBalance)
	} else if choice == 2 {
		keepRunning = true
		var depositValue float64
		fmt.Print(`Input the value: `)
		fmt.Scan(&depositValue)
		accountBalance += depositValue
		fmt.Println("Your Balance is: ", accountBalance)
	} else if choice == 4 {
		keepRunning = false
		fmt.Println("End!")
	} else {
		keepRunning = true
		fmt.Println("Invalid option!")
	}
	return keepRunning
}
