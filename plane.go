package math

type PlaneSide int

const (
	PlaneSide_OnPlane PlaneSide = iota
	PlaneSide_Back
	PlaneSide_Front
)

type Plane struct {
	Normal Vector3
	D      float32
}

func NewPlane(normal Vector3, d float32) *Plane {
	return &Plane{normal.Cpy().Nor(), d}
}

func (p *Plane) Cpy() *Plane {
	return &Plane{p.Normal.Cpy(), p.D}
}

// Sets the plane normal and distance to the origin based on the three given points which are considered to be on the plane.
// The normal is calculated via a cross product between (point1-point2)x(point2-point3)
func (p *Plane) Set(p1, p2, p3 Vector3) {
	l := p1.Sub(p2)
	r := p2.Sub(p3)
	p.Normal = l.Crs(r).Nor()
	p.D = -p1.Dot(p.Normal)
}

// Calculates the shortest signed distance between the plane and the given point.
func (p *Plane) Distance(vec Vector3) float32 {
	return p.Normal.Dot(vec) + p.D
}

// Returns on which side the given point lies relative to the plane and its normal.
// PlaneSide.Front refers to the side the plane normal points to.
func (p *Plane) PlaneSide(vec Vector3) PlaneSide {
	dist := p.Normal.Dot(vec) + p.D
	if dist == 0 {
		return PlaneSide_OnPlane
	}
	if dist < 0 {
		return PlaneSide_Back
	}
	return PlaneSide_Front
}

// Returns whether the plane is facing the direction vector. Think of the direction vector as the direction a camera looks in.
// This method will return true if the front side of the plane determined by its normal faces the camera.
func (p *Plane) IsFrontFacing(direction Vector3) bool {
	dot := p.Normal.Dot(direction)
	return dot <= 0
}
