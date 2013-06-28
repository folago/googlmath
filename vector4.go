package math

type Vector4 struct {
	X float32
	Y float32
	Z float32
	W float32
}

func Vec4(x, y, z, w float32) Vector4 {
	return Vector4{x, y, z, w}
}

func (vec *Vector4) Set(x, y, z, w float32) Vector4 {
	vec.X = x
	vec.Y = y
	vec.Z = z
	vec.W = w
	return *vec
}

func (vec *Vector4) SetVec2(v Vector2) Vector4 {
	vec.X = v.X
	vec.Y = v.Y
	return *vec
}

func (vec *Vector4) SetVec3(v Vector3) Vector4 {
	vec.X = v.X
	vec.Y = v.Y
	vec.Z = v.Z
	return *vec
}

func (vec Vector4) Vec2() Vector2 {
	return Vec2(vec.X, vec.Y)
}

func (vec Vector4) Vec3() Vector3 {
	return Vec3(vec.X, vec.Y, vec.Z)
}

func (vec Vector4) Cpy() Vector4 {
	return Vector4{vec.X, vec.Y, vec.Z, vec.W}
}

// Returns a zero vector
func (vec Vector4) Clr() Vector4 {
	vec.X = 0
	vec.Y = 0
	vec.Z = 0
	vec.W = 0
	return vec
}

func (vec Vector4) Add(vec2 Vector4) Vector4 {
	vec.X += vec2.X
	vec.Y += vec2.Y
	vec.Z += vec2.Z
	vec.W += vec2.W
	return vec
}

func (vec Vector4) Sub(vec2 Vector4) Vector4 {
	vec.X -= vec2.X
	vec.Y -= vec2.Y
	vec.Z -= vec2.Z
	vec.W -= vec2.W
	return vec
}

func (vec Vector4) Mul(vec2 Vector4) Vector4 {
	vec.X *= vec2.X
	vec.Y *= vec2.Y
	vec.Z *= vec2.Z
	vec.W *= vec2.W
	return vec
}

func (vec Vector4) Div(vec2 Vector4) Vector4 {
	vec.X /= vec2.X
	vec.Y /= vec2.Y
	vec.Z /= vec2.Z
	vec.W /= vec2.W
	return vec
}

func (vec Vector4) IsZero() bool {
	return vec.X == 0 && vec.Y == 0 && vec.Z == 0 && vec.W == 0
}

func (vec Vector4) Scale(scalar float32) Vector4 {
	vec.X *= scalar
	vec.Y *= scalar
	vec.Z *= scalar
	vec.W *= scalar
	return vec
}

func (vec Vector4) Invert() Vector4 {
	vec.X = -vec.X
	vec.Y = -vec.Y
	vec.Z = -vec.Z
	vec.W = -vec.W
	return vec
}
