package math

type Path2 interface {
	ValueAt(t float32) Vector2
	Approximate(vec Vector2) float32
}

type Path3 interface {
	ValueAt(t float32) Vector3
	Approximate(vec Vector3) float32
}
