package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const FILE_NAME_BALANCE string = `balance_file.txt`

var accountBalance float64 = 0

func main() {
	setAccountBalance()

	// for i := 0; i < 100; i++ {
	for {
		if !run() {
			break
		}
	}
}

func setAccountBalance() {
	balance, err := balanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------")
		// panic("Can't continue, sorry!")
	}

	accountBalance = balance
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
		balanceToFile(accountBalance)
	} else if choice == 4 {
		keepRunning = false
		fmt.Println("End!")
	} else {
		keepRunning = true
		fmt.Println("Invalid option!")
	}
	return keepRunning
}

func balanceFromFile() (float64, error) {
	data, err := os.ReadFile(FILE_NAME_BALANCE)
	if err != nil {
		return 0, errors.New("failed to read file")
	}

	balanceText := string(data)
	balance64, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 0, errors.New("failed to convert")
	}
	return balance64, nil
}

func balanceToFile(balance float64) {
	balanceByte := fmt.Sprint(balance)
	os.WriteFile(FILE_NAME_BALANCE, []byte(balanceByte), 0644)
}
