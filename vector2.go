package math

type Vector2 struct {
	X float32
	Y float32
}

func Vec2(x, y float32) Vector2 {
	return Vector2{X: x, Y: y}
}

func (vec Vector2) Cpy() Vector2 {
	return Vector2{X: vec.X, Y: vec.Y}
}

// The euclidian length
func (vec Vector2) Len() float32 {
	return Sqrt(vec.X*vec.X + vec.Y*vec.Y)
}

// The squared euclidian length
func (vec Vector2) Len2() float32 {
	return vec.X*vec.X + vec.Y*vec.Y
}

func (vec *Vector2) Set(x, y float32) Vector2 {
	vec.X = x
	vec.Y = y
	return *vec
}

func (vec *Vector2) SetVec2(vec2 Vector2) Vector2 {
	vec.X = vec2.X
	vec.Y = vec2.Y
	return *vec
}

func (vec *Vector2) SetVec3(vec2 Vector3) Vector2 {
	vec.X = vec2.X
	vec.Y = vec2.Y
	return *vec
}

func (vec Vector2) Vec3() Vector3 {
	return Vec3(vec.X, vec.Y, 0)
}

func (vec Vector2) Sub(vec2 Vector2) Vector2 {
	vec.X -= vec2.X
	vec.Y -= vec2.Y
	return vec
}

// Returns a zero vector
func (vec Vector2) Clr() Vector2 {
	vec.X = 0
	vec.Y = 0
	return vec
}

// Returns the normalized vector
func (vec Vector2) Nor() Vector2 {
	len := vec.Len()
	if len != 0 {
		vec.X /= len
		vec.Y /= len
	}
	return vec
}

func (vec Vector2) Add(vec2 Vector2) Vector2 {
	vec.X += vec2.X
	vec.Y += vec2.Y
	return vec
}

// Dot returns the dot product of this vector and the given vector.
func (vec Vector2) Dot(vec2 Vector2) float32 {
	return vec.X*vec2.X + vec.Y*vec2.Y
}

func (vec Vector2) Mul(vec2 Vector2) Vector2 {
	vec.X *= vec2.X
	vec.Y *= vec2.Y
	return vec
}

func (vec Vector2) Div(vec2 Vector2) Vector2 {
	vec.X /= vec2.X
	vec.Y /= vec2.Y
	return vec
}

func (vec Vector2) Scale(scale float32) Vector2 {
	vec.X *= scale
	vec.Y *= scale
	return vec
}

// Distance returns the distance between this and the given vector.
func (vec Vector2) Distance(vec2 Vector2) float32 {
	xd := vec2.X - vec.X
	yd := vec2.Y - vec.Y
	return Sqrt(xd*xd + yd*yd)
}

// Distance2 returns the squared distance between this and the given vector.
func (vec Vector2) Distance2(vec2 Vector2) float32 {
	xd := vec2.X - vec.X
	yd := vec2.Y - vec.Y
	return xd*xd + yd*yd
}

// Returns a vector limited to given value based on this vector
func (vec Vector2) Limit(limit float32) Vector2 {
	if vec.Len2() > limit*limit {
		vec = vec.Nor()
		vec = vec.Scale(limit)
	}
	return vec
}

func (vec Vector2) MulMatrix(m *Matrix3) Vector2 {
	vec.X = vec.X*m.M11 + vec.Y*m.M21 + m.M31
	vec.Y = vec.X*m.M12 + vec.Y*m.M22 + m.M32
	return vec
}

// Cross returns the cross product of this vector ang the given vector.
func (vec Vector2) Cross(vec2 Vector2) float32 {
	return vec.X*vec2.Y - vec.Y*vec2.X
}

func (vec Vector2) Angle() float32 {
	angle := Atan2(vec.Y, vec.X) * RadiansToDegrees
	if angle < 0 {
		angle += 360
	}
	return angle
}

func (vec *Vector2) SetAngle(angle float32) Vector2 {
	vec.X = vec.Len()
	vec.Y = 0.0
	v := vec.Rotate(angle)
	return vec.SetVec2(v)
}

// Returns the rotated Vector2 by the given angle, counter-clockwise.
func (vec Vector2) Rotate(degrees float32) Vector2 {
	rad := degrees * DegreeToRadians
	cos := Cos(rad)
	sin := Sin(rad)

	x := vec.X*cos - vec.Y*sin
	y := vec.X*sin + vec.Y*cos
	vec.X = x
	vec.Y = y
	return vec
}

// Lerp returns the linearly interpolates between this vector and the target vector by alpha which is in the range [0,1].
func (vec Vector2) Lerp(target Vector2, alpha float32) Vector2 {
	invAlpha := 1.0 - alpha
	vec.X = vec.X*invAlpha + target.X*alpha
	vec.Y = vec.Y*invAlpha + target.Y*alpha
	return vec
}

// Faceforward returns this vector if n,Dot(i) < 0, otherwise, returns the negative of this vector.
func (vec Vector2) Faceforward(i, n Vector2) Vector2 {
	if n.Dot(i) < 0 {
		return vec
	}
	return Vector2{-vec.X, -vec.Y}
}
