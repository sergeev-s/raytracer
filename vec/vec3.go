package vec

import "math"

type Vec3 struct {
	X, Y, Z float64
}

type Point3 = Vec3
type Color = Vec3

func NewVec3(c [3]float64) Vec3 {
	return Vec3{c[0], c[1], c[2]}
}

// func (v Vec3) GetX() float64 {
// 	return v.X
// }

// func (v Vec3) GetY() float64 {
// 	return v.Y
// }

// func (v Vec3) GetZ() float64 {
// 	return v.Z
// }

func (v Vec3) Negate() Vec3 {
	return Vec3{
		X: -v.X,
		Y: -v.Y,
		Z: -v.Z,
	}
}

func (v Vec3) Add(u Vec3) Vec3 {
	return Vec3{v.X + u.X, v.Y + u.Y, v.Z + u.Z}
}

func (v Vec3) Sub(u Vec3) Vec3 {
	return Vec3{v.X - u.X, v.Y - u.Y, v.Z - u.Z}
}

func (v Vec3) Scale(s float64) Vec3 {
	return Vec3{v.X * s, v.Y * s, v.Z * s}
}

func (v Vec3) Divide(s float64) Vec3 {
	return v.Scale(1 / s)
}

func (v Vec3) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vec3) Dot(u Vec3) float64 {
	return u.X*v.X + u.Y*v.Y + u.Z*v.Z
}

func (v Vec3) Cross(u Vec3) Vec3 {
	return Vec3{
		v.Y*u.Z - v.Z*u.Y,
		v.Z*u.X - v.X*u.Z,
		v.X*u.Y - v.Y*u.X,
	}
}

func (v Vec3) Unit() Vec3 {
	if v.NearZero() {
		return Vec3{}
	}
	return v.Scale(1 / v.Length())
}

func (v Vec3) NearZero() bool {
	const eps = 1e-8
	return math.Abs(v.X) < eps && math.Abs(v.Y) < eps && math.Abs(v.Z) < eps
}
