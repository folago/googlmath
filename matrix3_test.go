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

type Matrix3TestValue struct {
	Matrix   *Matrix3
	Expected *Matrix3
}

type ArrayMatrix3TestValue struct {
	Matrix   *Matrix3
	Expected []float32
}

type FloatMatrix3TestValue struct {
	Value    float32
	Expected *Matrix3
}

type Matrix3TestSuite struct {
	newXRotationTestTable []FloatMatrix3TestValue
	newYRotationTestTable []FloatMatrix3TestValue
	newZRotationTestTable []FloatMatrix3TestValue
	mulTestTable          []MulMatrix3TestValue
	determinantTestTable  []DeterminantMatrix3TestValue
	inverseTestTable      []Matrix3TestValue
	toArrayTestTable      []ArrayMatrix3TestValue
	transposeTestTable    []Matrix3TestValue
}

var matrixTest3Suite = Suite(&Matrix3TestSuite{})

func (test *Matrix3TestSuite) SetUpTest(c *C) {
	test.newXRotationTestTable = []FloatMatrix3TestValue{
		FloatMatrix3TestValue{
			30,
			&Matrix3{1.0, 0, 0, 0, 0.866, -0.5, 0, 0.5, 0.866},
		},
		FloatMatrix3TestValue{
			-99,
			&Matrix3{1.0, 0, 0, 0, -Sin(Pi / 20), Cos(Pi / 20), 0, -Cos(Pi / 20), -Sin(Pi / 20)},
		},
	}

	test.newYRotationTestTable = []FloatMatrix3TestValue{
		FloatMatrix3TestValue{
			-99,
			&Matrix3{-Sin(Pi / 20), 0, -Cos(Pi / 20), 0, 1, 0, Cos(Pi / 20), 0, -Sin(Pi / 20)},
		},
		FloatMatrix3TestValue{
			11,
			&Matrix3{Cos(11 * Pi / 180), 0, Sin(11 * Pi / 180), 0, 1, 0, -Sin(11 * Pi / 180), 0, Cos(11 * Pi / 180)},
		},
	}

	test.newZRotationTestTable = []FloatMatrix3TestValue{
		FloatMatrix3TestValue{
			-99,
			&Matrix3{-Sin(Pi / 20), Cos(Pi / 20), 0, -Cos(Pi / 20), -Sin(Pi / 20), 0, 0, 0, 1},
		},
		FloatMatrix3TestValue{
			11,
			&Matrix3{Cos(11 * Pi / 180), -Sin(11 * Pi / 180), 0, Sin(11 * Pi / 180), Cos(11 * Pi / 180), 0, 0, 0, 1},
		},
	}

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

	test.inverseTestTable = []Matrix3TestValue{
		Matrix3TestValue{
			&Matrix3{11.0, 0.0, 3.0, -3.0, 2.5, 0.0, 2.2, 0.3, 12.0},
			&Matrix3{0.096525, 0.002896, -0.024131, 0.11583, 0.403475, -0.028958, -0.020592, -0.010618, 0.088481},
		},
	}

	test.toArrayTestTable = []ArrayMatrix3TestValue{
		ArrayMatrix3TestValue{
			&Matrix3{1, 0, 0, 0, 1, 0, 0, 0, 1},
			[]float32{1, 0, 0, 0, 1, 0, 0, 0, 1},
		},
		ArrayMatrix3TestValue{
			&Matrix3{1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]float32{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	test.transposeTestTable = []Matrix3TestValue{
		Matrix3TestValue{
			&Matrix3{11.0, 0.0, 3.0, -3.0, 2.5, 0.0, 2.2, 0.3, 12.0},
			&Matrix3{11.0, -3.0, 2.2, 0.0, 2.5, 0.3, 3.0, 0.0, 12.0},
		},
	}
}

func (test *Matrix3TestSuite) TestNewMatrix3(c *C) {
	var m11, m12, m13, m21, m22, m23, m31, m32, m33 float32 = 1, 2, 3, 4, 5, 6, 7, 8, 9
	m := NewMatrix3(m11, m12, m13, m21, m22, m23, m31, m32, m33)
	m2 := &Matrix3{m11, m12, m13, m21, m22, m23, m31, m32, m33}
	c.Check(m, Matrix3Check, m2)
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

func (test *Matrix3TestSuite) TestNewXRotationMatrix3(c *C) {
	for _, value := range test.newXRotationTestTable {
		matrix := NewXRotationMatrix3(value.Value)
		c.Check(matrix, Matrix3Check, value.Expected)
	}
}

func (test *Matrix3TestSuite) TestNewYRotationMatrix3(c *C) {
	for _, value := range test.newYRotationTestTable {
		matrix := NewYRotationMatrix3(value.Value)
		c.Check(matrix, Matrix3Check, value.Expected)
	}
}

func (test *Matrix3TestSuite) TestNewZRotationMatrix3(c *C) {
	for _, value := range test.newZRotationTestTable {
		matrix := NewZRotationMatrix3(value.Value)
		c.Check(matrix, Matrix3Check, value.Expected)
	}
}

// TODO Test NewRotationMatrix3
// TODO Test NewTranslationMatrix3
// TODO Test NewScaleMatrix3

func (test *Matrix3TestSuite) TestSet(c *C) {
	m := &Matrix3{}
	m2 := &Matrix3{0, 1, 2, 3, 4, 5, 6, 7, 8}
	m.Set(m2)
	c.Check(m, Matrix3Check, m2)
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

func (test *Matrix3TestSuite) TestInverse(c *C) {
	for _, value := range test.inverseTestTable {
		inv, err := value.Matrix.Inverse()
		c.Check(err, IsNil)
		c.Check(inv, Matrix3Check, value.Expected)
	}
}

func (test *Matrix3TestSuite) TestToArray(c *C) {
	for _, value := range test.toArrayTestTable {
		a := value.Matrix.ToArray()
		c.Check(a, DeepEquals, value.Expected)
	}
}

func (test *Matrix3TestSuite) TestTranspose(c *C) {
	for _, value := range test.transposeTestTable {
		m := value.Matrix.Transpose()
		c.Check(m, DeepEquals, value.Expected)
	}
}

// TODO Test Proj2D
// TODO Test ShearX2D
// TODO Test ShearY2D
