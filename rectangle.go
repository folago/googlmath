package math

type Rectangle struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

func Rect(x, y, width, height float32) *Rectangle {
	return &Rectangle{x, y, width, height}
}

func (r *Rectangle) ContainsRec(rec *Rectangle) bool {
	xmin := rec.X
	xmax := xmin + rec.Width
	ymin := rec.Y
	ymax := ymin + rec.Height

	return ((xmin > r.X && xmin < r.X+r.Width) && (xmax > r.X && xmax < r.X+r.Width)) && ((ymin > r.Y && ymin < r.Y+r.Height) && (ymax > r.Y && ymax < r.Y+r.Height))
}

func (r *Rectangle) Overlaps(rec Rectangle) bool {
	return !(r.X > rec.X+rec.Width || r.X+r.Width < rec.X || r.Y > rec.Y+rec.Height || r.Y+r.Height < rec.Y)
}

func (r *Rectangle) Set(x, y, width, height float32) *Rectangle {
	r.X = x
	r.Y = y
	r.Width = width
	r.Height = height
	return r
}

func (r *Rectangle) ContainsVec2(x, y float32) bool {
	return r.X < x && r.X+r.Width > x && r.Y < y && r.Y+r.Height > y
}

// Merges this rectangle with the other rectangle.
func (r *Rectangle) Merge(rect *Rectangle) *Rectangle {
	minX := Min(r.X, rect.X)
	maxX := Max(r.X+r.Width, rect.X+rect.Width)
	r.X = minX
	r.Width = maxX - minX

	minY := Min(r.Y, rect.Y)
	maxY := Max(r.Y+r.Height, rect.Y+rect.Height)
	r.Y = minY
	r.Height = maxY - minY
	return r
}
