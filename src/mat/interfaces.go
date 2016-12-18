package mat

import (
	"fmt"
)

type OperationType int
type DimensionType int

const (
	AdditionOperation OperationType = iota
	SubtractionOperation
	MultiplicationOperation
	DivisionOperation
)

const (
	BroadcastNone    = 0
	BroadcastColumns = 1
	BroadcastRows    = 2
)

const (
	ColumnDimension DimensionType = iota
	RowDimension
	RowAndColumnDimension
)

//
// Element wise  Matrix operation : a (op) b and return as a new matrix with the result
//
func ElementwiseOperation(a [][]float64, b [][]float64, op OperationType) ([][]float64, error) {

	var err error

	nrows := len(a)
	ncols := len(a[0])

	broadcast := BroadcastNone
	if nrows != len(b) || ncols != len(b[0]) {
		// check for a broadcast opportunity
		if nrows == len(b) && len(b[0]) == 1 {
			broadcast = BroadcastColumns
		} else {
			if nrows == 1 && ncols == len(b[0]) {
				broadcast = BroadcastColumns
			} else {
				err = fmt.Errorf("Number of Columns and Rows must be the same between the two matrix")
				return nil, err
			}
		}

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

	switch broadcast {
	case BroadcastRows:
		for i := 0; i < nrows; i++ {
			for j := 0; j < ncols; j++ {
				newarray[i][j] = opfn(a[i][j], b[0][j])
			}
		}
	case BroadcastColumns:
		for i := 0; i < nrows; i++ {
			for j := 0; j < ncols; j++ {
				newarray[i][j] = opfn(a[i][j], b[i][0])
			}
		}
	default:
		for i := 0; i < nrows; i++ {
			for j := 0; j < ncols; j++ {
				newarray[i][j] = opfn(a[i][j], b[i][j])
			}
		}
	}

	return newarray, err
}

//
// Element wise  Matrix operation : a (op) b and return as a new matrix with the result
//
func ScalerOperation(a [][]float64, b float64, op OperationType) ([][]float64, error) {

	var err error

	nrows := len(a)
	ncols := len(a[0])

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
			newarray[i][j] = opfn(a[i][j], b)
		}
	}
	return newarray, err
}

//
// Element wise  Matrix operation : a (op) b and return as a new matrix with the result
//
func ScalerArrayOperation(a []float64, b float64, op OperationType) ([]float64, error) {

	var err error

	nrows := len(a)

	// creat a new target matrix
	newarray := make([]float64, nrows)

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
		newarray[i] = opfn(a[i], b)
	}
	return newarray, err
}

func SumArray(a [][]float64, dim DimensionType) ([][]float64, error) {

	var err error
	nrows := len(a)
	ncols := len(a[0])

	switch dim {
	case ColumnDimension:
		nrows = 1
	case RowDimension:
		ncols = 1
	case RowAndColumnDimension:
	default:

	}
	// creat a new target matrix
	newarray := make([][]float64, nrows)
	frow := make([]float64, nrows*(ncols))
	for i := range newarray {
		newarray[i], frow = frow[:ncols], frow[ncols:]
	}

	for i := range a {
		for j := range a[i] {
			newarray[0][j] += a[i][j]
		}
	}

	return newarray, err
}

func Transpose(a [][]float64) ([][]float64, error) {

	var err error
	nrows := len(a)
	ncols := len(a[0])

	trows := ncols
	tcols := nrows

	// creat a new target matrix
	newarray := make([][]float64, trows)
	frow := make([]float64, trows*(tcols))
	for i := range newarray {
		newarray[i], frow = frow[:tcols], frow[tcols:]
	}

	for i := range a {
		for j := range a[i] {
			newarray[j][i] += a[i][j]
		}
	}

	return newarray, err

}
