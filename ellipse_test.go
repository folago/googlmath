package math

import (
	. "launchpad.net/gocheck"
)

type EllipseContainTestValue struct {
	X, Y     float32
	Expected bool
}

type EllipseTestSuite struct {
	e                *Ellipse
	containTestTable []EllipseContainTestValue
}

var _ = Suite(EllipseTestSuite{})

func (s *EllipseTestSuite) SetUpTest(c *C) {
	s.e = NewEllipse(0, 0, 1, 1)
	s.containTestTable = []EllipseContainTestValue{
		EllipseContainTestValue{0, 0, true},
		EllipseContainTestValue{1, 0, true},
		EllipseContainTestValue{-2, -2, false},
		EllipseContainTestValue{2, 0, false},
	}
}

func (s *EllipseTestSuite) TestEllipse(c *C) {
	// Contains
	for i := range s.containTestTable {
		c.Assert(s.e.Contains(s.containTestTable[i].X, s.containTestTable[i].Y), Equals, s.containTestTable[i].Expected)
	}

	// Set
	c.Assert(s.e.Set(1, 1.1, 2.0, 2.2), Equals, NewEllipse(1, 1.1, 2.0, 2.2))
}
