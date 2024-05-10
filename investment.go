package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	investmentAmount := 1000.0
	years := 10.0
	expectedReturnRate := 5.5

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futereAdjustedInflation := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futereAdjustedInflation)
}

// func main() {
// 	investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5

// 	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

// 	fmt.Println(futureValue)
// }

// func main() {
// 	var investmentAmount float64 = 1000
// 	var expectedReturnRate float64 = 5.5
// 	var years float64 = 10

// 	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)

// 	fmt.Println(futureValue)
// }

// func main() {
// 	var investmentAmount = 1000
// 	var expectedReturnRate = 5.5
// 	var years = 10

// 	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))

// 	fmt.Println(futureValue)
// }
