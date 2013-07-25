package math

import (
	"errors"
)

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

func NewFrustumMatrix4(left, right, bottom, top, near, far float32) *Matrix4 {

	//x:= 2.0 * near / ( right - left )
	//y:= 2.0 * near / ( top - bottom )

	//a:= ( right + left ) / ( right - left )
	//b:= ( top + bottom ) / ( top - bottom )
	//c:= - ( far + near ) / ( far - near )
	//d:= - 2.0 * far * near / ( far - near )

	return &Matrix4{2.0 * near / (right - left), 0, 0, 0,
		0, 2.0 * near / (top - bottom), 0, 0,
		(right + left) / (right - left), (top + bottom) / (top - bottom), -(far + near) / (far - near), -1,
		0, 0, -2.0 * far * near / (far - near), 0}

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

func (m *Matrix4) MulVec4(vec Vector4) Vector4 {
	vec.X = vec.X*m.M11 + vec.Y*m.M21 + vec.Z*m.M31 + vec.W*m.M41
	vec.Y = vec.X*m.M12 + vec.Y*m.M22 + vec.Z*m.M32 + vec.W*m.M42
	vec.Z = vec.X*m.M13 + vec.Y*m.M23 + vec.Z*m.M33 + vec.W*m.M43
	vec.W = vec.X*m.M14 + vec.Y*m.M24 + vec.Z*m.M34 + vec.W*m.M44
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

//return the inverted matrix
func (m *Matrix4) Inverted() (*Matrix4, error) {
	det := m.Determinant()
	if det == 0 {
		return nil, errors.New("non-invertible matrix")
	}

	tmp := &Matrix4{}

	tmp.M11 = m.M32*m.M43*m.M24 - m.M42*m.M33*m.M24 + m.M42*m.M23*m.M34 - m.M22*m.M43*m.M34 - m.M32*m.M23*m.M44 + m.M22*m.M33*m.M44
	tmp.M21 = m.M41*m.M33*m.M24 - m.M31*m.M43*m.M24 - m.M41*m.M23*m.M34 + m.M21*m.M43*m.M34 + m.M31*m.M23*m.M44 - m.M21*m.M33*m.M44
	tmp.M31 = m.M31*m.M42*m.M24 - m.M41*m.M32*m.M24 + m.M41*m.M22*m.M34 - m.M21*m.M42*m.M34 - m.M31*m.M22*m.M44 + m.M21*m.M32*m.M44
	tmp.M41 = m.M41*m.M32*m.M23 - m.M31*m.M42*m.M23 - m.M41*m.M22*m.M33 + m.M21*m.M42*m.M33 + m.M31*m.M22*m.M43 - m.M21*m.M32*m.M43
	tmp.M12 = m.M42*m.M33*m.M14 - m.M32*m.M43*m.M14 - m.M42*m.M13*m.M34 + m.M12*m.M43*m.M34 + m.M32*m.M13*m.M44 - m.M12*m.M33*m.M44
	tmp.M22 = m.M31*m.M43*m.M14 - m.M41*m.M33*m.M14 + m.M41*m.M13*m.M34 - m.M11*m.M43*m.M34 - m.M31*m.M13*m.M44 + m.M11*m.M33*m.M44
	tmp.M32 = m.M41*m.M32*m.M14 - m.M31*m.M42*m.M14 - m.M41*m.M12*m.M34 + m.M11*m.M42*m.M34 + m.M31*m.M12*m.M44 - m.M11*m.M32*m.M44
	tmp.M42 = m.M31*m.M42*m.M13 - m.M41*m.M32*m.M13 + m.M41*m.M12*m.M33 - m.M11*m.M42*m.M33 - m.M31*m.M12*m.M43 + m.M11*m.M32*m.M43
	tmp.M13 = m.M22*m.M43*m.M14 - m.M42*m.M23*m.M14 + m.M42*m.M13*m.M24 - m.M12*m.M43*m.M24 - m.M22*m.M13*m.M44 + m.M12*m.M23*m.M44
	tmp.M23 = m.M41*m.M23*m.M14 - m.M21*m.M43*m.M14 - m.M41*m.M13*m.M24 + m.M11*m.M43*m.M24 + m.M21*m.M13*m.M44 - m.M11*m.M23*m.M44
	tmp.M33 = m.M21*m.M42*m.M14 - m.M41*m.M22*m.M14 + m.M41*m.M12*m.M24 - m.M11*m.M42*m.M24 - m.M21*m.M12*m.M44 + m.M11*m.M22*m.M44
	tmp.M43 = m.M41*m.M22*m.M13 - m.M21*m.M42*m.M13 - m.M41*m.M12*m.M23 + m.M11*m.M42*m.M23 + m.M21*m.M12*m.M43 - m.M11*m.M22*m.M43
	tmp.M14 = m.M32*m.M23*m.M14 - m.M22*m.M33*m.M14 - m.M32*m.M13*m.M24 + m.M12*m.M33*m.M24 + m.M22*m.M13*m.M34 - m.M12*m.M23*m.M34
	tmp.M24 = m.M21*m.M33*m.M14 - m.M31*m.M23*m.M14 + m.M31*m.M13*m.M24 - m.M11*m.M33*m.M24 - m.M21*m.M13*m.M34 + m.M11*m.M23*m.M34
	tmp.M34 = m.M31*m.M22*m.M14 - m.M21*m.M32*m.M14 - m.M31*m.M12*m.M24 + m.M11*m.M32*m.M24 + m.M21*m.M12*m.M34 - m.M11*m.M22*m.M34
	tmp.M44 = m.M21*m.M32*m.M13 - m.M31*m.M22*m.M13 + m.M31*m.M12*m.M23 - m.M11*m.M32*m.M23 - m.M21*m.M12*m.M33 + m.M11*m.M22*m.M33

	inv_det := 1.0 / det
	tmp.M11 *= inv_det
	tmp.M21 *= inv_det
	tmp.M31 *= inv_det
	tmp.M41 *= inv_det
	tmp.M12 *= inv_det
	tmp.M22 *= inv_det
	tmp.M32 *= inv_det
	tmp.M42 *= inv_det
	tmp.M13 *= inv_det
	tmp.M23 *= inv_det
	tmp.M33 *= inv_det
	tmp.M43 *= inv_det
	tmp.M14 *= inv_det
	tmp.M24 *= inv_det
	tmp.M34 *= inv_det
	tmp.M44 *= inv_det

	return tmp, nil
}

//inverts the matrix and returns it
func (m *Matrix4) Invert() (*Matrix4, error) {
	det := m.Determinant()
	if det == 0 {
		return nil, errors.New("non-invertible matrix")
	}

	tmp := &Matrix4{}

	tmp.M11 = m.M32*m.M43*m.M24 - m.M42*m.M33*m.M24 + m.M42*m.M23*m.M34 - m.M22*m.M43*m.M34 - m.M32*m.M23*m.M44 + m.M22*m.M33*m.M44
	tmp.M21 = m.M41*m.M33*m.M24 - m.M31*m.M43*m.M24 - m.M41*m.M23*m.M34 + m.M21*m.M43*m.M34 + m.M31*m.M23*m.M44 - m.M21*m.M33*m.M44
	tmp.M31 = m.M31*m.M42*m.M24 - m.M41*m.M32*m.M24 + m.M41*m.M22*m.M34 - m.M21*m.M42*m.M34 - m.M31*m.M22*m.M44 + m.M21*m.M32*m.M44
	tmp.M41 = m.M41*m.M32*m.M23 - m.M31*m.M42*m.M23 - m.M41*m.M22*m.M33 + m.M21*m.M42*m.M33 + m.M31*m.M22*m.M43 - m.M21*m.M32*m.M43
	tmp.M12 = m.M42*m.M33*m.M14 - m.M32*m.M43*m.M14 - m.M42*m.M13*m.M34 + m.M12*m.M43*m.M34 + m.M32*m.M13*m.M44 - m.M12*m.M33*m.M44
	tmp.M22 = m.M31*m.M43*m.M14 - m.M41*m.M33*m.M14 + m.M41*m.M13*m.M34 - m.M11*m.M43*m.M34 - m.M31*m.M13*m.M44 + m.M11*m.M33*m.M44
	tmp.M32 = m.M41*m.M32*m.M14 - m.M31*m.M42*m.M14 - m.M41*m.M12*m.M34 + m.M11*m.M42*m.M34 + m.M31*m.M12*m.M44 - m.M11*m.M32*m.M44
	tmp.M42 = m.M31*m.M42*m.M13 - m.M41*m.M32*m.M13 + m.M41*m.M12*m.M33 - m.M11*m.M42*m.M33 - m.M31*m.M12*m.M43 + m.M11*m.M32*m.M43
	tmp.M13 = m.M22*m.M43*m.M14 - m.M42*m.M23*m.M14 + m.M42*m.M13*m.M24 - m.M12*m.M43*m.M24 - m.M22*m.M13*m.M44 + m.M12*m.M23*m.M44
	tmp.M23 = m.M41*m.M23*m.M14 - m.M21*m.M43*m.M14 - m.M41*m.M13*m.M24 + m.M11*m.M43*m.M24 + m.M21*m.M13*m.M44 - m.M11*m.M23*m.M44
	tmp.M33 = m.M21*m.M42*m.M14 - m.M41*m.M22*m.M14 + m.M41*m.M12*m.M24 - m.M11*m.M42*m.M24 - m.M21*m.M12*m.M44 + m.M11*m.M22*m.M44
	tmp.M43 = m.M41*m.M22*m.M13 - m.M21*m.M42*m.M13 - m.M41*m.M12*m.M23 + m.M11*m.M42*m.M23 + m.M21*m.M12*m.M43 - m.M11*m.M22*m.M43
	tmp.M14 = m.M32*m.M23*m.M14 - m.M22*m.M33*m.M14 - m.M32*m.M13*m.M24 + m.M12*m.M33*m.M24 + m.M22*m.M13*m.M34 - m.M12*m.M23*m.M34
	tmp.M24 = m.M21*m.M33*m.M14 - m.M31*m.M23*m.M14 + m.M31*m.M13*m.M24 - m.M11*m.M33*m.M24 - m.M21*m.M13*m.M34 + m.M11*m.M23*m.M34
	tmp.M34 = m.M31*m.M22*m.M14 - m.M21*m.M32*m.M14 - m.M31*m.M12*m.M24 + m.M11*m.M32*m.M24 + m.M21*m.M12*m.M34 - m.M11*m.M22*m.M34
	tmp.M44 = m.M21*m.M32*m.M13 - m.M31*m.M22*m.M13 + m.M31*m.M12*m.M23 - m.M11*m.M32*m.M23 - m.M21*m.M12*m.M33 + m.M11*m.M22*m.M33

	inv_det := 1.0 / det
	m.M11 = tmp.M11 * inv_det
	m.M21 = tmp.M21 * inv_det
	m.M31 = tmp.M31 * inv_det
	m.M41 = tmp.M41 * inv_det
	m.M12 = tmp.M12 * inv_det
	m.M22 = tmp.M22 * inv_det
	m.M32 = tmp.M32 * inv_det
	m.M42 = tmp.M42 * inv_det
	m.M13 = tmp.M13 * inv_det
	m.M23 = tmp.M23 * inv_det
	m.M33 = tmp.M33 * inv_det
	m.M43 = tmp.M43 * inv_det
	m.M14 = tmp.M14 * inv_det
	m.M24 = tmp.M24 * inv_det
	m.M34 = tmp.M34 * inv_det
	m.M44 = tmp.M44 * inv_det

	return m, nil
}

// The determinant of this matrix.
func (m *Matrix4) Determinant() float32 {
	return m.M14*m.M23*m.M32*m.M41 -
		m.M13*m.M24*m.M32*m.M41 -
		m.M14*m.M22*m.M33*m.M41 +
		m.M12*m.M24*m.M33*m.M41 +
		m.M13*m.M22*m.M34*m.M41 -
		m.M12*m.M23*m.M34*m.M41 -
		m.M14*m.M23*m.M31*m.M41 +
		m.M13*m.M24*m.M31*m.M41 +
		m.M14*m.M21*m.M33*m.M41 -
		m.M11*m.M24*m.M33*m.M41 -
		m.M13*m.M21*m.M34*m.M41 +
		m.M11*m.M23*m.M34*m.M41 +
		m.M14*m.M22*m.M31*m.M43 -
		m.M12*m.M24*m.M31*m.M43 -
		m.M14*m.M21*m.M32*m.M43 +
		m.M11*m.M24*m.M32*m.M43 +
		m.M12*m.M21*m.M34*m.M43 -
		m.M11*m.M22*m.M34*m.M43 -
		m.M13*m.M22*m.M31*m.M44 +
		m.M12*m.M23*m.M31*m.M44 +
		m.M13*m.M21*m.M32*m.M44 -
		m.M11*m.M23*m.M32*m.M44 -
		m.M12*m.M21*m.M33*m.M44 +
		m.M11*m.M22*m.M33*m.M44
}

// Equal to gluProject
func Project(obj Vector3, modelview, projection *Matrix4, viewport Vector4) Vector3 {
	// Modelview transform
	ft0 := modelview.M11*obj.X + modelview.M21*obj.Y + modelview.M31*obj.Z + modelview.M41
	ft1 := modelview.M12*obj.X + modelview.M22*obj.Y + modelview.M32*obj.Z + modelview.M42
	ft2 := modelview.M13*obj.X + modelview.M23*obj.Y + modelview.M33*obj.Z + modelview.M43
	ft3 := modelview.M14*obj.X + modelview.M24*obj.Y + modelview.M34*obj.Z + modelview.M44

	// Projection transform, the final row of projection matrix is always [0,0,-1,0]
	// so we optimize for that.
	ft4 := projection.M11*ft0 + projection.M21*ft1 + projection.M31*ft2 + projection.M41*ft3
	ft5 := projection.M12*ft0 + projection.M22*ft1 + projection.M32*ft2 + projection.M42*ft3
	ft6 := projection.M13*ft0 + projection.M23*ft1 + projection.M33*ft2 + projection.M43*ft3
	ft7 := -ft2
	// The result normalizes between -1 and 1
	if ft7 == 0.0 { // The w value
		return Vec3(0, 0, 0)
	}
	ft7 = 1.0 / ft7

	// Perspective division
	ft4 *= ft7
	ft5 *= ft7
	ft6 *= ft7

	// Window coordinates
	// Map x, y to range 0-1
	x := (ft4*0.5+0.5)*viewport.Z + viewport.X
	y := (ft5*0.5+0.5)*viewport.W + viewport.Y
	z := (1.0 + ft6) * 0.5
	return Vec3(x, y, z)
}

func UnProject(window Vector3, modelview, projection *Matrix4, viewport Vector4) (Vector3, error) {
	a := projection.Mul(modelview)

	// Compute the inverse of matrix a
	inverse, err := a.Invert()
	if err != nil {
		return Vec3(0, 0, 0), err
	}

	tmp := Vec4(window.X, window.Y, window.Z, 1)
	tmp.X = (tmp.X - viewport.X) / viewport.Z
	tmp.Y = (tmp.Y - viewport.Y) / viewport.W
	tmp = tmp.Scale(2).Sub(Vec4(1, 1, 1, 1))

	obj := inverse.MulVec4(tmp)
	obj = obj.Scale(1.0 / obj.W)
	return Vec3(obj.X, obj.Y, obj.Z), nil
}

//compose a Matrix4 given a position (Vector3) a rotation (Quaternion) and a scale (Vector3)
//shameless inspired from three.js
func (m *Matrix4) Compose(position Vector3, rotation Quaternion, scale Vector3) *Matrix4 {
	//set rotation
	m = rotation.Matrix()
	//set scale
	m.M11 *= scale.X
	m.M12 *= scale.X
	m.M13 *= scale.X
	m.M21 *= scale.Y
	m.M22 *= scale.Y
	m.M23 *= scale.Y
	m.M31 *= scale.Z
	m.M32 *= scale.Z
	m.M33 *= scale.Z
	m.M41 = position.X
	m.M42 = position.Y
	m.M43 = position.Z
	return m
}

//Extract the pure rotation matrix filling the one passed as argument or
//alocating anotehr one
func (m *Matrix4) RotMatrix(rot *Matrix4) *Matrix4 {
	if rot == nil {
		rot = new(Matrix4)
	}

	v := Vector3{m.M11, m.M12, m.M13}
	//v.X, v.Y, v.Z = m.M11, m.M12, m.M13
	scaleX := 1.0 / v.Len()
	v.X, v.Y, v.Z = m.M21, m.M22, m.M23
	scaleY := 1.0 / v.Len()
	v.X, v.Y, v.Z = m.M31, m.M32, m.M33
	scaleZ := 1.0 / v.Len()

	rot.M11 = m.M11 * scaleX
	rot.M12 = m.M12 * scaleX
	rot.M13 = m.M13 * scaleX

	rot.M21 = m.M21 * scaleY
	rot.M22 = m.M22 * scaleY
	rot.M23 = m.M23 * scaleY

	rot.M31 = m.M31 * scaleZ
	rot.M32 = m.M32 * scaleZ
	rot.M33 = m.M33 * scaleZ

	rot.M44 = 1 //not sure about this...

	return rot

}

//extract the scale factors into a Vector3
func (m *Matrix4) ScaleVector(v *Vector3) *Vector3 {
	if v == nil {
		v = new(Vector3)
	}
	v.X, v.Y, v.Z = m.M11, m.M12, m.M13
	scaleX := v.Len()
	v.X, v.Y, v.Z = m.M21, m.M22, m.M23
	scaleY := v.Len()
	v.X, v.Y, v.Z = m.M31, m.M32, m.M33
	scaleZ := v.Len()
	v.X, v.Y, v.Z = scaleX, scaleY, scaleZ

	return v
}
