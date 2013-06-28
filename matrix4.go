package math

import (
	"errors"
)

const (
	Matrix4M00 = 0  // 0
	Matrix4M01 = 4  // 1
	Matrix4M02 = 8  // 2
	Matrix4M03 = 12 // 3
	Matrix4M10 = 1  // 4
	Matrix4M11 = 5  // 5
	Matrix4M12 = 9  // 6
	Matrix4M13 = 13 // 7
	Matrix4M20 = 2  // 8
	Matrix4M21 = 6  // 9
	Matrix4M22 = 10 // 10
	Matrix4M23 = 14 // 11
	Matrix4M30 = 3  // 12
	Matrix4M31 = 7  // 13
	Matrix4M32 = 11 // 14
	Matrix4M33 = 15 // 15
)

// Encapsulates a column major 4 by 4 matrix.
type Matrix4 []float32

func NewIdentityMatrix4() Matrix4 {
	m := make(Matrix4, 16)
	m[Matrix4M00] = 1
	m[Matrix4M11] = 1
	m[Matrix4M22] = 1
	m[Matrix4M33] = 1
	return m

}

func NewRotationMatrix4(axis Vector3, angle float32) Matrix4 {
	if angle == 0 {
		return NewIdentityMatrix4()
	}
	m := Matrix4{}
	return m.SetQuaternion(NewQuaternion(axis.X, axis.Y, axis.Z, angle))
}

func NewTranslationMatrix4(vec Vector3) Matrix4 {
	m := NewIdentityMatrix4()
	m[Matrix4M03] = vec.X
	m[Matrix4M13] = vec.Y
	m[Matrix4M23] = vec.Z
	return m
}

func (m Matrix4) Set(mat Matrix4) Matrix4 {
	copy(m, mat)
	return m
}

func (m Matrix4) Cpy() Matrix4 {
	mat := make(Matrix4, 16)
	copy(mat, m)
	return mat
}

// Adds a translational component to the matrix in the 4th column. The other columns are untouched.
func (m Matrix4) Trn(x, y, z float32) Matrix4 {
	m[Matrix4M03] += x
	m[Matrix4M13] += y
	m[Matrix4M23] += z
	return m
}

func (m Matrix4) Translate(vec Vector3) Matrix4 {
	m[Matrix4M03] += vec.X
	m[Matrix4M13] += vec.Y
	m[Matrix4M23] += vec.Z
	return m
}

func (m Matrix4) Rotate(axis Vector3, angle float32) Matrix4 {
	if angle == 0 {
		return m
	}
	quat := NewQuaternion(axis.X, axis.Y, axis.Z, angle)
	m2 := quat.Matrix()
	return m.Mul(m2)
}

func (m Matrix4) Scale(scale Vector3) Matrix4 {
	tmp := NewIdentityMatrix4()
	tmp[Matrix4M00] = scale.X
	tmp[Matrix4M01] = 0
	tmp[Matrix4M02] = 0
	tmp[Matrix4M03] = 0
	tmp[Matrix4M10] = 0
	tmp[Matrix4M11] = scale.Y
	tmp[Matrix4M12] = 0
	tmp[Matrix4M13] = 0
	tmp[Matrix4M20] = 0
	tmp[Matrix4M21] = 0
	tmp[Matrix4M22] = scale.Z
	tmp[Matrix4M23] = 0
	tmp[Matrix4M30] = 0
	tmp[Matrix4M31] = 0
	tmp[Matrix4M32] = 0
	tmp[Matrix4M33] = 1
	return m.Mul(tmp)
}

func (m Matrix4) Values() []float32 {
	return m
}

