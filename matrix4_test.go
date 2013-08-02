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

type Vector4Matrix4TestValue struct {
	Matrix   *Matrix4
	Value    Vector4
	Expected Vector4
}

type ProjectMatrix4TestValue struct {
	Value    Vector3
	Model    *Matrix4
	Proj     *Matrix4
	Viewport Vector4
	Expected Vector3
}

type Matrix4TestSuite struct {
	perspectiveTestTable []MatrixPerspectiveTestValue
	lookAtTestTable      []MatrixLookAtTestValue
	translateTestTable   []MatrixTranslateTestValue
	rotationTestTable    []MatrixRotationTestValue
	orthoTestTable       []MatrixOrthoTestValue
	mulTestTable         []MatrixMulTestValue
	mulVec4TestTable     []Vector4Matrix4TestValue
	setTestTable         []MatrixSetTestValue
	scaleTestTable       []MatrixScaleTestValue
	invertTestTable      []MatrixInvertTestValue
	determinantTestTable []MatrixDeterminantTestValue
	projectTestTable     []ProjectMatrix4TestValue
	unProjectTestTable   []ProjectMatrix4TestValue
}

var matrixTestSuite = Suite(&Matrix4TestSuite{})

func (test *Matrix4TestSuite) SetUpTest(c *C) {
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

	test.mulVec4TestTable = []Vector4Matrix4TestValue{
		Vector4Matrix4TestValue{
			&Matrix4{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, -3.0, 2.2, 15.0, 1.0},
			Vec4(1.0, 2.0, 3.0, 4.0),
			Vec4(-11.0, 10.8, 63.0, 4.0),
		},
		Vector4Matrix4TestValue{
			&Matrix4{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, -3.3, 2.2, 1.3, 1.0},
			Vec4(-1.0, 2.0, 3.0, -4.0),
			Vec4(12.2, -6.8, -2.2, -4.0),
		},
		Vector4Matrix4TestValue{
			&Matrix4{0.5625, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, -0.0, -1.0, 0.0, 0.0, 1.0, 0.0},
			Vec4(1.2, 0.2, -3.3, 4.0),
			Vec4(0.675, 0.2, 4.0, 3.3),
		},
		Vector4Matrix4TestValue{
			&Matrix4{0.5625, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, -0.0, -1.0, 0.0, 0.0, 1.0, 0.0},
			Vec4(-2.2, 0.0, -3.3, 4.0),
			Vec4(-1.2375, 0.0, 4.0, 3.3),
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

	test.projectTestTable = []ProjectMatrix4TestValue{
		ProjectMatrix4TestValue{
			Value:    Vec3(-1.0, 2.0, 3.0),
			Model:    &Matrix4{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, -3.3, 2.2, 1.3, 1.0},
			Proj:     &Matrix4{0.5625, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, -0.0, -1.0, 0.0, 0.0, 1.0, 0.0},
			Viewport: Vec4(0.0, 0.0, 1.0, 1.0),
			Expected: Vec3(0.78125, 0.011628, 0.383721),
		},
	}

	test.unProjectTestTable = []ProjectMatrix4TestValue{
		ProjectMatrix4TestValue{
			Value:    Vec3(1.0, 2.0, 3.0),
			Model:    &Matrix4{1.0, 0.0, 0.0, 0.0, 0.0, 0.965926, 0.258819, 0.0, 0.0, -0.258819, 0.965926, 0.0, 0.0, 0.0, 0.0, 1.0},
			Proj:     &Matrix4{0.562500, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, -0.0, -1.0, 0.0, 0.0, 1.0, 0.0},
			Viewport: Vec4(0.0, 0.0, 16.0, 9.0),
			Expected: Vec3(-0.311111, -0.159089, -0.164428),
		},
		ProjectMatrix4TestValue{
			Value:    Vec3(1.0, 2.0, 3.0),
			Model:    &Matrix4{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0},
			Proj:     &Matrix4{0.5625, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, -0.0, -1.0, 0.0, 0.0, 1.0, 0.0},
			Viewport: Vec4(0.0, 0.0, 16.0, 9.0),
			Expected: Vec3(-0.311111, -0.111111, -0.2),
		},
	}
}

func (test *Matrix4TestSuite) TestMatrixPerspective(c *C) {
	for i := range test.perspectiveTestTable {
		value := test.perspectiveTestTable[i]
		matrix := NewPerspectiveMatrix4(value.Fov, value.AspectRatio, value.Near, value.Far)
		c.Check(matrix, Matrix4Check, value.Expected)
	}
}

func (test *Matrix4TestSuite) TestMatrixLookAt(c *C) {
	for i := range test.lookAtTestTable {
		value := test.lookAtTestTable[i]
		matrix := NewLookAtMatrix4(value.Eye, value.Center, value.Up)
		c.Check(matrix, Matrix4Check, value.Expected)
	}
}

func (test *Matrix4TestSuite) TestMatrixTranslation(c *C) {
	for i := range test.translateTestTable {
		value := test.translateTestTable[i]
		translation := value.Translation
		matrix := NewTranslationMatrix4(translation.X, translation.Y, translation.Z)
		c.Check(matrix, Matrix4Check, value.Expected)
	}
}

func (test *Matrix4TestSuite) TestMatrixRotation(c *C) {
	for i := range test.rotationTestTable {
		value := test.rotationTestTable[i]
		matrix := NewRotationMatrix4(value.Axis, value.Angle)
		c.Check(matrix, Matrix4Check, value.Expected)
	}
}

func (test *Matrix4TestSuite) TestMatrixOrtho(c *C) {
	for i := range test.orthoTestTable {
		value := test.orthoTestTable[i]
		matrix := NewOrthoMatrix4(value.Left, value.Right, value.Bottom, value.Top, value.Near, value.Far)
		c.Check(matrix, Matrix4Check, value.Expected)
	}
}

func (test *Matrix4TestSuite) TestMatrixMul(c *C) {
	for i := range test.mulTestTable {
		value := test.mulTestTable[i]
		matrix := value.M1.Mul(value.M2)
		c.Check(matrix, Matrix4Check, value.Expected)
	}
}

func (test *Matrix4TestSuite) TestMatrixMulVec4(c *C) {
	for i := range test.mulVec4TestTable {
		value := test.mulVec4TestTable[i]
		v := value.Matrix.MulVec4(value.Value)
		c.Check(v, Vector4Check, value.Expected)
	}
}

func (test *Matrix4TestSuite) TestMatrixSet(c *C) {
	for i := range test.setTestTable {
		value := test.setTestTable[i]
		matrix := value.M1.Set(value.M2)
		c.Check(matrix, Matrix4Check, value.Expected)
		c.Check(matrix, Matrix4Check, value.M1)
		c.Check(matrix, Matrix4Check, value.M2)
		c.Check(value.M1, Matrix4Check, value.M2)
	}
}

func (test *Matrix4TestSuite) TestMatrixScale(c *C) {
	for i := range test.scaleTestTable {
		value := test.scaleTestTable[i]
		matrix := value.Matrix.Scale(value.Scalar)
		c.Check(matrix, Matrix4Check, value.Expected)
	}
}

func (test *Matrix4TestSuite) TestMatrixInvert(c *C) {
	for i := range test.invertTestTable {
		value := test.invertTestTable[i]
		matrix, err := value.Matrix.Invert()
		c.Check(err, IsNil)
		c.Check(matrix, Matrix4Check, value.Expected)
	}
}

func (test *Matrix4TestSuite) TestMatrixDeterminant(c *C) {
	for i := range test.determinantTestTable {
		value := test.determinantTestTable[i]
		det := value.Matrix.Determinant()
		c.Check(det, EqualsFloat32, value.Expected)
	}
}

func (test *Matrix4TestSuite) TestProject(c *C) {
	for i := range test.projectTestTable {
		value := test.projectTestTable[i]
		prj := Project(value.Value, value.Model, value.Proj, value.Viewport)
		c.Check(prj, Vector3Check, value.Expected)
	}
}

func (test *Matrix4TestSuite) TestUnProject(c *C) {
	for i := range test.unProjectTestTable {
		value := test.unProjectTestTable[i]
		unProj, err := UnProject(value.Value, value.Model, value.Proj, value.Viewport)
		c.Check(err, IsNil)
		c.Check(unProj, Vector3Check, value.Expected)
	}
}
