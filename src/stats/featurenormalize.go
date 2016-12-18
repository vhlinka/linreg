package stats

import (
	"math"
)

//
// FeatureNormalize(a [][]float64)
// Normalizes the two dimentional slice by applying the formula:
//		(X - Xavg) / stddev(X)
//
// note: the following perform the calculations column wise - similar to the matrix calculations
// 		 in Octave - i.e.  mean(X) ...
//
//	returns the calculated mean and stdev slices
//
//
func FeatureNormalize(a [][]float64) ([]float64, []float64) {

	nrows := len(a)
	ncols := len(a[0])

	// =============== sumColumns
	stdtotal := make([]float64, ncols)
	average := make([]float64, ncols)
	stddev := make([]float64, ncols)

	for i := range average {
		average[i] = 0.0
		stdtotal[i] = 0.0
	}
	//	first calculte the mean(X)
	for i := range a {
		for j := range a[i] {
			average[j] += a[i][j]
		}
	}
	for i := range average {
		average[i] = average[i] / float64(nrows)
	}

	// --- now calculate the stddev for each column
	for i := range a {
		for j := range a[i] {
			stdtotal[j] += math.Pow(a[i][j]-average[j], 2)
		}
	}

	for i := range stdtotal {
		stddev[i] = stdtotal[i] / float64(nrows-1)
		stddev[i] = math.Sqrt(stddev[i])
	}

	// === now apply the values to normalize the data

	for i := range a {
		for j := range a[i] {
			a[i][j] = (a[i][j] - average[j]) / stddev[j]
		}
	}

	return average, stddev
}

//
//
//
func ApplyNormalizeParameters(a [][]float64, mu []float64, sigma []float64) {

	for i := range a {
		for j := range a[i] {
			a[i][j] = (a[i][j] - mu[j]) / sigma[j]
		}
	}

}
