package sphere

import (
	"github.com/sergeev-s/raytracer/hittable"
	"github.com/sergeev-s/raytracer/ray"
	"github.com/sergeev-s/raytracer/vec"
	"math"
)

type Sphere struct {
	Center vec.Point3
	Radius float64
}

func NewSphere(center vec.Point3, radius float64) Sphere {
	return Sphere{
		Center: center,
		Radius: math.Max(0, radius),
	}
}

func (sphere Sphere) Hit(r ray.Ray, tMin, tMax float64) (hittable.HitRecord, bool) {
	oc := sphere.Center.Sub(r.Origin)
	a := r.Direction.LengthSquared()
	b := -2 * r.Direction.Dot(oc)
	c := oc.LengthSquared() - sphere.Radius*sphere.Radius
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return hittable.HitRecord{}, false
	}

	root := (-b - math.Sqrt(discriminant)) / (2 * a)

	if root <= tMin || root >= tMax {
		root = (-b + math.Sqrt(discriminant)) / (2 * a)
		if root <= tMin || root >= tMax {
			return hittable.HitRecord{}, false
		}
	}

	p := r.At(root)
	outwardNormal := p.Sub(sphere.Center).Divide(sphere.Radius)
	inFrontFace := r.Direction.Dot(outwardNormal) < 0
	normal := outwardNormal
	if !inFrontFace {
		normal = outwardNormal.Negate()
	}

	return hittable.HitRecord{
		P:      p,
		Normal: normal,
		T:      root,
	}, true
}
