package math

type Point3 struct {
	X, Y, Z int
}

func Pt3(x, y, z int) Point3 {
	return Point3{X: x, Y: y, Z: z}
}

func (p Point3) Add(q Point3) Point3 {
	p.X += q.X
	p.Y += q.Y
	p.Z += q.Z
	return p
}

func (p Point3) Sub(q Point3) Point3 {
	p.X -= q.X
	p.Y -= q.Y
	p.Z -= q.Z
	return p
}

func (p Point3) Mul(q Point3) Point3 {
	p.X *= q.X
	p.Y *= q.Y
	p.Z *= q.Z
	return p
}

func (p Point3) Div(q Point3) Point3 {
	p.X /= q.X
	p.Y /= q.Y
	p.Z /= q.Z
	return p
}

func (p Point3) Equals(q Point3) bool {
	return p.X == q.X && p.Y == q.Y && p.Z == q.Z
}
