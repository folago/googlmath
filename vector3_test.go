package math

import (
	"testing"
)

func TestVec3(t *testing.T) {
	var x float32 = 1.23
	var y float32 = -3.21
	var z float32 = -0.0
	vec := Vec3(x, y, z)
	if x != vec.X {
		t.Errorf("%g should be equal to %g", x, vec.X)
	}
	if y != vec.Y {
		t.Errorf("%g should be equal to %g", y, vec.Y)
	}
	if z != vec.Z {
		t.Errorf("%g should be equal to %g", z, vec.Z)
	}
}

func TestVec3Set(t *testing.T) {
	var x float32 = -2
	var y float32 = 0
	var z float32 = 1
	vec := Vec3(x, y, z)
	vec.Set(1, 1, 1)
	if vec.X != 1 || vec.Y != 1 || vec.Z != 1 {
		t.Errorf("%+v should be equal to Vector3(1,1,1)", vec)
	}
}

func TestVec3SetVec3(t *testing.T) {
	var x float32 = -2
	var y float32 = 0
	var z float32 = 1
	vec := Vec3(x, y, z)
	vec.SetVec3(Vec3(1, 1, 1))
	if vec.X != 1 || vec.Y != 1 || vec.Z != 1 {
		t.Errorf("%+v should be equal to Vector3(1,1,1)", vec)
	}
}

func TestVec3Cpy(t *testing.T) {
	var x float32 = 1.23
	var y float32 = -3.21
	var z float32 = 0
	a := Vec3(x, y, z)
	b := a.Cpy()
	if a != b {
		t.Errorf("%+v should be equal to %+v", a, b)
	}
}

func TestVec3Clr(t *testing.T) {
	var x float32 = 1.23
	var y float32 = -3.21
	var z float32 = 0
	a := Vec3(x, y, z)
	b := a.Clr()
	if b.X != 0 || b.Y != 0 || b.Z != 0 {
		t.Errorf("%+v should be equal to Vector3{0,0,0}", b)
	}
	if a.X != x || a.Y != y || a.Z != z {
		t.Errorf("%+v should be equal to %+v", a, b)
	}
}

func TestVec3Add(t *testing.T) {
	var x float32 = 1.23
	var y float32 = -3.21
	var z float32 = 0
	a := Vec3(x, y, z)
	b := Vec3(0.77, 3.21, 0)
	c := a.Add(b)
	if a.X != x || a.Y != y || a.Z != z {
		t.Errorf("%+v should be equal to Vector3{%g,%g,%g}", a, x, y, z)
	}
	if b.X != 0.77 || b.Y != 3.21 || b.Z != 0 {
		t.Errorf("%+v should be equal to Vector3{0.77,3.21,0}", b)
	}
	if c.X != 2.0 || c.Y != 0 || c.Z != 0 {
		t.Errorf("%+v should be equal to Vector3{2.0,0,0}", b)
	}
}

func TestVec3Sub(t *testing.T) {
	var x float32 = 1.23
	var y float32 = -3.21
	var z float32 = 0
	a := Vec3(x, y, z)
	b := Vec3(0.23, 3.21, 0)
	c := a.Sub(b)
	if a.X != x || a.Y != y || a.Z != z {
		t.Errorf("%+v should be equal to Vector3{%g,%g,%g}", a, x, y, z)
	}
	if b.X != 0.23 || b.Y != 3.21 || b.Z != 0 {
		t.Errorf("%+v should be equal to Vector3{0.23,3.21,0}", b)
	}
	if c.X != 1.0 || c.Y != -6.42 || c.Z != 0 {
		t.Errorf("%+v should be equal to Vector3{1.0,-6.42,0}", b)
	}
}

func TestVec3Mul(t *testing.T) {
	var x float32 = 1.25
	var y float32 = -2.0
	var z float32 = 0
	a := Vec3(x, y, z)
	b := Vec3(4, 2.5, -2)
	c := a.Mul(b)
	if a.X != x || a.Y != y || a.Z != z {
		t.Errorf("%+v should be equal to Vector3{%g,%g,%g}", a, x, y, z)
	}
	if b.X != 4 || b.Y != 2.5 || b.Z != -2 {
		t.Errorf("%+v should be equal to Vector3{4,2.5,-2}", b)
	}
	if c.X != 5.0 || c.Y != -5.0 || c.Z != 0 {
		t.Errorf("%+v should be equal to Vector3{5.0,-5.0,0}", b)
	}
}

func TestVec3Div(t *testing.T) {
	var x float32 = 2
	var y float32 = -2.0
	var z float32 = 0
	a := Vec3(x, y, z)
	b := Vec3(2, 2, -2)
	c := a.Div(b)
	if a.X != x || a.Y != y || a.Z != z {
		t.Errorf("%+v should be equal to Vector3{%g,%g,%g}", a, x, y, z)
	}
	if b.X != 2 || b.Y != 2 || b.Z != -2 {
		t.Errorf("%+v should be equal to Vector3{2,2,-2}", b)
	}
	if c.X != 1 || c.Y != -1.0 || c.Z != 0 {
		t.Errorf("%+v should be equal to Vector3{1.0,-1.0,0}", b)
	}
}

func TestVec3Len(t *testing.T) {
	var x float32 = 2
	var y float32 = 0
	var z float32 = 0
	a := Vec3(x, y, z)
	length := a.Len()
	if a.X != x || a.Y != y || a.Z != z {
		t.Errorf("%+v should be equal to Vector3{%g,%g,%g}", a, x, y, z)
	}
	if length != 2 {
		t.Errorf("length %g should be equal to length 2", length)
	}
}

