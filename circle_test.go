package math

import (
	"testing"
)

func TestCircle(t *testing.T) {
	circle := Circ(0, 0, 1.0)

	// Contains
	if circle.Contains(0, 0) == false {
		t.Errorf("%v should contain Vector2{0,0}", circle)
	}
	if circle.Contains(1, 0) == false {
		t.Errorf("%v should contain Vector2{1,0}", circle)
	}
	if circle.Contains(-2, -2) == true {
		t.Errorf("%v should not contain Vector2{-2,-2}", circle)
	}
	if circle.Contains(2, 0) == true {
		t.Errorf("%v should not contain Vector2{2,0}", circle)
	}

	// Set
	circle.Set(1, 1.1, 0.0)
	if circle.X != 1 && circle.Y != 1.1 && circle.Radius != 0.0 {
		t.Errorf("%v should be Circle{1,1.1,0.0}", circle)
	}
}
