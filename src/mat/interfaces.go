package mat

import (
	"fmt"
)

type OperationType int

const (
	AdditionOperation OperationType = iota
	SubtractionOperation
	MultiplicationOperation
	DivisionOperation
)

//
// Subtract Matrix a from b and return as a new matrix
//
func ElementwiseOperation(a [][]float64, b [][]float64, op OperationType) ([][]float64, error) {

	var err error

	nrows := len(a)
	ncols := len(a[0])

	if nrows != len(b) || ncols != len(b[0]) {
		err = fmt.Errorf("Number of Columns and Rows must be the same between the two matrix")
		return nil, err
	}

	// creat a new target matrix
	newarray := make([][]float64, nrows)
	frow := make([]float64, nrows*(ncols))
	for i := range newarray {
		newarray[i], frow = frow[:ncols], frow[ncols:]
	}

	opfn := func(x, y float64) float64 { return 0 }
	// determine which operation to apply to each element
	switch op {
	case AdditionOperation:
		opfn = func(x, y float64) float64 {
			return x + y
		}
	case SubtractionOperation:
		opfn = func(x, y float64) float64 {
			return x - y
		}
	case MultiplicationOperation:
		opfn = func(x, y float64) float64 {
			return x * y
		}
	case DivisionOperation:
		opfn = func(x, y float64) float64 {
			return x / y
		}
	default:
		err = fmt.Errorf("Unsupported Operation Specified")
		return nil, err
	}

	for i := 0; i < nrows; i++ {
		for j := 0; j < ncols; j++ {
			newarray[i][j] = opfn(a[i][j], b[i][j])
		}
	}

	return newarray, err
}
