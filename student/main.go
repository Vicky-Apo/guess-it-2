package main

import (
	"fmt"
	"math"
	"os"
)

// Variables to store our prediction data
var trendPoint float64
var isFirstInput bool = true

func adjustPrediction(newValue float64, lastValue float64) float64 {
	// First number becomes our starting point
	if isFirstInput {
		isFirstInput = false
		trendPoint = newValue
		return trendPoint
	}

	// Find how much the numbers changed
	change := math.Abs(newValue - lastValue)

	// Pick speed of adjustment
	var changeSpeed float64
	if change > 20 {
		changeSpeed = 0.25 // Fast speed for big changes
	} else if change > 10 {
		changeSpeed = 0.15 // Medium speed for medium changes
	} else {
		changeSpeed = 0.10 // Slow speed for small changes
	}

	// Calculate and apply the adjustment
	adjustment := change * changeSpeed
	if newValue > trendPoint {
		trendPoint += adjustment // Move up
	} else {
		trendPoint -= adjustment // Move down
	}

	return trendPoint
}

func main() {
	var lastValue float64

	// Read numbers until input ends
	for {
		var newNumber float64
		_, err := fmt.Fscan(os.Stdin, &newNumber)
		if err != nil {
			break
		}

		// Skip prediction for first number
		if isFirstInput {
			adjustPrediction(newNumber, newNumber)
			lastValue = newNumber
			continue
		}

		// Get next prediction
		predictedValue := adjustPrediction(newNumber, lastValue)

		// Create range of ±6 around prediction
		lowerBound := predictedValue - 6
		upperBound := predictedValue + 6

		// Show prediction range
		fmt.Printf("%.2f %.2f\n", lowerBound, upperBound)

		lastValue = newNumber
	}
}
