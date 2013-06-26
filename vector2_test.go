package math

import (
	"testing"
)

func TestVec2(t *testing.T) {
	var x float32 = 1.23
	var y float32 = -3.21
	vec := Vec2(x, y)
	if x != vec.X {
		t.Errorf("%+v should be equal to %+v", x, vec.X)
	}
	if y != vec.Y {
		t.Errorf("%+v should be equal to %+v", y, vec.Y)
	}
}

func TestVec2SetVec2(t *testing.T) {
	var x float32 = -2
	var y float32 = 0
	vec := Vec2(x, y)
	vec.SetVec2(Vec2(1, 1))
	if vec.X != 1 || vec.Y != 1 {
		t.Errorf("%+v should be equal to Vector2(1,1)", vec)
	}
}

func TestVec2Set(t *testing.T) {
	var x float32 = -2
	var y float32 = 0
	vec := Vec2(x, y)
	vec.Set(1, 1)
	if vec.X != 1 || vec.Y != 1 {
		t.Errorf("%+v should be equal to Vector2(1,1)", vec)
	}
}

func TestVec2Cpy(t *testing.T) {
	var x float32 = 1.23
	var y float32 = -3.21
	vec := Vec2(x, y)
	vec2 := vec.Cpy()
	if vec != vec2 {
		t.Errorf("%+v should be equal to %+v", vec, vec2)
	}
}

func TestVec2Len(t *testing.T) {
	var x float32 = -2
	var y float32 = 0
	vec := Vec2(x, y)
	length := vec.Len()
	if length != 2 {
		t.Errorf("Length %+v should be equal to 2", length)
	}
}

func TestVec2Len2(t *testing.T) {
	var x float32 = -2
	var y float32 = 0
	vec := Vec2(x, y)
	length := vec.Len2()
	if length != 4 {
		t.Errorf("Length %+v should be equal to 4", length)
	}
}

func TestVec2Sub(t *testing.T) {
	var x float32 = -2
	var y float32 = 0
	vec := Vec2(x, y)
	vec2 := Vec2(-2, -1)
	vec = vec.Sub(vec2)
	if vec.X != 0 || vec.Y != 1 {
		t.Errorf("%+v should be equal to Vector2(0,1)", vec)
	}
}

func TestVec2Clr(t *testing.T) {
	var x float32 = -2
	var y float32 = 0
	vec := Vec2(x, y)
	vec = vec.Clr()
	if vec.X != 0 || vec.Y != 0 {
		t.Errorf("%+v should be equal to Vector2(0,0)", vec)
	}
}

func TestVec2Nor(t *testing.T) {
	var x float32 = -2
	var y float32 = 0
	vec := Vec2(x, y)
	vec = vec.Nor()
	if vec.X != -1 || vec.Y != 0 {
		t.Errorf("%+v should be equal to Vector2(-1,0)", vec)
	}
}

func TestVec2Add(t *testing.T) {
	var x float32 = -2
	var y float32 = 0
	vec := Vec2(x, y)
	vec2 := Vec2(2, 2)
	vec = vec.Add(vec2)
	if vec.X != 0 || vec.Y != 2 {
		t.Errorf("%+v should be equal to Vector2(0,2)", vec)
	}
}

func TestVec2Dot(t *testing.T) {
	a := Vec2(2, 5)
	b := Vec2(-3, 1)
	dot := a.Dot(b)
	if dot != -1.0 {
		t.Errorf("%g should be equal to %g", dot, -1.0)
	}
}

func TestVec2Mul(t *testing.T) {
	var x float32 = -2
	var y float32 = 0
	vec := Vec2(x, y)
	vec2 := Vec2(-2, 2)
	vec = vec.Mul(vec2)
	if vec.X != 4 || vec.Y != 0 {
		t.Errorf("%+v should be equal to Vector2(4,0)", vec)
	}
}

func TestVec2Div(t *testing.T) {
	var x float32 = -2
	var y float32 = 0
	vec := Vec2(x, y)
	vec2 := Vec2(-2, 2)
	vec = vec.Div(vec2)
	if vec.X != 1 || vec.Y != 0 {
		t.Errorf("%+v should be equal to Vector2(1,0)", vec)
	}
}

func TestVec2Scale(t *testing.T) {
	var x float32 = -2
	var y float32 = 0
	vec := Vec2(x, y)
	vec = vec.Scale(2)
	if vec.X != -4 || vec.Y != 0 {
		t.Errorf("%+v should be equal to Vector2(-4,0)", vec)
	}
}

