package math

import (
	. "launchpad.net/gocheck"
)

type Vector2TestSuite struct {
	vec Vector2
}

var _ = Suite(&Vector2TestSuite{})

func (s *Vector2TestSuite) Vec2(c *C) {
	s.vec = Vec2(1.23, -3.21)
	c.Assert(s.vec, Equals, Vector2{1.23, -3.21})
}

func (s *Vector2TestSuite) Vector2Set(c *C) {
	s.vec.Set(-2, 0)
	c.Assert(s.vec, Equals, Vec2(-2, 0))
}

func (s *Vector2TestSuite) Vector2SetVec(c *C) {
	s.vec.SetVec2(Vec2(-2, 0))
	c.Assert(s.vec, Equals, Vec2(-2, 0))
}

func (s *Vector2TestSuite) Vector2Cpy(c *C) {
	vec2 := s.vec.Cpy()
	c.Assert(s.vec, Equals, vec2)
}

func (s *Vector2TestSuite) Vector2Len(c *C) {
	s.vec.Set(-2, 0)
	c.Assert(s.vec.Len(), Equals, 2)
}

func (s *Vector2TestSuite) Vector2Len2(c *C) {
	s.vec.Set(-2, 0)
	c.Assert(s.vec.Len2(), Equals, 4)
}

func (s *Vector2TestSuite) Vector2Sub(c *C) {
	s.vec.Set(-2, 0)
	c.Assert(s.vec.Sub(Vec2(-2, -1)), Equals, Vec2(0, 1))
}

func (s *Vector2TestSuite) Vector2Clr(c *C) {
	s.vec.Set(-2, 0)
	c.Assert(s.vec.Clr(), Equals, Vec2(0, 0))
	c.Assert(s.vec, Equals, Vec2(-2, 0))
}

func (s *Vector2TestSuite) Vector2Nor(c *C) {
	s.vec.Set(-2, 0)
	c.Assert(s.vec.Nor(), Equals, Vec2(-1, 0))
}

func (s *Vector2TestSuite) Vector2Add(c *C) {
	s.vec.Set(-2, 0)
	c.Assert(s.vec.Add(Vec2(2, 2)), Equals, Vec2(0, 2))
}

func (s *Vector2TestSuite) Vector2Dot(c *C) {
	s.vec.Set(-2, 0)
	c.Assert(s.vec.Dot(Vec2(-3, 1)), Equals, -1.0)
}

func (s *Vector2TestSuite) Vector2Mul(c *C) {
	s.vec.Set(-2, 0)
	c.Assert(s.vec.Mul(Vec2(-2, 2)), Equals, Vec2(4, 0))
}

func (s *Vector2TestSuite) Vector2Div(c *C) {
	s.vec.Set(-2, 0)
	c.Assert(s.vec.Div(Vec2(-2, 2)), Equals, Vec2(1, 0))
}

func (s *Vector2TestSuite) Vector2Scale(c *C) {
	s.vec.Set(-2, 0)
	c.Assert(s.vec.Scale(2), Equals, Vec2(-4, 0))
}

func (s *Vector2TestSuite) Vector2Distance(c *C) {
	s.vec.Set(-2, 0)
	c.Assert(s.vec.Distance(Vec2(0, 0)), Equals, 2)
}

func (s *Vector2TestSuite) Vector2Distance2(c *C) {
	s.vec.Set(-2, 0)
	c.Assert(s.vec.Distance2(Vec2(0, 0)), Equals, 4)
}

func (s *Vector2TestSuite) Vector2Limit(c *C) {
	s.vec.Set(-2, 0)
	c.Assert(s.vec.Limit(1), Equals, Vec2(-1, 0))
}

// TODO MulMatrix

func (s *Vector2TestSuite) Vector2Cross(c *C) {
	s.vec.Set(5, 1)
	c.Assert(s.vec.Cross(Vec2(-1, 0)), Equals, 1.0)
}

func (s *Vector2TestSuite) Vector2Angle(c *C) {
	s.vec.Set(1, 1)
	c.Assert(s.vec.Angle(), Equals, 45.0)
}

func (s *Vector2TestSuite) Vector2SetAngle(c *C) {
	s.vec.Set(2, 0)
	s.vec.SetAngle(90)
	c.Assert(int(s.vec.X*1000)/1000, Equals, 0)
	c.Assert(s.vec.Y, Equals, 2)
}

func (s *Vector2TestSuite) Vector2Rotate(c *C) {
	var angle float32 = 45
	var x float32 = 5
	var y float32 = -2
	xResult := Cos(angle*DegreeToRadians)*x - Sin(angle*DegreeToRadians)*y
	yResult := Sin(angle*DegreeToRadians)*x + Cos(angle*DegreeToRadians)*y
	s.vec.Set(x, y)

	vec := s.vec.Rotate(angle)
	c.Assert(vec.X, Equals, xResult)
	c.Assert(vec.Y, Equals, yResult)
}

func (s *Vector2TestSuite) Vector2Lerp(c *C) {
	var alpha float32 = 0.5
	v1 := Vec2(1, 1)
	v2 := Vec2(-2, 0)
	v3 := v1.Lerp(v2, alpha)

	xResult := v1.X*(1-alpha) + v2.X*alpha
	yResult := v1.Y*(1-alpha) + v2.Y*alpha

	c.Assert(v3.X, Equals, xResult)
	c.Assert(v3.Y, Equals, yResult)
}

func (s *Vector2TestSuite) Vector2Faceforward(c *C) {
	v := Vec2(1.0, -2.0)
	n := Vec2(0.0, 0.0)
	i := Vec2(2.2, 0.3)

	expected := Vec2(-1.0, 2.0)
	result := v.Faceforward(i, n)
	c.Check(result, EqualsFloat32, expected)
}

// ### Benchmarks ###

func (s *Vector2TestSuite) BenchmarkVector2Add(c *C) {
	vec1 := Vec2(0, 0)
	vec2 := Vec2(1, 1)
	c.ResetTimer()
	for i := 0; i < c.N; i++ {
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

func (s *Vector2TestSuite) BenchmarkVector2NoPointerAdd(c *C) {
	vec1 := Vector2NoPointer{0, 0}
	vec2 := Vector2NoPointer{1, 1}
	c.ResetTimer()
	for i := 0; i < c.N; i++ {
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

func (s *Vector2TestSuite) BenchmarkVector2PointerAdd(c *C) {
	vec1 := &Vector2Pointer{0, 0}
	vec2 := &Vector2Pointer{1, 1}
	c.ResetTimer()
	for i := 0; i < c.N; i++ {
		vec1.Add(vec2)
	}
}
