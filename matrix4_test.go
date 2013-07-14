package math

import (
	. "launchpad.net/gocheck"
)

type MatrixPerspectiveTestValue struct {
	Fov         float32
	AspectRatio float32
	Near        float32
	Far         float32
	Expected    *Matrix4
}

type MatrixLookAtTestValue struct {
	Eye, Center, Up Vector3
	Expected        *Matrix4
}

type MatrixTranslateTestValue struct {
	Translation Vector3
	Expected    *Matrix4
}

type MatrixRotationTestValue struct {
	Axis     Vector3
	Angle    float32
	Expected *Matrix4
}

type MatrixOrthoTestValue struct {
	Left, Right, Bottom, Top, Near, Far float32
	Expected                            *Matrix4
}

type MatrixMulTestValue struct {
	M1       *Matrix4
	M2       *Matrix4
	Expected *Matrix4
}

type MatrixSetTestValue struct {
	M1       *Matrix4
	M2       *Matrix4
	Expected *Matrix4
}

type MatrixScaleTestValue struct {
	Scalar   Vector3
	Matrix   *Matrix4
	Expected *Matrix4
}

type MatrixInvertTestValue struct {
	Matrix   *Matrix4
	Expected *Matrix4
}

type MatrixDeterminantTestValue struct {
	Matrix   *Matrix4
	Expected float32
}

type MatrixTestSuite struct {
	perspectiveTestTable []MatrixPerspectiveTestValue
	lookAtTestTable      []MatrixLookAtTestValue
	translateTestTable   []MatrixTranslateTestValue
	rotationTestTable    []MatrixRotationTestValue
	orthoTestTable       []MatrixOrthoTestValue
	mulTestTable         []MatrixMulTestValue
	setTestTable         []MatrixSetTestValue
	scaleTestTable       []MatrixScaleTestValue
	invertTestTable      []MatrixInvertTestValue
	determinantTestTable []MatrixDeterminantTestValue
}

var matrixTestSuite = Suite(&MatrixTestSuite{})

func (test *MatrixTestSuite) SetUpTest(c *C) {
	test.perspectiveTestTable = []MatrixPerspectiveTestValue{
		MatrixPerspectiveTestValue{45.0, 4.0 / 3.0, 0.1, 100.0, &Matrix4{1.810660, 0.0, 0.0, 0.0, 0.0, 2.4142134, 0.0, 0.0, 0.0, 0.0, -1.002002, -1.0, 0.0, 0.0, -0.2002002, 0.0}},
		MatrixPerspectiveTestValue{90.0, 16.0 / 9.0, -1.0, 1.0, &Matrix4{0.562500, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, -0.0, -1.0, 0.0, 0.0, 1.0, 0.0}},
	}

	test.lookAtTestTable = []MatrixLookAtTestValue{
		MatrixLookAtTestValue{Vec3(4, 3, 3), Vec3(0, 0, 0), Vec3(0, 1, 0), &Matrix4{0.600000, -0.411597, 0.685994, 0.0, 0.0, 0.857493, 0.514496, 0.0, -0.800000, -0.308697, 0.514496, 0.0, 0.0, 0.0, -5.830953, 1.0}},
	}

	test.translateTestTable = []MatrixTranslateTestValue{
		MatrixTranslateTestValue{Vec3(0, 0, 15), &Matrix4{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 15.0, 1.0}},
		MatrixTranslateTestValue{Vec3(-3.0, 2.2, 15), &Matrix4{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, -3.0, 2.200000, 15.0, 1.0}},
	}

	test.rotationTestTable = []MatrixRotationTestValue{
		MatrixRotationTestValue{Vec3(0, 2.5, 0), 25.0, &Matrix4{0.906308, 0.0, -0.422618, 0.0, 0.0, 1.0, 0.0, 0.0, 0.422618, 0.0, 0.906308, 0.0, 0.0, 0.0, 0.0, 1.0}},
		MatrixRotationTestValue{Vec3(2.0, 2.5, 0.0), -45.0, &Matrix4{0.821407, 0.142875, 0.552158, 0.0, 0.142875, 0.885700, -0.441726, 0.0, -0.552158, 0.441726, 0.707107, 0.0, 0.0, 0.0, 0.0, 1.0}},
	}

	test.orthoTestTable = []MatrixOrthoTestValue{
		MatrixOrthoTestValue{-10.0, 10.0, -10.0, 10.0, 0.0, 100.0, &Matrix4{.1, 0.0, 0.0, 0.0, 0.0, 0.1, 0.0, 0.0, 0.0, 0.0, -0.02, 0.0, -0.0, -0.0, -1.0, 1.0}},
		MatrixOrthoTestValue{0.0, 10.0, 0.0, 10.0, 0.0, 100.0, &Matrix4{0.2, 0.0, 0.0, 0.0, 0.0, 0.2, 0.0, 0.0, 0.0, 0.0, -0.02, 0.0, -1.0, -1.0, -1.0, 1.0}},
	}

	test.mulTestTable = []MatrixMulTestValue{
		MatrixMulTestValue{
			M1:       &Matrix4{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0},
			M2:       &Matrix4{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, -3.0, 2.200000, 15.0, 1.0},
			Expected: &Matrix4{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, -3.0, 2.200000, 15.0, 1.0},
		},
	}

	test.setTestTable = []MatrixSetTestValue{
		MatrixSetTestValue{
			M1:       &Matrix4{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0},
			M2:       &Matrix4{-1.0, 2.2, 0.33, 0.22, 0.2, 1.5, 0.1, 0.3, 4.0, 0.2, 1.1, -2.0, 0.1, 2.2, 32.0, 0.0},
			Expected: &Matrix4{-1.0, 2.2, 0.33, 0.22, 0.2, 1.5, 0.1, 0.3, 4.0, 0.2, 1.1, -2.0, 0.1, 2.2, 32.0, 0.0},
		},
	}

	test.scaleTestTable = []MatrixScaleTestValue{
		MatrixScaleTestValue{Vec3(2.0, 3.3, -2.2), NewIdentityMatrix4(), &Matrix4{2.0, 0.0, 0.0, 0.0, 0.0, 3.3, 0.0, 0.0, -0.0, -0.0, -2.2, -0.0, 0.0, 0.0, 0.0, 1.0}},
		MatrixScaleTestValue{Vec3(2.0, 3.3, -2.2), &Matrix4{0.2, 0.0, 0.0, 0.0, 0.0, 0.2, 0.0, 0.0, 0.0, 0.0, -0.02, 0.0, -1.0, -1.0, -1.0, 1.0}, &Matrix4{0.4, 0.0, 0.0, 0.0, 0.0, 0.66, 0.0, 0.0, -0.0, -0.0, 0.044, -0.0, -1.0, -1.0, -1.0, 1.0}},
	}

	test.invertTestTable = []MatrixInvertTestValue{
		MatrixInvertTestValue{NewIdentityMatrix4(), NewIdentityMatrix4()},
		MatrixInvertTestValue{&Matrix4{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, -3.0, 2.2, 15.0, 1.0}, &Matrix4{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 3.0, -2.2, -15.0, 1.0}},
	}

	test.determinantTestTable = []MatrixDeterminantTestValue{
		MatrixDeterminantTestValue{&Matrix4{0.2, 0.0, 0.0, 0.0, 0.0, 0.2, 0.0, 0.0, 0.0, 0.0, -0.02, 0.0, -1.0, -1.0, -1.0, 1.0}, -0.0008},
		MatrixDeterminantTestValue{NewIdentityMatrix4(), 1.0},
		MatrixDeterminantTestValue{&Matrix4{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, -3.0, 2.2, 15.0, 1.0}, 1.0},
	}
}

