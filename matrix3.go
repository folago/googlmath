package math

import (
	"errors"
)

type Matrix3 []float32

const (
	Matrix3M00 = 0
	Matrix3M01 = 3
	Matrix3M02 = 6
	Matrix3M10 = 1
	Matrix3M11 = 4
	Matrix3M12 = 7
	Matrix3M20 = 2
	Matrix3M21 = 5
	Matrix3M22 = 8
)

func IdentityMatrix3() Matrix3 {
	m := make(Matrix3, 9)
	m[Matrix3M00] = 1
	m[Matrix3M10] = 0
	m[Matrix3M20] = 0
	m[Matrix3M01] = 0
	m[Matrix3M11] = 1
	m[Matrix3M21] = 0
	m[Matrix3M02] = 0
	m[Matrix3M12] = 0
	m[Matrix3M22] = 1
	return m
}

// Multiplies this matrix with the provided matrix and stores the result in this matrix.
func (m Matrix3) Mul(mat Matrix3) Matrix3 {
	m[Matrix3M00] = m[Matrix3M00]*mat[Matrix3M00] + m[Matrix3M01]*mat[Matrix3M10] + m[Matrix3M02]*mat[Matrix3M20]
	m[Matrix3M10] = m[Matrix3M00]*mat[Matrix3M01] + m[Matrix3M01]*mat[Matrix3M11] + m[Matrix3M02]*mat[Matrix3M21]
	m[Matrix3M20] = m[Matrix3M00]*mat[Matrix3M02] + m[Matrix3M01]*mat[Matrix3M12] + m[Matrix3M02]*mat[Matrix3M22]
	m[Matrix3M01] = m[Matrix3M10]*mat[Matrix3M00] + m[Matrix3M11]*mat[Matrix3M10] + m[Matrix3M12]*mat[Matrix3M20]
	m[Matrix3M11] = m[Matrix3M10]*mat[Matrix3M01] + m[Matrix3M11]*mat[Matrix3M11] + m[Matrix3M12]*mat[Matrix3M21]
	m[Matrix3M21] = m[Matrix3M10]*mat[Matrix3M02] + m[Matrix3M11]*mat[Matrix3M12] + m[Matrix3M12]*mat[Matrix3M22]
	m[Matrix3M02] = m[Matrix3M20]*mat[Matrix3M00] + m[Matrix3M21]*mat[Matrix3M10] + m[Matrix3M22]*mat[Matrix3M20]
	m[Matrix3M12] = m[Matrix3M20]*mat[Matrix3M01] + m[Matrix3M21]*mat[Matrix3M11] + m[Matrix3M22]*mat[Matrix3M21]
	m[Matrix3M22] = m[Matrix3M20]*mat[Matrix3M02] + m[Matrix3M21]*mat[Matrix3M12] + m[Matrix3M22]*mat[Matrix3M22]
	return m
}

// Sets this matrix to a rotation matrix that will rotate any vector in counter-clockwise order around the z-axis.
func (m Matrix3) SetToRotation(degrees float32) Matrix3 {
	angle := DegreeToRadians * degrees
	cos := Cos(angle)
	sin := Sin(angle)

	m[Matrix3M00] = cos
	m[Matrix3M10] = sin
	m[Matrix3M20] = 0

	m[Matrix3M01] = -sin
	m[Matrix3M11] = cos
	m[Matrix3M21] = 0

	m[Matrix3M02] = 0
	m[Matrix3M12] = 0
	m[Matrix3M22] = 1

	return m
}

// Sets this matrix to a translation matrix.
func (m Matrix3) SetToTranslation(x, y float32) Matrix3 {
	m[Matrix3M00] = 1
	m[Matrix3M10] = 0
	m[Matrix3M20] = 0

	m[Matrix3M01] = 0
	m[Matrix3M11] = 1
	m[Matrix3M21] = 0

	m[Matrix3M02] = x
	m[Matrix3M12] = y
	m[Matrix3M22] = 1

	return m
}

// Sets this matrix to a scaling matrix.
func (m Matrix3) SetToScaling(scaleX, scaleY float32) Matrix3 {
	m[Matrix3M00] = scaleX
	m[Matrix3M10] = 0
	m[Matrix3M20] = 0
	m[Matrix3M01] = 0
	m[Matrix3M11] = scaleY
	m[Matrix3M21] = 0
	m[Matrix3M02] = 0
	m[Matrix3M12] = 0
	m[Matrix3M22] = 1
	return m
}

// The determinant of this matrix
func (m Matrix3) Det() float32 {
	return m[Matrix3M00]*m[Matrix3M11]*m[Matrix3M22] + m[Matrix3M01]*m[Matrix3M12]*m[Matrix3M20] + m[Matrix3M02]*m[Matrix3M10]*m[Matrix3M21] - m[Matrix3M00]*m[Matrix3M12]*m[Matrix3M21] - m[Matrix3M01]*m[Matrix3M10]*m[Matrix3M22] - m[Matrix3M02]*m[Matrix3M11]*m[Matrix3M20]
}

