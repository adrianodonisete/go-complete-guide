package main

import "fmt"

func main() {
	age := 32

	agePointer := &age
	fmt.Println("Age: ", *agePointer)

	calcYears := getCalcYears(age)
	fmt.Println("Calc Age: ", calcYears)

	calcYearsP := getCalcYearsPointer(agePointer)
	fmt.Println("Calc Age Pointer: ", calcYearsP)
}

func getCalcYears(age int) int {
	return age - 18
}

func getCalcYearsPointer(age *int) int {
	return *age - 18
}
