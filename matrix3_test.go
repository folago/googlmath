package math

import (
	. "launchpad.net/gocheck"
)

type Matrix3TestSuite struct{}

var matrixTest3Suite = Suite(&Matrix3TestSuite{})

func (test *Matrix3TestSuite) SetUpTest(c *C) {}

func (test *Matrix3TestSuite) TestNewIdentityMatrix3(c *C) {
	expected := &Matrix3{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}
	obtained := NewIdentityMatrix3()
	c.Check(obtained, Matrix3Check, expected)
}
