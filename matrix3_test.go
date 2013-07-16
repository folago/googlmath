package math

import (
	. "launchpad.net/gocheck"
)

type MulMatrix3TestValue struct {
	Matrix   *Matrix3
	Matrix2  *Matrix3
	Expected *Matrix3
}

type DeterminantMatrix3TestValue struct {
	Matrix   *Matrix3
	Expected float32
}

type Matrix3TestSuite struct {
	mulTestTable         []MulMatrix3TestValue
	determinantTestTable []DeterminantMatrix3TestValue
}

var matrixTest3Suite = Suite(&Matrix3TestSuite{})

func (test *Matrix3TestSuite) SetUpTest(c *C) {
	test.mulTestTable = []MulMatrix3TestValue{
		MulMatrix3TestValue{
			&Matrix3{11.0, 0.0, 3.0, -3.0, 2.5, 0.0, 2.2, 0.3, 12.0},
			&Matrix3{1.0, 0.0, 0.0, 0.0, 2.5, 0.0, 2.2, 0.2, 1.0},
			&Matrix3{11.0, 0.0, 3.0, -7.5, 6.25, 0.0, 25.800001, 0.8, 18.6},
		},
		MulMatrix3TestValue{
			&Matrix3{11.0, 0.34, -3.0, -3.0, 2.5, 0.0, 32.200001, 3.32, 12.0},
			&Matrix3{1.0, 0.0, 3.0, 5.0, -1.5, -2.0, -2.2, 0.2, 2.0},
			&Matrix3{107.600006, 10.3, 33.0, -4.900002, -8.69, -39.0, 39.599998, 6.392, 30.6},
		},
	}

	test.determinantTestTable = []DeterminantMatrix3TestValue{
		DeterminantMatrix3TestValue{
			&Matrix3{11.0, 0.34, -3.0, -3.0, 2.5, 0.0, 32.200001, 3.32, 12.0},
			613.619995,
		},
		DeterminantMatrix3TestValue{
			&Matrix3{1.0, 0.0, 3.0, 5.0, -1.5, -2.0, -2.2, 0.2, 2.0},
			-9.5,
		},
	}
}

func (test *Matrix3TestSuite) TestNewIdentityMatrix3(c *C) {
	expected := &Matrix3{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}
	obtained := NewIdentityMatrix3()
	c.Check(obtained, Matrix3Check, expected)
}

func (test *Matrix3TestSuite) TestMul(c *C) {
	for _, value := range test.mulTestTable {
		matrix := value.Matrix.Mul(value.Matrix2)
		c.Check(value.Matrix, Not(Matrix3Check), value.Matrix2)
		c.Check(value.Matrix, Not(Matrix3Check), matrix)
		c.Check(value.Matrix2, Not(Matrix3Check), matrix)
		c.Check(matrix, Matrix3Check, value.Expected)

	}
}

func (test *Matrix3TestSuite) TestDeterminant(c *C) {
	for _, value := range test.determinantTestTable {
		det := value.Matrix.Determinant()
		c.Check(det, EqualsFloat32, value.Expected)
	}
}