func (test *MatrixTestSuite) TestMatrixPerspective(c *C) {
	for i := range test.perspectiveTestTable {
		value := test.perspectiveTestTable[i]
		matrix := NewPerspectiveMatrix4(value.Fov, value.AspectRatio, value.Near, value.Far)
		c.Check(matrix, Matrix4Check, value.Expected)
	}
}

func (test *MatrixTestSuite) TestMatrixLookAt(c *C) {
	for i := range test.lookAtTestTable {
		value := test.lookAtTestTable[i]
		matrix := NewLookAtMatrix4(value.Eye, value.Center, value.Up)
		c.Check(matrix, Matrix4Check, value.Expected)
	}
}

func (test *MatrixTestSuite) TestMatrixTranslation(c *C) {
	for i := range test.translateTestTable {
		value := test.translateTestTable[i]
		translation := value.Translation
		matrix := NewTranslationMatrix4(translation.X, translation.Y, translation.Z)
		c.Check(matrix, Matrix4Check, value.Expected)
	}
}

func (test *MatrixTestSuite) TestMatrixRotation(c *C) {
	for i := range test.rotationTestTable {
		value := test.rotationTestTable[i]
		matrix := NewRotationMatrix4(value.Axis, value.Angle)
		c.Check(matrix, Matrix4Check, value.Expected)
	}
}

func (test *MatrixTestSuite) TestMatrixOrtho(c *C) {
	for i := range test.orthoTestTable {
		value := test.orthoTestTable[i]
		matrix := NewOrthoMatrix4(value.Left, value.Right, value.Bottom, value.Top, value.Near, value.Far)
		c.Check(matrix, Matrix4Check, value.Expected)
	}
}

func (test *MatrixTestSuite) TestMatrixMul(c *C) {
	for i := range test.mulTestTable {
		value := test.mulTestTable[i]
		matrix := value.M1.Mul(value.M2)
		c.Check(matrix, Matrix4Check, value.Expected)
	}
}

func (test *MatrixTestSuite) TestMatrixSet(c *C) {
	for i := range test.setTestTable {
		value := test.setTestTable[i]
		matrix := value.M1.Set(value.M2)
		c.Check(matrix, Matrix4Check, value.Expected)
		c.Check(matrix, Matrix4Check, value.M1)
		c.Check(matrix, Matrix4Check, value.M2)
		c.Check(value.M1, Matrix4Check, value.M2)
	}
}

func (test *MatrixTestSuite) TestMatrixScale(c *C) {
	for i := range test.scaleTestTable {
		value := test.scaleTestTable[i]
		matrix := value.Matrix.Scale(value.Scalar)
		c.Check(matrix, Matrix4Check, value.Expected)
	}
}

func (test *MatrixTestSuite) TestMatrixInvert(c *C) {
	for i := range test.invertTestTable {
		value := test.invertTestTable[i]
		matrix, err := value.Matrix.Invert()
		c.Check(err, IsNil)
		c.Check(matrix, Matrix4Check, value.Expected)
	}
}

func (test *MatrixTestSuite) TestMatrixDeterminant(c *C) {
	for i := range test.determinantTestTable {
		value := test.determinantTestTable[i]
		det := value.Matrix.Determinant()
		c.Check(Float32Equals(det, value.Expected), Equals, true)
	}
}
