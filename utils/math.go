package utils

import (
	"math"
	"sort"
)

// The Variance function calculates the variance of a set of values given the mean.
func Variance(values []float64, mean float64) float64 {
	sum := 0.
	for _, e := range values {
		sum += (mean - e) * (mean - e)
	}

	return sum / float64(len(values))

}

func EcartType(values []float64, mean float64) float64 {
	return math.Sqrt(Variance(values, mean))

}

func NthTile(values []float64, k int, nth int) float64 {
	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})

	return values[k*(len(values)-1)/nth]

}
