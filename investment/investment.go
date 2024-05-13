package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	outputText(`Investment Amount: `)
	fmt.Scan(&investmentAmount)

	outputText(`Expected Return Rate: `)
	fmt.Scan(&expectedReturnRate)

	outputText(`Total Years: `)
	fmt.Scan(&years)

	futureValue, futereAdjustedInflation := calculateFutureValue(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value adjusted: %.2f\n", futereAdjustedInflation)

	// fmt.Println("Future value: ", futureValue)
	fmt.Print(formattedFV, formattedRFV)

	fmt.Printf(`Future value: %.2f
	Future Value adjusted: %.2f
	`, futureValue, futereAdjustedInflation)

	futureValue, futereAdjustedInflation = calculateFutureValueAlt(investmentAmount, expectedReturnRate, years)
	fmt.Print(futureValue, futereAdjustedInflation)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fvi := fv / math.Pow(1+inflationRate/100, years)
	return fv, fvi
}

func calculateFutureValueAlt(investmentAmount, expectedReturnRate, years float64) (fv float64, fvi float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fvi = fv / math.Pow(1+inflationRate/100, years)
	return
}