// Multiplies this matrix with the given matrix, storing the result in this matrix.
func (m Matrix4) Mul(mat Matrix4) Matrix4 {
	tmp := make(Matrix4, 16)
	tmp[Matrix4M00] = m[Matrix4M00]*mat[Matrix4M00] + m[Matrix4M01]*mat[Matrix4M10] + m[Matrix4M02]*mat[Matrix4M20] + m[Matrix4M03]*mat[Matrix4M30]
	tmp[Matrix4M01] = m[Matrix4M00]*mat[Matrix4M01] + m[Matrix4M01]*mat[Matrix4M11] + m[Matrix4M02]*mat[Matrix4M21] + m[Matrix4M03]*mat[Matrix4M31]
	tmp[Matrix4M02] = m[Matrix4M00]*mat[Matrix4M02] + m[Matrix4M01]*mat[Matrix4M12] + m[Matrix4M02]*mat[Matrix4M22] + m[Matrix4M03]*mat[Matrix4M32]
	tmp[Matrix4M03] = m[Matrix4M00]*mat[Matrix4M03] + m[Matrix4M01]*mat[Matrix4M13] + m[Matrix4M02]*mat[Matrix4M23] + m[Matrix4M03]*mat[Matrix4M33]
	tmp[Matrix4M10] = m[Matrix4M10]*mat[Matrix4M00] + m[Matrix4M11]*mat[Matrix4M10] + m[Matrix4M12]*mat[Matrix4M20] + m[Matrix4M13]*mat[Matrix4M30]
	tmp[Matrix4M11] = m[Matrix4M10]*mat[Matrix4M01] + m[Matrix4M11]*mat[Matrix4M11] + m[Matrix4M12]*mat[Matrix4M21] + m[Matrix4M13]*mat[Matrix4M31]
	tmp[Matrix4M12] = m[Matrix4M10]*mat[Matrix4M02] + m[Matrix4M11]*mat[Matrix4M12] + m[Matrix4M12]*mat[Matrix4M22] + m[Matrix4M13]*mat[Matrix4M32]
	tmp[Matrix4M13] = m[Matrix4M10]*mat[Matrix4M03] + m[Matrix4M11]*mat[Matrix4M13] + m[Matrix4M12]*mat[Matrix4M23] + m[Matrix4M13]*mat[Matrix4M33]
	tmp[Matrix4M20] = m[Matrix4M20]*mat[Matrix4M00] + m[Matrix4M21]*mat[Matrix4M10] + m[Matrix4M22]*mat[Matrix4M20] + m[Matrix4M23]*mat[Matrix4M30]
	tmp[Matrix4M21] = m[Matrix4M20]*mat[Matrix4M01] + m[Matrix4M21]*mat[Matrix4M11] + m[Matrix4M22]*mat[Matrix4M21] + m[Matrix4M23]*mat[Matrix4M31]
	tmp[Matrix4M22] = m[Matrix4M20]*mat[Matrix4M02] + m[Matrix4M21]*mat[Matrix4M12] + m[Matrix4M22]*mat[Matrix4M22] + m[Matrix4M23]*mat[Matrix4M32]
	tmp[Matrix4M23] = m[Matrix4M20]*mat[Matrix4M03] + m[Matrix4M21]*mat[Matrix4M13] + m[Matrix4M22]*mat[Matrix4M23] + m[Matrix4M23]*mat[Matrix4M33]
	tmp[Matrix4M30] = m[Matrix4M30]*mat[Matrix4M00] + m[Matrix4M31]*mat[Matrix4M10] + m[Matrix4M32]*mat[Matrix4M20] + m[Matrix4M33]*mat[Matrix4M30]
	tmp[Matrix4M31] = m[Matrix4M30]*mat[Matrix4M01] + m[Matrix4M31]*mat[Matrix4M11] + m[Matrix4M32]*mat[Matrix4M21] + m[Matrix4M33]*mat[Matrix4M31]
	tmp[Matrix4M32] = m[Matrix4M30]*mat[Matrix4M02] + m[Matrix4M31]*mat[Matrix4M12] + m[Matrix4M32]*mat[Matrix4M22] + m[Matrix4M33]*mat[Matrix4M32]
	tmp[Matrix4M33] = m[Matrix4M30]*mat[Matrix4M03] + m[Matrix4M31]*mat[Matrix4M13] + m[Matrix4M32]*mat[Matrix4M23] + m[Matrix4M33]*mat[Matrix4M33]
	return m.Set(tmp)
}

func (m Matrix4) MulVec3(vec Vector3) Vector3 {
	vec.X = vec.X*m[Matrix4M00] + vec.Y*m[Matrix4M01] + vec.Z*m[Matrix4M02] + m[Matrix4M03]
	vec.Y = vec.X*m[Matrix4M10] + vec.Y*m[Matrix4M11] + vec.Z*m[Matrix4M12] + m[Matrix4M13]
	vec.Z = vec.X*m[Matrix4M20] + vec.Y*m[Matrix4M21] + vec.Z*m[Matrix4M22] + m[Matrix4M23]
	return vec
}

