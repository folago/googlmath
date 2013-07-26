package math

import (
	. "launchpad.net/gocheck"
)

type BB2Vec3TestValue struct {
	Min, Max Vector3
	Expected *BoundingBox
}

type BBBoolTestValue struct {
	Value    *BoundingBox
	Expected bool
}

type BoundingBoxTestSuite struct {
	newBBTestTable   []BB2Vec3TestValue
	isValidTestTable []BBBoolTestValue
}

var _ = Suite(&BoundingBoxTestSuite{})

func (s *BoundingBoxTestSuite) SetUpTest(c *C) {
	s.newBBTestTable = []BB2Vec3TestValue{
		BB2Vec3TestValue{
			Min:      Vec3(0, 0, 0),
			Max:      Vec3(1, 2, 3),
			Expected: &BoundingBox{Min: Vec3(0, 0, 0), Max: Vec3(1, 2, 3)},
		},
		BB2Vec3TestValue{
			Min:      Vec3(-1, -2.2, 0),
			Max:      Vec3(2, 3, 3),
			Expected: &BoundingBox{Min: Vec3(-1, -2.2, 0), Max: Vec3(2, 3, 3)},
		},
	}

	s.isValidTestTable = []BBBoolTestValue{
		BBBoolTestValue{
			NewBoundingBox(Vec3(0, 0, 0), Vec3(1, 1, 1)),
			true,
		},
		BBBoolTestValue{
			NewBoundingBox(Vec3(0, 0, 0), Vec3(1.2, 2.3, 4.3)),
			true,
		},
		BBBoolTestValue{
			NewBoundingBox(Vec3(-2, -1, 1), Vec3(1, 0, 3)),
			true,
		},
		BBBoolTestValue{
			NewBoundingBox(Vec3(0, 0, 0), Vec3(0, 0, 0)),
			false,
		},
		BBBoolTestValue{
			NewBoundingBox(Vec3(2, 2, 2), Vec3(1, 1, 1)), // It is true, because Set swaps min<->max
			true,
		},
		BBBoolTestValue{
			NewBoundingBox(Vec3(0, 2, 1), Vec3(0, 2, 1)),
			false,
		},
	}
}

func (s *BoundingBoxTestSuite) TestNewBoundingBox(c *C) {
	for _, value := range s.newBBTestTable {
		obtained := NewBoundingBox(value.Min, value.Max)
		c.Check(obtained, BoundingBoxCheck, value.Expected)
	}
}

func (s *BoundingBoxTestSuite) TestSet(c *C) {
	for _, value := range s.newBBTestTable {
		bb := NewBoundingBox(Vec3(0, 0, 0), Vec3(0, 0, 0))
		obtained := bb.Set(value.Min, value.Max)
		c.Check(bb, BoundingBoxCheck, obtained)
		c.Check(obtained, BoundingBoxCheck, value.Expected)
	}
}

func (s *BoundingBoxTestSuite) TestCpy(c *C) {
	bb := NewBoundingBox(Vec3(-1, 2, -3), Vec3(1, 0.2, 3.3))
	obtained := bb.Cpy()
	c.Check(obtained, BoundingBoxCheck, bb)
}

func (s *BoundingBoxTestSuite) TestIsValid(c *C) {
	for _, value := range s.isValidTestTable {
		obtained := value.Value.IsValid()
		c.Check(obtained, Equals, value.Expected, Commentf("%v.IsValid()==%t epected %t", value.Value, obtained, value.Expected))
	}
}
