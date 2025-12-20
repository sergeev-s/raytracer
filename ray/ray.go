package ray

import (
	"github.com/sergeev-s/raytracer/vec"
)

type Ray struct {
	Origin    vec.Point3
	Direction vec.Vec3
}

func NewRay(origin vec.Point3, direction vec.Vec3) Ray {
	return Ray{
		Origin:    origin,
		Direction: direction,
	}
}

func (r Ray) At(t float64) vec.Point3 {
	return r.Origin.Add(r.Direction.Scale(t))
}
