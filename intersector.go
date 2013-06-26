package math

// Returns whether the given point is inside the triangle.
// This assumes that the point is on the plane of the triangle.
// No check is performed that this is the case.
func IsPointInTriangle(point, t1, t2, t3 Vector3) bool {
	v0 := t1.Sub(point)
	v1 := t2.Sub(point)
	v2 := t3.Sub(point)

	ab := v0.Dot(v1)
	ac := v0.Dot(v2)
	bc := v1.Dot(v2)
	cc := v2.Dot(v2)

	if bc*ac-cc*ab < 0 {
		return false
	}
	bb := v1.Dot(v1)
	if ab*bc-ac*bb < 0 {
		return false
	}
	return true
}
