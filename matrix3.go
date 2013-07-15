package math

import (
	"errors"
)

type Matrix3 struct {
	M11, M12, M13 float32
	M21, M22, M23 float32
	M31, M32, M33 float32
}

func NewIdentityMatrix3() *Matrix3 {
	return &Matrix3{M11: 1.0, M22: 1.0, M33: 1.0}
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

// Sets this matrix to a rotation matrix that will rotate any vector in counter-clockwise order around the z-axis.
func (m *Matrix3) SetToRotation(degrees float32) *Matrix3 {
	angle := DegreeToRadians * degrees
	cos := Cos(angle)
	sin := Sin(angle)

	m.M11 = cos
	m.M12 = sin
	m.M13 = 0

	m.M21 = -sin
	m.M22 = cos
	m.M23 = 0

	m.M31 = 0
	m.M32 = 0
	m.M33 = 1

	return m
}

// Sets this matrix to a translation matrix.
func (m *Matrix3) SetToTranslation(x, y float32) *Matrix3 {
	m.M11 = 1
	m.M12 = 0
	m.M13 = 0

	m.M21 = 0
	m.M22 = 1
	m.M23 = 0

	m.M31 = x
	m.M32 = y
	m.M33 = 1

	return m
}

// Sets this matrix to a scaling matrix.
func (m *Matrix3) SetToScaling(scaleX, scaleY float32) *Matrix3 {
	m.M11 = scaleX
	m.M12 = 0
	m.M13 = 0
	m.M21 = 0
	m.M22 = scaleY
	m.M23 = 0
	m.M31 = 0
	m.M32 = 0
	m.M33 = 1
	return m
}

// The determinant of this matrix
func (m *Matrix3) Det() float32 {
	return m.M11*m.M22*m.M33 + m.M21*m.M32*m.M13 + m.M31*m.M12*m.M23 - m.M11*m.M32*m.M23 - m.M21*m.M12*m.M33 - m.M31*m.M22*m.M13
}

// Inverts this matrix given that the determinant is != 0
func (m *Matrix3) Inv() (*Matrix3, error) {
	det := m.Det()
	if det == 0 {
		return nil, errors.New("Can't invert a singular matrix")
	}

	invDet := 1.0 / det

	m.M11 = invDet * (m.M22*m.M33 - m.M23*m.M32)
	m.M12 = invDet * (m.M13*m.M32 - m.M12*m.M33)
	m.M13 = invDet * (m.M12*m.M23 - m.M13*m.M22)
	m.M21 = invDet * (m.M23*m.M31 - m.M21*m.M33)
	m.M22 = invDet * (m.M11*m.M33 - m.M13*m.M31)
	m.M23 = invDet * (m.M13*m.M21 - m.M11*m.M23)
	m.M31 = invDet * (m.M21*m.M32 - m.M22*m.M31)
	m.M32 = invDet * (m.M12*m.M31 - m.M11*m.M32)
	m.M33 = invDet * (m.M11*m.M22 - m.M12*m.M21)

	return m, nil
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

// Adds a translational component to the matrix in the 3rd column. The other columns are untouched.
func (m *Matrix3) Trn(x, y float32) *Matrix3 {
	m.M31 += x
	m.M32 += y
	return m
}

// Postmultiplies this matrix by a translation matrix. Postmultiplication is also used by OpenGL ES' 1.x glTranslate/glRotate/glScale.
func (m *Matrix3) Translate(x, y float32) *Matrix3 {
	mat := &Matrix3{}
	m.M11 = 1
	m.M12 = 0
	m.M13 = 0

	m.M21 = 0
	m.M22 = 1
	m.M23 = 0

	m.M31 = x
	m.M32 = y
	m.M33 = 1
	m.Mul(mat)
	return m
}

// Postmultiplies this matrix with a (counter-clockwise) rotation matrix. Postmultiplication is also used by OpenGL ES' 1.x glTranslate/glRotate/glScale.
func (m *Matrix3) Rotate(angle float32) *Matrix3 {
	if angle == 0 {
		return m
	}
	angle = DegreeToRadians * angle
	cos := Cos(angle)
	sin := Sin(angle)

	mat := &Matrix3{}
	m.M11 = cos
	m.M12 = sin
	m.M13 = 0

	m.M21 = -sin
	m.M22 = cos
	m.M23 = 0

	m.M31 = 0
	m.M32 = 0
	m.M33 = 1
	m.Mul(mat)
	return m
}

// Postmultiplies this matrix with a scale matrix. Postmultiplication is also used by OpenGL ES' 1.x glTranslate/glRotate/glScale.
func (m *Matrix3) Scale(scaleX, scaleY float32) *Matrix3 {
	mat := &Matrix3{}
	m.M11 = scaleX
	m.M12 = 0
	m.M13 = 0

	m.M21 = 0
	m.M22 = scaleY
	m.M23 = 0

	m.M31 = 0
	m.M32 = 0
	m.M33 = 1
	m.Mul(mat)
	return m
}

func (m *Matrix3) Values() []float32 {
	return []float32{m.M11, m.M12, m.M13, m.M21, m.M22, m.M23, m.M31, m.M32, m.M33}
}

// Scale the matrix in the both the x, y, z components by the scalar value.
func (m *Matrix3) Scl(scale float32) *Matrix3 {
	m.M11 *= scale
	m.M22 *= scale
	m.M33 *= scale
	return m
}

// Transposes the current matrix.
func (m *Matrix3) Transposes() *Matrix3 {
	m.M21 = m.M12
	m.M31 = m.M13
	m.M12 = m.M21
	m.M32 = m.M23
	m.M13 = m.M31
	m.M23 = m.M32
	return m
}
