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

func (box *BoundingBox) Dx() float32 {
	return box.Max.X - box.Min.X
}

func (box *BoundingBox) Dy() float32 {
	return box.Max.Y - box.Min.Y
}

func (box *BoundingBox) Dz() float32 {
	return box.Max.Z - box.Min.Z
}

func (box *BoundingBox) Center() Vector3 {
	dimension := box.Max.Sub(box.Min)
	dimension = dimension.Scale(0.5)
	return box.Min.Add(dimension)
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

func (box *BoundingBox) Corners() []Vector3 {
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
	return !box.IsValid() || (box.Min.X <= bounds.Min.X && box.Min.Y <= bounds.Min.Y && box.Min.Z <= bounds.Min.Z && box.Max.X >= bounds.Max.X && box.Max.Y >= bounds.Max.Y && box.Max.Z >= bounds.Max.Z)
}

func (box *BoundingBox) Overlaps(bounds *BoundingBox) bool {
	if bounds.ContainsVec(box.Min) || bounds.ContainsVec(box.Max) {
		return true
	}
	if bounds.Max.X < box.Max.X && bounds.Max.Y < box.Max.Y && bounds.Max.Z < box.Max.Z {
		if bounds.Min.X > box.Min.X && bounds.Min.Y > box.Min.Y && bounds.Min.Z > box.Min.Z {
			return true
		}
	}
	return false
}

func (box *BoundingBox) ContainsVec(v Vector3) bool {
	return box.Min.X <= v.X && box.Max.X >= v.X && box.Min.Y <= v.Y && box.Max.Y >= v.Y && box.Min.Z <= v.Z && box.Max.Z >= v.Z
}

func (box *BoundingBox) Inf() *BoundingBox {
	box.Min.Set(math.MaxFloat32, math.MaxFloat32, math.MaxFloat32)
	box.Max.Set(math.SmallestNonzeroFloat32, math.SmallestNonzeroFloat32, math.SmallestNonzeroFloat32)
	return box
}

func (box *BoundingBox) Offset(off Vector3) *BoundingBox {
	return &BoundingBox{Min: box.Min.Add(off),
		Max: box.Max.Add(off),
	}
}

func (box *BoundingBox) Clr() *BoundingBox {
	box.Min = box.Min.Clr()
	box.Max = box.Max.Clr()
	return box
}
