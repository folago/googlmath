package math

import (
	"testing"
)

func TestIntersectorIsPointInTriangle(t *testing.T) {
	point := Vec3(0.5, 0.5, 0)
	t1 := Vec3(0, 0, 0)
	t2 := Vec3(1, 1, 0)
	t3 := Vec3(0, 1, 0)

	if IsPointInTriangle(point, t1, t2, t3) == false {
		t.Errorf("Point %+v should be in Triangle(%+v,%+v,%+v)", point, t1, t2, t3)
	}

	point.X = 2
	if IsPointInTriangle(point, t1, t2, t3) == true {
		t.Errorf("Point %+v should not be in Triangle(%+v,%+v,%+v)", point, t1, t2, t3)
	}
}
