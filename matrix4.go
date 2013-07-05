package math

type Matrix4 struct {
	M11, M12, M13, M14 float32
	M21, M22, M23, M24 float32
	M31, M32, M33, M34 float32
	M41, M42, M43, M44 float32
}

func NewMatrix4() *Matrix4 {
	return &Matrix4{}
}

func NewIdentityMatrix4() *Matrix4 {
	return &Matrix4{
		M11: 1.0,
		M22: 1.0,
		M33: 1.0,
		M44: 1.0,
	}
}

func NewPerspectiveMatrix4(fovy, aspectRatio, near, far float32) *Matrix4 {
	fovy = fovy * DegreeToRadians
	nmf := near - far
	f := 1.0 / Tan(fovy/2)
	return &Matrix4{
		f / aspectRatio, 0, 0, 0,
		0, f, 0, 0,
		0, 0, (near + far) / nmf, -1,
		0, 0, (2 * far * near) / nmf, 0,
	}
}

func NewTranslationMatrix4(x, y, z float32) *Matrix4 {
	return &Matrix4{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, x, y, z, 1}
}

// LookAt Matrix right hand
func NewLookAtMatrix4(eye, center, up Vector3) *Matrix4 {
	zAxis := (eye.Sub(center)).Nor()
	xAxis := (up.Crs(zAxis)).Nor()
	yAxis := zAxis.Crs(xAxis)

	return &Matrix4{
		xAxis.X, yAxis.X, zAxis.X, 0,
		xAxis.Y, yAxis.Y, zAxis.Y, 0,
		xAxis.Z, yAxis.Z, zAxis.Z, 0,
		-(xAxis.Dot(eye)), -(yAxis.Dot(eye)), -(zAxis.Dot(eye)), 1,
	}
}

func NewRotationMatrix4(axis Vector3, angle float32) *Matrix4 {
	axis = axis.Nor()
	angle = DegreeToRadians * angle

	c := Cos(angle)
	s := Sin(angle)
	k := 1 - c

	return &Matrix4{axis.X*axis.X*k + c, axis.X*axis.Y*k + axis.Z*s, axis.X*axis.Z*k - axis.Y*s, 0,
		axis.X*axis.Y*k - axis.Z*s, axis.Y*axis.Y*k + c, axis.Y*axis.Z*k + axis.X*s, 0,
		axis.X*axis.Z*k + axis.Y*s, axis.Y*axis.Z*k - axis.X*s, axis.Z*axis.Z*k + c, 0,
		0, 0, 0, 1}

}

func NewOrthoMatrix4(left, right, bottom, top, near, far float32) *Matrix4 {
	xOrtho := 2 / (right - left)
	yOrtho := 2 / (top - bottom)
	zOrtho := -2 / (far - near)

	tx := -(right + left) / (right - left)
	ty := -(top + bottom) / (top - bottom)
	tz := -(far + near) / (far - near)
	return &Matrix4{M11: xOrtho, M22: yOrtho, M33: zOrtho, M41: tx, M42: ty, M43: tz, M44: 1}
}

func (m1 *Matrix4) Set(m2 *Matrix4) *Matrix4 {
	(*m1) = (*m2)
	return m1
}

// Multiplicates this matrix with m2 matrix and returns the new matrix.
func (m1 *Matrix4) Mul(m2 *Matrix4) *Matrix4 {
	temp := &Matrix4{
		m1.M11*m2.M11 + m1.M21*m2.M12 + m1.M31*m2.M13 + m1.M41*m2.M14,
		m1.M12*m2.M11 + m1.M22*m2.M12 + m1.M32*m2.M13 + m1.M42*m2.M14,
		m1.M13*m2.M11 + m1.M23*m2.M12 + m1.M33*m2.M13 + m1.M43*m2.M14,
		m1.M14*m2.M11 + m1.M24*m2.M12 + m1.M34*m2.M13 + m1.M44*m2.M14,
		m1.M11*m2.M21 + m1.M21*m2.M22 + m1.M31*m2.M23 + m1.M41*m2.M24,
		m1.M12*m2.M21 + m1.M22*m2.M22 + m1.M32*m2.M23 + m1.M42*m2.M24,
		m1.M13*m2.M21 + m1.M23*m2.M22 + m1.M33*m2.M23 + m1.M43*m2.M24,
		m1.M14*m2.M21 + m1.M24*m2.M22 + m1.M34*m2.M23 + m1.M44*m2.M24,
		m1.M11*m2.M31 + m1.M21*m2.M32 + m1.M31*m2.M33 + m1.M41*m2.M34,
		m1.M12*m2.M31 + m1.M22*m2.M32 + m1.M32*m2.M33 + m1.M42*m2.M34,
		m1.M13*m2.M31 + m1.M23*m2.M32 + m1.M33*m2.M33 + m1.M43*m2.M34,
		m1.M14*m2.M31 + m1.M24*m2.M32 + m1.M34*m2.M33 + m1.M44*m2.M34,
		m1.M11*m2.M41 + m1.M21*m2.M42 + m1.M31*m2.M43 + m1.M41*m2.M44,
		m1.M12*m2.M41 + m1.M22*m2.M42 + m1.M32*m2.M43 + m1.M42*m2.M44,
		m1.M13*m2.M41 + m1.M23*m2.M42 + m1.M33*m2.M43 + m1.M43*m2.M44,
		m1.M14*m2.M41 + m1.M24*m2.M42 + m1.M34*m2.M43 + m1.M44*m2.M44}
	return temp
}

func (m *Matrix4) MulVec3(vec Vector3) Vector3 {
	vec.X = vec.X*m.M11 + vec.Y*m.M21 + vec.Z*m.M31 + m.M41
	vec.Y = vec.X*m.M12 + vec.Y*m.M22 + vec.Z*m.M32 + m.M42
	vec.Z = vec.X*m.M13 + vec.Y*m.M23 + vec.Z*m.M33 + m.M43
	return vec
}

func (m *Matrix4) Scale(scalar Vector3) *Matrix4 {
	s := &Matrix4{
		M11: scalar.X,
		M22: scalar.Y,
		M33: scalar.Z,
		M44: 1,
	}
	return m.Mul(s)
}
