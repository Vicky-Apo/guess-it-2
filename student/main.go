package main

import (
	"fmt"
	"math"
	"os"
)

// Keeps track of the trend we use for prediction
var trendPoint float64
var isFirstInput bool = true

// Function to adjust the predicted trend based on the latest number
func adjustPrediction(newValue float64, lastValue float64) float64 {
	// If it's the first number, set it as the trend and return
	if isFirstInput {
		isFirstInput = false
		trendPoint = newValue
		return trendPoint
	}

	// Measure how much the value changed
	difference := math.Abs(newValue - lastValue)

	// Define how fast the trend should adjust
	var changeSpeed float64
	if difference > 20 { // Big jump → adjust quickly
		changeSpeed = 0.25
	} else if difference > 10 { // Medium change → moderate adjustment
		changeSpeed = 0.15
	} else { // Small change → slow adaptation
		changeSpeed = 0.10
	}

	// Calculate how much to shift the trend
	shiftAmount := difference * changeSpeed
	if newValue > trendPoint {
		trendPoint += shiftAmount
	} else {
		trendPoint -= shiftAmount
	}

	return trendPoint
}

func main() {
	var lastValue float64

	// Read numbers one by one
	for {
		var newNumber float64
		_, err := fmt.Fscan(os.Stdin, &newNumber)
		if err != nil {
			break
		}

		// Handle the first input separately
		if isFirstInput {
			fmt.Println("No prediction available")
			adjustPrediction(newNumber, newNumber)
			lastValue = newNumber
			continue
		}

		// Adjust prediction based on the new number
		predictedValue := adjustPrediction(newNumber, lastValue)

		// Dynamic range: Increase if recent jumps were large
		dynamicRange := 6.0 + (math.Abs(newNumber-lastValue) * 0.2)

		// Define prediction boundaries
		lowerBound := predictedValue - dynamicRange
		upperBound := predictedValue + dynamicRange

		// Output the predicted range
		fmt.Printf("%.2f %.2f\n", lowerBound, upperBound)

		// Update last known value
		lastValue = newNumber
	}
}
