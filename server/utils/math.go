package utils

import (
	"math"
	"math/rand"
)

func Ceil(num float64, n int) float64 {
	pow := math.Pow(10, float64(n))
	return math.Ceil(num*pow) / pow
}

func WeightedRandomIndex(weights []float64) (int, float64) {
	// Step 1: Calculate the sum of all weights
	var sum float64
	for _, weight := range weights {
		sum += weight
	}

	// Step 2: Normalize the weights to get probabilities
	probabilities := make([]float64, len(weights))
	for i, weight := range weights {
		probabilities[i] = weight / sum
	}

	// Step 3: Generate a random number between 0 and 1
	randomValue := rand.Float64()

	// Step 4: Determine which weight the random value corresponds to
	var cumulativeProbability float64
	for i, probability := range probabilities {
		cumulativeProbability += probability
		if randomValue < cumulativeProbability {
			return i, probabilities[i]
		}
	}

	// Should never reach here if the weights are properly normalized
	return len(weights) - 1, probabilities[len(weights)-1]
}
