package stats

import (
	"math"
)

func StdDev(numbers []float64, mean float64) float64 {
	total := 0.0
	for _, number := range numbers {
		total += math.Pow(number-mean, 2)
	}
	variance := total / float64(len(numbers)-1)
	return math.Sqrt(variance)
}

func Mean(numbers []float64) float64 {
	average := 0.0

	for _, x := range numbers {
		average += x
	}
	average = average / float64(len(numbers))
	return average
}