func (m Matrix4) Transpose() Matrix4 {
	tmp := make(Matrix4, 16)
	tmp[Matrix4M00] = m[Matrix4M00]
	tmp[Matrix4M01] = m[Matrix4M10]
	tmp[Matrix4M02] = m[Matrix4M20]
	tmp[Matrix4M03] = m[Matrix4M30]
	tmp[Matrix4M10] = m[Matrix4M01]
	tmp[Matrix4M11] = m[Matrix4M11]
	tmp[Matrix4M12] = m[Matrix4M21]
	tmp[Matrix4M13] = m[Matrix4M31]
	tmp[Matrix4M20] = m[Matrix4M02]
	tmp[Matrix4M21] = m[Matrix4M12]
	tmp[Matrix4M22] = m[Matrix4M22]
	tmp[Matrix4M23] = m[Matrix4M32]
	tmp[Matrix4M30] = m[Matrix4M03]
	tmp[Matrix4M31] = m[Matrix4M13]
	tmp[Matrix4M32] = m[Matrix4M23]
	tmp[Matrix4M33] = m[Matrix4M33]
	return m.Set(tmp)
}

// Sets the matrix to an identity matrix.
func (m Matrix4) Identity() Matrix4 {
	m[Matrix4M00] = 1
	m[Matrix4M01] = 0
	m[Matrix4M02] = 0
	m[Matrix4M03] = 0
	m[Matrix4M10] = 0
	m[Matrix4M11] = 1
	m[Matrix4M12] = 0
	m[Matrix4M13] = 0
	m[Matrix4M20] = 0
	m[Matrix4M21] = 0
	m[Matrix4M22] = 1
	m[Matrix4M23] = 0
	m[Matrix4M30] = 0
	m[Matrix4M31] = 0
	m[Matrix4M32] = 0
	m[Matrix4M33] = 1
	return m
}

// The determinant of this matrix.
func (m Matrix4) Det() float32 {
	return m[Matrix4M30]*m[Matrix4M21]*m[Matrix4M12]*m[Matrix4M03] -
		m[Matrix4M20]*m[Matrix4M31]*m[Matrix4M12]*m[Matrix4M03] -
		m[Matrix4M30]*m[Matrix4M11]*m[Matrix4M22]*m[Matrix4M03] +
		m[Matrix4M10]*m[Matrix4M31]*m[Matrix4M22]*m[Matrix4M03] +
		m[Matrix4M20]*m[Matrix4M11]*m[Matrix4M32]*m[Matrix4M03] -
		m[Matrix4M10]*m[Matrix4M21]*m[Matrix4M32]*m[Matrix4M03] -
		m[Matrix4M30]*m[Matrix4M21]*m[Matrix4M02]*m[Matrix4M13] +
		m[Matrix4M20]*m[Matrix4M31]*m[Matrix4M02]*m[Matrix4M13] +
		m[Matrix4M30]*m[Matrix4M01]*m[Matrix4M22]*m[Matrix4M13] -
		m[Matrix4M00]*m[Matrix4M31]*m[Matrix4M22]*m[Matrix4M13] -
		m[Matrix4M20]*m[Matrix4M01]*m[Matrix4M32]*m[Matrix4M13] +
		m[Matrix4M00]*m[Matrix4M21]*m[Matrix4M32]*m[Matrix4M13] +
		m[Matrix4M30]*m[Matrix4M11]*m[Matrix4M02]*m[Matrix4M23] -
		m[Matrix4M10]*m[Matrix4M31]*m[Matrix4M02]*m[Matrix4M23] -
		m[Matrix4M30]*m[Matrix4M01]*m[Matrix4M12]*m[Matrix4M23] +
		m[Matrix4M00]*m[Matrix4M31]*m[Matrix4M12]*m[Matrix4M23] +
		m[Matrix4M10]*m[Matrix4M01]*m[Matrix4M32]*m[Matrix4M23] -
		m[Matrix4M00]*m[Matrix4M11]*m[Matrix4M32]*m[Matrix4M23] -
		m[Matrix4M20]*m[Matrix4M11]*m[Matrix4M02]*m[Matrix4M33] +
		m[Matrix4M10]*m[Matrix4M21]*m[Matrix4M02]*m[Matrix4M33] +
		m[Matrix4M20]*m[Matrix4M01]*m[Matrix4M12]*m[Matrix4M33] -
		m[Matrix4M00]*m[Matrix4M21]*m[Matrix4M12]*m[Matrix4M33] -
		m[Matrix4M10]*m[Matrix4M01]*m[Matrix4M22]*m[Matrix4M33] +
		m[Matrix4M00]*m[Matrix4M11]*m[Matrix4M22]*m[Matrix4M33]
}

