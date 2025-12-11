package vec

import "math"

type Vec3 struct {
	X, Y, Z float64
}

func NewVec3(c [3]float64) Vec3 {
	return Vec3{c[0], c[1], c[2]}
}

func (vec *Vec3) GetX() float64 {
	return vec.X
}

func (vec *Vec3) GetY() float64 {
	return vec.Y
}

func (vec *Vec3) GetZ() float64 {
	return vec.Z
}

func (vec Vec3) Negate() Vec3 {
	return Vec3{
		X: -vec.X,
		Y: -vec.Y,
		Z: -vec.Z,
	}
}

func (vec1 Vec3) Add(vec2 Vec3) Vec3 {
	return Vec3{vec1.X + vec2.X, vec1.Y + vec2.Y, vec1.Z + vec2.Z}
}

func (vec1 Vec3) Sub(vec2 Vec3) Vec3 {
	return Vec3{vec2.X - vec1.X, vec2.Y - vec1.Y, vec2.Z - vec1.Z}
}

func (vec1 Vec3) Scale(s float64) Vec3 {
	return Vec3{vec1.X * s, vec1.Y * s, vec1.Z * s}
}

func Divide(vec *Vec3, s float64) Vec3 {
	return vec.Scale(1 / s)
}

func (vec Vec3) LengthSquared() float64 {
	return vec.X*vec.X + vec.Y*vec.Y + vec.Z*vec.Z
}

func (vec Vec3) Length() float64 {
	return math.Sqrt(vec.LengthSquared())
}

func (vec1 *Vec3) Dot(vec2 *Vec3) float64 {
	return vec2.X*vec1.X + vec2.Y*vec1.Y + vec2.Z*vec1.Z
}

func (vec1 *Vec3) Cross(vec2 *Vec3) Vec3 {
	return Vec3{
		vec2.Y*vec1.Z - vec2.Z*vec1.Y,
		vec2.Z*vec1.X - vec2.X*vec1.Z,
		vec2.X*vec1.Y - vec2.Y*vec1.X,
	}
}

func (vec Vec3) Unit() Vec3 {
	return vec.Scale(1 / vec.Length())
}
