package math

import (
	"errors"
)

type Matrix3 struct {
	M11, M12, M13 float32
	M21, M22, M23 float32
	M31, M32, M33 float32
}

func NewMatrix3(m11, m12, m13, m21, m22, m23, m31, m32, m33 float32) *Matrix3 {
	return &Matrix3{m11, m12, m13, m21, m22, m23, m31, m32, m33}
}

func NewIdentityMatrix3() *Matrix3 {
	return &Matrix3{M11: 1.0, M22: 1.0, M33: 1.0}
}

func NewXRotationMatrix3(angle float32) *Matrix3 {
	angle = DegreeToRadians * angle

	c := Cos(angle)
	s := Sin(angle)

	return &Matrix3{
		1, 0, 0,
		0, c, -s,
		0, s, c,
	}
}

func NewYRotationMatrix3(angle float32) *Matrix3 {
	angle = DegreeToRadians * angle

	c := Cos(angle)
	s := Sin(angle)

	return &Matrix3{
		c, 0, s,
		0, 1, 0,
		-s, 0, c,
	}
}

// Returns a rotation matrix that will rotate any vector in counter-clockwise order around the z-axis.
func NewZRotationMatrix3(angle float32) *Matrix3 {
	angle = DegreeToRadians * angle

	c := Cos(angle)
	s := Sin(angle)

	return &Matrix3{
		c, -s, 0,
		s, c, 0,
		0, 0, 1,
	}
}

func NewRotationMatrix3(axis Vector3, angle float32) *Matrix3 {
	axis = axis.Nor()
	angle = DegreeToRadians * angle

	c := Cos(angle)
	s := Sin(angle)
	k := 1 - c

	return &Matrix3{axis.X*axis.X*k + c, axis.X*axis.Y*k + axis.Z*s, axis.X*axis.Z*k - axis.Y*s,
		axis.X*axis.Y*k - axis.Z*s, axis.Y*axis.Y*k + c, axis.Y*axis.Z*k + axis.X*s,
		axis.X*axis.Z*k + axis.Y*s, axis.Y*axis.Z*k - axis.X*s, axis.Z*axis.Z*k + c}
}

func NewTranslationMatrix3(x, y float32) *Matrix3 {
	return &Matrix3{
		1, 0, 0,
		0, 1, 0,
		x, y, 1,
	}
}

func NewScaleMatrix3(scaleX, scaleY float32) *Matrix3 {
	return &Matrix3{
		scaleX, 0, 0,
		0, scaleY, 0,
		0, 0, 1,
	}
}

// Copies the values from the provided matrix to this matrix.
func (m *Matrix3) Set(mat *Matrix3) *Matrix3 {
	m.M11 = mat.M11
	m.M12 = mat.M12
	m.M13 = mat.M13
	m.M21 = mat.M21
	m.M22 = mat.M22
	m.M23 = mat.M23
	m.M31 = mat.M31
	m.M32 = mat.M32
	m.M33 = mat.M33
	return m
}

// Multiplies this matrix with the provided matrix and returns a new matrix.
func (m *Matrix3) Mul(mat *Matrix3) *Matrix3 {
	temp := &Matrix3{}
	temp.M11 = m.M11*mat.M11 + m.M21*mat.M12 + m.M31*mat.M13
	temp.M12 = m.M12*mat.M11 + m.M22*mat.M12 + m.M32*mat.M13
	temp.M13 = m.M13*mat.M11 + m.M23*mat.M12 + m.M33*mat.M13

	temp.M21 = m.M11*mat.M21 + m.M21*mat.M22 + m.M31*mat.M23
	temp.M22 = m.M12*mat.M21 + m.M22*mat.M22 + m.M32*mat.M23
	temp.M23 = m.M13*mat.M21 + m.M23*mat.M22 + m.M33*mat.M23

	temp.M31 = m.M11*mat.M31 + m.M21*mat.M32 + m.M31*mat.M33
	temp.M32 = m.M12*mat.M31 + m.M22*mat.M32 + m.M32*mat.M33
	temp.M33 = m.M13*mat.M31 + m.M23*mat.M32 + m.M33*mat.M33
	return temp
}

// Returns tThe determinant of this matrix
func (m *Matrix3) Determinant() float32 {
	return m.M11*m.M22*m.M33 + m.M21*m.M32*m.M13 + m.M31*m.M12*m.M23 - m.M11*m.M32*m.M23 - m.M21*m.M12*m.M33 - m.M31*m.M22*m.M13
}

// Returns the inverse matrix given that the determinant is != 0
func (m *Matrix3) Inverse() (*Matrix3, error) {
	det := m.Determinant()
	if det == 0 {
		return nil, errors.New("Can't invert a singular matrix")
	}

	invDet := 1.0 / det

	return &Matrix3{
		invDet * (m.M22*m.M33 - m.M23*m.M32), invDet * (m.M13*m.M32 - m.M12*m.M33), invDet * (m.M12*m.M23 - m.M13*m.M22),
		invDet * (m.M23*m.M31 - m.M21*m.M33), invDet * (m.M11*m.M33 - m.M13*m.M31), invDet * (m.M13*m.M21 - m.M11*m.M23),
		invDet * (m.M21*m.M32 - m.M22*m.M31), invDet * (m.M12*m.M31 - m.M11*m.M32), invDet * (m.M11*m.M22 - m.M12*m.M21),
	}, nil
}

func (m *Matrix3) ToArray() []float32 {
	return []float32{m.M11, m.M12, m.M13, m.M21, m.M22, m.M23, m.M31, m.M32, m.M33}
}

// Returns this matrix transposed.
func (m *Matrix3) Transpose() *Matrix3 {
	return &Matrix3{
		m.M11, m.M21, m.M31,
		m.M12, m.M22, m.M32,
		m.M13, m.M23, m.M33,
	}
}

// Build planar projection matrix along normal axis.
func (m *Matrix3) Proj2D(normal Vector3) *Matrix3 {
	r := &Matrix3{
		1 - normal.X*normal.X, -normal.X * normal.Y, 0,
		-normal.X * normal.Y, 1 - normal.Y*normal.Y, 0,
		0, 0, 1,
	}
	return m.Mul(r)
}

// Returns a transformed matrix with a shearing on X axis.
// TODO ShearX2D(y float32) *Matrix3

// Returns a transformed matrix with a shearing on Y axis.
// TODO ShearY2D(x float32) *Matrix3
