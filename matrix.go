// Package matrix provides simple matrix struct with basic matrix operations
package matrix

// Matrix represents base matrix struct
type Matrix struct {
	//rows num
	m int

	//columns num
	n int

	//matrix data
	data [][]float64
}

// Dimensions returns dimensions of matrix
func (matrix *Matrix) Dimensions() (int, int) {
	return matrix.m, matrix.n
}

// Size returns matrix size
func (matrix *Matrix) Size() int {
	return matrix.m * matrix.n
}

// Equals compares two matrices
func (matrix *Matrix) Equals(matrixRight *Matrix) bool {
	if matrix.m != matrixRight.m || matrix.n != matrixRight.n {
		return false
	}

	for i := 0; i < matrix.m; i++ {
		for j := 0; j < matrix.n; j++ {
			if matrix.data[i][j] != matrixRight.data[i][j] {
				return false
			}
		}
	}

	return true
}

//Add returns (matrix + matrixRight)
func (matrix *Matrix) Add(matrixRight *Matrix) *Matrix {
	if matrix.m != matrixRight.m || matrix.n != matrixRight.n {
		panic("matrix should have same sizes")
	}

	return matrix.Clone().Map(func(val float64, i int, j int) float64 {
		return val + matrixRight.data[i][j]
	})
}

//At returns ceil value
func (matrix *Matrix) At(i int, j int) float64 {
	if i >= matrix.m || j >= matrix.n {
		panic("out of bounds")
	}

	return matrix.data[i][j]
}

//Multiplication  returns (matrix · matrixRight)
func (matrix *Matrix) Multiplication(matrixRight *Matrix) *Matrix {
	if matrix.n != matrixRight.m {
		panic("the number of columns of the left matrix should be the same as the number of rows of the right matrix")
	}

	resultMatrix := NewZeroMatrix(matrix.m, matrixRight.n)

	resultMatrix.Map(func(val float64, i int, j int) float64 {
		summ := 0.0
		for n := 0; n < matrix.n; n++ {
			summ += matrix.data[i][n] * matrixRight.data[n][j]
		}

		return summ
	})

	return resultMatrix
}

//ScalarMultiplication returns (scalar · matrix)
func (matrix *Matrix) ScalarMultiplication(scalar float64) *Matrix {
	return matrix.Clone().Map(func(val float64, i int, j int) float64 {
		return scalar * val
	})
}

// Transpose returns transposed matrix
func (matrix *Matrix) Transpose() *Matrix {
	transposed := NewZeroMatrix(matrix.n, matrix.m)
	transposed.Map(func(val float64, i int, j int) float64 {
		return matrix.data[j][i]
	})

	return transposed
}

// Map applies the callback to the elements of the given matrix
func (matrix *Matrix) Map(callback func(val float64, i int, j int) float64) *Matrix {
	for i := 0; i < matrix.m; i++ {
		for j := 0; j < matrix.n; j++ {
			matrix.data[i][j] = callback(matrix.data[i][j], i, j)
		}
	}

	return matrix
}

// Walk through matrix
func (matrix *Matrix) Walk(callback func(val float64, i int, j int)) *Matrix {
	for i := 0; i < matrix.m; i++ {
		for j := 0; j < matrix.n; j++ {
			callback(matrix.data[i][j], i, j)
		}
	}

	return matrix
}

//Clone returns copy of matrix
func (matrix *Matrix) Clone() *Matrix {
	rawData := make([][]float64, matrix.m)

	for i := range rawData {
		rawData[i] = append([]float64(nil), matrix.data[i]...)
	}

	return &Matrix{m: matrix.m, n: matrix.n, data: rawData}
}

// ZeroClone clones zeroed matrix
func (matrix *Matrix) ZeroClone() *Matrix {
	return NewZeroMatrix(matrix.m, matrix.n)
}

// GetRow returns Row vector
func (matrix *Matrix) GetRow(i int) *Matrix {
	if i >= matrix.m {
		panic("i is out of bounds")
	}

	rawData := append([]float64(nil), matrix.data[i]...)

	return NewMatrix(1, matrix.n, rawData)
}

// GetColumn returns Column vector
func (matrix *Matrix) GetColumn(j int) *Matrix {
	if j >= matrix.n {
		panic("j is out of bounds")
	}

	rawData := make([]float64, matrix.m)

	for i := 0; i < matrix.m; i++ {
		rawData[i] = matrix.data[i][j]
	}

	return NewMatrix(matrix.m, 1, rawData)
}

// NewMatrix creates new Matrix struct
func NewMatrix(m int, n int, rawData []float64) *Matrix {
	if m*n != len(rawData) {
		panic("rawData data size")
	}

	data := make([][]float64, m)

	for i := 0; i < m; i++ {
		row := make([]float64, n)
		for j := 0; j < n; j++ {
			row[j] = rawData[i*n+j]
		}

		data[i] = row
	}

	return &Matrix{m: m, n: n, data: data}
}

//NewZeroMatrix creates new zero Matrix
func NewZeroMatrix(m int, n int) *Matrix {
	rawData := make([]float64, m*n)

	return NewMatrix(m, n, rawData)
}