// Inverts this matrix given that the determinant is != 0
func (m Matrix3) Inv() (Matrix3, error) {
	det := m.Det()
	if det == 0 {
		return nil, errors.New("Can't invert a singular matrix")
	}

	invDet := 1.0 / det

	m[Matrix3M00] = invDet * (m[Matrix3M11]*m[Matrix3M22] - m[Matrix3M21]*m[Matrix3M12])
	m[Matrix3M10] = invDet * (m[Matrix3M20]*m[Matrix3M12] - m[Matrix3M10]*m[Matrix3M22])
	m[Matrix3M20] = invDet * (m[Matrix3M10]*m[Matrix3M21] - m[Matrix3M20]*m[Matrix3M11])
	m[Matrix3M01] = invDet * (m[Matrix3M21]*m[Matrix3M02] - m[Matrix3M01]*m[Matrix3M22])
	m[Matrix3M11] = invDet * (m[Matrix3M00]*m[Matrix3M22] - m[Matrix3M20]*m[Matrix3M02])
	m[Matrix3M21] = invDet * (m[Matrix3M20]*m[Matrix3M01] - m[Matrix3M00]*m[Matrix3M21])
	m[Matrix3M02] = invDet * (m[Matrix3M01]*m[Matrix3M12] - m[Matrix3M11]*m[Matrix3M02])
	m[Matrix3M12] = invDet * (m[Matrix3M10]*m[Matrix3M02] - m[Matrix3M00]*m[Matrix3M12])
	m[Matrix3M22] = invDet * (m[Matrix3M00]*m[Matrix3M11] - m[Matrix3M10]*m[Matrix3M01])

	return m, nil
}

// Copies the values from the provided matrix to this matrix.
func (m Matrix3) Set(mat Matrix3) Matrix3 {
	copy(m, mat)
	return m
}

// Adds a translational component to the matrix in the 3rd column. The other columns are untouched.
func (m Matrix3) Trn(x, y float32) Matrix3 {
	m[Matrix3M02] += x
	m[Matrix3M12] += y
	return m
}

// Postmultiplies this matrix by a translation matrix. Postmultiplication is also used by OpenGL ES' 1.x glTranslate/glRotate/glScale.
func (m Matrix3) Translate(x, y float32) Matrix3 {
	mat := Matrix3{}
	mat[Matrix3M00] = 1
	mat[Matrix3M10] = 0
	mat[Matrix3M20] = 0

	mat[Matrix3M01] = 0
	mat[Matrix3M11] = 1
	mat[Matrix3M21] = 0

	mat[Matrix3M02] = x
	mat[Matrix3M12] = y
	mat[Matrix3M22] = 1
	m.Mul(mat)
	return m
}

// Postmultiplies this matrix with a (counter-clockwise) rotation matrix. Postmultiplication is also used by OpenGL ES' 1.x glTranslate/glRotate/glScale.
func (m Matrix3) Rotate(angle float32) Matrix3 {
	if angle == 0 {
		return m
	}
	angle = DegreeToRadians * angle
	cos := Cos(angle)
	sin := Sin(angle)

	mat := Matrix3{}
	mat[Matrix3M00] = cos
	mat[Matrix3M10] = sin
	mat[Matrix3M20] = 0

	mat[Matrix3M01] = -sin
	mat[Matrix3M11] = cos
	mat[Matrix3M21] = 0

	mat[Matrix3M02] = 0
	mat[Matrix3M12] = 0
	mat[Matrix3M22] = 1
	m.Mul(mat)
	return m
}

// Postmultiplies this matrix with a scale matrix. Postmultiplication is also used by OpenGL ES' 1.x glTranslate/glRotate/glScale.
func (m Matrix3) Scale(scaleX, scaleY float32) Matrix3 {
	mat := Matrix3{}
	mat[Matrix3M00] = scaleX
	mat[Matrix3M10] = 0
	mat[Matrix3M20] = 0

	mat[Matrix3M01] = 0
	mat[Matrix3M11] = scaleY
	mat[Matrix3M21] = 0

	mat[Matrix3M02] = 0
	mat[Matrix3M12] = 0
	mat[Matrix3M22] = 1
	m.Mul(mat)
	return m
}

func (m Matrix3) Values() []float32 {
	return m
}

// Scale the matrix in the both the x and y components by the scalar value.
func (m Matrix3) Scl(scale float32) Matrix3 {
	m[Matrix3M00] *= scale
	m[Matrix3M11] *= scale
	return m
}

// Transposes the current matrix.
func (m Matrix3) Transposes() Matrix3 {
	// Where MXY you do not have to change MXX
	m[Matrix3M01] = m[Matrix3M10]
	m[Matrix3M02] = m[Matrix3M20]
	m[Matrix3M10] = m[Matrix3M01]
	m[Matrix3M12] = m[Matrix3M21]
	m[Matrix3M20] = m[Matrix3M02]
	m[Matrix3M21] = m[Matrix3M12]
	return m
}
