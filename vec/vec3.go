package vec

import (
	"math"
	"math/rand"
)

type Vec3 struct {
	X, Y, Z float64
}

type Point3 = Vec3
type Color = Vec3

func Reflect(v, n Vec3) Vec3 {
	return v.Sub(n.Scale(v.Dot(n) * 2))
}

func Random() Vec3 {
	return Vec3{
		X: RandomInterval(-1, 1),
		Y: RandomInterval(-1, 1),
		Z: RandomInterval(-1, 1),
	}
}

func RandomUnitVector() Vec3 {
	for {
		p := Random()
		lenSq := p.LengthSquared()
		if lenSq >= 1e-160 && lenSq <= 1 {
			return p.Divide(math.Sqrt(lenSq))
		}
	}	
}

func RandomOnHemisphere(normal Vec3) Vec3 {
	randomUnit := RandomUnitVector()
    if normal.Dot(randomUnit) > 0.0 {
		return randomUnit
	} else {
		return randomUnit.Negate()
	}
}

func NewVec3(c [3]float64) Vec3 {
	return Vec3{c[0], c[1], c[2]}
}

func (v Vec3) Mult(u Vec3) Vec3 {
	return Vec3{v.X * u.X, v.Y * u.Y, v.Z * u.Z}
}

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

func RandomInterval(min, max float64) float64 {
	return min + (max-min)*rand.Float64()
}

func (v Vec3) GetUnitVec() Vec3 {
	vLen := v.Length()
	return v.Divide(vLen)
}

func (uv Vec3) Refract(n Vec3, etaiOverEtat float64) Vec3 {
	cosTheta := math.Min(uv.Negate().Dot(n), 1.0)
	rOutPerp := uv.Add(n.Scale(cosTheta)).Scale(etaiOverEtat)
	rOutParallel := n.Scale(-math.Sqrt(math.Abs(1.0 - rOutPerp.LengthSquared())))
	return rOutPerp.Add(rOutParallel)	
}