// Sets the matrix to a rotation matrix representing the quaternion.
func (m Matrix4) SetQuaternion(quaternion *Quaternion) Matrix4 {
	// Compute quaternion factors
	l_xx := quaternion.X * quaternion.X
	l_xy := quaternion.X * quaternion.Y
	l_xz := quaternion.X * quaternion.Z
	l_xw := quaternion.X * quaternion.W
	l_yy := quaternion.Y * quaternion.Y
	l_yz := quaternion.Y * quaternion.Z
	l_yw := quaternion.Y * quaternion.W
	l_zz := quaternion.Z * quaternion.Z
	l_zw := quaternion.Z * quaternion.W

	// Set matrix from quaternion
	m[Matrix4M00] = 1 - 2*(l_yy+l_zz)
	m[Matrix4M01] = 2 * (l_xy - l_zw)
	m[Matrix4M02] = 2 * (l_xz + l_yw)
	m[Matrix4M03] = 0
	m[Matrix4M10] = 2 * (l_xy + l_zw)
	m[Matrix4M11] = 1 - 2*(l_xx+l_zz)
	m[Matrix4M12] = 2 * (l_yz - l_xw)
	m[Matrix4M13] = 0
	m[Matrix4M20] = 2 * (l_xz - l_yw)
	m[Matrix4M21] = 2 * (l_yz + l_xw)
	m[Matrix4M22] = 1 - 2*(l_xx+l_yy)
	m[Matrix4M23] = 0
	m[Matrix4M30] = 0
	m[Matrix4M31] = 0
	m[Matrix4M32] = 0
	m[Matrix4M33] = 1

	return m
}

// Sets the matrix to a projection matrix with a near- and far plane, a field of view in degrees and an aspect ratio.
func (m Matrix4) SetToProjection(near, far, fov, aspectRatio float32) Matrix4 {
	m.Identity()
	l_fd := 1.0 / Tan((fov*(Pi/180))/2.0)
	l_a1 := (far + near) / (near - far)
	l_a2 := (2 * far * near) / (near - far)
	m[Matrix4M00] = l_fd / aspectRatio
	m[Matrix4M10] = 0
	m[Matrix4M20] = 0
	m[Matrix4M30] = 0
	m[Matrix4M01] = 0
	m[Matrix4M11] = l_fd
	m[Matrix4M21] = 0
	m[Matrix4M31] = 0
	m[Matrix4M02] = 0
	m[Matrix4M12] = 0
	m[Matrix4M22] = l_a1
	m[Matrix4M32] = -1
	m[Matrix4M03] = 0
	m[Matrix4M13] = 0
	m[Matrix4M23] = l_a2
	m[Matrix4M33] = 0
	return m
}

// Sets the matrix to an orthographic projection like glOrtho (http://www.opengl.org/sdk/docs/man/xhtml/glOrtho.xml) following the OpenGL equivalent
func (m Matrix4) SetToOrtho2D(left, right, bottom, top, near, far float32) Matrix4 {
	m.Identity()

	x_orth := 2 / (right - left)
	y_orth := 2 / (top - bottom)
	z_orth := -2 / (far - near)

	tx := -(right + left) / (right - left)
	ty := -(top + bottom) / (top - bottom)
	tz := -(far + near) / (far - near)

	m[Matrix4M00] = x_orth
	m[Matrix4M10] = 0
	m[Matrix4M20] = 0
	m[Matrix4M30] = 0
	m[Matrix4M01] = 0
	m[Matrix4M11] = y_orth
	m[Matrix4M21] = 0
	m[Matrix4M31] = 0
	m[Matrix4M02] = 0
	m[Matrix4M12] = 0
	m[Matrix4M22] = z_orth
	m[Matrix4M32] = 0
	m[Matrix4M03] = tx
	m[Matrix4M13] = ty
	m[Matrix4M23] = tz
	m[Matrix4M33] = 1
	return m
}

// Sets this matrix to a translation matrix, overwriting it first by an identity matrix and then setting the 4th column to the translation vector.
func (m Matrix4) SetToTranslation(x, y, z float32) Matrix4 {
	m.Identity()
	m[Matrix4M03] = x
	m[Matrix4M13] = y
	m[Matrix4M23] = z
	return m
}

