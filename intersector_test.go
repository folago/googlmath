package math

import (
	. "launchpad.net/gocheck"
)

type IsPointInTriangleTestValue struct {
	Point      Vector3
	T1, T2, T3 Vector3
	Expected   bool
}

type IntersectorTestSuite struct {
	containTestTable []IsPointInTriangleTestValue
}

var _ = Suite(IntersectorTestSuite{})

func (s *IntersectorTestSuite) SetUpTest(c *C) {
	s.containTestTable = []IsPointInTriangleTestValue{
		IsPointInTriangleTestValue{Vec3(0.5, 0.5, 0), Vec3(0, 0, 0), Vec3(1, 1, 0), Vec3(0, 1, 0), true},
		IsPointInTriangleTestValue{Vec3(2, 0.5, 0), Vec3(0, 0, 0), Vec3(1, 1, 0), Vec3(0, 1, 0), false},
	}
}

func (s *IntersectorTestSuite) TestIsPointInTriangle(c *C) {

	for i := range s.containTestTable {
		value := s.containTestTable[i]
		c.Assert(IsPointInTriangle(value.Point, value.T1, value.T2, value.T3), Equals, value.Expected)
	}
}
