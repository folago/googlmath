package math

type Ray struct {
	Origin    Vector3
	Direction Vector3
}

func NewRay(origin, direction Vector3) *Ray {
	ray := &Ray{Origin: Vec3(0, 0, 0), Direction: Vec3(0, 0, 0)}
	ray.Origin.SetVec3(origin)
	ray.Direction.SetVec3(direction)
	return ray
}

func (r *Ray) Cpy() *Ray {
	return &Ray{Origin: r.Origin.Cpy(), Direction: r.Direction.Cpy()}
}

func (r *Ray) Set(origin, direction Vector3) *Ray {
	r.Origin.SetVec3(origin)
	r.Direction.SetVec3(direction)
	return r
}

func (r *Ray) GetEndPoint(distance float32) Vector3 {
	return r.Origin.Cpy().Add(r.Direction.Cpy().Scale(distance))
}

// Multiplies the ray by the given matrix. Use this to transform a ray into another coordinate system.
func (r *Ray) Mul(matrix Matrix4) *Ray {
	tmp := r.Origin.Cpy().Add(r.Direction)
	tmp.MulMatrix(matrix)
	r.Origin.MulMatrix(matrix)
	tmp.Sub(r.Origin)
	r.Direction.Set(tmp.X, tmp.Y, tmp.Z)
	return r
}
