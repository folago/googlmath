package math

import (
	. "launchpad.net/gocheck"
)

type Vector3TestSuite struct {
	vec Vector3
}

var _ = Suite(&Vector3TestSuite{})

func (s *Vector3TestSuite) Vec3(c *C) {
	s.vec = Vec3(1.23, -3.21, -0.0)
	c.Assert(s.vec, Equals, Vector3{1.23, -3.21, -0.0})
}

func (s *Vector3TestSuite) Vector3Set(c *C) {
	s.vec.Set(-2, 0, 1)
	c.Assert(s.vec, Equals, Vector3{-2, 0, 1})
}

func (s *Vector3TestSuite) Vector3SetVec3(c *C) {
	s.vec.SetVec3(Vec3(-2, 0, 1))
	c.Assert(s.vec, Equals, Vector3{-2, 0, 1})
}

func (s *Vector3TestSuite) Vector3Cpy(c *C) {
	s.vec.SetVec3(Vec3(-2, 0, 1))
	vec := s.vec.Cpy()
	c.Assert(s.vec, Equals, vec)
}

func (s *Vector3TestSuite) Vector3Clr(c *C) {
	s.vec.SetVec3(Vec3(-2, 0, 1))
	vec := s.vec.Clr()
	c.Assert(vec, Equals, Vector3{0, 0, 0})
}

func (s *Vector3TestSuite) Vector3Add(c *C) {
	s.vec.SetVec3(Vec3(-2, 0, 1))
	c.Assert(s.vec.Add(Vec3(2, 2, 1)), Equals, Vector3{0, 2, 2})
}

func (s *Vector3TestSuite) Vector3Sub(c *C) {
	s.vec.SetVec3(Vec3(-2, 0, 1))
	c.Assert(s.vec.Sub(Vec3(2, 2, 1)), Equals, Vector3{-4, -2, -1})
}

func (s *Vector3TestSuite) Vector3Mul(c *C) {
	s.vec.SetVec3(Vec3(-2, 0, 1))
	c.Assert(s.vec.Mul(Vec3(2, 2, -1)), Equals, Vector3{-4, 0, -1})
}

func (s *Vector3TestSuite) Vector3Div(c *C) {
	s.vec.SetVec3(Vec3(-2, 2, 1))
	c.Assert(s.vec.Div(Vec3(2, 2, -1)), Equals, Vector3{-1, 1, -1})
}

func (s *Vector3TestSuite) Vector3Len(c *C) {
	s.vec.SetVec3(Vec3(2, 0, 0))
	c.Assert(s.vec.Len(), Equals, 2)
}

func (s *Vector3TestSuite) Vector3Len2(c *C) {
	s.vec.SetVec3(Vec3(2, 0, 0))
	c.Assert(s.vec.Len2(), Equals, 4)
}

func (s *Vector3TestSuite) Vector3Dst(c *C) {
	s.vec.SetVec3(Vec3(2, 0, 0))
	c.Assert(s.vec.Dst(Vec3(0, 0, 0)), Equals, 2)
}

func (s *Vector3TestSuite) Vector3Dst2(c *C) {
	s.vec.SetVec3(Vec3(2, 0, 0))
	c.Assert(s.vec.Dst2(Vec3(0, 0, 0)), Equals, 4)
}

func (s *Vector3TestSuite) Vector3Nor(c *C) {
	s.vec.SetVec3(Vec3(2, 0, 0))
	c.Assert(s.vec.Nor(), Equals, Vec3(1, 0, 0))
}

func (s *Vector3TestSuite) Vector3Dot(c *C) {
	s.vec.SetVec3(Vec3(2, 0, 0))
	c.Assert(s.vec.Dot(Vec3(2, 0, 0)), Equals, 4)
}

func (s *Vector3TestSuite) Vector3Crs(c *C) {
	s.vec.SetVec3(Vec3(2, 1, 4))
	c.Assert(s.vec.Crs(Vec3(2, -3, 0)), Equals, Vec3(12, 6, -9))
}

// TODO MulMatrix
// TODO Prj
// TODO Rot
// TODO IsUnit
// TODO IsZero
// TODO Lerp
// TODO Slerp

func (s *Vector3TestSuite) Vector3Limit(c *C) {
	s.vec.SetVec3(Vec3(4, 0, 0))
	c.Assert(s.vec.Limit(3.3), Equals, Vec3(3.3, 0, 0))
}

func (s *Vector3TestSuite) Vector3Scale(c *C) {
	s.vec.SetVec3(Vec3(4, 0, -2))
	c.Assert(s.vec.Scale(3), Equals, Vec3(12, 0, -6))
}

func (s *Vector3TestSuite) Vector3Invert(c *C) {
	s.vec.SetVec3(Vec3(4, 0, -2))
	c.Assert(s.vec.Invert(), Equals, Vec3(-4, 0, -2))
}