// Sets this matrix to a translation and scaling matrix by first overwritting it with an identity and then setting the translation vector in the 4th column and the scaling vector in the diagonal.
func (m Matrix4) SetToTranslationAndScaling(translationX, translationY, translationZ, scaleX, scaleY, scaleZ float32) Matrix4 {
	m.Identity()
	m[Matrix4M03] = translationX
	m[Matrix4M13] = translationY
	m[Matrix4M23] = translationZ
	m[Matrix4M00] = scaleX
	m[Matrix4M11] = scaleY
	m[Matrix4M22] = scaleZ
	return m
}

// Sets the matrix to a rotation matrix around the given axis.
func (m Matrix4) SetToRotation(axisX, axisY, axisZ, angle float32) Matrix4 {
	if angle == 0 {
		return m.Identity()
	}
	return m.SetQuaternion(NewQuaternion(axisX, axisY, axisZ, angle))
}

func (m Matrix4) SetFromEulerAngles(yaw, pitch, roll float32) Matrix4 {
	quat := &Quaternion{}
	quat.SetEulerAngles(yaw, pitch, roll)
	return m.SetQuaternion(quat)
}

func (m Matrix4) SetToScaling(x, y, z float32) Matrix4 {
	m.Identity()
	m[Matrix4M00] = x
	m[Matrix4M11] = y
	m[Matrix4M22] = z
	return m
}

func (m Matrix4) SetToLookAt(direction, up Vector3) Matrix4 {
	l_vez := direction.Cpy().Nor()
	l_vex := direction.Cpy().Nor()
	l_vex.Crs(up).Nor()
	l_vey := l_vex.Cpy().Crs(l_vez).Nor()
	m.Identity()

	m[Matrix4M00] = l_vex.X
	m[Matrix4M01] = l_vex.Y
	m[Matrix4M02] = l_vex.Z
	m[Matrix4M10] = l_vey.X
	m[Matrix4M11] = l_vey.Y
	m[Matrix4M12] = l_vey.Z
	m[Matrix4M20] = -l_vez.X
	m[Matrix4M21] = -l_vez.Y
	m[Matrix4M22] = -l_vez.Z

	return m
}

func (m Matrix4) SetToLookAtTarget(position, target, up Vector3) Matrix4 {
	tmpVec := target.Cpy().Sub(position)
	m.SetToLookAt(tmpVec, up)
	return m.Mul(NewTranslationMatrix4(position.Cpy().Scale(-1)))
}

func (m Matrix4) Prj(vec Vector3) Vector3 {
	invW := 1.0 / (vec.X*m[Matrix4M30] + vec.Y*m[Matrix4M31] + vec.Z*m[Matrix4M32] + m[Matrix4M33])
	x := (vec.X*m[Matrix4M00] + vec.Y*m[Matrix4M01] + vec.Z*m[Matrix4M02] + m[Matrix4M03]) * invW
	y := (vec.X*m[Matrix4M10] + vec.Y*m[Matrix4M11] + vec.Z*m[Matrix4M12] + m[Matrix4M13]) * invW
	z := (vec.X*m[Matrix4M20] + vec.Y*m[Matrix4M21] + vec.Z*m[Matrix4M22] + m[Matrix4M23]) * invW
	return Vec3(x, y, z)
}

