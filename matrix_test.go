package matrix

import "testing"

func TestSize(t *testing.T) {
	matrices := []struct {
		matrix Matrix
		size   int
	}{
		{Matrix{m: 0, n: 0}, 0},
		{Matrix{m: 1, n: 0}, 0},
		{Matrix{m: 0, n: 1}, 0},
		{Matrix{m: 1, n: 1}, 1},
		{Matrix{m: 1, n: 2}, 2},
		{Matrix{m: 2, n: 1}, 2},
		{Matrix{m: 2, n: 2}, 4},
	}

	for _, matrixData := range matrices {
		if matrixData.matrix.Size() != matrixData.size {
			t.Errorf("matrix %v should have size %d, but has %d", matrixData.matrix, matrixData.size, matrixData.matrix.Size())
		}
	}
}

func TestEquals(t *testing.T) {
	testData := []struct {
		matrix      *Matrix
		matrixRight *Matrix
		eq          bool
	}{
		{NewMatrix(0, 0, []float64{}), NewMatrix(0, 0, []float64{}), true},
		{NewMatrix(0, 1, []float64{}), NewMatrix(0, 0, []float64{}), false},
	}

	for _, test := range testData {
		if test.matrix.Equals(test.matrixRight) != test.eq {
			if test.eq {
				t.Errorf("matrix \n%v should Equals \n%v", test.matrix, test.matrixRight)
			} else {
				t.Errorf("matrix \n%v should not Equals \n%v", test.matrix, test.matrixRight)
			}
		}
	}
}
func TestAdd(t *testing.T) {
	testData := []struct {
		matrix      *Matrix
		matrixRight *Matrix
		add         *Matrix
	}{
		{NewMatrix(0, 0, []float64{}), NewMatrix(0, 0, []float64{}), NewMatrix(0, 0, []float64{})},
		{NewMatrix(1, 1, []float64{1}), NewMatrix(1, 1, []float64{1}), NewMatrix(1, 1, []float64{2})},
		{NewMatrix(1, 2, []float64{1, 2}), NewMatrix(1, 2, []float64{1, 2}), NewMatrix(1, 2, []float64{2, 4})},
	}

	for _, test := range testData {
		add := test.matrix.Add(test.matrixRight)

		if !add.Equals(test.add) {
			t.Errorf("matrix \n%v should Equals \n%v", add, test.add)
		}
	}
}

func TestScalarMultiplication(t *testing.T) {
	testData := []struct {
		matrix *Matrix
		scalar float64
		result *Matrix
	}{
		{NewMatrix(2, 2, []float64{1, 1, 1, 1}), 0, NewMatrix(2, 2, []float64{0, 0, 0, 0})},
		{NewMatrix(2, 2, []float64{1, 1, 1, 1}), 1, NewMatrix(2, 2, []float64{1, 1, 1, 1})},
		{NewMatrix(2, 2, []float64{1, 1, 1, 1}), 2, NewMatrix(2, 2, []float64{2, 2, 2, 2})},
	}

	for _, test := range testData {
		rtest := test.matrix.ScalarMultiplication(test.scalar)

		if !rtest.Equals(test.result) {
			t.Errorf("matrix \n%v should Equals \n%v", rtest, test.result)
		}
	}
}

func TestTranspose(t *testing.T) {
	testData := []struct {
		matrix *Matrix
		result *Matrix
	}{
		{NewMatrix(2, 2, []float64{1, 1, 1, 1}), NewMatrix(2, 2, []float64{1, 1, 1, 1})},
		{NewMatrix(2, 2, []float64{1, 1, 2, 2}), NewMatrix(2, 2, []float64{1, 2, 1, 2})},
		{NewMatrix(2, 2, []float64{1, 2, 3, 4}), NewMatrix(2, 2, []float64{1, 3, 2, 4})},
	}

	for _, test := range testData {
		tr := test.matrix.Transpose()
		if !tr.Equals(test.result) {
			t.Errorf("result matrix %v should equals transposed %v", test.result, tr)
		}
	}
}

func TestMultiplication(t *testing.T) {
	testData := []struct {
		left   *Matrix
		right  *Matrix
		result *Matrix
	}{
		{NewMatrix(2, 3, []float64{2, 3, 4, 1, 0, 0}), NewMatrix(3, 2, []float64{0, 1000, 1, 100, 0, 10}), NewMatrix(2, 2, []float64{3, 2340, 0, 1000})},
	}

	for _, test := range testData {
		result := test.left.Multiplication(test.right)
		if !result.Equals(test.result) {
			t.Errorf("result matrix %v should equals %v", test.result, result)
		}
	}
}

func TestNewZeroMatrix(t *testing.T) {
	testData := []struct {
		m int
		n int
	}{
		{0, 0},
		{1, 0},
		{2, 2},
		{100, 100},
	}

	for _, test := range testData {
		matrix := NewZeroMatrix(test.m, test.n)

		for i := 0; i < test.m; i++ {
			for j := 0; j < test.n; j++ {
				if matrix.data[i][j] != 0.0 {
					t.Errorf("matrix should be empty")
				}
			}
		}
	}
}

func TestDimensions(t *testing.T) {
	testData := []struct {
		m int
		n int
	}{
		{0, 0},
		{1, 0},
		{2, 2},
		{100, 100},
	}

	for _, test := range testData {
		matrix := NewZeroMatrix(test.m, test.n)

		rM, rN := matrix.Dimensions()

		if rM != test.m || rN != test.n {
			t.Errorf("incorect returned dimentions")
		}
	}
}

func TestGetRow(t *testing.T) {
	testData := []struct {
		matrix *Matrix
		i      int
		row    *Matrix
	}{
		{NewMatrix(2, 2, []float64{2, 3, 4, 1}), 1, NewMatrix(1, 2, []float64{4, 1})},
	}

	for _, test := range testData {
		result := test.matrix.GetRow(test.i)
		if !result.Equals(test.row) {
			t.Errorf("result matrix %v should equals %v", test.row, result)
		}
	}
}

func TestGetColumn(t *testing.T) {
	testData := []struct {
		matrix *Matrix
		i      int
		row    *Matrix
	}{
		{NewMatrix(2, 2, []float64{2, 3, 4, 1}), 1, NewMatrix(2, 1, []float64{3, 1})},
	}

	for _, test := range testData {
		result := test.matrix.GetColumn(test.i)
		if !result.Equals(test.row) {
			t.Errorf("result matrix %v should equals %v", test.row, result)
		}
	}
}
