package math

var (
	clipSpacePlanePoints = []Vector3{
		Vec3(-1, -1, -1),
		Vec3(1, -1, -1),
		Vec3(1, 1, -1),
		Vec3(-1, 1, -1),
		Vec3(-1, -1, 1),
		Vec3(1, -1, 1),
		Vec3(1, 1, 1),
		Vec3(-1, 1, 1)}
)

// A truncated rectangular pyramid.
// Used to define the viewable region and it's projection onto the screen.
type Frustum struct {
	// The six clipping planes, near, far, left, right, top, bottom
	Left, Right *Plane
	Top, Bottom *Plane
	Near, Far   *Plane

	planePoints []Vector3
}

func NewFrustum() *Frustum {
	zeroPlane := NewPlane(Vec3(0, 0, 0), 0)
	return &Frustum{Left: zeroPlane, Right: zeroPlane.Cpy(),
		Top: zeroPlane.Cpy(), Bottom: zeroPlane.Cpy(),
		Near: zeroPlane.Cpy(), Far: zeroPlane.Cpy(),
		planePoints: make([]Vector3, len(clipSpacePlanePoints))}
}

func (f *Frustum) Update(invProjectionView Matrix4) {
	for i := range clipSpacePlanePoints {
		f.planePoints[i] = invProjectionView.Project(clipSpacePlanePoints[i])
	}
	f.Near.Set(f.planePoints[1], f.planePoints[0], f.planePoints[2])
	f.Far.Set(f.planePoints[4], f.planePoints[5], f.planePoints[7])
	f.Left.Set(f.planePoints[0], f.planePoints[4], f.planePoints[3])
	f.Right.Set(f.planePoints[5], f.planePoints[1], f.planePoints[6])
	f.Top.Set(f.planePoints[2], f.planePoints[3], f.planePoints[6])
	f.Bottom.Set(f.planePoints[4], f.planePoints[0], f.planePoints[1])
}

// Returns whether the point is in the frustum.
func (f *Frustum) PointInFrustum(point Vector3) bool {
	if f.Left.PlaneSide(point) == PlaneSide_Back {
		return false
	}
	if f.Right.PlaneSide(point) == PlaneSide_Back {
		return false
	}
	if f.Top.PlaneSide(point) == PlaneSide_Back {
		return false
	}
	if f.Bottom.PlaneSide(point) == PlaneSide_Back {
		return false
	}
	if f.Near.PlaneSide(point) == PlaneSide_Back {
		return false
	}
	if f.Far.PlaneSide(point) == PlaneSide_Back {
		return false
	}
	return true
}

// Returns whether the given sphere is in the frustum.
func (f *Frustum) SphereInFrustum(center Vector3, radius float32) bool {
	if (f.Left.Normal.X*center.X + f.Left.Normal.Y*center.Y + f.Left.Normal.Z*center.Z) < (-radius - f.Left.D) {
		return false
	}
	if (f.Right.Normal.X*center.X + f.Right.Normal.Y*center.Y + f.Right.Normal.Z*center.Z) < (-radius - f.Right.D) {
		return false
	}
	if (f.Top.Normal.X*center.X + f.Top.Normal.Y*center.Y + f.Top.Normal.Z*center.Z) < (-radius - f.Top.D) {
		return false
	}
	if (f.Bottom.Normal.X*center.X + f.Bottom.Normal.Y*center.Y + f.Bottom.Normal.Z*center.Z) < (-radius - f.Bottom.D) {
		return false
	}
	if (f.Near.Normal.X*center.X + f.Near.Normal.Y*center.Y + f.Near.Normal.Z*center.Z) < (-radius - f.Near.D) {
		return false
	}
	if (f.Far.Normal.X*center.X + f.Far.Normal.Y*center.Y + f.Far.Normal.Z*center.Z) < (-radius - f.Far.D) {
		return false
	}
	return true
}

// Returns whether the given sphere is in the frustum not checking whether it is behind the near and far clipping plane.
func (f *Frustum) SphereInFrustumWithoutNearFar(center Vector3, radius float32) bool {
	if (f.Left.Normal.X*center.X + f.Left.Normal.Y*center.Y + f.Left.Normal.Z*center.Z) < (-radius - f.Left.D) {
		return false
	}
	if (f.Right.Normal.X*center.X + f.Right.Normal.Y*center.Y + f.Right.Normal.Z*center.Z) < (-radius - f.Right.D) {
		return false
	}
	if (f.Top.Normal.X*center.X + f.Top.Normal.Y*center.Y + f.Top.Normal.Z*center.Z) < (-radius - f.Top.D) {
		return false
	}
	if (f.Bottom.Normal.X*center.X + f.Bottom.Normal.Y*center.Y + f.Bottom.Normal.Z*center.Z) < (-radius - f.Bottom.D) {
		return false
	}
	return true
}

// Returns whether the given {@link BoundingBox} is in the frustum.
func (f *Frustum) BoundsInFrustum(bounds *BoundingBox) bool {
	corners := bounds.GetCorners()

	out := 0
	for i := 0; i < len(corners); i++ {
		if f.Left.PlaneSide(corners[i]) == PlaneSide_Back {
			out++
		}
	}
	if out == 8 {
		return false
	}

	out = 0
	for i := 0; i < len(corners); i++ {
		if f.Right.PlaneSide(corners[i]) == PlaneSide_Back {
			out++
		}
	}
	if out == 8 {
		return false
	}

	out = 0
	for i := 0; i < len(corners); i++ {
		if f.Top.PlaneSide(corners[i]) == PlaneSide_Back {
			out++
		}
	}
	if out == 8 {
		return false
	}

	out = 0
	for i := 0; i < len(corners); i++ {
		if f.Bottom.PlaneSide(corners[i]) == PlaneSide_Back {
			out++
		}
	}
	if out == 8 {
		return false
	}

	out = 0
	for i := 0; i < len(corners); i++ {
		if f.Near.PlaneSide(corners[i]) == PlaneSide_Back {
			out++
		}
	}
	if out == 8 {
		return false
	}

	out = 0
	for i := 0; i < len(corners); i++ {
		if f.Far.PlaneSide(corners[i]) == PlaneSide_Back {
			out++
		}
	}
	if out == 8 {
		return false
	}

	return true
}
