package math

import (
	"errors"
)

type Polygon struct {
	localVertices []float32
	worldVertices []float32
	dirty         bool
	origin        Vector2
	position      Vector2
	rotation      float32
	scalar        Vector2
	bounds        *Rectangle
}

func NewPolygon(vertices []float32) (*Polygon, error) {
	if len(vertices) < 6 {
		return nil, errors.New("Polygon must contain at least three points.")
	}
	return &Polygon{localVertices: vertices, dirty: true}, nil
}

func (p *Polygon) Vertices() []float32 {
	return p.localVertices
}

func (p *Polygon) TransformedVertices() []float32 {
	if p.dirty == false {
		return p.worldVertices
	}

	p.dirty = false
	localVertices := p.localVertices
	if p.worldVertices == nil || len(p.worldVertices) < len(localVertices) {
		p.worldVertices = make([]float32, len(localVertices))
	}
	cos := Cos(p.rotation) * DegreeToRadians
	sin := Sin(p.rotation) * DegreeToRadians

	for i := 0; i < len(localVertices); i += 2 {
		x := localVertices[i] - p.origin.X
		y := localVertices[i+1] - p.origin.Y

		if p.scalar.X != 1 || p.scalar.Y != 1 {
			x *= p.scalar.X
			y *= p.scalar.Y
		}

		if p.rotation != 0 {
			oldX := x
			x = cos*x - sin*y
			y = sin*oldX + cos*y
		}

		p.worldVertices[i] = p.position.X + x + p.origin.X
		p.worldVertices[i+1] = p.position.Y + y + p.origin.Y
	}
	return p.worldVertices
}

func (p *Polygon) SetOrigin(origin Vector2) {
	p.origin = origin
	p.dirty = true
}

func (p *Polygon) SetPosition(position Vector2) {
	p.position = position
	p.dirty = true
}

func (p *Polygon) Translate(vec Vector2) {
	p.position = p.position.Add(vec)
	p.dirty = true
}

func (p *Polygon) SetRotation(degrees float32) {
	p.rotation = degrees
	p.dirty = true
}

func (p *Polygon) Rotate(degrees float32) {
	p.rotation += degrees
	p.dirty = true
}

func (p *Polygon) SetScale(scalar Vector2) {
	p.scalar = scalar
	p.dirty = true
}

func (p *Polygon) Scale(amount float32) {
	p.scalar = p.scalar.Scale(amount)
	p.dirty = true
}

func (p *Polygon) Dirty() {
	p.dirty = true
}

func (p *Polygon) Area() float32 {
	var area float32

	vertices := p.TransformedVertices()

	var x1, y1, x2, y2 int
	for i := 0; i < len(vertices); i += 2 {
		x1 = i
		y1 = i + 1
		x2 = (i + 2) % len(vertices)
		y2 = (i + 3) % len(vertices)

		area += vertices[x1] * vertices[y2]
		area -= vertices[x2] * vertices[y1]
	}
	area *= 0.5
	return area
}

// Returns an axis-aligned bounding box of this polygon.
func (p *Polygon) BoundingRectangle() *Rectangle {
	vertices := p.TransformedVertices()
	minX := vertices[0]
	minY := vertices[1]
	maxX := vertices[0]
	maxY := vertices[1]

	for i := 0; i < len(vertices); i += 2 {
		if minX > vertices[i] {
			minX = vertices[i]
		}
		if minY > vertices[i+1] {
			minY = vertices[i+1]
		}
		if maxX < vertices[i] {
			maxX = vertices[i]
		}
		if maxY < vertices[i+1] {
			maxY = vertices[i+1]
		}
	}

	if p.bounds == nil {
		p.bounds = Rect(minX, minY, maxX, maxY)
	} else {
		p.bounds.X = minX
		p.bounds.Y = minY
		p.bounds.Width = maxX
		p.bounds.Height = maxY
	}
	return p.bounds
}

func (p *Polygon) Contains(vec Vector2) bool {
	vertices := p.TransformedVertices()
	intersects := 0

	for i := 0; i < len(vertices); i += 2 {
		x1 := vertices[i]
		y1 := vertices[i+1]
		x2 := vertices[(i+2)%len(vertices)]
		y2 := vertices[(i+3)%len(vertices)]
		if ((y1 <= p.position.Y && p.position.Y < y2) || (y2 <= p.position.Y && p.position.Y < y1)) && p.position.X < ((x2-x1)/(y2-y1)*(p.position.Y-y1)+x1) {
			intersects++
		}
	}
	return (intersects & 1) == 1
}

func (p *Polygon) Position() Vector2 { return p.position }
func (p *Polygon) Origin() Vector2   { return p.origin }
func (p *Polygon) Rotation() float32 { return p.rotation }
func (p *Polygon) Scalar() Vector2   { return p.scalar }
