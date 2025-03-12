package main

import (
	"fmt"
	"math"
	"os"
)

// Global variables for tracking prediction state
var currentTrend float64
var isFirstInput = true

// updateTrend calculates the next prediction based on new input
func updateTrend(newValue, lastValue float64) float64 {
	// Handle first input
	if isFirstInput {
		isFirstInput = false
		currentTrend = newValue
		return currentTrend
	}

	// Calculate change between values
	difference := math.Abs(newValue - lastValue)

	// Determine adjustment rate based on change size
	adjustmentRate := 0.10 // default for small changes
	if difference > 20 {
		adjustmentRate = 0.25 // large changes
	} else if difference > 10 {
		adjustmentRate = 0.15 // medium changes
	}

	// Update trend direction
	if newValue > currentTrend {
		currentTrend += difference * adjustmentRate
	} else {
		currentTrend -= difference * adjustmentRate
	}

	return currentTrend
}

func main() {
	var lastValue float64
	fmt.Println("Enter numbers (one per line, Ctrl+D to stop):")

	for {
		// Get input
		var newNumber float64
		if _, err := fmt.Fscan(os.Stdin, &newNumber); err != nil {
			fmt.Println("\nProgram finished!")
			break
		}

		// Handle first number
		if isFirstInput {
			fmt.Println("First number received. No prediction available yet.")
			updateTrend(newNumber, newNumber)
			lastValue = newNumber
			continue
		}

		// Make prediction
		predictedTrend := updateTrend(newNumber, lastValue)

		// Calculate adaptive range based on volatility
		predictionRange := 6.0 + (math.Abs(newNumber-lastValue) * 0.2)

		// Print prediction
		fmt.Printf("Next prediction: %.2f to %.2f\n",
			predictedTrend-predictionRange,
			predictedTrend+predictionRange)

		lastValue = newNumber
	}
}
