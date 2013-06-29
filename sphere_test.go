package math

import (
	. "launchpad.net/gocheck"
)

type SphereOverlapsTestValue struct {
	Sphere, Sphere2 *Sphere
	Expected        bool
}

type SphereTestSuite struct {
	containTestTable []SphereOverlapsTestValue
}

var _ = Suite(SphereTestSuite{})

func (s *SphereTestSuite) SetUpTest(c *C) {
	s.containTestTable = []SphereOverlapsTestValue{
		SphereOverlapsTestValue{NewSphere(Vec3(1, -2, 0), 12.0), NewSphere(Vec3(0, 2, 0), 30.0), true},
	}
}

func (s *SphereTestSuite) Overlaps(c *C) {

	for i := range s.containTestTable {
		value := s.containTestTable[i]
		c.Assert(value.Sphere.Overlaps(value.Sphere2), Equals, value.Expected)
	}
}
