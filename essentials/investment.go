package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print(`Investment Amount: `)
	fmt.Scan(&investmentAmount)

	fmt.Print(`Expected Return Rate: `)
	fmt.Scan(&expectedReturnRate)

	fmt.Print(`Total Yeaer: `)
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futereAdjustedInflation := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value adjusted: %.2f\n", futereAdjustedInflation)

	// fmt.Println("Future value: ", futureValue)
	fmt.Printf(formattedFV, formattedRFV)

	fmt.Printf(`Future value: %.2f
	Future Value adjusted: %.2f
	`, futureValue, futereAdjustedInflation)
}

// func main() {
// 	investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5

// 	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

// 	fmt.Println(futureValue)
// }

// func main() {
// 	var investmentAmount float64 = 1000
// 	expectedReturnRate := 5.5
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
