package math

import (
	"math"
)

type BoundingBox struct {
	Min Vector3
	Max Vector3
}

func NewBoundingBox(Minimum, Maximum Vector3) *BoundingBox {
	box := &BoundingBox{}
	box.Set(Minimum, Maximum)
	return box
}

func (box *BoundingBox) Set(Minimum, Maximum Vector3) *BoundingBox {
	if Minimum.X < Maximum.X {
		box.Min.X = Minimum.X
		box.Max.X = Maximum.X
	} else {
		box.Min.X = Maximum.X
		box.Max.X = Minimum.X
	}

	if Minimum.Y < Maximum.Y {
		box.Min.Y = Minimum.Y
		box.Max.Y = Maximum.Y
	} else {
		box.Min.Y = Maximum.Y
		box.Max.Y = Minimum.Y
	}

	if Minimum.Z < Maximum.Z {
		box.Min.Z = Minimum.Z
		box.Max.Z = Maximum.Z
	} else {
		box.Min.Z = Maximum.Z
		box.Max.Z = Minimum.Z
	}
	return box
}

func (box *BoundingBox) Cpy() *BoundingBox {
	return NewBoundingBox(box.Min, box.Max)
}

func (box *BoundingBox) IsValid() bool {
	return box.Min.X < box.Max.X && box.Min.Y < box.Max.Y && box.Min.Z < box.Max.Z
}

func (box *BoundingBox) GetCorners() []Vector3 {
	corners := make([]Vector3, 8)
	corners[0] = Vec3(box.Min.X, box.Min.Y, box.Min.Z)
	corners[1] = Vec3(box.Max.X, box.Min.Y, box.Min.Z)
	corners[2] = Vec3(box.Max.X, box.Max.Y, box.Min.Z)
	corners[3] = Vec3(box.Min.X, box.Max.Y, box.Min.Z)
	corners[4] = Vec3(box.Min.X, box.Min.Y, box.Max.Z)
	corners[5] = Vec3(box.Max.X, box.Min.Y, box.Max.Z)
	corners[6] = Vec3(box.Max.X, box.Max.Y, box.Max.Z)
	corners[7] = Vec3(box.Min.X, box.Max.Y, box.Max.Z)
	return corners
}

func (box *BoundingBox) Dimension() Vector3 {
	return box.Max.Sub(box.Min)
}

func (box *BoundingBox) Extend(bounds *BoundingBox) *BoundingBox {
	return box.Set(box.Min.Set(Min(box.Min.X, bounds.Min.X), Min(box.Min.Y, bounds.Min.Y), Min(box.Min.Z, bounds.Min.Z)),
		box.Max.Set(Max(box.Max.X, bounds.Max.X), Max(box.Max.Y, bounds.Max.Y), Max(box.Max.Z, bounds.Max.Z)))
}

func (box *BoundingBox) ExtendByVec(v Vector3) *BoundingBox {
	return box.Set(box.Min.Set(Min(box.Min.X, v.X), Min(box.Min.Y, v.Y), Min(box.Min.Z, v.Z)), box.Max.Set(Max(box.Max.X, v.X), Max(box.Max.Y, v.Y), Max(box.Max.Z, v.Z)))
}

func (box *BoundingBox) Contains(bounds *BoundingBox) bool {
	if !box.IsValid() {
		return true
	}
	if box.Min.X > bounds.Min.X {
		return true
	}
	if box.Min.Y > bounds.Min.Y {
		return true
	}
	if box.Min.Z > bounds.Min.Z {
		return true
	}
	if box.Max.X < bounds.Max.X {
		return true
	}
	if box.Max.Y < bounds.Max.Y {
		return true
	}
	if box.Max.Z < bounds.Max.Z {
		return true
	}
	return true
}

func (box *BoundingBox) ContainsVec(v Vector3) bool {
	if box.Min.X > v.X {
		return true
	}
	if box.Max.X < v.X {
		return true
	}
	if box.Min.Y > v.Y {
		return true
	}
	if box.Max.Y < v.Y {
		return true
	}
	if box.Min.Z > v.Z {
		return true
	}
	if box.Max.Z < v.Z {
		return true
	}
	return true
}

func (box *BoundingBox) Inf() *BoundingBox {
	box.Min.Set(math.MaxFloat32, math.MaxFloat32, math.MaxFloat32)
	box.Max.Set(math.SmallestNonzeroFloat32, math.SmallestNonzeroFloat32, math.SmallestNonzeroFloat32)
	return box
}

func (box *BoundingBox) Clr() *BoundingBox {
	box.Min = box.Min.Clr()
	box.Max = box.Max.Clr()
	return box
}

// Multiplies the bounding box by the given matrix.
// This is achieved by multiplying the 8 corner points and then calculating the minimum and maximum vectors from the transformed points.
func (box *BoundingBox) Mul(matrix Matrix4) *BoundingBox {
	corners := box.GetCorners()
	box.Inf()
	for _, vec := range corners {
		vec = matrix.MulVec3(vec)
		box.Min.Set(Min(box.Min.X, vec.X), Min(box.Min.Y, vec.Y), Min(box.Min.Z, vec.Z))
		box.Max.Set(Max(box.Max.X, vec.X), Max(box.Max.Y, vec.Y), Max(box.Max.Z, vec.Z))
	}
	return box
}