func TestVec3Len2(t *testing.T) {
	var x float32 = 2
	var y float32 = 0
	var z float32 = 0
	a := Vec3(x, y, z)
	length2 := a.Len2()
	if a.X != x || a.Y != y || a.Z != z {
		t.Errorf("%+v should be equal to Vector3{%g,%g,%g}", a, x, y, z)
	}
	if length2 != a.Len()*a.Len() {
		t.Errorf("length2 %g should be equal to length*length", length2)
	}
	if length2 != 4 {
		t.Errorf("length2 %g should be equal to 4", length2)
	}
}

func TestVec3Dst(t *testing.T) {
	var x float32 = 2
	var y float32 = 0
	var z float32 = 0
	a := Vec3(x, y, z)
	b := Vec3(0, 0, 0)
	dst := a.Dst(b)
	if a.X != x || a.Y != y || a.Z != z {
		t.Errorf("%+v should be equal to Vector3{%g,%g,%g}", a, x, y, z)
	}
	if dst != 2 {
		t.Errorf("Distance %g should be equal to 2", dst)
	}
}

func TestVec3Dst2(t *testing.T) {
	var x float32 = 2
	var y float32 = 0
	var z float32 = 0
	a := Vec3(x, y, z)
	b := Vec3(0, 0, 0)
	dst2 := a.Dst2(b)
	if a.X != x || a.Y != y || a.Z != z {
		t.Errorf("%+v should be equal to Vector3{%g,%g,%g}", a, x, y, z)
	}
	if dst2 != a.Dst(b)*a.Dst(b) {
		t.Errorf("Distance2 %g should be equal to Distance*Distance", dst2)
	}
	if dst2 != 4 {
		t.Errorf("Distance2 %g should be equal to 4", dst2)
	}
}

func TestVec3Nor(t *testing.T) {
	var x float32 = 2
	var y float32 = 0
	var z float32 = 0
	a := Vec3(x, y, z)
	nor := a.Nor()
	if a.X != x || a.Y != y || a.Z != z {
		t.Errorf("%+v should be equal to Vector3{%g,%g,%g}", a, x, y, z)
	}
	if nor.X != 1 || nor.Y != 0 || nor.Z != 0 {
		t.Errorf("%+v should be equal to Vector3{%g,%g,%g}", nor, 1, 0, 0)
	}
}

func TestVec3Dot(t *testing.T) {
	var x float32 = 2
	var y float32 = 0
	var z float32 = 0
	a := Vec3(x, y, z)
	b := Vec3(x, y, z)
	dot := a.Dot(b)
	if a.X != x || a.Y != y || a.Z != z {
		t.Errorf("%+v should be equal to Vector3{%g,%g,%g}", a, x, y, z)
	}
	if dot != 4 {
		t.Errorf("Dot product %g should be equal to 4", dot)
	}
}

func TestVec3Cross(t *testing.T) {
	var x float32 = 2
	var y float32 = 1
	var z float32 = 4
	a := Vec3(x, y, z)
	b := Vec3(2, -3, 0)
	cross := a.Crs(b)
	if a.X != x || a.Y != y || a.Z != z {
		t.Errorf("%+v should be equal to Vector3{%g,%g,%g}", a, x, y, z)
	}
	ox := a.Y*b.Z - b.Y*a.Z
	oy := a.Z*b.X - b.Z*a.X
	oz := a.X*b.Y - b.X*a.Y
	if cross.X != ox || cross.Y != oy || cross.Z != oz {
		t.Errorf("Cross product %g should be equal to Vector3(%g,%g,%g)", cross, ox, oy, oz)
	}
}

// TODO MulMatrix
// TODO Prj
// TODO Rot
// TODO IsUnit
// TODO IsZero
// TODO Lerp
// TODO Slerp

func TestVec3Limit(t *testing.T) {
	var x float32 = 4
	var y float32 = 0
	var z float32 = 0
	a := Vec3(x, y, z)
	b := a.Limit(3.3)
	if a.X != x || a.Y != y || a.Z != z {
		t.Errorf("%+v should be equal to Vector3{%g,%g,%g}", a, x, y, z)
	}
	if b.X != 3.3 || b.Y != 0 || b.Z != 0 {
		t.Errorf("%+v should be equal to Vector3{%g,%g,%g}", b, 3.3, 0.0, 0.0)
	}
}

func TestVec3Scale(t *testing.T) {
	var x float32 = 2
	var y float32 = 0
	var z float32 = 0
	var scale float32 = 2
	a := Vec3(x, y, z)
	b := a.Scale(scale)
	if a.X != x || a.Y != y || a.Z != z {
		t.Errorf("%+v should be equal to Vector3{%g,%g,%g}", a, x, y, z)
	}
	if b.X != x*scale || b.Y != y*scale || b.Z != z*scale {
		t.Errorf("%+v should be equal to Vector3{%g,%g,%g}", b, x*scale, y*scale, z*scale)
	}
}

func TestVec3Invert(t *testing.T) {
	var x float32 = 1.2
	var y float32 = -2.1
	var z float32 = 0
	a := Vec3(x, y, z)
	b := a.Invert()
	if a.X != x || a.Y != y || a.Z != z {
		t.Errorf("%+v should be equal to Vector3{%g,%g,%g}", a, x, y, z)
	}
	if b.X != -x || b.Y != -y || b.Z != -z {
		t.Errorf("%+v should be equal to Vector3{%g,%g,%g}", b, -x, -y, -z)
	}
}
