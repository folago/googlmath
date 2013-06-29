package math

import (
	. "launchpad.net/gocheck"
)

type CircleContainTestValue struct {
	X, Y     float32
	Expected bool
}

type CircleTestSuite struct {
	circle           Circle
	containTestTable []CircleContainTestValue
}

var _ = Suite(CircleTestSuite{})

func (s *CircleTestSuite) SetUpTest(c *C) {
	s.circle = Circ(0, 0, 1.0)
	s.containTestTable = []CircleContainTestValue{
		CircleContainTestValue{0, 0, true},
		CircleContainTestValue{1, 0, true},
		CircleContainTestValue{-2, -2, false},
		CircleContainTestValue{2, 0, false}}
}

func (s *CircleTestSuite) TestCircle(c *C) {
	// Contains
	for i := range s.containTestTable {
		c.Assert(s.circle.Contains(s.containTestTable[i].X, s.containTestTable[i].Y), Equals, s.containTestTable[i].Expected)
	}

	// Set
	c.Assert(s.circle.Set(1, 1.1, 0.0), Equals, Circ(1, 1.1, 0.0))
}