func TestVec2Dst(t *testing.T) {
	var x float32 = -2
	var y float32 = 0
	vec := Vec2(x, y)
	vec2 := Vec2(0, 0)
	dst := vec.Dst(vec2)
	if dst != 2 {
		t.Errorf("The distance between %+v and %+v should be %+v", vec, vec2, 2)
	}
}

func TestVec2Dst2(t *testing.T) {
	var x float32 = -2
	var y float32 = 0
	vec := Vec2(x, y)
	vec2 := Vec2(0, 0)
	dst := vec.Dst2(vec2)
	if dst != 4 {
		t.Errorf("The distance between %+v and %+v should be %+v", vec, vec2, 4)
	}
}

func TestVec2Limit(t *testing.T) {
	var x float32 = -2
	var y float32 = 0
	vec := Vec2(x, y)
	vec = vec.Limit(1)
	if vec.X != -1 || vec.Y != 0 {
		t.Errorf("%+v should be equal to Vector2(-1,0)", vec)
	}
}

// TODO MulMatrix

func TestVec2Crs(t *testing.T) {
	a := Vec2(5, 1)
	b := Vec2(-1, 0)
	crs := a.Crs(b)
	if crs != 1.0 {
		t.Errorf("%+v should be equal to %g", crs, 1.0)
	}
}

func TestVec2Angle(t *testing.T) {
	a := Vec2(1, 1)
	angle := a.Angle()
	if angle != 45.0 {
		t.Errorf("%g should be equal to %g", angle, 45.0)
	}
}

func TestVec2SetAngle(t *testing.T) {
	a := Vec2(2, 0)
	b := a.SetAngle(90)
	if int(b.X*1000)/1000 != 0 || b.Y != 2 {
		t.Errorf("%+v should be equal to Vector2(X:%g,Y:%g)", b, 0, 2)
	}
}

func TestVec2Rotate(t *testing.T) {
	var angle float32 = 45
	var x float32 = 5
	var y float32 = -2
	xResult := Cos(angle*DegreeToRadians)*x - Sin(angle*DegreeToRadians)*y
	yResult := Sin(angle*DegreeToRadians)*x + Cos(angle*DegreeToRadians)*y

	a := Vec2(x, y)
	b := a.Rotate(angle)
	if b.X != xResult || b.Y != yResult {
		t.Errorf("%+v should be equal to Vector2(X:%g,Y:%g)", b, xResult, yResult)
	}
}

func TestVec2Lerp(t *testing.T) {
	var alpha float32 = 0.5
	a := Vec2(1, 1)
	b := Vec2(-2, 0)
	c := a.Lerp(b, alpha)

	xResult := a.X*(1-alpha) + b.X*alpha
	yResult := a.Y*(1-alpha) + b.Y*alpha
	if c.X != xResult || c.Y != yResult {
		t.Errorf("%+v should be equal to Vector2(X:%g,Y:%g)", c, xResult, yResult)
	}
}

// ### Benchmarks ###

func BenchmarkVector2Add(b *testing.B) {
	vec1 := Vec2(0, 0)
	vec2 := Vec2(1, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		vec1.Add(vec2)
	}
}

// Benchmarking pointer vs no pointer for Vector2
// NOTE: A pointer uses 64bit on a 64bit system.
// NoPointer is faster since less code is copied into the function (only the vector2 value)
// but Pointer requires less memory usage.

type Vector2NoPointer struct {
	X, Y float32
}

func (vec Vector2NoPointer) Add(vec2 Vector2NoPointer) Vector2NoPointer {
	vec.X += vec2.X
	vec.Y += vec2.Y
	return vec
}

func BenchmarkVector2NoPointerAdd(b *testing.B) {
	vec1 := Vector2NoPointer{0, 0}
	vec2 := Vector2NoPointer{1, 1}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		vec1.Add(vec2)
	}
}

type Vector2Pointer struct {
	X, Y float32
}

func (vec *Vector2Pointer) Add(vec2 *Vector2Pointer) *Vector2Pointer {
	vec.X += vec2.X
	vec.Y += vec2.Y
	return vec
}

func BenchmarkVector2PointerAdd(b *testing.B) {
	vec1 := &Vector2Pointer{0, 0}
	vec2 := &Vector2Pointer{1, 1}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		vec1.Add(vec2)
	}
}
