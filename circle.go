package math

// A two dimension circle.
type Circle struct {
	X      float32
	Y      float32
	Radius float32
}

func Circ(x, y, radius float32) Circle {
	return Circle{x, y, radius}
}

func (c Circle) Contains(x, y float32) bool {
	x = c.X - x
	y = c.Y - y
	return x*x+y*y <= c.Radius*c.Radius
}

func (c *Circle) Set(x, y, radius float32) Circle {
	c.X = x
	c.Y = y
	c.Radius = radius
	return *c
}
