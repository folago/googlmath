package math

import (
	"testing"
)

func TestNewSphere(t *testing.T) {
	var radius float32 = 12.0
	var center Vector3 = Vec3(1, 2, 3)
	sphere := NewSphere(center, radius)
	if sphere.Radius != radius {
		t.Errorf("Radius %g should be equal to %g", sphere.Radius, radius)
	}
	if sphere.Center != center {
		t.Errorf("Center %g should be equal to %g", sphere.Center, center)
	}
}

func TestSphereOverlaps(t *testing.T) {
	var radius float32 = 12.0
	var center Vector3 = Vec3(1, -2, 0)
	sphere := NewSphere(center, radius)

	var radius2 float32 = 30.0
	var center2 Vector3 = Vec3(0, 2, 0)
	sphere2 := NewSphere(center2, radius2)

	if sphere.Overlaps(sphere2) != true {
		t.Errorf("Sphere %+v should overlaps Sphere %+v", sphere, sphere2)
	}
}
