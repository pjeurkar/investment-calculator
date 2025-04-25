package main

import (
	"fmt" // imported math package to use for the power of number in the formula for futureValue.
	"math"
)

func main(){
	/*
	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10

	//var futureValue = investmentAmount * (1 + expectedReturnRate / 100)
	// got error saying mismatched typed int and floadt64! 
	var futureValue = float64(investmentAmount) * math.Pow(1 + expectedReturnRate / 100, float64(years))
	*/
	// type conversion to make the above formula more readable. ⬇️

	/*------------------------------------------------------------
	var investmentAmount float64 = 1000		// explicitly set the int to float64
	var expectedReturnRate = 5.5
	var years float64 = 10

	var futureValue = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)

	fmt.Println(futureValue)		// to get the output
	------------------------------------------------------------*/
	/*
	// now to make the code more concise, doing things according to Go ways.

	//var investmentAmount, years float64 = 1000, 10	// explicit way
	investmentAmount, years := 1000.0, 10.0		// implicit way
	expectedReturnRate := 5.5		// := Short assignment operator is a shortcut to assign variables acc. to Go
	// we can do all the above in 1 line too! ⬇️
	// investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	------------------------------------------------------------*/

	const inflationRate = 6.5		// can't ever change value in const.
	investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5
	/*
	// here we can just write it like this if we want, without assigning any values since we are asking for user input.

	var invinvestmentAmount float64		
	var years float64
	var expectedReturnRate float64
	*/
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)		// for input (we need to pass the pointer to Scan for it to work using (&) )
	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Years: ")
	fmt.Scan(&years)
			

	futureValue := investmentAmount * math.Pow(1+ expectedReturnRate/ 100, years)

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)		// to get the output
	fmt.Println(futureRealValue)	

}