package math

// Encapsulates a 3D sphere with a center and a radius
type Sphere struct {
	Radius float32
	Center Vector3
}

func NewSphere(center Vector3, radius float32) *Sphere {
	return &Sphere{radius, center}
}

func (s *Sphere) Overlaps(sphere *Sphere) bool {
	return s.Center.Distance2(sphere.Center) < (s.Radius+sphere.Radius)*(s.Radius+sphere.Radius)
}
