package math

type Point struct {
	X, Y int
}

func Pt(x, y int) Point {
	return Point{X: x, Y: y}
}

func (p Point) Cpy() Point {
	return Point{X: p.X, Y: p.Y}
}

func (p Point) Add(q Point) Point {
	p.X += q.X
	p.Y += q.Y
	return p
}

func (p Point) Sub(q Point) Point {
	p.X -= q.X
	p.Y -= q.Y
	return p
}

func (p Point) Mul(q Point) Point {
	p.X *= q.X
	p.Y *= q.Y
	return p
}

func (p Point) Div(q Point) Point {
	p.X /= q.X
	p.Y /= q.Y
	return p
}

func (p Point) Equals(q Point) bool {
	return p.X == q.X && p.Y == q.Y
}
