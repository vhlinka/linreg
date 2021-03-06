package matrix

import ()

//
// return sum by column and sum by row
//
//
//
func SumAll(a [][]float64) ([]float64, []float64) {

	nrows := len(a)
	ncols := len(a[0])

	// =============== sumColumns
	colsums := make([]float64, ncols)
	rowsums := make([]float64, nrows)

	for i := range a {
		for j := range a[i] {
			colsums[j] += a[i][j]
			rowsums[i] += a[i][j]
		}
	}

	return colsums, rowsums

}