func (m Matrix4) Inv() (Matrix4, error) {
	det := m.Det()
	if det == 0 {
		return nil, errors.New("non-invertible matrix")
	}

	tmp := m.Cpy()

	tmp[Matrix4M00] = m[Matrix4M12]*m[Matrix4M23]*m[Matrix4M31] - m[Matrix4M13]*m[Matrix4M22]*m[Matrix4M31] + m[Matrix4M13]*m[Matrix4M21]*m[Matrix4M32] - m[Matrix4M11]*m[Matrix4M23]*m[Matrix4M32] - m[Matrix4M12]*m[Matrix4M21]*m[Matrix4M33] + m[Matrix4M11]*m[Matrix4M22]*m[Matrix4M33]
	tmp[Matrix4M01] = m[Matrix4M03]*m[Matrix4M22]*m[Matrix4M31] - m[Matrix4M02]*m[Matrix4M23]*m[Matrix4M31] - m[Matrix4M03]*m[Matrix4M21]*m[Matrix4M32] + m[Matrix4M01]*m[Matrix4M23]*m[Matrix4M32] + m[Matrix4M02]*m[Matrix4M21]*m[Matrix4M33] - m[Matrix4M01]*m[Matrix4M22]*m[Matrix4M33]
	tmp[Matrix4M02] = m[Matrix4M02]*m[Matrix4M13]*m[Matrix4M31] - m[Matrix4M03]*m[Matrix4M12]*m[Matrix4M31] + m[Matrix4M03]*m[Matrix4M11]*m[Matrix4M32] - m[Matrix4M01]*m[Matrix4M13]*m[Matrix4M32] - m[Matrix4M02]*m[Matrix4M11]*m[Matrix4M33] + m[Matrix4M01]*m[Matrix4M12]*m[Matrix4M33]
	tmp[Matrix4M03] = m[Matrix4M03]*m[Matrix4M12]*m[Matrix4M21] - m[Matrix4M02]*m[Matrix4M13]*m[Matrix4M21] - m[Matrix4M03]*m[Matrix4M11]*m[Matrix4M22] + m[Matrix4M01]*m[Matrix4M13]*m[Matrix4M22] + m[Matrix4M02]*m[Matrix4M11]*m[Matrix4M23] - m[Matrix4M01]*m[Matrix4M12]*m[Matrix4M23]
	tmp[Matrix4M10] = m[Matrix4M13]*m[Matrix4M22]*m[Matrix4M30] - m[Matrix4M12]*m[Matrix4M23]*m[Matrix4M30] - m[Matrix4M13]*m[Matrix4M20]*m[Matrix4M32] + m[Matrix4M10]*m[Matrix4M23]*m[Matrix4M32] + m[Matrix4M12]*m[Matrix4M20]*m[Matrix4M33] - m[Matrix4M10]*m[Matrix4M22]*m[Matrix4M33]
	tmp[Matrix4M11] = m[Matrix4M02]*m[Matrix4M23]*m[Matrix4M30] - m[Matrix4M03]*m[Matrix4M22]*m[Matrix4M30] + m[Matrix4M03]*m[Matrix4M20]*m[Matrix4M32] - m[Matrix4M00]*m[Matrix4M23]*m[Matrix4M32] - m[Matrix4M02]*m[Matrix4M20]*m[Matrix4M33] + m[Matrix4M00]*m[Matrix4M22]*m[Matrix4M33]
	tmp[Matrix4M12] = m[Matrix4M03]*m[Matrix4M12]*m[Matrix4M30] - m[Matrix4M02]*m[Matrix4M13]*m[Matrix4M30] - m[Matrix4M03]*m[Matrix4M10]*m[Matrix4M32] + m[Matrix4M00]*m[Matrix4M13]*m[Matrix4M32] + m[Matrix4M02]*m[Matrix4M10]*m[Matrix4M33] - m[Matrix4M00]*m[Matrix4M12]*m[Matrix4M33]
	tmp[Matrix4M13] = m[Matrix4M02]*m[Matrix4M13]*m[Matrix4M20] - m[Matrix4M03]*m[Matrix4M12]*m[Matrix4M20] + m[Matrix4M03]*m[Matrix4M10]*m[Matrix4M22] - m[Matrix4M00]*m[Matrix4M13]*m[Matrix4M22] - m[Matrix4M02]*m[Matrix4M10]*m[Matrix4M23] + m[Matrix4M00]*m[Matrix4M12]*m[Matrix4M23]
	tmp[Matrix4M20] = m[Matrix4M11]*m[Matrix4M23]*m[Matrix4M30] - m[Matrix4M13]*m[Matrix4M21]*m[Matrix4M30] + m[Matrix4M13]*m[Matrix4M20]*m[Matrix4M31] - m[Matrix4M10]*m[Matrix4M23]*m[Matrix4M31] - m[Matrix4M11]*m[Matrix4M20]*m[Matrix4M33] + m[Matrix4M10]*m[Matrix4M21]*m[Matrix4M33]
	tmp[Matrix4M21] = m[Matrix4M03]*m[Matrix4M21]*m[Matrix4M30] - m[Matrix4M01]*m[Matrix4M23]*m[Matrix4M30] - m[Matrix4M03]*m[Matrix4M20]*m[Matrix4M31] + m[Matrix4M00]*m[Matrix4M23]*m[Matrix4M31] + m[Matrix4M01]*m[Matrix4M20]*m[Matrix4M33] - m[Matrix4M00]*m[Matrix4M21]*m[Matrix4M33]
	tmp[Matrix4M22] = m[Matrix4M01]*m[Matrix4M13]*m[Matrix4M30] - m[Matrix4M03]*m[Matrix4M11]*m[Matrix4M30] + m[Matrix4M03]*m[Matrix4M10]*m[Matrix4M31] - m[Matrix4M00]*m[Matrix4M13]*m[Matrix4M31] - m[Matrix4M01]*m[Matrix4M10]*m[Matrix4M33] + m[Matrix4M00]*m[Matrix4M11]*m[Matrix4M33]
	tmp[Matrix4M23] = m[Matrix4M03]*m[Matrix4M11]*m[Matrix4M20] - m[Matrix4M01]*m[Matrix4M13]*m[Matrix4M20] - m[Matrix4M03]*m[Matrix4M10]*m[Matrix4M21] + m[Matrix4M00]*m[Matrix4M13]*m[Matrix4M21] + m[Matrix4M01]*m[Matrix4M10]*m[Matrix4M23] - m[Matrix4M00]*m[Matrix4M11]*m[Matrix4M23]
	tmp[Matrix4M30] = m[Matrix4M12]*m[Matrix4M21]*m[Matrix4M30] - m[Matrix4M11]*m[Matrix4M22]*m[Matrix4M30] - m[Matrix4M12]*m[Matrix4M20]*m[Matrix4M31] + m[Matrix4M10]*m[Matrix4M22]*m[Matrix4M31] + m[Matrix4M11]*m[Matrix4M20]*m[Matrix4M32] - m[Matrix4M10]*m[Matrix4M21]*m[Matrix4M32]
	tmp[Matrix4M31] = m[Matrix4M01]*m[Matrix4M22]*m[Matrix4M30] - m[Matrix4M02]*m[Matrix4M21]*m[Matrix4M30] + m[Matrix4M02]*m[Matrix4M20]*m[Matrix4M31] - m[Matrix4M00]*m[Matrix4M22]*m[Matrix4M31] - m[Matrix4M01]*m[Matrix4M20]*m[Matrix4M32] + m[Matrix4M00]*m[Matrix4M21]*m[Matrix4M32]
	tmp[Matrix4M32] = m[Matrix4M02]*m[Matrix4M11]*m[Matrix4M30] - m[Matrix4M01]*m[Matrix4M12]*m[Matrix4M30] - m[Matrix4M02]*m[Matrix4M10]*m[Matrix4M31] + m[Matrix4M00]*m[Matrix4M12]*m[Matrix4M31] + m[Matrix4M01]*m[Matrix4M10]*m[Matrix4M32] - m[Matrix4M00]*m[Matrix4M11]*m[Matrix4M32]
	tmp[Matrix4M33] = m[Matrix4M01]*m[Matrix4M12]*m[Matrix4M20] - m[Matrix4M02]*m[Matrix4M11]*m[Matrix4M20] + m[Matrix4M02]*m[Matrix4M10]*m[Matrix4M21] - m[Matrix4M00]*m[Matrix4M12]*m[Matrix4M21] - m[Matrix4M01]*m[Matrix4M10]*m[Matrix4M22] + m[Matrix4M00]*m[Matrix4M11]*m[Matrix4M22]

	inv_det := 1.0 / det
	m[Matrix4M00] = tmp[Matrix4M00] * inv_det
	m[Matrix4M01] = tmp[Matrix4M01] * inv_det
	m[Matrix4M02] = tmp[Matrix4M02] * inv_det
	m[Matrix4M03] = tmp[Matrix4M03] * inv_det
	m[Matrix4M10] = tmp[Matrix4M10] * inv_det
	m[Matrix4M11] = tmp[Matrix4M11] * inv_det
	m[Matrix4M12] = tmp[Matrix4M12] * inv_det
	m[Matrix4M13] = tmp[Matrix4M13] * inv_det
	m[Matrix4M20] = tmp[Matrix4M20] * inv_det
	m[Matrix4M21] = tmp[Matrix4M21] * inv_det
	m[Matrix4M22] = tmp[Matrix4M22] * inv_det
	m[Matrix4M23] = tmp[Matrix4M23] * inv_det
	m[Matrix4M30] = tmp[Matrix4M30] * inv_det
	m[Matrix4M31] = tmp[Matrix4M31] * inv_det
	m[Matrix4M32] = tmp[Matrix4M32] * inv_det
	m[Matrix4M33] = tmp[Matrix4M33] * inv_det

	return m, nil
}
