package math

import (
	"testing"
)

func TestEllipse(t *testing.T) {
	e := NewEllipse(0, 0, 1, 1)

	// Contains
	if e.Contains(0, 0) == false {
		t.Errorf("%v should contain Vector2{0,0}", e)
	}
	if e.Contains(1, 0) == false {
		t.Errorf("%v should contain Vector2{1,0}", e)
	}
	if e.Contains(-2, -2) == true {
		t.Errorf("%v should not contain Vector2{-2,-2}", e)
	}
	if e.Contains(2, 0) == true {
		t.Errorf("%v should not contain Vector2{2,0}", e)
	}

	// Set
	e.Set(1, 1.1, 2.0, 2.2)
	if e.X != 1 && e.Y != 1.1 && e.Width != 2.0 && e.Height != 2.2 {
		t.Errorf("%v should be Circle{1,1.1,2.0,2.2}", e)
	}
}
